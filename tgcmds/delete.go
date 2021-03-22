package tgcmds

import (
	"tgnotify"
	"tgnotify/dao"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	Regist("/delete", &CMDDelete{})
}

//CMDDelete delete userinfo
type CMDDelete struct {
}

//OnCallback oncall
func (c *CMDDelete) OnCallback(bot *tgnotify.TGBot, update *tgbotapi.Update, cmd string, params []string) error {
	sql := "delete from tbl_tgnotify where chatid = ? limit 1"
	_, err := dao.GetEngine().Exec(sql, update.Message.Chat.ID)
	if err != nil {
		return err
	}
	bot.WriteBot(update.Message.Chat.ID, "Delete user succ")
	return nil
}
