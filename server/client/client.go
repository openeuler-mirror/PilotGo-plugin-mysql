/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-mysql licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: guozhengxin <guozhengxin@kylinos.cn>
 * Date: Fri Jul 7 15:26:19 2023 +0800
 */
package client

import (
	"github.com/gin-gonic/gin"

	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
)

var globalClient *client.Client

func Init(info *client.PluginInfo) error {
	globalClient = client.DefaultClient(info)

	return nil
}

func RegisterHandlers(engine *gin.Engine) {
	globalClient.RegisterHandlers(engine)
}
