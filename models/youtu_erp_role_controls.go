package models

import (
	"time"
)

type YoutuErpRoleControls struct {
	Id        int64     `xorm:"pk autoincr BIGINT(20)"`
	Name      string    `xorm:"default 'all' comment('角色名') VARCHAR(64)"`
	UserId    int64     `xorm:"index BIGINT(20)"`
	DeletedAt time.Time `xorm:"DATETIME"`
}
