//Package main mainpkg
package main

import (
	"context"
	"flag"
	"log"
	"tgnotify/config"
	"tgnotify/server"

	"github.com/xxxsen/common/logger"
	"github.com/xxxsen/common/logutil"
	"go.uber.org/zap"
)

var cfg = flag.String("config", "", "config file")

func main() {
	flag.Parse()

	c, err := config.Parse(*cfg)
	if err != nil {
		log.Panicf("parse config fail, err:%v", err)
	}

	//init log
	logger.Init(c.Log.File, c.Log.Level, int(c.Log.FileCount), int(c.Log.FileSize), int(c.Log.KeepDays), c.Log.Console)
	svr, err := server.New(server.WithBind(c.Listen), server.WithBotConfig(c.ChatID, c.Token), server.WithUser(c.User))
	if err != nil {
		logutil.GetLogger(context.Background()).With(zap.Error(err)).Fatal("init notify server fail")
	}
	if err := svr.Run(); err != nil {
		logutil.GetLogger(context.Background()).With(zap.Error(err)).Fatal("run notify server fail")
	}
}
