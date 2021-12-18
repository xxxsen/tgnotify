package tgnotify

import (
	"context"
	"flag"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//TGCallback 回调相关的处理逻辑
type TGCallback interface {
	GetFlags() *flag.FlagSet                                                   //返回自身的flagset
	OnCallback(ctx context.Context, bot *TGBot, update *tgbotapi.Update) error //
}

//TGCallbackBuilder 由具体的子类实现
type TGCallbackBuilder func() TGCallback

//TGMSGCallback 回调集合
type TGMSGCallback struct {
	tgcmdmp map[string]TGCallbackBuilder
}

func NewTGMSGCallback() *TGMSGCallback {
	cmds := &TGMSGCallback{tgcmdmp: make(map[string]TGCallbackBuilder)}
	return cmds
}

func (cb *TGMSGCallback) RegistMAP(full map[string]TGCallbackBuilder) {
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

func (cb *TGMSGCallback) OnCallback(ctx context.Context, bot *TGBot, update *tgbotapi.Update) error {
	cmd, params, err := cb.decodeText(update.Message.Text)
	if err != nil {
		return err
	}
	builder, ok := cb.tgcmdmp[cmd]
	if !ok {
		return fmt.Errorf("unsupported cmd:%s", cmd)
	}
	caller := builder()
	fg := caller.GetFlags()
	if fg != nil {
		err = fg.Parse(params)
		if err != nil {
			return err
		}
	}
	return caller.OnCallback(ctx, bot, update)
}
