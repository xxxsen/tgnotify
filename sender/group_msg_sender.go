package sender

import (
	"context"

	"github.com/xxxsen/common/errs"
)

var glgp *GroupMsgSender

func InitGlobalGroupMessageSender(chsm map[string]IMessageSender, def IMessageSender) {
	glgp = NewGroupMsgSender(chsm, def)
}

type GroupMsgSender struct {
	chs map[string]IMessageSender
	def IMessageSender
}

func NewGroupMsgSender(chsm map[string]IMessageSender, def IMessageSender) *GroupMsgSender {
	return &GroupMsgSender{
		chs: chsm, def: def,
	}
}

func (c *GroupMsgSender) SendMessage(ctx context.Context, name string, mode string, message string) error {
	s, ok := c.chs[name]
	if !ok {
		s = c.def
	}
	if s == nil {
		return errs.New(errs.ErrServiceInternal, "no sender found")
	}
	return s.SendMessage(ctx, mode, message)
}

func SendMessageByChannel(ctx context.Context, name string, mode string, message string) error {
	return glgp.SendMessage(ctx, name, mode, message)
}
