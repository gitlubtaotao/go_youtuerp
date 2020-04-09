package models

import (
	"time"
)

type UserShareData struct {
	Id        int64     `xorm:"pk autoincr BIGINT(20)"`
	UserId    int       `xorm:"comment('所属员工') index INT(11)"`
	ShareId   int       `xorm:"comment('被共享的id') INT(11)"`
	DeletedAt time.Time `xorm:"DATETIME"`
}
