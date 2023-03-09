package system

import "github.com/puppet4/monitor/utils"

type SystemConfigService struct{}

func (systemConfigService *SystemConfigService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); nil != err {
		return &s, err
	}
	if s.Ram, err = utils.InitRAM(); nil != err {
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); nil != err {
		return &s, err
	}
	return &s, nil
}
