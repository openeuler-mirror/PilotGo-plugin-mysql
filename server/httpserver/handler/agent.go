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
