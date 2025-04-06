package sender

import (
	"context"
	"fmt"
	"tgnotify/message"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type botMessageSender struct {
	chatid int64
	bot    *tgbotapi.BotAPI
}

func NewBotMessageSender(chatid int64, token string) (*botMessageSender, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("create bot failed, err:%w", err)
	}
	return &botMessageSender{chatid: chatid, bot: bot}, nil
}

func (c *botMessageSender) Name() string {
	return "tgbot"
}

func (c *botMessageSender) kind2kind(kind string) string {
	switch kind {
	case message.MKindText:
		return ""
	case message.MKindHTML:
		return tgbotapi.ModeHTML
	case message.MKindMarkdown, message.MKindTGMarkdown:
		return tgbotapi.ModeMarkdownV2
	}
	return ""
}

func (c *botMessageSender) SendMessage(ctx context.Context, msg message.IMessage) error {
	res, err := msg.Marshal()
	if err != nil {
		return fmt.Errorf("marshal msg failed, err:%w", err)
	}
	bm := tgbotapi.NewMessage(c.chatid, res)
	bm.ParseMode = c.kind2kind(msg.Kind())
	if _, err := c.bot.Send(bm); err != nil {
		return fmt.Errorf("send bot msg failed, err:%w", err)
	}
	return nil
}
