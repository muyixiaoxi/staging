package main

import (
	"fmt"
	"staging/logger"
	"staging/pkg/settings"
	"staging/service"
)

func main() {
	// 加载配置
	app, err := settings.Init()
	if err != nil {
		fmt.Printf("init setting failed,error: %v\n", err)
	}

	// 初始化日志
	err = logger.Init(app.LogConfig, app.Mode)
	if err != nil {
		fmt.Printf("init logger failed,error: %v\n", err)
	}
	service.Init(app)
}
