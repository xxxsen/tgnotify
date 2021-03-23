package tgcmds

import (
	"flag"
	"fmt"
	"tgnotify"
	"tgnotify/dao"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	Regist("/chg", NewCMDChange)
}

//CMDChange change userinfo
type CMDChange struct {
	user *string
	code *string
	fg   *flag.FlagSet
}

//NewCMDChange 生成新的Change
func NewCMDChange() tgnotify.TGCallback {
	c := &CMDChange{fg: flag.NewFlagSet("chg", flag.ContinueOnError)}
	c.user = c.fg.String("user", "", "用户名")
	c.code = c.fg.String("code", "", "密钥")
	return c
}

//GetFlags 获取flag
func (c *CMDChange) GetFlags() *flag.FlagSet {
	return c.fg
}

//OnCallback oncall
func (c *CMDChange) OnCallback(bot *tgnotify.TGBot, update *tgbotapi.Update) error {
	chatid := update.Message.Chat.ID
	if len(*c.user) == 0 || len(*c.code) == 0 {
		return fmt.Errorf("invalid params, exp: /chg --user=${user} --code=${code}")
	}
	sql := "update tbl_tgnotify set user = ?, code = ? where chatid = ? limit 1"
	rs, err := dao.GetEngine().Exec(sql, *c.user, *c.code, chatid)
	if err != nil {
		return err
	}
	cnt, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if cnt == 0 {
		return fmt.Errorf("user may not exist")
	}
	bot.WriteBot(chatid, fmt.Sprintf("Change info succ, new info: user:%s, code:%s", *c.user, *c.code))
	return nil
}
