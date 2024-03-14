package secret

import "kubeimooc/convert"

type SeviceGroup struct {
	SecretService
}

var secretConvert = convert.ConvertGroupApp.SecretConvert
