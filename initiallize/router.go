package initiallize

import (
	"github.com/gin-gonic/gin"
	"kubeimooc/middleware"
	"kubeimooc/router"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors)
	examplGroup := router.RouterGroupApp.ExampleRouterGroup
	k8sGroup := router.RouterGroupApp.K8SRouterGroup
	examplGroup.InitExample(r)
	k8sGroup.InitK8SRouter(r)
	return r
	//r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
