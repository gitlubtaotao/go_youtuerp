package models

import (
	"time"
)

type BaseDataDestinations struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	City           string    `xorm:"comment('城市') VARCHAR(255)"`
	Address        string    `xorm:"comment('地址') VARCHAR(255)"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedAt      time.Time `xorm:"not null DATETIME"`
	CityPya        string    `xorm:"comment('城市拼音全写') VARCHAR(255)"`
	CityPyf        string    `xorm:"comment('城市拼音简写') VARCHAR(255)"`
	AddressPya     string    `xorm:"comment('地址拼音全写') VARCHAR(255)"`
	AddressPyf     string    `xorm:"comment('地址拼音简写') VARCHAR(255)"`
	DeletedAt      time.Time `xorm:"index DATETIME"`
	IsLocalChanged int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
}
