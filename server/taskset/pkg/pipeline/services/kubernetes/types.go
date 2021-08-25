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
	"context"
	"scheduler/pkg/common/list"
	api "scheduler/pkg/pipeline/apis/module"
	"scheduler/pkg/pipeline/app"
	config "scheduler/pkg/pipeline/config"
	"sync"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	typeJob "volcano.sh/volcano/pkg/apis/batch/v1alpha1"
	vcclientset "volcano.sh/volcano/pkg/client/clientset/versioned"
	batchinformer "volcano.sh/volcano/pkg/client/informers/externalversions/batch/v1alpha1"
	batchlister "volcano.sh/volcano/pkg/client/listers/batch/v1alpha1"
)

type KubeInterface interface {
	Run(chan struct{})
	Shutdown()
	Create(*typeJob.Job) error
	Get(namespace, jobID string) (*typeJob.Job, error)
	Delete(namespace, jobID string) error
	AddEventHandler(cache.ResourceEventHandler)
	GetClient() *kubernetes.Clientset
	GetVcClient() *vcclientset.Clientset
}

type eventBox struct {
	name      string
	workerNum int
	box       *list.List
	handler   api.KubeEventListener
	ctx       context.Context
	cancel    context.CancelFunc
	event     chan *api.JobEvent
	dispatch  chan *api.JobEvent
	ack       chan int
	wg        sync.WaitGroup
}

type Service struct {
	config           *rest.Config
	evictConfig      *config.EvictConfig
	userCenterConfig *config.UserCenterConfig
	chargeConfig     *config.ChargeConfig
	app              *app.App
	mutex            sync.RWMutex
	kube             KubeInterface
	mailboxes        map[string]*eventBox
	stopChan         chan struct{}
	logger           *zap.Logger
}

type kubeImp struct {
	config           *rest.Config
	evictConfig      *config.EvictConfig
	userCenterConfig *config.UserCenterConfig
	chargeConfig     *config.ChargeConfig
	client           *vcclientset.Clientset
	k8sClient        *kubernetes.Clientset
	jobInformer      batchinformer.JobInformer
	jobLister        batchlister.JobLister
}
