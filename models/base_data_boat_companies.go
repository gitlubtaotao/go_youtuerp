package models

import (
	"time"
)

type BaseDataBoatCompanies struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	ChinaName      string    `xorm:"comment('中文名') VARCHAR(255)"`
	Name           string    `xorm:"comment('船运公司名称') unique VARCHAR(255)"`
	Url            string    `xorm:"comment('官网地址') VARCHAR(255)"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedAt      time.Time `xorm:"not null DATETIME"`
	ChinaNamePya   string    `xorm:"comment('船运公司拼音全写') VARCHAR(255)"`
	ChinaNamePyf   string    `xorm:"VARCHAR(255)"`
	DeletedAt      time.Time `xorm:"index DATETIME"`
	IsLocalChanged int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
}
