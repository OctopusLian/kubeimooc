package k8s

import (
	"github.com/gin-gonic/gin"
	secretreq "kubeimooc/model/secret/request"
	"kubeimooc/response"
)

//

type SecretApi struct {
}

func (*SecretApi) CreateOrUpdateSecret(ctx *gin.Context) {
	var secretReq secretreq.Secret
	if err := ctx.ShouldBind(&secretReq); err != nil {
		response.FailWithMessage(ctx, "参数绑定失败！")
		return
	}
	if err := secretService.CreateOrUpdateSecret(secretReq); err != nil {
		response.FailWithMessage(ctx, "创建或更新Secret失败："+err.Error())
		return
	}
	response.Success(ctx)
}
func (*SecretApi) GetSecretDetailOrList(ctx *gin.Context) {
	name := ctx.Query("name")
	namespace := ctx.Param("namespace")
	keyword := ctx.Query("keyword")
	var data any
	var err error
	if name != "" {
		data, err = secretService.GetSecretDetail(namespace, name)
	} else {
		data, err = secretService.GetSecretList(namespace, keyword)
	}
	if err != nil {
		response.SuccessWithMessage(ctx, "获取Secret失败："+err.Error())
	} else {
		response.SuccessWithDetailed(ctx, "获取Secret成功", data)
	}
}
func (*SecretApi) DeleteSecret(ctx *gin.Context) {
	if err := secretService.DeleteSecret(ctx.Param("namespace"), ctx.Param("name")); err != nil {
		response.SuccessWithMessage(ctx, "删除Secret失败："+err.Error())
	} else {
		response.Success(ctx)
	}
}
