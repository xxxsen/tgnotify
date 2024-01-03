package handler

import (
	"tgnotify/model"
	"tgnotify/server/handler/middleware"
	"tgnotify/server/handler/msg"

	"github.com/gin-gonic/gin"
	"github.com/xxxsen/common/cgi"
	"github.com/xxxsen/common/cgi/codec"
)

func OnRegist(engine *gin.Engine) {
	msgGroup := engine.Group("/")
	msgGroup.POST("/", middleware.AuthMiddleware(), cgi.WrapHandler(nil, codec.CustomCodec(codec.JsonCodec, codec.NopCodec), msg.SendMessage))
	msgGroup.POST("/json", middleware.AuthMiddleware(), cgi.WrapHandler(&model.SendMessageRequest{}, codec.JsonCodec, msg.SendMessageJson))
}
