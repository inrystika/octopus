package framework

import (
	"scheduler/pkg/applet/conf"
	typeJob "volcano.sh/volcano/pkg/apis/batch/v1alpha1"
	api "scheduler/pkg/pipeline/apis/common"
)

type AppletDelegator interface {
	Applet
	Linker
}

type Applet interface {
	ExecTemplateTranslator(packet *AppletPacket) (*typeJob.Job, error)
	ExecFactorGenerator(packet *AppletPacket) (*Factor, error)
	ExecAccessGate(packet *AppletPacket) (*Accessor, error)
	ExecTemplateDecorator(packet *AppletPacket) (*typeJob.Job, error)
	ExecSchedulerBinder(packet *AppletPacket) (*typeJob.Job, error)
	ExecLifeHook(packet *AppletPacket) ([]byte, error)
}

type Linker interface {
	RegisterApplet(feature *api.Feature, applet Applet, conf *conf.AppletConfiguration) error
	Start(fn WrapperFeatureFn) error
	Stop() error
}
