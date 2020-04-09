package models

import (
	"time"
)

type BaseDataVessels struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	Name          string    `xorm:"VARCHAR(255)"`
	BoatCompanyId string    `xorm:"VARCHAR(255)"`
	Enabled       int       `xorm:"default 1 TINYINT(1)"`
	CreatedAt     time.Time `xorm:"not null DATETIME"`
	UpdatedAt     time.Time `xorm:"not null DATETIME"`
	DeletedAt     time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
