package models

import (
	"time"
)

type Invoices struct {
	Id                int64     `xorm:"pk autoincr BIGINT(20)"`
	UserCompaniesId   int64     `xorm:"index BIGINT(20)"`
	FinanceBanksId    int64     `xorm:"index BIGINT(20)"`
	Name              string    `xorm:"not null comment('发票名称') index VARCHAR(64)"`
	TaxpayerNumber    string    `xorm:"not null comment('纳税人识别号') VARCHAR(32)"`
	Address           string    `xorm:"not null comment('地址') TEXT"`
	PhoneNumber       string    `xorm:"not null comment('电话') VARCHAR(32)"`
	BankName          string    `xorm:"not null VARCHAR(64)"`
	BankNumber        string    `xorm:"not null comment('银行账号') VARCHAR(32)"`
	CreatedAt         time.Time `xorm:"not null DATETIME"`
	UpdatedAt         time.Time `xorm:"not null DATETIME"`
	FinanceCurrencyId int64     `xorm:"index BIGINT(20)"`
	DeletedAt         time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
	LockVersion       int       `xorm:"default 0 INT(11)"`
}
