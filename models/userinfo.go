package models

type UserInfo struct {
	Chatid uint64 `json:"chatid"`
	Code   string `json:"code"`
}

type FileStorage struct {
	Users []*UserInfo `json:"users"`
}
