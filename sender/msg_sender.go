package sender

import (
	"context"
	"tgnotify/message"
)

type IMessageSender interface {
	Name() string
	SendMessage(ctx context.Context, msg message.IMessage) error
}

var (
	defaultInst IMessageSender
)

func SetSenderImpl(s IMessageSender) {
	defaultInst = s
}

func SendMessage(ctx context.Context, msg message.IMessage) error {
	return defaultInst.SendMessage(ctx, msg)
}
