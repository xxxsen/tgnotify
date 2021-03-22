package tgnotify

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TGCallback interface {
	OnCallback(bot *TGBot, update *tgbotapi.Update, cmd string, params []string) error
}

type TGMSGCallback struct {
	tgcmdmp map[string]TGCallback
}

func NewTGMSGCallback() *TGMSGCallback {
	cmds := &TGMSGCallback{tgcmdmp: make(map[string]TGCallback)}
	return cmds
}

func (cb *TGMSGCallback) RegistMAP(full map[string]TGCallback) {
	for cmd, proc := range full {
		cb.tgcmdmp[cmd] = proc
	}
}

//decodeText decode text
func (cb *TGMSGCallback) decodeText(txt string) (string, []string, error) {
	txt = strings.TrimSpace(txt)
	arrs := strings.Split(txt, " ")
	if len(arrs) < 1 || !strings.HasPrefix(arrs[0], "/") {
		return "", nil, fmt.Errorf("invalid command")
	}
	return arrs[0], arrs[1:], nil
}

func (cb *TGMSGCallback) OnCallback(bot *TGBot, update *tgbotapi.Update) error {
	cmd, params, err := cb.decodeText(update.Message.Text)
	if err != nil {
		return err
	}
	proc, ok := cb.tgcmdmp[cmd]
	if !ok {
		return fmt.Errorf("unsupported cmd:%s", cmd)
	}
	return proc.OnCallback(bot, update, cmd, params)
}
