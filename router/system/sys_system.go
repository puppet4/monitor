package system

import (
	"github.com/gin-gonic/gin"
	"github.com/puppet4/monitor/api"
)

type SysRouter struct{}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("system")
	systemApi := api.ApiGroupApp.SystemApiGroup.SystemApi
	{
		//获取服务器信息
		sysRouter.GET("getServerInfo", systemApi.GetServerInfo)
	}
}
