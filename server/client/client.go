package client

import (
	"github.com/gin-gonic/gin"

	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
)

var global_client *client.Client

func Init(info *client.PluginInfo) error {
	global_client = client.DefaultClient(info)

	return nil
}

func RegisterHandlers(engine *gin.Engine) {
	global_client.RegisterHandlers(engine)
}
