package tgcmds

import (
	"context"
	"flag"
	"fmt"
	"tgnotify"
	"tgnotify/dao"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
func (c *CMDInfo) OnCallback(ctx context.Context, bot *tgnotify.TGBot, update *tgbotapi.Update) error {
	info, ok := dao.GetFileStorage().QueryUserByChatid(ctx, uint64(update.Message.Chat.ID))
	if !ok {
		return fmt.Errorf("not found")
	}
	bot.WriteBotf(update.Message.Chat.ID, "read info succ, chatid:%d, code:%s",
		update.Message.Chat.ID, info.Code)
	return nil
}
