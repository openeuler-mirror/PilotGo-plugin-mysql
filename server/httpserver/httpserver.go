/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-mysql licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: guozhengxin <guozhengxin@kylinos.cn>
 * Date: Fri Jul 7 11:13:04 2023 +0800
 */
package httpserver

import (
	"context"
	"errors"
	"net/http"
	"os"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"github.com/gin-gonic/gin"

	"gitee.com/openeuler/PilotGo-plugin-mysql/client"
	"gitee.com/openeuler/PilotGo-plugin-mysql/config"
	"gitee.com/openeuler/PilotGo-plugin-mysql/httpserver/handler"
)

var httpServer *http.Server

func Start() error {
	conf := config.Config().HttpServerConf
	engine := gin.New()
	logger.Info("start http service on: http://%s", conf.Addr)

	registerHandlers(engine)

	httpServer = &http.Server{
		Addr:    conf.Addr,
		Handler: engine,
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

// 注册http handler
func registerHandlers(engine *gin.Engine) {

	// 注册client默认的handlers
	client.RegisterHandlers(engine)

	api := engine.Group("plugin/mysql/api")

	// 监控agent相关接口
	{
		api.GET("/info", handler.AgentInfoHandler)
		api.PUT("/install", handler.AgentInstallHandler)
		api.PUT("/update", handler.AgentInstallHandler)
		api.DELETE("/uninstall", handler.AgentUninstallHandler)
	}
}
