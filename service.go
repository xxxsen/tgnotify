package tgnotify

import (
	"log"
	"net/http"
	"tgnotify/dao"
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
		log.Printf("Recv unimplement cmd from user, cmd:%s\n", cmd)
		return
	}
	user := req.Header.Get("user")
	code := req.Header.Get("code")
	u, succ, err := fc.OnAuth(sc.svr, user, code)
	if err != nil {
		rsp.WriteHeader(http.StatusInternalServerError)
		log.Printf("Recv cmd:%s but auth err, err:%v\n", cmd, err)
		return
	}
	if !succ {
		rsp.WriteHeader(http.StatusUnauthorized)
		log.Printf("Recv cmd:%s, auth fail, user:%s, code:%s\n", cmd, user, code)
		return
	}

	err = fc.OnDo(sc.svr, cmd, u, req, sc.svr.bot)
	if err != nil {
		log.Printf("Do cmd:%s cause internal err, user:%s, header:%+v, err:%v\n", cmd, user, req.Header, err)
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}
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
