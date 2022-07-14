/*
Copyright 2021.

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

package agent

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	agentv1 "nodeagent/apis/agent/v1"
	actSvc "nodeagent/controllers/agent/service"
)

// NodeActionReconciler reconciles a NodeAction object
type NodeActionReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	NodeName string
	Svc      actSvc.ActionService
}

//+kubebuilder:rbac:groups=agent.octopus.openi.org.cn,resources=nodeactions,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=agent.octopus.openi.org.cn,resources=nodeactions/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=agent.octopus.openi.org.cn,resources=nodeactions/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the NodeAction object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *NodeActionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	na := agentv1.NodeAction{}
	if err := r.Get(ctx, req.NamespacedName, &na); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	// only handle NodeAction for local node
	// if Reconciler.NodeName is "", it can work in local development
	if na.Spec.NodeName == "" || r.NodeName != na.Spec.NodeName {
		if r.NodeName != "" {
			return ctrl.Result{}, nil
		}
	}
	nodeActionState := na.Status.State
	switch nodeActionState {
	case "":
		// start action
		if err := r.doAction(ctx, na); err != nil {
			return ctrl.Result{}, err
		}
	default:
		// if resource has Status.state, it mean action is running, return reconcile
		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, nil
}

func (r *NodeActionReconciler) doAction(ctx context.Context, nodeAction agentv1.NodeAction) error {
	logger := log.FromContext(ctx)
	// first update resource state, due to stop event handler after now
	nodeAction.Status.State = agentv1.ActionRunningState
	if err := r.Status().Update(ctx, &nodeAction); err != nil {
		logger.Error(err, err.Error())
		return err
	}

	// run g to work service
	go func() {
		results := r.Svc.Do(ctx, nodeAction.Spec.Actions)
		commandStatuses := []*agentv1.CommandStatus{}
		for _, result := range results {
			cs := agentv1.CommandStatus{
				Name: result.Name,
			}
			if result.Error != nil {
				cs.Result = agentv1.CommandFailedResult
				cs.Reason = result.Error.Error()

				logger.Error(result.Error, "action error", "CommandName", result.Name, "ActionName", nodeAction.Name)
			} else {
				cs.Result = agentv1.CommandSucceedResult
			}
			commandStatuses = append(commandStatuses, &cs)
		}
		nodeAction.Status.Actions = commandStatuses
		nodeAction.Status.State = agentv1.ActionCompletedState
		if err := r.Status().Update(ctx, &nodeAction); err != nil {
			logger.Error(err, err.Error())
		}
	}()
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NodeActionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&agentv1.NodeAction{}).
		Complete(r)
}
