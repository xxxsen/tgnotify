package models

import (
	"time"
)

type TblTgnotify struct {
	Id     int       `xorm:"not null pk autoincr INT(10)"`
	User   string    `xorm:"not null unique(uniq_uinfo) VARCHAR(32)"`
	Code   string    `xorm:"not null unique(uniq_uinfo) VARCHAR(32)"`
	Chatid int64     `xorm:"not null unique unique(uniq_uinfo) BIGINT(20)"`
	Ts     time.Time `xorm:"not null default 'current_timestamp()' TIMESTAMP"`
}
