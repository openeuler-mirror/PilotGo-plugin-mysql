/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-mysql licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: guozhengxin <guozhengxin@kylinos.cn>
 * Date: Mon Jul 10 11:34:43 2023 +0800
 */
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AgentInfoHandler(ctx *gin.Context) {

	ctx.String(http.StatusOK, "monitor agent info")
}

func AgentInstallHandler(ctx *gin.Context) {

	ctx.String(http.StatusOK, "install monitor agent")
}

func AgentUpdateHandler(ctx *gin.Context) {

	ctx.String(http.StatusOK, "update monitor agent")
}

func AgentUninstallHandler(ctx *gin.Context) {

	ctx.String(http.StatusOK, "uninstall monitor agent")
}
