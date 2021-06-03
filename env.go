package main

import (
	"fmt"
	"os"
)

// 兼容Drone插件和普通使用
// 优先使用普通模式
// 没有配置再加载Drone配置
func env(env string) (config string) {
	if config = os.Getenv(env); "" != config {
		return
	}
	config = os.Getenv(fmt.Sprintf("PLUGIN_%s", env))

	return
}
