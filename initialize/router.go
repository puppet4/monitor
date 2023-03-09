package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/puppet4/monitor/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System
	PrivateGroup := Router.Group("system")
	{
		systemRouter.InitSystemRouter(PrivateGroup)
	}
	return Router
}
