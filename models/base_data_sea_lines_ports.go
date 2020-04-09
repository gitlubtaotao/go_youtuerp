package models

import (
	"time"
)

type BaseDataSeaLinesPorts struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	SeaLineId      int       `xorm:"comment('航线') INT(11)"`
	SeaPortId      int       `xorm:"comment('港口') INT(11)"`
	IsLocalChanged int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedAt      time.Time `xorm:"not null DATETIME"`
}
