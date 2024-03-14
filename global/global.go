package global

import (
	"k8s.io/client-go/kubernetes"
	"kubeimooc/config"
)

var (
	CONF          config.Server
	KubeConfigSet *kubernetes.Clientset
)
