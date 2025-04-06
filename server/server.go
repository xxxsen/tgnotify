package server

import (
	"tgnotify/server/handler"

	"github.com/gin-gonic/gin"
	"github.com/xxxsen/common/webapi"
	"github.com/xxxsen/common/webapi/auth"
	"github.com/xxxsen/common/webapi/middleware"
)

type NotifyServer struct {
	c      *config
	engine webapi.IWebEngine
}

func New(bind string, opts ...Option) (*NotifyServer, error) {
	c := &config{}
	for _, opt := range opts {
		opt(c)
	}
	s := &NotifyServer{c: c}
	engine, err := webapi.NewEngine("/", bind,
		webapi.WithAuth(auth.MapUserMatch(c.users)),
		webapi.WithRegister(s.register),
		webapi.WithExtraMiddlewares(middleware.MustAuthMiddleware()),
	)
	if err != nil {
		return nil, err
	}
	s.engine = engine
	return s, nil
}

func (s *NotifyServer) register(r *gin.RouterGroup) {
	router := r.Group("/")
	{
		router.POST("", handler.SendMessage)
	}
}

func (s *NotifyServer) Run() error {
	return s.engine.Run()
}
