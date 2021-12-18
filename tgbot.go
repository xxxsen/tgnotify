//Package tgnotify tg robot
package tgnotify

import (
	"context"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/xxxsen/log"
)

//TGBot tg robot struct
type TGBot struct {
	bot   *tgbotapi.BotAPI
	token string
	cb    MSGCallback
}

type MSGCallback interface {
	OnCallback(ctx context.Context, bot *TGBot, update *tgbotapi.Update) error
}

func (bot *TGBot) asyncUpdate() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		ctx := context.Background()
		var err error
		if bot.cb != nil {
			err = bot.cb.OnCallback(ctx, bot, &update)
		}
		log.Tracef("recv message from remote, sender:%d, msg:%s, proc result:%v",
			update.Message.Chat.ID, update.Message.Text, err)
		if err != nil {
			bot.WriteBot(update.Message.Chat.ID, fmt.Sprintf("[ERROR]internal err, msg:%s", err))
		}
	}
	return nil
}

//NewBot create new robot
func NewBot(token string) (*TGBot, error) {
	//parse config
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	tg := &TGBot{bot: bot, token: token}
	return tg, nil
}

//RegistMSGCallback regist callback
func (bot *TGBot) RegistMSGCallback(cb MSGCallback) {
	bot.cb = cb
}

func (bot *TGBot) Start() {
	go bot.asyncUpdate()
}

func (bot *TGBot) WriteBotf(id int64, formatter string, args ...interface{}) error {
	return bot.WriteBot(id, fmt.Sprintf(formatter, args...))
}

//WriteBot write a bot message
func (bot *TGBot) WriteBot(id int64, message string) error {
	return bot.WriteModeBot(id, "", message)
}

func (bot *TGBot) detectMode(mode string) string {
	if strings.EqualFold(mode, "html") {
		return tgbotapi.ModeHTML
	} else if strings.EqualFold(mode, "markdown") {
		return tgbotapi.ModeMarkdown
	}
	return ""
}

func (bot *TGBot) WriteModeBot(id int64, mode string, message string) error {
	msg := tgbotapi.NewMessage(id, message)
	mode = bot.detectMode(mode)
	if len(mode) != 0 {
		msg.ParseMode = mode
	}
	_, err := bot.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
