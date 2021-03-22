//Package tgnotify tg robot
package tgnotify

import (
	"fmt"
	"log"
	"tgnotify/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//TGBot tg robot struct
type TGBot struct {
	bot *tgbotapi.BotAPI
	cfg *config.NotifyConfig
	cb  MSGCallback
}

type MSGCallback interface {
	OnCallback(bot *TGBot, update *tgbotapi.Update) error
}

func (bot *TGBot) GetConf() *config.NotifyConfig {
	return bot.cfg
}

func (bot *TGBot) asyncUpdate() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		var err error
		if bot.cb != nil {
			err = bot.cb.OnCallback(bot, &update)
		}
		log.Printf("Recv message from remote, sender:%d, msg:%s, proc result:%v\n",
			update.Message.Chat.ID, update.Message.Text, err)
		if err != nil {
			bot.WriteBot(update.Message.Chat.ID, fmt.Sprintf("[ERROR]internal err, msg:%s", err))
		}
	}
	return nil
}

//NewBot create new robot
func NewBot(cfg *config.NotifyConfig) (*TGBot, error) {
	//parse config
	bot, err := tgbotapi.NewBotAPI(cfg.Bot.Token)
	if err != nil {
		return nil, err
	}
	tg := &TGBot{bot: bot, cfg: cfg}
	return tg, nil
}

//RegistMSGCallback regist callback
func (bot *TGBot) RegistMSGCallback(cb MSGCallback) {
	bot.cb = cb
}

func (bot *TGBot) Start() {
	go bot.asyncUpdate()
}

//WriteBot write a bot message
func (bot *TGBot) WriteBot(id int64, message string) error {
	_, err := bot.bot.Send(tgbotapi.NewMessage(id, message))
	if err != nil {
		return err
	}
	return nil
}
