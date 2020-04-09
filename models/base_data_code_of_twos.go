package models

import (
	"time"
)

type BaseDataCodeOfTwos struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	Code           string    `xorm:"comment('编码') VARCHAR(255)"`
	Airport        string    `xorm:"comment('机场') VARCHAR(255)"`
	Name           string    `xorm:"comment('机场名') VARCHAR(255)"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedAt      time.Time `xorm:"not null DATETIME"`
	AirportPya     string    `xorm:"comment('机场拼音全写') VARCHAR(255)"`
	AirportPyf     string    `xorm:"comment('机场拼音简写') VARCHAR(255)"`
	NamePya        string    `xorm:"comment('机场名全写') VARCHAR(255)"`
	NamePyf        string    `xorm:"comment('机场名简写') VARCHAR(255)"`
	DeletedAt      time.Time `xorm:"index DATETIME"`
	IsLocalChanged int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
	Website        string    `xorm:"comment('网站') VARCHAR(255)"`
}
