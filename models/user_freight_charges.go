package models

import (
	"time"
)

type UserFreightCharges struct {
	Id              int64     `xorm:"pk autoincr BIGINT(20)"`
	UserId          int64     `xorm:"index BIGINT(20)"`
	FreightChargeId int64     `xorm:"index BIGINT(20)"`
	Enabled         int       `xorm:"default 1 TINYINT(1)"`
	CreatedAt       time.Time `xorm:"not null DATETIME"`
	UpdatedAt       time.Time `xorm:"not null DATETIME"`
}
