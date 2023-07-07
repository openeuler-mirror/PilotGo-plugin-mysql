package httpserver

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"gitee.com/openeuler/PilotGo-plugin-mysql/config"
	"gitee.com/openeuler/PilotGo-plugins/sdk/logger"
)

var httpServer *http.Server

func Start() error {
	conf := config.Config().HttpServerConf
	router := gin.New()
	logger.Info("start http service on: http://%s", conf.Addr)

	httpServer = &http.Server{
		Addr:    conf.Addr,
		Handler: router,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("failed to start http service, error: %v", err)
			os.Exit(-1)
		}
	}()
	return nil
}

func Stop() error {
	ctx := context.Background()
	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Error("stop http server error: %s", err.Error())
		return err
	}
	return nil
}
