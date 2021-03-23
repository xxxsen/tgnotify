package tgcmds

import "tgnotify"

var mp = make(map[string]tgnotify.TGCallbackBuilder)

func Regist(cmd string, cb tgnotify.TGCallbackBuilder) {
	mp[cmd] = cb
}

func GetTGCMDS() map[string]tgnotify.TGCallbackBuilder {
	return mp
}
