package sender

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xxxsen/common/errs"
)

type BotMsgSender struct {
	name   string
	chatid int64
	bot    *tgbotapi.BotAPI
}

func NewBotMsgSender(name string, chatid int64, token string) (*BotMsgSender, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, errs.Wrap(errs.ErrUnknown, "init bot fail", err)
	}
	return &BotMsgSender{name, chatid, bot}, nil
}

func (c *BotMsgSender) Name() string {
	return c.name
}

func (c *BotMsgSender) SendMessage(ctx context.Context, mode string, message string) error {
	msg := tgbotapi.NewMessage(c.chatid, message)
	if len(mode) != 0 {
		msg.ParseMode = mode
	}
	_, err := c.bot.Send(msg)
	if err != nil {
		return errs.Wrap(errs.ErrIO, "send msg fail", err)
	}
	return nil
}
