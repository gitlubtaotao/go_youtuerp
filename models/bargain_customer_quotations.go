package models

import (
	"time"
)

type BargainCustomerQuotations struct {
	Id                       int64     `xorm:"pk autoincr BIGINT(20)"`
	CompanyInstructionId     int       `xorm:"default 0 comment('客户') index INT(11)"`
	CustomerQuotationNo      string    `xorm:"default '0' comment('编号') VARCHAR(255)"`
	OrderMasterId            int       `xorm:"default 0 comment('订单') index INT(11)"`
	FinishedTime             time.Time `xorm:"comment('完成时间') DATETIME"`
	Status                   string    `xorm:"default 'init' comment('状态') VARCHAR(64)"`
	Remarks                  string    `xorm:"comment('状态') VARCHAR(255)"`
	UserCompanyId            int       `xorm:"default 0 index INT(11)"`
	AirFees                  string    `xorm:"default 0.0000 DECIMAL(15,4)"`
	LclFees                  string    `xorm:"default 0.0000 DECIMAL(15,4)"`
	BargainSupplyQuotationId int       `xorm:"index INT(11)"`
	CreatedAt                time.Time `xorm:"not null DATETIME"`
	UpdatedAt                time.Time `xorm:"not null DATETIME"`
	BargainMainId            int64     `xorm:"index BIGINT(20)"`
	DeletedAt                time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
	IsDirectBargain          int       `xorm:"default 0 comment('直接生成客户报价') TINYINT(1)"`
	LockVersion              int       `xorm:"default 0 INT(11)"`
}
