package models

import (
	"time"
)

type BaseDataPackageTypes struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	Name           string    `xorm:"comment('包装单位简称') VARCHAR(255)"`
	NameCn         string    `xorm:"comment('中文名称') VARCHAR(255)"`
	NameEn         string    `xorm:"comment('英文名称') VARCHAR(255)"`
	EdiCode        string    `xorm:"comment('EDI代码') VARCHAR(255)"`
	Remark         string    `xorm:"comment('备注') VARCHAR(255)"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedAt      time.Time `xorm:"not null DATETIME"`
	DeletedAt      time.Time `xorm:"index DATETIME"`
	IsLocalChanged int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
}
