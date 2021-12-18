package tgcmds

import (
	"context"
	"flag"
	"tgnotify"
	"tgnotify/dao"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
	Regist("/delete", NewCMDDelete)
}

//CMDDelete delete userinfo
type CMDDelete struct {
	fg *flag.FlagSet
}

//GetFlags 获取flag
func (c *CMDDelete) GetFlags() *flag.FlagSet {
	return c.fg
}

//NewCMDDelete 新的delete命令字
func NewCMDDelete() tgnotify.TGCallback {
	return &CMDDelete{fg: flag.NewFlagSet("delete", flag.ContinueOnError)}
}

//OnCallback oncall
func (c *CMDDelete) OnCallback(ctx context.Context, bot *tgnotify.TGBot, update *tgbotapi.Update) error {
	chatid := uint64(update.Message.Chat.ID)
	if err := dao.GetFileStorage().DeleteByChatid(context.Background(), chatid); err != nil {
		return err
	}
	bot.WriteBotf(update.Message.Chat.ID, "delete account succ, chatid:%d", chatid)
	return nil
}
