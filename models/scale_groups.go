package models

import (
	"time"
)

type ScaleGroups struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	Name          string    `xorm:"comment('群组名') VARCHAR(128)"`
	Enabled       int       `xorm:"default 1 comment('有效状态') TINYINT(1)"`
	CreatedAt     time.Time `xorm:"not null DATETIME"`
	UpdatedAt     time.Time `xorm:"not null DATETIME"`
	UserCompanyId int       `xorm:"index INT(11)"`
}
