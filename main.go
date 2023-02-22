package main

import (
	"fmt"
	"github.com/puppet4/monitor/core"
	"github.com/puppet4/monitor/global"
)

func main() {
	//初始化Viper
	global.CONST_VP = core.Viper()
	fmt.Println(global.CONST_CONFIG.Mysql)
}
