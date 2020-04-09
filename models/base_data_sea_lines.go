package models

import (
	"time"
)

type BaseDataSeaLines struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	Name           string    `xorm:"comment('航线名') VARCHAR(255)"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedAt      time.Time `xorm:"not null DATETIME"`
	EnName         string    `xorm:"comment('航线英文名') VARCHAR(255)"`
	DeletedAt      time.Time `xorm:"index DATETIME"`
	IsLocalChanged int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
}
