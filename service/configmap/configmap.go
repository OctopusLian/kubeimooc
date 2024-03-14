package configmap

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeimooc/global"
	configmapreq "kubeimooc/model/configmap/request"
	configmapres "kubeimooc/model/configmap/response"
	"strings"
)

//

type ConfigMapService struct {
}

// 查询详情
func (ConfigMapService) GetConfigMapDetail(namespace string, name string) (cm configmapres.ConfigMap, err error) {
	configMapK8s, err := global.KubeConfigSet.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return
	}
	cm = configConvert.GeCmReqDetail(*configMapK8s)
	return
}

// 查询configmap list
func (ConfigMapService) GetConfigMapList(namespace string, keyword string) (cmList []configmapres.ConfigMap, err error) {
	//1 从k8s查询
	list, errGetK8s := global.KubeConfigSet.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		err = errGetK8s
		return
	}
	//2 转换为res(filter)
	configMapList := make([]configmapres.ConfigMap, 0)
	for _, item := range list.Items {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		configMapList = append(configMapList, configConvert.GeCmReqItem(item))
	}
	return configMapList, nil
}
func (*ConfigMapService) DeleteConfigMap(ns string, name string) error {
	return global.KubeConfigSet.CoreV1().ConfigMaps(ns).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (*ConfigMapService) CreateOrUpdateConfigMap(configReq configmapreq.ConfigMap) error {
	// 将 request 转为 k8s 结构
	configMap := configConvert.CmReq2K8sConvert(configReq)
	ctx := context.TODO()
	//判断是否存在
	_, errSearch := global.KubeConfigSet.CoreV1().ConfigMaps(configReq.Namespace).Get(ctx, configReq.Name, metav1.GetOptions{})
	if errSearch == nil {
		_, err := global.KubeConfigSet.CoreV1().ConfigMaps(configMap.Namespace).Update(ctx, configMap, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
	} else {
		_, err := global.KubeConfigSet.CoreV1().ConfigMaps(configMap.Namespace).Create(ctx, configMap, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
