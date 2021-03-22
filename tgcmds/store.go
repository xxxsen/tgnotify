package tgcmds

import "tgnotify"

var mp = make(map[string]tgnotify.TGCallback)

func Regist(cmd string, cb tgnotify.TGCallback) {
	mp[cmd] = cb
}

func GetTGCMDS() map[string]tgnotify.TGCallback {
	return mp
}
