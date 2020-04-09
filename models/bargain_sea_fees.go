package models

import (
	"time"
)

type BargainSeaFees struct {
	Id                         int64     `xorm:"pk autoincr BIGINT(20)"`
	UnitPrice                  string    `xorm:"default 0.0000 DECIMAL(15,4)"`
	Count                      int       `xorm:"default 0 comment('数量') INT(11)"`
	FinanceCurrency            string    `xorm:"default 'CNY' comment('币种') VARCHAR(64)"`
	SupplyId                   int       `xorm:"default 0 comment('供应商') index INT(11)"`
	BargainSupplyQuotationId   int       `xorm:"default 0 comment('供应商报价') index INT(11)"`
	BargainCustomerQuotationId int       `xorm:"default 0 comment('客户报价单') index INT(11)"`
	BoxSizeId                  int       `xorm:"default 0 INT(11)"`
	CreatedAt                  time.Time `xorm:"not null DATETIME"`
	UpdatedAt                  time.Time `xorm:"not null DATETIME"`
	BargainMainId              int64     `xorm:"index BIGINT(20)"`
	BargainBoxSizeCountId      int64     `xorm:"index BIGINT(20)"`
	DeletedAt                  time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
	LockVersion                int       `xorm:"default 0 INT(11)"`
}
