package httpserver

import (
	"os"

	"github.com/gin-gonic/gin"

	"gitee.com/openeuler/PilotGo-plugin-mysql/config"
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
)

func Start() error {
	go func() {
		conf := config.Config().HttpServerConf
		router := gin.New()
		logger.Info("start http service on: http://%s", conf.Addr)
		if err := router.Run(conf.Addr); err != nil {
			logger.Error("failed to start http service, error: %v", err)
			os.Exit(-1)
		}
	}()
	return nil
}
