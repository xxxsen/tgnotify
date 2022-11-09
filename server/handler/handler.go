package handler

import (
	"tgnotify/server/handler/middleware"
	"tgnotify/server/handler/msg"

	"github.com/gin-gonic/gin"
	"github.com/xxxsen/common/naivesvr"
	"github.com/xxxsen/common/naivesvr/codec"
)

func OnRegist(engine *gin.Engine) {
	msgGroup := engine.Group("/")
	msgGroup.POST("/", naivesvr.WrapHandler(nil, codec.CustomCodec(codec.JsonCodec, codec.NopCodec), msg.SendMessage), middleware.AuthMiddleware())
	msgGroup.POST("/json", naivesvr.WrapHandler(&msg.SendMessageRequest{}, codec.JsonCodec, msg.SendMessageJson), middleware.AuthMiddleware())
}
