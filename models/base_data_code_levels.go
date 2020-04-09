package models

import (
	"time"
)

type BaseDataCodeLevels struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedAt      time.Time `xorm:"not null DATETIME"`
	Name           string    `xorm:"comment('层级名称') VARCHAR(255)"`
	CodeName       string    `xorm:"VARCHAR(255)"`
	Status         int       `xorm:"default 1 comment('是否有效') TINYINT(1)"`
	SourceType     string    `xorm:"comment('来源') index VARCHAR(255)"`
	Remark         string    `xorm:"comment('说明') TEXT"`
	LockVersion    int       `xorm:"default 0 comment('行级乐观锁') INT(11)"`
	DeletedAt      time.Time `xorm:"index DATETIME"`
	IsLocalChanged int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
}
