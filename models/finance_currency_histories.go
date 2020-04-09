package models

import (
	"time"
)

type FinanceCurrencyHistories struct {
	Id                int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt         time.Time `xorm:"not null DATETIME"`
	UpdatedAt         time.Time `xorm:"not null DATETIME"`
	FinanceCurrencyId int64     `xorm:"index BIGINT(20)"`
	RateRealtime      string    `xorm:"comment('实时汇率') DECIMAL(15,4)"`
	RateFix           string    `xorm:"comment('固定汇率') DECIMAL(15,4)"`
	RateResult        string    `xorm:"comment('实际汇率') DECIMAL(15,4)"`
	ValidTime         time.Time `xorm:"comment('汇率生效时间') DATETIME"`
	InvalidTime       time.Time `xorm:"comment('汇率失效时间') DATETIME"`
	UserId            int64     `xorm:"index BIGINT(20)"`
	LockVersion       int       `xorm:"default 0 INT(11)"`
	DeletedAt         time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
