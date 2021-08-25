/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package privilege

import (
	"k8s.io/klog"
	"volcano.sh/volcano/pkg/scheduler/api"
	"volcano.sh/volcano/pkg/scheduler/framework"
	"volcano.sh/volcano/pkg/scheduler/util"
	"volcano.sh/volcano/pkg/apis/scheduling"
)

type Action struct{}

func New() *Action {
	return &Action{}
}

func (alloc *Action) Name() string {
	return "privilege"
}

func (alloc *Action) Initialize() {}

func (alloc *Action) Execute(ssn *framework.Session) {
	klog.V(3).Infof("Enter Privilege ...")
	defer klog.V(3).Infof("Leaving Privilege ...")

	privilegesMap := map[api.QueueID]*util.PriorityQueue{}
	privilegeTasks := map[api.JobID]*util.PriorityQueue{}

	var underRequest []*api.JobInfo
	queues := map[api.QueueID]*api.QueueInfo{}

	for _, job := range ssn.Jobs {
		
		if vr := ssn.JobValid(job); vr != nil && !vr.Pass {
			klog.V(4).Infof("Job <%s/%s> Queue <%s> skip privilege, reason: %v, message %v", job.Namespace, job.Name, job.Queue, vr.Reason, vr.Message)
			continue
		}
		if job.PodGroup.Status.Phase == scheduling.PodGroupPending{
			continue
		}

		if !job.IsPrivilegeJob() {
			continue
		}
		if job.HasPrivileged() {
			continue
		}
		if !ssn.PrivilegeJobEnqueueable(job) {
			continue
		}
		if queue, found := ssn.Queues[job.Queue]; !found {
			continue
		} else if _, existed := queues[queue.UID]; !existed {
			klog.V(3).Infof("Added Queue <%s> for Job <%s/%s>",
				queue.Name, job.Namespace, job.Name)
			queues[queue.UID] = queue
		}
		if len(job.TaskStatusIndex[api.Pending]) != 0 && !ssn.JobPrivileged(job) {
			if _, found := privilegesMap[job.Queue]; !found {
				privilegesMap[job.Queue] = util.NewPriorityQueue(ssn.JobOrderFn)
			}
			privilegesMap[job.Queue].Push(job)
			underRequest = append(underRequest, job)
			privilegeTasks[job.UID] = util.NewPriorityQueue(ssn.TaskOrderFn)
			for _, task := range job.TaskStatusIndex[api.Pending] {
				privilegeTasks[job.UID].Push(task)
			}
		}
	}

	// Privilege between Jobs within Queue.
	for _, queue := range queues {
		for {
			privilegers := privilegesMap[queue.UID]

			// If no privilegers, no privilege.
			if privilegers == nil || privilegers.Empty() {
				klog.V(4).Infof("No privilegers in Queue <%s>, break.", queue.Name)
				break
			}

			privilegeJob := privilegers.Pop().(*api.JobInfo)
			
			if privilegeJob.PodGroup.Status.Phase == scheduling.PodGroupPending {
				privilegeJob.PodGroup.Status.Phase = scheduling.PodGroupPrivilege
				ssn.Jobs[privilegeJob.UID] = privilegeJob
			}
			
			if privilegeJob.HasPrivileged() {
				break
			}

			stmt := framework.NewStatement(ssn)
			assigned := false
			for {
				if privilegeJob.HasPrivileged() {
					break
				}

				// If not privilege tasks, next job.
				if privilegeTasks[privilegeJob.UID].Empty() {
					klog.V(3).Infof("No preemptor task in job <%s/%s>.",
						privilegeJob.Namespace, privilegeJob.Name)
					break
				}

				privileger := privilegeTasks[privilegeJob.UID].Pop().(*api.TaskInfo)
				if privileged, _ := privilege(ssn, stmt, privileger, func(task *api.TaskInfo) bool {
					// Ignore non running task.
					if task.Status != api.Running {
						return false
					}
					if task.WillEvict {
						return false
					}
					// Ignore task with empty resource request.
					if task.Resreq.IsEmpty() {
						return false
					}
					job, found := ssn.Jobs[task.Job]
					if !found {
						return false
					}
					// Privilege other jobs within queue
					return job.Queue == privilegeJob.Queue && privileger.Job != task.Job
				}); privileged {
					assigned = true
				}
			}
			if ssn.JobPrivileged(privilegeJob) {
				privilegeJob.SetPrivileged()
				stmt.Commit()
			} else {
				stmt.Discard()
				continue
			}

			if assigned {
				privilegers.Push(privilegeJob)
			}
		}

	}
}

func (alloc *Action) UnInitialize() {}

func privilege(
	ssn *framework.Session,
	stmt *framework.Statement,
	privileger *api.TaskInfo,
	filter func(*api.TaskInfo) bool,
) (bool, error) {
	assigned := false

	allNodes := util.GetNodeList(ssn.Nodes)

	predicateNodes, _ := util.PredicateNodes(privileger, allNodes, ssn.PredicateFn)

	nodeScores := util.PrioritizeNodes(privileger, predicateNodes, ssn.BatchNodeOrderFn, ssn.NodeOrderMapFn, ssn.NodeOrderReduceFn)

	selectedNodes := util.SortNodes(nodeScores)
	for _, node := range selectedNodes {
		klog.V(3).Infof("Considering Task <%s/%s> on Node <%s>.",
			privileger.Namespace, privileger.Name, node.Name)

		var prvilegees []*api.TaskInfo
		for _, task := range node.Tasks {

			if filter == nil {
				prvilegees = append(prvilegees, task.Clone())
			} else if filter(task) {
				prvilegees = append(prvilegees, task.Clone())
			}
		}
		victims := ssn.Preemptable(privileger, prvilegees)

		if err := util.ValidatePrivilegeVictims(privileger, node, victims); err != nil {
			klog.V(3).Infof("No validated victims on Node <%s>: %v", node.Name, err)
			continue
		}

		victimsQueue := util.NewPriorityQueue(func(l, r interface{}) bool {
			return !ssn.TaskOrderFn(l, r)
		})
		for _, victim := range victims {
			victimsQueue.Push(victim)
		}

		job, found := ssn.Jobs[privileger.Job]
		if !found {
			klog.Errorf("Failed to privilege Task <%s/%s> on Node <%s> because job not found",
				privileger.Namespace, privileger.Name, node.Name)
		}

		// Preempt victims for tasks, pick lowest priority task first.
		preempted := api.EmptyResource()

		for !victimsQueue.Empty() {
			// If reclaimed enough resources, break loop to avoid Sub panic.
			if privileger.InitResreq.LessEqual(node.PrivilegeIdle()) && ssn.CurPrivilegeJobEnqueueable(job) {
				break
			}
			privilegee := victimsQueue.Pop().(*api.TaskInfo)
			klog.V(3).Infof("Try to privilege Task <%s/%s> for Tasks <%s/%s>",
				privilegee.Namespace, privilegee.Name, privileger.Namespace, privileger.Name)
			if err := stmt.LableJobEvict(privileger, privilegee, "preemptByPrivileger"); err != nil {
				klog.Errorf("Failed to privilege Task <%s/%s> for Tasks <%s/%s>: %v",
					privilegee.Namespace, privilegee.Name, privileger.Namespace, privileger.Name, err)
				continue
			}
			preempted.Add(privilegee.Resreq)
		}

		if privileger.InitResreq.LessEqual(node.PrivilegeIdle()) && ssn.CurPrivilegeJobEnqueueable(job) {
			if err := stmt.Privilege(privileger, node.Name); err != nil {
				klog.Errorf("Failed to privilege Task <%s/%s> on Node <%s>",
					privileger.Namespace, privileger.Name, node.Name)
			}
			assigned = true
			break
		}
	
	}

	return assigned, nil
}