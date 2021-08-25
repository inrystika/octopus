// MIT License
//
// Copyright (c) PCL. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE
//

package kubernetes

import (
	"fmt"

	typeJob "volcano.sh/volcano/pkg/apis/batch/v1alpha1"
	"k8s.io/client-go/tools/cache"

	"scheduler/pkg/pipeline/constants/jobstate"
)

func deltaFIFOObjToTaskSet(obj interface{}) (*typeJob.Job, string) {

	taskset, ok := obj.(*typeJob.Job)

	if ok {
		return taskset, ""
	}

	deletedFinalStateUnknown, ok := obj.(cache.DeletedFinalStateUnknown)

	if !ok {
		return nil, fmt.Sprintf("Failed to convert obj to TaskSet or DeletedFinalStateUnknown: %#v", obj)
	}

	taskset, ok = deletedFinalStateUnknown.Obj.(*typeJob.Job)

	if !ok {

		return nil, fmt.Sprintf("Failed to convert DeletedFinalStateUnknown.Obj to TaskSet: %#v", deletedFinalStateUnknown)
	}

	return taskset, ""
}

func MapPhaseToState(phase typeJob.JobPhase) string {

	if phase == typeJob.Pending || phase == typeJob.Restarting {
		return jobstate.PENDING
	}

	if phase == typeJob.Running {
		return jobstate.RUNNING
	}

	if phase == typeJob.Failed || phase == typeJob.Aborting || phase == typeJob.Aborted{
		return jobstate.FAILED
	}
	if phase == typeJob.Completed {
		return jobstate.SUCCEEDED
	}
	return jobstate.UNKNOWN
}

func getJobHeader(taskset *typeJob.Job) []byte {

	if taskset.Annotations["header"] != "" {
		return []byte(taskset.Annotations["header"])
	}

	return nil
}
