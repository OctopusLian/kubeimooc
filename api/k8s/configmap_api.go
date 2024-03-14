package k8s

import (
	"github.com/gin-gonic/gin"
	configmapreq "kubeimooc/model/configmap/request"
	"kubeimooc/response"
)

type ConfigMapApi struct {
}

// 创建或修改
func (ConfigMapApi) CreateOrUpdateConfigMap(c *gin.Context) {
	var configMapReq configmapreq.ConfigMap
	err := c.ShouldBind(&configMapReq)
	if err != nil {
		response.FailWithMessage(c, "ConfigMap参数解析失败！")
		return
	}
	err = configMapService.CreateOrUpdateConfigMap(configMapReq)
	if err != nil {
		response.FailWithMessage(c, err.Error())
	} else {
		response.Success(c)
	}
}

// 查询configmap和详情或列表
func (ConfigMapApi) GetConfigMapDetailOrList(c *gin.Context) {
	name := c.Query("name")
	namespace := c.Param("namespace")
	keyword := c.Query("keyword")
	if name == "" {
		list, err := configMapService.GetConfigMapList(namespace, keyword)
		if err != nil {
			response.FailWithMessage(c, "查询ConfigMap列表失败！")
			return
		} else {
			response.SuccessWithDetailed(c, "查询ConfigMap列表成功！", list)
		}
	} else {
		detail, err := configMapService.GetConfigMapDetail(namespace, name)
		if err != nil {
			response.FailWithMessage(c, "查询ConfigMap失败！")
			return
		} else {
			response.SuccessWithDetailed(c, "查询ConfigMap成功！", detail)
		}
	}
}

// 删除
func (ConfigMapApi) DeleteConfigMap(c *gin.Context) {
	err := configMapService.DeleteConfigMap(c.Param("namespace"), c.Param("name"))
	if err != nil {
		response.FailWithMessage(c, err.Error())
	} else {
		response.Success(c)
	}
}
