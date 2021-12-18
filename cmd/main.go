//Package main mainpkg
package main

import (
	"tgnotify"
	"tgnotify/dao"
	"tgnotify/tgcmds"

	flag "github.com/xxxsen/envflag"
	"github.com/xxxsen/log"
)

var listen = flag.String("listen", ":8333", "listen address")
var savefile = flag.String("save_file", "./user.data", "file for saving user data")
var logLevel = flag.String("log_level", "trace", "log level")
var token = flag.String("token", "", "bot token")

func main() {
	flag.Parse()

	//init log
	log.Init("", log.StringToLevel(*logLevel),
		0, 0, 7, true)

	log.Infof("LISTEN:%v", *listen)
	log.Infof("SAVE_FILE:%v", *savefile)
	log.Infof("LOG_LEVEL:%v", *logLevel)
	log.Infof("TOKEN:%v", *token)

	if len(*token) == 0 {
		log.Fatal("invalid tg token")
	}

	err := dao.Init(*savefile)
	if err != nil {
		log.Fatal(err)
	}

	nt, err := tgnotify.NewBot(*token)
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
	svr.Regist(msg, "msg", "")

	svr.ServeForever(*listen)
}
