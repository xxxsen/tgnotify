//Package main mainpkg
package main

import (
	"flag"
	"tgnotify"
	"tgnotify/config"
	"tgnotify/dao"
	"tgnotify/log"
	"tgnotify/tgcmds"
)

var conf = flag.String("conf", "./config.json", "")

func main() {
	flag.Parse()
	cfg := config.Global()
	err := cfg.Parse(*conf)
	if err != nil {
		panic(err)
	}

	//init log
	log.Init(cfg.Log.File, log.StringToLevel(cfg.Log.Level),
		int(cfg.Log.Count), int(cfg.Log.Size), int(cfg.Log.KeepDays), cfg.Log.WriteConsole)

	log.Infof("Read config finish, config:%+v", *cfg)

	err = dao.Init(&dao.DBParams{
		Host: cfg.DB.Host,
		Port: int(cfg.DB.Port),
		User: cfg.DB.User,
		Pwd:  cfg.DB.Pwd,
		DB:   cfg.DB.DBName,
	})
	if err != nil {
		log.Fatal(err)
	}

	nt, err := tgnotify.NewBot(cfg)
	if err != nil {
		log.Fatal(err)
	}

	//regist callback
	cb := tgnotify.NewTGMSGCallback()
	cb.RegistMAP(tgcmds.GetTGCMDS())
	nt.RegistMSGCallback(cb)

	//start bot
	nt.Start()
	svr, err := tgnotify.NewService(nt)
	if err != nil {
		log.Fatal(err)
	}

	//regist service msg
	msg := &tgnotify.ServiceDoMSG{}
	svr.Regist("msg", msg)
	svr.Regist("", msg)

	svr.ServeForever(cfg.Server.Listen)
}
