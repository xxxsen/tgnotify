package tgcmds

import (
	"fmt"
	"tgnotify"
	"tgnotify/dao"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	Regist("/chg", &CMDChange{})
}

//CMDChange change userinfo
type CMDChange struct {
}

//OnCallback oncall
func (c *CMDChange) OnCallback(bot *tgnotify.TGBot, update *tgbotapi.Update, cmd string, params []string) error {
	chatid := update.Message.Chat.ID
	if len(params) != 2 {
		return fmt.Errorf("need more params, exp: /chg ${user} ${code}")
	}

	sql := "update tbl_tgnotify set user = ?, code = ? where chatid = ? limit 1"
	rs, err := dao.GetEngine().Exec(sql, params[0], params[1], chatid)
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
	bot.WriteBot(chatid, fmt.Sprintf("Change info succ, new info: user:%s, code:%s", params[0], params[1]))
	return nil
}
