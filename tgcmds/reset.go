package tgcmds

import (
	"context"
	"flag"
	"tgnotify"
	"tgnotify/dao"
	"tgnotify/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

func init() {
	Regist("/reset", NewCMDReset)
}

//CMDReset change userinfo
type CMDReset struct {
	fg *flag.FlagSet
}

//NewCMDReset 生成新的Change
func NewCMDReset() tgnotify.TGCallback {
	c := &CMDReset{fg: flag.NewFlagSet("reset", flag.ContinueOnError)}
	return c
}

//GetFlags 获取flag
func (c *CMDReset) GetFlags() *flag.FlagSet {
	return c.fg
}

//OnCallback oncall
func (c *CMDReset) OnCallback(ctx context.Context, bot *tgnotify.TGBot, update *tgbotapi.Update) error {
	chatid := update.Message.Chat.ID
	uinfo := &models.UserInfo{
		Chatid: uint64(chatid),
		Code:   uuid.NewString(),
	}
	if err := dao.GetFileStorage().UpdateUser(ctx, uinfo); err != nil {
		return err
	}
	bot.WriteBotf(chatid, "reset info succ, new info: chatid:%d, code:%s", chatid, uinfo.Code)
	return nil
}
