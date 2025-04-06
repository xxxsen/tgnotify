// Package main mainpkg
package main

import (
	"flag"
	"log"
	"tgnotify/config"
	"tgnotify/sender"
	"tgnotify/server"

	"github.com/xxxsen/common/logger"
	"go.uber.org/zap"
)

var cfg = flag.String("config", "", "config file")

func main() {
	flag.Parse()

	c, err := config.Parse(*cfg)
	if err != nil {
		log.Fatalf("parse config fail, err:%v", err)
	}
	//init log
	logkit := logger.Init(c.LogConfig.File, c.LogConfig.Level, int(c.LogConfig.FileCount), int(c.LogConfig.FileSize), int(c.LogConfig.KeepDays), c.LogConfig.Console)
	bot, err := sender.NewBotMessageSender(c.ChatID, c.Token)
	if err != nil {
		logkit.Fatal("init bot sender failed", zap.Error(err))
	}
	sender.SetSenderImpl(bot)
	svr, err := server.New(c.Listen, server.WithUser(c.User))
	if err != nil {
		logkit.Fatal("init notify server fail", zap.Error(err))
	}
	logkit.Info("init server succ, start it...")
	if err := svr.Run(); err != nil {
		logkit.Fatal("run notify server fail", zap.Error(err))
	}
}
