package common

type OnDeploymentAdd func(obj interface{})
type OnDeploymentUpdate func(old, obj interface{})
type OnDeploymentDelete func(obj interface{})
