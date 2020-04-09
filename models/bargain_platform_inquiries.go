package models

import (
	"time"
)

type BargainPlatformInquiries struct {
	Id                  int64     `xorm:"pk autoincr BIGINT(20)"`
	InquiryNo           string    `xorm:"default '0' comment('询价编号') VARCHAR(255)"`
	Status              string    `xorm:"default 'init' comment('状态') VARCHAR(64)"`
	LastedQuotationTime time.Time `xorm:"comment('最近报价时间') DATETIME"`
	SupplyId            int       `xorm:"default 0 comment('供应商') index INT(11)"`
	UserCompanyId       int       `xorm:"default 0 index INT(11)"`
	Remarks             string    `xorm:"comment('询价备注') VARCHAR(255)"`
	CreatedAt           time.Time `xorm:"not null DATETIME"`
	UpdatedAt           time.Time `xorm:"not null DATETIME"`
	BargainMainId       int64     `xorm:"index BIGINT(20)"`
	DeletedAt           time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
}
