package tgnotify

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"tgnotify/models"
)

//ServiceDoMSG service msg
type ServiceDoMSG struct {
}

//OnDo do logic
func (sd *ServiceDoMSG) OnDo(ctx context.Context, svr *Service, cmd string, user *models.UserInfo, req *http.Request, tgnt *TGBot) error {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	mode := strings.TrimSpace(req.Header.Get("mode"))
	err = tgnt.WriteModeBot(int64(user.Chatid), mode, string(data))
	if err != nil {
		return err
	}
	return nil
}

//OnAuth auth logic
func (sd *ServiceDoMSG) OnAuth(ctx context.Context, svr *Service, chatid uint64, code string) (*models.UserInfo, bool, error) {
	return svr.CommonAuth(ctx, chatid, code)
}
