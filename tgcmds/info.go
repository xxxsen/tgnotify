package tgcmds

import (
	"flag"
	"fmt"
	"tgnotify"
	"tgnotify/dao"
	"tgnotify/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	Regist("/info", NewCMDInfo)
}

//CMDInfo get userinfo
type CMDInfo struct {
	fg *flag.FlagSet
}

//NewCMDInfo 新的info命令自
func NewCMDInfo() tgnotify.TGCallback {
	return &CMDInfo{fg: flag.NewFlagSet("info", flag.ContinueOnError)}
}

//GetFlags 获取flag
func (c *CMDInfo) GetFlags() *flag.FlagSet {
	return c.fg
}

//OnCallback oncall
func (c *CMDInfo) OnCallback(bot *tgnotify.TGBot, update *tgbotapi.Update) error {
	rs := &[]models.TblTgnotify{}
	err := dao.GetEngine().SQL("select * from tbl_tgnotify where chatid = ? limit 1", update.Message.Chat.ID).Find(rs)
	if err != nil {
		return err
	}
	if len(*rs) == 0 {
		return fmt.Errorf("not found user info")
	}
	bot.WriteBot(update.Message.Chat.ID, fmt.Sprintf("User:%s, Code:%s, ChatID:%d",
		(*rs)[0].User, (*rs)[0].Code, update.Message.Chat.ID))
	return nil
}
