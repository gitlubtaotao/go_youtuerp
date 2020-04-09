package models

import (
	"time"
)

type BargainExtendInfos struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	Remarks1      string    `xorm:"comment('存放优途价格（经加价），客户议价，供应商价（原价）') TEXT"`
	Remarks2      string    `xorm:"TEXT"`
	Remarks3      string    `xorm:"TEXT"`
	Remarks4      string    `xorm:"TEXT"`
	Remarks5      string    `xorm:"TEXT"`
	CreatedAt     time.Time `xorm:"not null DATETIME"`
	UpdatedAt     time.Time `xorm:"not null DATETIME"`
	BargainMainId int64     `xorm:"index BIGINT(20)"`
	DeletedAt     time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
}
