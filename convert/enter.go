package convert

import (
	"kubeimooc/convert/configmap"
	"kubeimooc/convert/node"
	"kubeimooc/convert/pod"
	"kubeimooc/convert/secret"
)

type ConvertGroup struct {
	PodConvert       pod.PodConvertGroup
	NodeConvert      node.Group
	ConfigMapConvert configmap.ConvertGroup
	SecretConvert    secret.ConvertGroup
}

var ConvertGroupApp = new(ConvertGroup)
