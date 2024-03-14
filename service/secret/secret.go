package secret

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeimooc/global"
	secretres "kubeimooc/model/secret/reponse"
	secretreq "kubeimooc/model/secret/request"
	"strings"
)

type SecretService struct {
}

func (SecretService) GetSecretDetail(namespace string, name string) (*secretres.Secret, error) {
	secretK8s, err := global.KubeConfigSet.CoreV1().Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	secretRes := secretConvert.SecretK8s2ResDetailConvert(*secretK8s)
	return &secretRes, err
}
func (SecretService) GetSecretList(namespace string, keyword string) ([]secretres.Secret, error) {
	list, err := global.KubeConfigSet.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	secretResList := make([]secretres.Secret, 0)
	for _, item := range list.Items {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		secretRes := secretConvert.SecretK8s2ResItemConvert(item)
		secretResList = append(secretResList, secretRes)
	}
	return secretResList, err
}

func (SecretService) DeleteSecret(namespace string, name string) error {
	return global.KubeConfigSet.CoreV1().Secrets(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (SecretService) CreateOrUpdateSecret(secret secretreq.Secret) (err error) {
	ctx := context.TODO()
	secretK8s := secretConvert.SecretReq2K8sConvert(secret)
	//查询是否存在
	secretApi := global.KubeConfigSet.CoreV1().Secrets(secretK8s.Namespace)
	_, errGet := secretApi.Get(ctx, secret.Name, metav1.GetOptions{})
	if errGet == nil {
		_, err = secretApi.Update(ctx, &secretK8s, metav1.UpdateOptions{})
	} else {
		_, err = secretApi.
			Create(ctx, &secretK8s, metav1.CreateOptions{})
	}
	return
}
