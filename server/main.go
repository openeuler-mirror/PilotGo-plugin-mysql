package main

import (
	"fmt"

	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
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

}
