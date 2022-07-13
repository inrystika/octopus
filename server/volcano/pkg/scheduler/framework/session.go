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

package framework

import (
	typeCache "server/volcano/pkg/scheduler/cache"

	typeApi "server/volcano/pkg/scheduler/api"

	"volcano.sh/volcano/pkg/scheduler/api"
)

// Session information for the current session
type Session struct {
	// typedef queue
	TypeQueues map[api.QueueID]*typeApi.QueueInfo
	// typedef plugins
	TypePlugins map[string]Plugin
}

func openSession(cache typeCache.Cache) *Session {
	ssn := &Session{
		TypeQueues:  map[api.QueueID]*typeApi.QueueInfo{},
		TypePlugins: map[string]Plugin{},
	}

	snapshot := cache.Snapshot()
	ssn.TypeQueues = snapshot.TypeQueues

	return ssn
}

func closeSession(ssn *Session) {
	ssn.TypePlugins = nil
}
