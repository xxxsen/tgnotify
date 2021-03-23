package tgcmds

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
	"tgnotify"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	Regist("/help", NewCMDHelp)
}

//NewCMDHelp 新建帮助
func NewCMDHelp() tgnotify.TGCallback {
	hp := &CMDHelp{fg: flag.NewFlagSet("help", flag.ContinueOnError)}
	hp.cmd = hp.fg.String("cmd", "/help", "需要帮助信息的命令")
	return hp
}

//CMDHelp 帮助命令
type CMDHelp struct {
	cmd *string
	fg  *flag.FlagSet
}

//GetFlags 获取flags
func (c *CMDHelp) GetFlags() *flag.FlagSet {
	return c.fg
}

//OnCallback oncall
func (c *CMDHelp) OnCallback(bot *tgnotify.TGBot, update *tgbotapi.Update) error {
	if len(*c.cmd) == 0 {
		return fmt.Errorf("exp: /help --cmd=${cmdname}")
	}
	mp := GetTGCMDS()
	cl := *c.cmd
	if !strings.HasPrefix(cl, "/") {
		cl = "/" + cl
	}
	builder, ok := mp[cl]
	if !ok {
		return fmt.Errorf("not found cmd:%s", cl)
	}
	cmd := builder()
	buf := bytes.NewBuffer(nil)
	fg := cmd.GetFlags()
	fg.SetOutput(buf)
	fg.Usage()
	return bot.WriteBot(update.Message.Chat.ID, buf.String())
}
