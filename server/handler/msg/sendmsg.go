package msg

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"tgnotify/constant"
	"tgnotify/model"
	"tgnotify/sender"

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
	ch := c.GetHeader(constant.KeyChannelHeader)
	return SendMessageJson(c, &model.SendMessageRequest{
		Channel:     ch,
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
	if err := sendMessageInternal(c, req.Channel, mode, req.Message); err != nil {
		return http.StatusOK, errs.Wrap(errs.ErrIO, "send internal fail", err), nil
	}
	return http.StatusOK, errs.ErrOK, nil
}

func sendMessageInternal(ctx context.Context, ch string, mode string, message string) error {
	err := sender.SendMessageByChannel(ctx, ch, mode, message)
	if err != nil {
		return errs.Wrap(errs.ErrIO, "send msg fail", err)
	}
	return nil
}
