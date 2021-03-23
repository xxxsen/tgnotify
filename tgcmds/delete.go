package tgcmds

import (
	"flag"
	"tgnotify"
	"tgnotify/dao"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	Regist("/delete", NewCMDDelete)
}

//CMDDelete delete userinfo
type CMDDelete struct {
}

//GetFlags 获取flag
func (c *CMDDelete) GetFlags() *flag.FlagSet {
	return nil
}

//NewCMDDelete 新的delete命令字
func NewCMDDelete() tgnotify.TGCallback {
	return &CMDDelete{}
}

//OnCallback oncall
func (c *CMDDelete) OnCallback(bot *tgnotify.TGBot, update *tgbotapi.Update) error {
	sql := "delete from tbl_tgnotify where chatid = ? limit 1"
	_, err := dao.GetEngine().Exec(sql, update.Message.Chat.ID)
	if err != nil {
		return err
	}
	bot.WriteBot(update.Message.Chat.ID, "Delete user succ")
	return nil
}
