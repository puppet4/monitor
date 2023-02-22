package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/puppet4/monitor/core/internal"
	"github.com/puppet4/monitor/global"
	"github.com/spf13/viper"
	"os"
)

func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		//判断命令行参数是否为空
		if config == "" {
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigTestFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				}
			} else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				config = configEnv
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigTestFile)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err.Error()))
	}
	v.WatchConfig()

	//添加配置文件变动监听事件
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config is changed", in.Name)
		if err := v.Unmarshal(&global.CONST_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.CONST_CONFIG); err != nil {
		fmt.Println(err)
	}
	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	//global.CONST_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}
