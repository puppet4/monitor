package service

import "github.com/puppet4/monitor/service/system"

type ServiceGroup struct {
	SystemServiceGroup system.SystemGroup
}

var ServiceGroupApp = new(ServiceGroup)
