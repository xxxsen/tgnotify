package tgnotify

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"tgnotify/dao"
	"tgnotify/models"

	"github.com/xxxsen/log"
)

//Service http service
type Service struct {
	bot *TGBot
	mp  map[string]ServiceDoIt
	sc  *serviceController
}

type serviceController struct {
	svr *Service
}

type ServiceDoIt interface {
	OnDo(ctx context.Context, svr *Service, cmd string, user *models.UserInfo, req *http.Request, tgnt *TGBot) error
	OnAuth(ctx context.Context, svr *Service, chatid uint64, code string) (*models.UserInfo, bool, error)
}

//NewService create new service
func NewService(tgnt *TGBot) (*Service, error) {
	svr := &Service{bot: tgnt, mp: make(map[string]ServiceDoIt)}
	sc := &serviceController{svr}
	svr.sc = sc
	return svr, nil
}

func (sc *serviceController) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	cmd := req.Header.Get("cmd")
	fc, ok := sc.svr.mp[cmd]
	if !ok {
		rsp.WriteHeader(http.StatusNotImplemented)
		log.Errorf("recv unimplement cmd from user, cmd:%s", cmd)
		return
	}
	user := strings.TrimSpace(req.Header.Get("user"))
	code := strings.TrimSpace(req.Header.Get("code"))

	chatid, err := strconv.ParseUint(user, 10, 64)
	if err != nil {
		rsp.WriteHeader(http.StatusBadRequest)
		log.Errorf("recv invalid user:%s, code:%s, err:%v", user, code, err)
		return
	}

	ctx := context.Background()
	u, succ, err := fc.OnAuth(ctx, sc.svr, chatid, code)
	if err != nil {
		rsp.WriteHeader(http.StatusInternalServerError)
		log.Errorf("recv cmd:%s but auth err, err:%v", cmd, err)
		return
	}
	if !succ {
		rsp.WriteHeader(http.StatusUnauthorized)
		log.Errorf("recv cmd:%s, auth fail, user:%s, code:%s", cmd, user, code)
		return
	}

	err = fc.OnDo(ctx, sc.svr, cmd, u, req, sc.svr.bot)
	if err != nil {
		log.Errorf("do cmd:%s cause internal err, user:%s, header:%+v, err:%v", cmd, user, req.Header, err)
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Tracef("cmd:%s exec success, user:%s, code:%s", cmd, user, code)
	rsp.WriteHeader(http.StatusOK)
}

//Regist regist new cb
func (svr *Service) Regist(cb ServiceDoIt, cmds ...string) {
	for _, cmd := range cmds {
		svr.mp[cmd] = cb
	}
}

//ServeForever start serving
func (svr *Service) ServeForever(ls string) error {
	err := http.ListenAndServe(ls, svr.sc)
	if err != nil {
		return err
	}
	return nil
}

//CommonAuth common auth
func (svr *Service) CommonAuth(ctx context.Context, chatid uint64, code string) (*models.UserInfo, bool, error) {
	info, ok := dao.GetFileStorage().QueryUserByChatid(ctx, chatid)
	if !ok {
		return nil, false, nil
	}
	if info.Code != code {
		return nil, false, nil
	}
	return info, true, nil
}
