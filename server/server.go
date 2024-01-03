package server

import (
	"context"
	"tgnotify/constant"
	"tgnotify/server/handler"

	"github.com/xxxsen/common/cgi"
	"github.com/xxxsen/common/errs"
	"github.com/xxxsen/common/logutil"
	"go.uber.org/zap"
)

type NotifyServer struct {
	c *config
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
	svr, err := cgi.NewServer(
		cgi.WithAddress(s.c.addr),
		cgi.WithHandlerRegister(handler.OnRegist),
		cgi.WithAttach(constant.KeyUserList, s.c.users), //TODO:
	)
	if err != nil {
		return errs.New(errs.ErrServiceInternal, "bind http server fail", err)
	}
	logutil.GetLogger(context.Background()).With(zap.String("addr", s.c.addr), zap.Int("user_count", len(s.c.users))).Info("start running server")
	return svr.Run()
}
