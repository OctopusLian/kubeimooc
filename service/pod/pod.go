package pod

import (
	"context"
	"errors"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	k8serror "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"kubeimooc/global"
	pod_req "kubeimooc/model/pod/request"
	pod_res "kubeimooc/model/pod/response"
	"strings"
	"time"
)

type PodService struct {
}

func (*PodService) GetPodList(namespace string, keyword string, nodeName string) (err error, _ []pod_res.PodListItem) {
	ctx := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return
	}
	podList := make([]pod_res.PodListItem, 0)
	for _, item := range list.Items {
		if nodeName != "" && item.Spec.NodeName != nodeName {
			continue
		}
		if strings.Contains(item.Name, keyword) {
			podItem := podConvert.PodK8s2ItemRes(item)
			podList = append(podList, podItem)
		}
	}
	return err, podList
}

func (*PodService) GetPodDetail(namespace string, name string) (podReq pod_req.Pod, err error) {
	ctx := context.TODO()
	podApi := global.KubeConfigSet.CoreV1().Pods(namespace)
	k8sGetPod, err := podApi.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]查询失败，detail：%s", namespace, name, err.Error())
		err = errors.New(errMsg)
		return
	}
	//将k8s pod 转为 pod request
	podReq = podConvert.K8s2ReqConvert.PodK8s2Req(*k8sGetPod)
	return
}

func (*PodService) DeletePod(namespace string, name string) error {
	background := metav1.DeletePropagationBackground
	var gracePeriodSeconds int64 = 0
	return global.KubeConfigSet.CoreV1().Pods(namespace).Delete(context.TODO(), name,
		metav1.DeleteOptions{
			GracePeriodSeconds: &gracePeriodSeconds,
			PropagationPolicy:  &background,
		})
}

func (*PodService) CreateOrUpdatePod(podReq pod_req.Pod) (msg string, err error) {
	//[no]update [no]patch [yes]delete+create
	k8sPod := podConvert.Req2K8sConvert.PodReq2K8s(podReq)
	ctx := context.TODO()
	podApi := global.KubeConfigSet.CoreV1().Pods(k8sPod.Namespace)
	if k8sGetPod, err := podApi.Get(ctx, k8sPod.Name, metav1.GetOptions{}); err == nil {
		//pod 参数是否合理
		k8sPodCopy := *k8sPod
		k8sPodCopy.Name = k8sPod.Name + "-validate"
		_, err := podApi.Create(ctx, &k8sPodCopy, metav1.CreateOptions{
			DryRun: []string{metav1.DryRunAll},
		})
		if err != nil {
			errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新失败，detail：%s", k8sPod.Namespace, k8sPod.Name, err.Error())
			return errMsg, err
		}
		//比如pod处于terminating状态 监听pod删除完毕之后 才开始创建pod
		var labelSelector []string
		for k, v := range k8sGetPod.Labels {
			labelSelector = append(labelSelector, fmt.Sprintf("%s=%s", k, v))
		}
		//label 格式 app=test,app2=test2
		watcher, err := podApi.Watch(ctx, metav1.ListOptions{
			LabelSelector: strings.Join(labelSelector, ","),
		})
		if err != nil {
			errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新失败，detail：%s", k8sPod.Namespace, k8sPod.Name, err.Error())
			return errMsg, err
		}
		//删除 -- 强制删除
		background := metav1.DeletePropagationBackground
		var gracePeriodSeconds int64 = 0
		err = podApi.Delete(ctx, k8sPod.Name, metav1.DeleteOptions{
			GracePeriodSeconds: &gracePeriodSeconds,
			PropagationPolicy:  &background,
		})
		if err != nil {
			errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新失败，detail：%s", k8sPod.Namespace, k8sPod.Name, err.Error())
			return errMsg, err
		}
	Loop:
		for {
			select {
			case event := <-watcher.ResultChan():
				k8sPodChan := event.Object.(*corev1.Pod)
				//查询k8s 判断是否已经删除 那么就不用判断删除事件了
				if _, err := podApi.Get(ctx, k8sPod.Name, metav1.GetOptions{}); k8serror.IsNotFound(err) {
					//重新创建
					if createdPod, err := podApi.Create(ctx, k8sPod, metav1.CreateOptions{}); err != nil {
						errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新失败，detail：%s", k8sPod.Namespace, k8sPod.Name, err.Error())
						return errMsg, err
					} else {
						successMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新成功", createdPod.Namespace, createdPod.Name)
						return successMsg, err
					}
				}
				switch event.Type {
				case watch.Deleted:
					if k8sPodChan.Name != k8sPod.Name {
						goto Loop
					}
					//重新创建
					if createdPod, err := podApi.Create(ctx, k8sPod, metav1.CreateOptions{}); err != nil {
						errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新失败，detail：%s", k8sPod.Namespace, k8sPod.Name, err.Error())
						return errMsg, err
					} else {
						successMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新成功", createdPod.Namespace, createdPod.Name)
						return successMsg, err
					}
				}
			case <-time.After(5 * time.Second):
				//重新创建
				if createdPod, err := podApi.Create(ctx, k8sPod, metav1.CreateOptions{}); err != nil {
					errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新失败，detail：%s", k8sPod.Namespace, k8sPod.Name, err.Error())
					return errMsg, err
				} else {
					successMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]更新成功", createdPod.Namespace, createdPod.Name)
					return successMsg, err
				}
			}
		}
	} else {
		if createdPod, err := podApi.Create(ctx, k8sPod, metav1.CreateOptions{}); err != nil {
			errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]创建失败，detail：%s", k8sPod.Namespace, k8sPod.Name, err.Error())
			return errMsg, err
		} else {
			successMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]创建成功", createdPod.Namespace, createdPod.Name)
			return successMsg, err
		}
	}
}
