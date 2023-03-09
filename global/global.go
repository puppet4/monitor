package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/puppet4/monitor/config"
	"github.com/spf13/viper"
)

var (
	CONST_VP     *viper.Viper
	CONST_CONFIG config.Server
	CONST_REDIS  *redis.Client
)
