package models

import (
	"time"
)

type InvoiceInfoFinanceFees struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	InvoiceInfoId int64     `xorm:"index BIGINT(20)"`
	FinanceFeeId  int64     `xorm:"index BIGINT(20)"`
	Enabled       int       `xorm:"default 1 TINYINT(1)"`
	CreatedAt     time.Time `xorm:"not null DATETIME"`
	UpdatedAt     time.Time `xorm:"not null DATETIME"`
	Amount        string    `xorm:"comment('开票金额') DECIMAL(13,4)"`
	DeletedAt     time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	LockVersion   int       `xorm:"default 0 INT(11)"`
}
