package configmap

import (
	corev1 "k8s.io/api/core/v1"
	configmapres "kubeimooc/model/configmap/response"
	"kubeimooc/utils"
)

//

type K82Res struct {
}

func (K82Res) GeCmReqItem(configMap corev1.ConfigMap) configmapres.ConfigMap {
	return configmapres.ConfigMap{
		Name:      configMap.Name,
		Namespace: configMap.Namespace,
		DataNum:   len(configMap.Data),
		Age:       configMap.CreationTimestamp.Unix(),
	}
}

func (this K82Res) GeCmReqDetail(configMap corev1.ConfigMap) configmapres.ConfigMap {
	detail := this.GeCmReqItem(configMap)
	detail.Labels = utils.ToList(configMap.Labels)
	detail.Data = utils.ToList(configMap.Data)
	return detail
}
