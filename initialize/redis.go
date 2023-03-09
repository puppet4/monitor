package initialize

import (
	"context"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"github.com/puppet4/monitor/global"
)

func Redis() {
	redisCfg := global.CONST_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {

	} else {
		fmt.Println(pong)
		global.CONST_REDIS = client
	}
}
