package core

import (
	"fmt"
	"github.com/puppet4/monitor/global"
	"github.com/puppet4/monitor/initialize"
	"go.uber.org/zap"
	"time"
)

type Server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.CONST_CONFIG.System.UseMultipoint || global.CONST_CONFIG.System.UseRedis {
		//InitializeRedis()
		initialize.Redis()
	}

	Router := initialize.Routers()
	Router.Static("/from-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.CONST_CONFIG.System.Addr)
	s := initServer(address, Router)

	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.CONST_LOG.Info("server run success on ", zap.String("address", address))

	s.ListenAndServe().Error()
}
