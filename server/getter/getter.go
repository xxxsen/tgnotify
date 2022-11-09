package getter

import (
	"context"
	"tgnotify/constant"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xxxsen/common/errs"
	"github.com/xxxsen/common/naivesvr"
)

func MustGetBot(ctx context.Context) *tgbotapi.BotAPI {
	v, ok := naivesvr.GetAttachKey(ctx, constant.KeyBot)
	if !ok {
		panic(errs.New(errs.ErrParam, "bot key not found"))
	}
	return v.(*tgbotapi.BotAPI)
}

func MustGetChatID(ctx context.Context) int64 {
	v, ok := naivesvr.GetAttachKey(ctx, constant.KeyChatID)
	if !ok {
		panic(errs.New(errs.ErrParam, "chatid key not found"))
	}
	return v.(int64)
}

func MustGetUserList(ctx context.Context) map[string]string {
	v, ok := naivesvr.GetAttachKey(ctx, constant.KeyUserList)
	if !ok {
		panic(errs.New(errs.ErrParam, "user list key not found"))
	}
	return v.(map[string]string)
}
