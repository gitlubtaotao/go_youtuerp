package models

import (
	"time"
)

type BaseDataCountries struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedAt      time.Time `xorm:"not null DATETIME"`
	Name           string    `xorm:"comment('国家中文名') VARCHAR(255)"`
	NameEn         string    `xorm:"comment('国家英文名') VARCHAR(255)"`
	NamePya        string    `xorm:"comment('国家中文名全拼') VARCHAR(255)"`
	NamePyf        string    `xorm:"VARCHAR(255)"`
	Code           string    `xorm:"comment('国家代码') VARCHAR(255)"`
	DeletedAt      time.Time `xorm:"index DATETIME"`
	IsLocalChanged int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
	NameEnShort    string    `xorm:"VARCHAR(255)"`
	NameCnShort    string    `xorm:"comment('中文简称') VARCHAR(255)"`
}
