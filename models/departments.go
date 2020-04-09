package models

import (
	"time"
)

type Departments struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	NameCn        string    `xorm:"comment('部门中文名') index VARCHAR(255)"`
	NameEn        string    `xorm:"comment('部门英文名') VARCHAR(255)"`
	CreatedAt     time.Time `xorm:"not null DATETIME"`
	UpdatedAt     time.Time `xorm:"not null DATETIME"`
	UserCompanyId int       `xorm:"index INT(11)"`
	DeletedAt     time.Time `xorm:"index DATETIME"`
	LockVersion   int       `xorm:"default 0 INT(11)"`
}
