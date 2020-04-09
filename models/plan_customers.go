package models

import (
	"time"
)

type PlanCustomers struct {
	Id                     int64     `xorm:"pk autoincr BIGINT(20)"`
	PlanMainId             int64     `xorm:"index BIGINT(20)"`
	CreatedAt              time.Time `xorm:"not null DATETIME"`
	UpdatedAt              time.Time `xorm:"not null DATETIME"`
	CompanyInstructionId   int       `xorm:"comment('委托单位') index INT(11)"`
	CompanyInstructionType string    `xorm:"comment('委托单位类型') VARCHAR(64)"`
	OceanChangesPaytypeId  int       `xorm:"comment('运费付款方式') INT(11)"`
	OtherChangesPaytypeId  int       `xorm:"comment('其他运费付款方式') INT(11)"`
	TradeTermId            int       `xorm:"comment('贸易条款') INT(11)"`
	TermId                 int       `xorm:"comment('装运条款') INT(11)"`
	BusinessTypeId         int       `xorm:"comment('业务类型') INT(11)"`
	BillProduceId          int       `xorm:"comment('出单方式') INT(11)"`
	HblNumber              string    `xorm:"comment('hbl编号') VARCHAR(64)"`
	BoxSizeCount           string    `xorm:"comment('柜型柜量') VARCHAR(255)"`
	IsFile                 int       `xorm:"default 0 comment('是否随机文件') TINYINT(1)"`
	Enabled                int       `xorm:"default 1 comment('有效状态') TINYINT(1)"`
	DeletedAt              time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
	LockVersion            int       `xorm:"default 0 INT(11)"`
}
