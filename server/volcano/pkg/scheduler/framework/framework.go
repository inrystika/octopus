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
	"time"

	typeCache "server/volcano/pkg/scheduler/cache"

	vcCache "volcano.sh/volcano/pkg/scheduler/cache"

	"volcano.sh/volcano/pkg/scheduler/conf"
	"volcano.sh/volcano/pkg/scheduler/framework"
	"volcano.sh/volcano/pkg/scheduler/metrics"
)

// OpenSession start the session
func OpenSession(cache vcCache.Cache, queueCache typeCache.Cache, tiers []conf.Tier, configurations []conf.Configuration) (*framework.Session, *Session) {
	ssn := framework.OpenSession(cache, tiers, configurations)
	typeSsn := openSession(queueCache)
	for _, tier := range tiers {
		for _, plugin := range tier.Plugins {
			if pb, found := GetPluginBuilder(plugin.Name); found {
				plugin := pb(plugin.Arguments)
				typeSsn.TypePlugins[plugin.Name()] = plugin
				onSessionOpenStart := time.Now()
				plugin.OnSessionOpen(ssn, typeSsn)
				metrics.UpdatePluginDuration(plugin.Name(), metrics.OnSessionOpen, metrics.Duration(onSessionOpenStart))
			}
		}
	}
	return ssn, typeSsn
}

// CloseSession close the session
func CloseSession(ssn *framework.Session, typeSsn *Session) {
	framework.CloseSession(ssn)
	for _, plugin := range typeSsn.TypePlugins {
		onSessionCloseStart := time.Now()
		plugin.OnSessionClose(ssn, typeSsn)
		metrics.UpdatePluginDuration(plugin.Name(), metrics.OnSessionClose, metrics.Duration(onSessionCloseStart))
	}
	closeSession(typeSsn)
}
