package models

import (
	"time"
)

type BaseDataCodes struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedAt      time.Time `xorm:"not null DATETIME"`
	CodeName       string    `xorm:"index VARCHAR(255)"`
	Name           string    `xorm:"VARCHAR(255)"`
	Remark         string    `xorm:"comment('备注') TEXT"`
	SourceType     string    `xorm:"comment('来源') VARCHAR(255)"`
	CodeLevelId    int64     `xorm:"index BIGINT(20)"`
	Status         int       `xorm:"default 1 TINYINT(1)"`
	LockVersion    int       `xorm:"default 0 comment('行级乐观锁') INT(11)"`
	DeletedAt      time.Time `xorm:"index DATETIME"`
	IsLocalChanged int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
}
