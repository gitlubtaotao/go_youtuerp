package models

import (
	"time"
)

type Verifications struct {
	Id                   int64     `xorm:"pk autoincr BIGINT(20)"`
	Number               string    `xorm:"not null comment('核销单号') VARCHAR(32)"`
	ClosingUnitName      string    `xorm:"not null comment('结算单位名称') VARCHAR(64)"`
	Currency             string    `xorm:"not null comment('币种') VARCHAR(16)"`
	Amount               string    `xorm:"comment('不含税金额') DECIMAL(13,4)"`
	TaxAmount            string    `xorm:"comment('税金') DECIMAL(13,4)"`
	Category             string    `xorm:"not null comment('类别[冲销，核销]') VARCHAR(255)"`
	ClosingUnitType      string    `xorm:"index(index_verifications_on_closing_unit_type_and_closing_unit_id) VARCHAR(255)"`
	ClosingUnitId        int64     `xorm:"comment('结算单位') index(index_verifications_on_closing_unit_type_and_closing_unit_id) BIGINT(20)"`
	CreatedAt            time.Time `xorm:"not null DATETIME"`
	UpdatedAt            time.Time `xorm:"not null DATETIME"`
	UserCompanyId        int64     `xorm:"index BIGINT(20)"`
	UserId               int64     `xorm:"index BIGINT(20)"`
	Remark               string    `xorm:"comment('备注') TEXT"`
	PayTypeId            int       `xorm:"comment('付款类型') index INT(11)"`
	CompanyAccountId     int       `xorm:"index INT(11)"`
	ClosingUnitAccountId int       `xorm:"index INT(11)"`
	AmountCost           string    `xorm:"comment('总的核销金额') DECIMAL(13,4)"`
	FinanceCurrencyId    int       `xorm:"index INT(11)"`
	PayOrReceive         string    `xorm:"VARCHAR(255)"`
	PayOrReceiveTime     time.Time `xorm:"DATETIME"`
	DeletedAt            time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
	LockVersion          int       `xorm:"default 0 INT(11)"`
}
