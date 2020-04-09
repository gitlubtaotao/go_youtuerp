package models

import (
	"time"
)

type FeeVerifications struct {
	Id              int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt       time.Time `xorm:"not null DATETIME"`
	UpdatedAt       time.Time `xorm:"not null DATETIME"`
	FinanceFeesId   int64     `xorm:"index BIGINT(20)"`
	VerificationsId int64     `xorm:"index BIGINT(20)"`
	AuditAmount     string    `xorm:"comment('核销金额') DECIMAL(13,4)"`
	DeletedAt       time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
}
