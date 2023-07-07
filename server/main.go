package main

import (
	"fmt"
	"os"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"

	"gitee.com/openeuler/PilotGo-plugin-mysql/config"
)

const Version = "0.0.1"

var PluginInfo = &client.PluginInfo{
	Name:        "mysql",
	Version:     Version,
	Description: "mysql监控管理插件",
	Author:      "guozhengxin",
	Email:       "guozhengxin@kylinos.cn",
}

func main() {
	fmt.Println("Thanks to choose PilotGo!")

	if err := config.ConfigInit("./config.yaml"); err != nil {
		fmt.Printf("failed to init config module: %s\n", err.Error())
		os.Exit(-1)
	}

	if err := logger.Init(config.Config().LogConf); err != nil {
		fmt.Printf("failed to init logger module: %s\n", err.Error())
		os.Exit(-1)
	}

}
