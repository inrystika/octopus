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
	"scheduler/pkg/pipeline/app"
	"scheduler/pkg/pipeline/config"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

func NewWithKubeImp(app *app.App, config *rest.Config, evictConfig *config.EvictConfig, userCenterConfig *config.UserCenterConfig, chargeConfig *config.ChargeConfig, kube KubeInterface) *Service {
	s := &Service{
		config:           config,
		evictConfig:      evictConfig,
		userCenterConfig: userCenterConfig,
		chargeConfig:     chargeConfig,
		app:              app,
		kube:             kube,
		mailboxes:        make(map[string]*eventBox, 10),
	}

	if app != nil && nil != app.Logger() {
		s.logger = app.Logger().Named("kube")
	}

	s.kube.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    s.onTaskSetAdd,
		UpdateFunc: s.onTaskSetUpdate,
		DeleteFunc: s.onTaskSetDelete,
	})

	return s
}

func New(app *app.App, k8sConfig *rest.Config, evictConfig *config.EvictConfig, userCenterConfig *config.UserCenterConfig, chargeConfig *config.ChargeConfig) *Service {
	kube := newkubeImp(k8sConfig, evictConfig, userCenterConfig, chargeConfig)
	return NewWithKubeImp(app, k8sConfig, evictConfig, userCenterConfig, chargeConfig, kube)
}

func (s *Service) Run() {
	go PrivilegeTimer(s)

	s.stopChan = make(chan struct{}, 0)

	defer func() {
		if nil != s.logger {
			s.logger.Sync()
		}
	}()

	s.kube.Run(s.stopChan)

	<-s.stopChan
}

func (s *Service) Shutdown() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, box := range s.mailboxes {
		if nil != box {
			box.shutdown()
		}
	}
	s.stopChan <- struct{}{}
	s.kube.Shutdown()

}
