package main

import (
	"github.com/puppet4/monitor/core"
	"github.com/puppet4/monitor/global"
)

func main() {
	//初始化Viper
	global.CONST_VP = core.Viper()
	//初始化zap日志库
	global.CONST_LOG = core.Zap()
	core.RunWindowsServer()
}
