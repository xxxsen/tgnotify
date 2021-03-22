package tgnotify

import (
	"io/ioutil"
	"net/http"
)

//ServiceDoMSG service msg
type ServiceDoMSG struct {
}

//OnDo do logic
func (sd *ServiceDoMSG) OnDo(svr *Service, cmd string, user *UserInfo, req *http.Request, tgnt *TGBot) error {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	mode := req.Header.Get("mode")
	err = tgnt.WriteModeBot(user.ChatID, mode, string(data))
	if err != nil {
		return err
	}
	return nil
}

//OnAuth auth logic
func (sd *ServiceDoMSG) OnAuth(svr *Service, user string, code string) (*UserInfo, bool, error) {
	return svr.CommonAuth(user, code)
}
