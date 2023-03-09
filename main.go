package main

import (
	"github.com/puppet4/monitor/core"
	"github.com/puppet4/monitor/global"
)

func main() {
	//初始化Viper
	global.CONST_VP = core.Viper()
	core.RunWindowsServer()
}
