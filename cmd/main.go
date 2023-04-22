// Package main mainpkg
package main

import (
	"context"
	"flag"
	"log"
	"tgnotify/config"
	"tgnotify/sender"
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
	mustInitGlobalMessageChannel(c)
	opts := []server.Option{
		server.WithBind(c.Listen),
		server.WithUser(c.User),
	}
	svr, err := server.New(opts...)
	if err != nil {
		logutil.GetLogger(context.Background()).With(zap.Error(err)).Fatal("init notify server fail")
	}
	if err := svr.Run(); err != nil {
		logutil.GetLogger(context.Background()).With(zap.Error(err)).Fatal("run notify server fail")
	}
}

func mustInitGlobalMessageChannel(c *config.Config) {
	ch, err := sender.NewBotMsgSender("default", c.ChatID, c.Token)
	if err != nil {
		panic(err)
	}
	chs := make(map[string]sender.IMessageSender)
	for name, cfg := range c.Channels {
		chitem, err := sender.NewBotMsgSender(name, cfg.ChatID, cfg.Token)
		if err != nil {
			panic(err)
		}
		chs[name] = chitem
	}
	sender.InitGlobalGroupMessageSender(chs, ch)
}
