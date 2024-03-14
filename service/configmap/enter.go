package configmap

import "kubeimooc/convert"

type ServiceGroup struct {
	ConfigMapService
}

var configConvert = convert.ConvertGroupApp.ConfigMapConvert
