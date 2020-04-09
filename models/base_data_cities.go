package models

import (
	"time"
)

type BaseDataCities struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	Name           string    `xorm:"comment('城市名') unique VARCHAR(255)"`
	ZoneId         int       `xorm:"comment('区域ID') index INT(11)"`
	NameEn         string    `xorm:"comment('城市英文名') VARCHAR(255)"`
	NamePya        string    `xorm:"comment('城市拼音全写') VARCHAR(255)"`
	NamePyf        string    `xorm:"comment('城市拼音简拼') VARCHAR(255)"`
	HasFreight     int       `xorm:"default 1 TINYINT(1)"`
	Position       int       `xorm:"default 1 INT(11)"`
	DeletedAt      time.Time `xorm:"index DATETIME"`
	IsLocalChanged int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
}
