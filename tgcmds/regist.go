package tgcmds

import (
	"flag"
	"fmt"
	"tgnotify"
	"tgnotify/dao"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	Regist("/reg", NewCMDRegist)
}

//CMDRegist regist
type CMDRegist struct {
	user *string
	code *string
	fg   *flag.FlagSet
}

//NewCMDRegist 新的注册命令字
func NewCMDRegist() tgnotify.TGCallback {
	reg := &CMDRegist{fg: flag.NewFlagSet("reg", flag.ContinueOnError)}
	reg.user = reg.fg.String("user", "", "用户名")
	reg.code = reg.fg.String("code", "", "密钥")
	return reg
}

//GetFlags 获取flags
func (c *CMDRegist) GetFlags() *flag.FlagSet {
	return c.fg
}

//OnCallback oncall
func (c *CMDRegist) OnCallback(bot *tgnotify.TGBot, update *tgbotapi.Update) error {
	chatid := update.Message.Chat.ID
	if len(*c.user) == 0 || len(*c.code) == 0 {
		return fmt.Errorf("invalid params, exp:/reg --user=${user} --code=${code}")
	}
	ts := time.Now()
	sql := "insert ignore into tbl_tgnotify(user, code, chatid, ts) values(?, ?, ?, ?)"
	sqlRs, err := dao.GetEngine().Exec(sql, *c.user, *c.code, chatid, ts)
	if err != nil {
		return err
	}
	cnt, err := sqlRs.RowsAffected()
	if err != nil {
		return err
	}
	if cnt == 0 {
		return fmt.Errorf("already regist, try /chg to change your info")
	}
	bot.WriteBot(chatid, fmt.Sprintf("Regist success, user:%s, code:%s", *c.user, *c.code))
	return nil
}
