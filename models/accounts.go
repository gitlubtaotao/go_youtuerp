package models

import (
	"time"
)

type Accounts struct {
	Id                  int64     `xorm:"pk autoincr BIGINT(20)"`
	FinanceBanksId      int64     `xorm:"index BIGINT(20)"`
	Name                string    `xorm:"not null comment('账户名称') index VARCHAR(64)"`
	UserName            string    `xorm:"not null comment('开户人姓名') VARCHAR(64)"`
	BankName            string    `xorm:"not null comment('开户行') VARCHAR(64)"`
	BankNumber          string    `xorm:"not null comment('银行账号') VARCHAR(32)"`
	Category            string    `xorm:"comment('账户类型') VARCHAR(32)"`
	CreatedAt           time.Time `xorm:"not null DATETIME"`
	UpdatedAt           time.Time `xorm:"not null DATETIME"`
	UserCompanyId       int64     `xorm:"index BIGINT(20)"`
	BankAddress         string    `xorm:"comment('开户行地址') VARCHAR(255)"`
	SwiftCode           string    `xorm:"comment('swift code') VARCHAR(255)"`
	TaxRegisterNumber   string    `xorm:"VARCHAR(255)"`
	BeneficiaryAddress  string    `xorm:"comment('开户人地址') TEXT"`
	BeneficiaryLocation string    `xorm:"comment('开户人位置') VARCHAR(255)"`
	DeletedAt           time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	LockVersion         int       `xorm:"default 0 INT(11)"`
}
