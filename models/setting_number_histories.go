package models

import (
	"time"
)

type SettingNumberHistories struct {
	Id              int64     `xorm:"pk autoincr BIGINT(20)"`
	Year            int       `xorm:"INT(11)"`
	Month           int       `xorm:"INT(11)"`
	Day             int       `xorm:"INT(11)"`
	SettingNumberId int64     `xorm:"index BIGINT(20)"`
	CurrentNumber   int       `xorm:"INT(11)"`
	CreatedAt       time.Time `xorm:"not null DATETIME"`
	UpdatedAt       time.Time `xorm:"not null DATETIME"`
}
