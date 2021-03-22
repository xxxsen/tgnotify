package tgcmds

import (
	"fmt"
	"tgnotify"
	"tgnotify/dao"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	Regist("/reg", &CMDRegist{})
}

//CMDRegist regist
type CMDRegist struct {
}

//OnCallback oncall
func (c *CMDRegist) OnCallback(bot *tgnotify.TGBot, update *tgbotapi.Update, cmd string, params []string) error {
	chatid := update.Message.Chat.ID
	if len(params) != 2 {
		return fmt.Errorf("invalid params, exp:/reg ${user} ${code}")
	}
	user := params[0]
	code := params[1]
	ts := time.Now()
	sql := "insert ignore into tbl_tgnotify(user, code, chatid, ts) values(?, ?, ?, ?)"
	sqlRs, err := dao.GetEngine().Exec(sql, user, code, chatid, ts)
	if err != nil {
		return err
	}
	cnt, err := sqlRs.RowsAffected()
	if err != nil {
		return err
	}
	if cnt == 0 {
		return fmt.Errorf("already regist, use /chg ${user} ${code} to change your info")
	}
	bot.WriteBot(chatid, fmt.Sprintf("Regist success, user:%s, code:%s", user, code))
	return nil
}
