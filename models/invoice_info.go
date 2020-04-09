package models

import (
	"time"
)

type InvoiceInfo struct {
	Id                 int64     `xorm:"pk autoincr BIGINT(20)"`
	InvoiceId          int64     `xorm:"index BIGINT(20)"`
	AddressId          int64     `xorm:"index BIGINT(20)"`
	Category           string    `xorm:"not null comment('发票类型') VARCHAR(64)"`
	Number             string    `xorm:"not null comment('发票号') VARCHAR(255)"`
	CourierCompanyId   int       `xorm:"comment('快递公司') INT(11)"`
	CourierCompanyName string    `xorm:"comment('快递公司名称') VARCHAR(255)"`
	CourierNumber      string    `xorm:"comment('快递单号') VARCHAR(128)"`
	Amount             string    `xorm:"comment('发票金额') DECIMAL(13,4)"`
	Status             string    `xorm:"default 'applied' comment('发票状态') VARCHAR(64)"`
	Enabled            int       `xorm:"default 1 comment('有效状态') TINYINT(1)"`
	Note               string    `xorm:"comment('备注') TEXT"`
	Date               time.Time `xorm:"comment('开票时间') DATETIME"`
	CreatedAt          time.Time `xorm:"not null DATETIME"`
	UpdatedAt          time.Time `xorm:"not null DATETIME"`
	ClosingUnitId      int       `xorm:"comment('结算单位') index INT(11)"`
	UserCompanyId      int       `xorm:"comment('所属公司') index INT(11)"`
	UserId             int       `xorm:"index INT(11)"`
	ApplyUserId        int       `xorm:"index INT(11)"`
	FinanceCurrencyId  int       `xorm:"index INT(11)"`
	DeletedAt          time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
	LockVersion        int       `xorm:"default 0 INT(11)"`
}
