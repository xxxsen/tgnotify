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

type botctx struct {
	chatid int64
	bot    **tgbotapi.BotAPI
}

type NotifyServer struct {
	c              *config
	bot            *botctx
	channelsBotMap map[string]*botctx
}

func New(opts ...Option) (*NotifyServer, error) {
	c := &config{}
	for _, opt := range opts {
		opt(c)
	}
	if len(c.addr) == 0 {
		return nil, errs.New(errs.ErrParam, "invalid param")
	}
	return &NotifyServer{c: c}, nil
}

func (s *NotifyServer) Run() error {
	svr, err := naivesvr.NewServer(
		naivesvr.WithAddress(s.c.addr),
		naivesvr.WithHandlerRegister(handler.OnRegist),
		naivesvr.WithAttach(constant.KeyUserList, s.c.users),
	)
	if err != nil {
		return errs.New(errs.ErrServiceInternal, "bind http server fail", err)
	}
	logutil.GetLogger(context.Background()).With(zap.String("addr", s.c.addr), zap.Int("user_count", len(s.c.users))).Info("start running server")
	return svr.Run()
}
