/*
Copyright 2020 The Volcano Authors.

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

package framework

import (
	vcclientset "server/apis/pkg/client/clientset/versioned"

	"volcano.sh/volcano/pkg/controllers/framework"
)

// ControllerOption is the main context object for the controllers.
type ControllerOption struct {
	framework.ControllerOption
	VolcanoClient vcclientset.Interface
}

// Controller is the interface of all controllers.
type Controller interface {
	Name() string
	Initialize(opt *ControllerOption) error
	// Run run the controller
	Run(stopCh <-chan struct{})
}
