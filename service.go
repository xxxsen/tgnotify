package tgnotify

import (
	"net/http"
	"tgnotify/dao"
	"tgnotify/log"
	"tgnotify/models"
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

//UserInfo user info
type UserInfo struct {
	User   string //user
	Code   string //code
	ChatID int64  //chatid
}

type ServiceDoIt interface {
	OnDo(svr *Service, cmd string, user *UserInfo, req *http.Request, tgnt *TGBot) error
	OnAuth(svr *Service, user string, code string) (*UserInfo, bool, error)
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
		log.Errorf("Recv unimplement cmd from user, cmd:%s", cmd)
		return
	}
	user := req.Header.Get("user")
	code := req.Header.Get("code")
	u, succ, err := fc.OnAuth(sc.svr, user, code)
	if err != nil {
		rsp.WriteHeader(http.StatusInternalServerError)
		log.Errorf("Recv cmd:%s but auth err, err:%v", cmd, err)
		return
	}
	if !succ {
		rsp.WriteHeader(http.StatusUnauthorized)
		log.Errorf("Recv cmd:%s, auth fail, user:%s, code:%s", cmd, user, code)
		return
	}

	err = fc.OnDo(sc.svr, cmd, u, req, sc.svr.bot)
	if err != nil {
		log.Errorf("Do cmd:%s cause internal err, user:%s, header:%+v, err:%v", cmd, user, req.Header, err)
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Tracef("Cmd:%s exec success, user:%s, code:%s", cmd, user, code)
	rsp.WriteHeader(http.StatusOK)
}

//Regist regist new cb
func (svr *Service) Regist(cmd string, cb ServiceDoIt) {
	svr.mp[cmd] = cb
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
func (svr *Service) CommonAuth(user, code string) (*UserInfo, bool, error) {
	sql := "select * from tbl_tgnotify where user = ? and code = ? limit 1"
	nt := &[]models.TblTgnotify{}
	err := dao.GetEngine().SQL(sql, user, code).Find(nt)
	if err != nil {
		return nil, false, err
	}
	if len(*nt) == 0 {
		return nil, false, nil
	}
	item := (*nt)[0]
	return &UserInfo{User: item.User, Code: item.Code, ChatID: item.Chatid}, true, nil
}
