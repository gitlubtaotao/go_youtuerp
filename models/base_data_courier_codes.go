package models

import (
	"time"
)

type BaseDataCourierCodes struct {
	Id          int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt   time.Time `xorm:"not null DATETIME"`
	UpdatedAt   time.Time `xorm:"not null DATETIME"`
	Name        string    `xorm:"VARCHAR(255)"`
	Code        string    `xorm:"VARCHAR(255)"`
	CourierType int       `xorm:"default 0 INT(11)"`
	DeletedAt   time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
