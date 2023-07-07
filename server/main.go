package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"

	mclient "gitee.com/openeuler/PilotGo-plugin-mysql/client"
	"gitee.com/openeuler/PilotGo-plugin-mysql/config"
	"gitee.com/openeuler/PilotGo-plugin-mysql/httpserver"
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

	if err := mclient.Init(PluginInfo); err != nil {
		fmt.Printf("failed to init plugin client: %s\n", err.Error())
		os.Exit(-1)
	}

	if err := httpserver.Start(); err != nil {
		fmt.Printf("failed to start http server: %s\n", err.Error())
		os.Exit(-1)
	}

	// 信号监听
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		s := <-c
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			logger.Info("signal interrupted: %s", s.String())
			// TODO: DO EXIT

			// http server exit
			logger.Info("stopping http server")
			httpserver.Stop()

			goto EXIT
		default:
			logger.Info("unknown signal: %s", s.String())
		}
	}

EXIT:
	logger.Info("exit system, bye~")

}
