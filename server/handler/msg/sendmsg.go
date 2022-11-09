package msg

import (
	"io/ioutil"
	"net/http"
	"strings"
	"tgnotify/model"
	"tgnotify/server/getter"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xxxsen/common/errs"
)

func type2mode(typ string) string {
	if len(typ) == 0 {
		return ""
	}
	if strings.EqualFold(typ, "html") {
		return tgbotapi.ModeHTML
	}
	if strings.EqualFold(typ, "markdown") {
		return tgbotapi.ModeMarkdown
	}
	if strings.EqualFold(typ, "markdownv2") {
		return tgbotapi.ModeMarkdownV2
	}
	return ""
}

func readmsg(c *gin.Context) (string, error) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return "", errs.Wrap(errs.ErrIO, "read msg fail", err)
	}
	return string(data), nil
}

func SendMessage(c *gin.Context, ireq interface{}) (int, errs.IError, interface{}) {
	msg, err := readmsg(c)
	if err != nil {
		return http.StatusOK, errs.Wrap(errs.ErrIO, "read msg fail", err), nil
	}
	return SendMessageJson(c, &model.SendMessageRequest{
		Message:     msg,
		MessageType: c.GetHeader("mode"),
	})
}

func SendMessageJson(c *gin.Context, ireq interface{}) (int, errs.IError, interface{}) {
	req := ireq.(*model.SendMessageRequest)
	if len(req.Message) == 0 {
		return http.StatusOK, errs.New(errs.ErrParam, "nil message"), nil
	}
	mode := type2mode(req.MessageType)
	bot := getter.MustGetBot(c)
	chatid := getter.MustGetChatID(c)
	if err := sendMessageInternal(bot, chatid, mode, req.Message); err != nil {
		return http.StatusOK, errs.Wrap(errs.ErrIO, "send internal fail", err), nil
	}
	return http.StatusOK, errs.ErrOK, nil
}

func sendMessageInternal(bot *tgbotapi.BotAPI, id int64, mode string, message string) error {
	msg := tgbotapi.NewMessage(id, message)
	if len(mode) != 0 {
		msg.ParseMode = mode
	}
	_, err := bot.Send(msg)
	if err != nil {
		return errs.Wrap(errs.ErrIO, "send msg fail", err)
	}
	return nil
}
