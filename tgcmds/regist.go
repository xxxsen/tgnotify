package tgcmds

import (
	"context"
	"flag"
	"fmt"
	"tgnotify"
	"tgnotify/dao"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

func init() {
	Regist("/reg", NewCMDRegist)
}

//CMDRegist regist
type CMDRegist struct {
	fg *flag.FlagSet
}

//NewCMDRegist 新的注册命令字
func NewCMDRegist() tgnotify.TGCallback {
	reg := &CMDRegist{fg: flag.NewFlagSet("reg", flag.ContinueOnError)}
	return reg
}

//GetFlags 获取flags
func (c *CMDRegist) GetFlags() *flag.FlagSet {
	return c.fg
}

//OnCallback oncall
func (c *CMDRegist) OnCallback(ctx context.Context, bot *tgnotify.TGBot, update *tgbotapi.Update) error {
	chatid := update.Message.Chat.ID
	code := uuid.NewString()
	if _, ok := dao.GetFileStorage().QueryUserByChatid(ctx, uint64(chatid)); ok {
		return fmt.Errorf("already exist")
	}
	if err := dao.GetFileStorage().NewUser(ctx, uint64(chatid), code); err != nil {
		return err
	}
	bot.WriteBotf(chatid, "regist success, chatid:%d, code:%s", chatid, code)
	return nil
}
