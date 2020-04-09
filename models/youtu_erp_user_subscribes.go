package models

import (
	"time"
)

type YoutuErpUserSubscribes struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	UserId     int       `xorm:"index INT(11)"`
	CreatedAt  time.Time `xorm:"not null DATETIME"`
	UpdatedAt  time.Time `xorm:"not null DATETIME"`
	OpenId     string    `xorm:"VARCHAR(255)"`
	Nickname   string    `xorm:"VARCHAR(255)"`
	Sex        int       `xorm:"INT(11)"`
	Province   string    `xorm:"VARCHAR(255)"`
	Country    string    `xorm:"VARCHAR(255)"`
	Headimgurl string    `xorm:"VARCHAR(255)"`
	Unionid    string    `xorm:"VARCHAR(255)"`
	DeletedAt  time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
