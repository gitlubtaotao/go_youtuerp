package models

import (
	"time"
)

type InvoiceApplyLists struct {
	Id                int64     `xorm:"pk autoincr BIGINT(20)"`
	InvoiceId         int       `xorm:"INT(11)"`
	AddressId         int       `xorm:"INT(11)"`
	Category          string    `xorm:"VARCHAR(255)"`
	Number            string    `xorm:"comment('流水号') index VARCHAR(255)"`
	Status            string    `xorm:"VARCHAR(255)"`
	UserCompanyId     int       `xorm:"INT(11)"`
	ClosingUnitId     int       `xorm:"index INT(11)"`
	ApplyUserId       int       `xorm:"INT(11)"`
	FinanceCurrencyId int       `xorm:"INT(11)"`
	Amount            string    `xorm:"default 0.0000 DECIMAL(13,4)"`
	FinanceFeeIds     string    `xorm:"VARCHAR(255)"`
	Note              string    `xorm:"VARCHAR(255)"`
	CreatedAt         time.Time `xorm:"not null DATETIME"`
	UpdatedAt         time.Time `xorm:"not null DATETIME"`
	DeletedAt         time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
	LockVersion       int       `xorm:"default 0 INT(11)"`
}
