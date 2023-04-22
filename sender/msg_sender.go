package sender

import (
	"context"
)

type IMessageSender interface {
	Name() string
	SendMessage(ctx context.Context, mode string, message string) error
}
