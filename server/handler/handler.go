package handler

import (
	"tgnotify/model"
	"tgnotify/server/handler/middleware"
	"tgnotify/server/handler/msg"

	"github.com/gin-gonic/gin"
	"github.com/xxxsen/common/naivesvr"
	"github.com/xxxsen/common/naivesvr/codec"
)

func OnRegist(engine *gin.Engine) {
	msgGroup := engine.Group("/")
	msgGroup.POST("/", middleware.AuthMiddleware(), naivesvr.WrapHandler(nil, codec.CustomCodec(codec.JsonCodec, codec.NopCodec), msg.SendMessage))
	msgGroup.POST("/json", middleware.AuthMiddleware(), naivesvr.WrapHandler(&model.SendMessageRequest{}, codec.JsonCodec, msg.SendMessageJson))
}
