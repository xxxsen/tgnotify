package server

import (
	"context"
	"tgnotify/constant"
	"tgnotify/server/handler"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xxxsen/common/errs"
	"github.com/xxxsen/common/logutil"
	"github.com/xxxsen/common/naivesvr"
	"go.uber.org/zap"
)

type NotifyServer struct {
	c   *config
	bot *tgbotapi.BotAPI
}

func New(opts ...Option) (*NotifyServer, error) {
	c := &config{}
	for _, opt := range opts {
		opt(c)
	}
	if len(c.addr) == 0 || len(c.token) == 0 || c.chatid == 0 {
		return nil, errs.New(errs.ErrParam, "invalid param")
	}
	bot, err := tgbotapi.NewBotAPI(c.token)
	if err != nil {
		return nil, errs.Wrap(errs.ErrUnknown, "init bot fail", err)
	}
	return &NotifyServer{c: c, bot: bot}, nil
}

func (s *NotifyServer) Run() error {
	svr, err := naivesvr.NewServer(
		naivesvr.WithAddress(s.c.addr),
		naivesvr.WithHandlerRegister(handler.OnRegist),
		naivesvr.WithAttach(constant.KeyBot, s.bot),
		naivesvr.WithAttach(constant.KeyChatID, s.c.chatid),
		naivesvr.WithAttach(constant.KeyUserList, s.c.users),
	)
	if err != nil {
		return errs.New(errs.ErrServiceInternal, "bind http server fail", err)
	}
	logutil.GetLogger(context.Background()).With(zap.String("addr", s.c.addr),
		zap.Int64("chatid", s.c.chatid), zap.Int("user_count", len(s.c.users))).Info("start running server")
	return svr.Run()
}
