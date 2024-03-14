package k8s

import (
	"github.com/gin-gonic/gin"
	"kubeimooc/api"
)

type K8sRouter struct {
}

func (*K8sRouter) InitK8SRouter(r *gin.Engine) {
	group := r.Group("/k8s")
	apiGroup := api.ApiGroupApp.K8SApiGroup
	group.POST("/pod", apiGroup.CreateOrUpdatePod)
	group.GET("/pod/:namespace", apiGroup.GetPodListOrDetail)
	group.DELETE("/pod/:namespace/:name", apiGroup.DeletePod)
	group.GET("/namespace", apiGroup.GetNamespaceList)
	///////////////////

	//nodeScheduling
	group.GET("/node", apiGroup.GetNodeDetailOrList)
	group.PUT("/node/label", apiGroup.UpdateNodeLabel)
	group.PUT("/node/taint", apiGroup.UpdateNodeTaint)

	//******************ConfigMap************************//
	group.POST("/configmap", apiGroup.CreateOrUpdateConfigMap)
	group.GET("/configmap/:namespace", apiGroup.GetConfigMapDetailOrList)
	group.DELETE("/configmap/:namespace/:name", apiGroup.DeleteConfigMap)

	//*******************Secret***********************//
	group.POST("/secret", apiGroup.CreateOrUpdateSecret)
	group.GET("/secret/:namespace", apiGroup.GetSecretDetailOrList)
	group.DELETE("/secret/:namespace/:name", apiGroup.DeleteSecret)

}
