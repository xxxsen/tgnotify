//Package main mainpkg
package main

import (
	"flag"
	"log"
	"tgnotify"
	"tgnotify/config"
	"tgnotify/dao"
	"tgnotify/tgcmds"
)

var conf = flag.String("conf", "./config.json", "")

func main() {
	flag.Parse()
	cfg := config.Global()
	err := cfg.Parse(*conf)
	if err != nil {
		log.Fatal(err)
	}

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
