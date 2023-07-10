package client

import (
	"github.com/gin-gonic/gin"

	"gitee.com/openeuler/PilotGo-plugins/sdk/plugin/client"
)

var globalClient *client.Client

func Init(info *client.PluginInfo) error {
	globalClient = client.DefaultClient(info)

	return nil
}

func RegisterHandlers(engine *gin.Engine) {
	globalClient.RegisterHandlers(engine)
}
