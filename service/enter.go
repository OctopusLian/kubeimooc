package service

import (
	"kubeimooc/service/configmap"
	"kubeimooc/service/node"
	"kubeimooc/service/pod"
	"kubeimooc/service/secret"
)

type ServiceGroup struct {
	PodServiceGroup       pod.PodServiceGroup
	NodeServiceGroup      node.Group
	ConfigMapServiceGroup configmap.ServiceGroup
	SecretServiceGroup    secret.SeviceGroup
}

var ServiceGroupApp = new(ServiceGroup)
