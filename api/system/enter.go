package system

import "github.com/puppet4/monitor/service"

type ApiGroup struct {
	SystemApi SystemApi
}

var (
	systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
)
