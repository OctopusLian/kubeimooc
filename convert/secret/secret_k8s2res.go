package secret

import (
	corev1 "k8s.io/api/core/v1"
	"kubeimooc/utils"
)
import secretres "kubeimooc/model/secret/reponse"

type K8s2Res struct {
}

func (K8s2Res) SecretK8s2ResItemConvert(secret corev1.Secret) secretres.Secret {
	return secretres.Secret{
		Name:      secret.Name,
		Namespace: secret.Namespace,
		Type:      secret.Type,
		DataNum:   len(secret.Data),
		Age:       secret.CreationTimestamp.Unix(),
	}
}

func (K8s2Res) SecretK8s2ResDetailConvert(secret corev1.Secret) secretres.Secret {
	return secretres.Secret{
		Name:      secret.Name,
		Namespace: secret.Namespace,
		Type:      secret.Type,
		DataNum:   len(secret.Data),
		Age:       secret.CreationTimestamp.Unix(),
		Data:      utils.ToListWithMapByte(secret.Data),
		Labels:    utils.ToList(secret.Labels),
	}
}
