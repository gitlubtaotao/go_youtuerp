package models

import (
	"time"
)

type BargainSupplyQuotations struct {
	Id                       int64     `xorm:"pk autoincr BIGINT(20)"`
	SupplyId                 int       `xorm:"default 0 comment('供应商') index INT(11)"`
	BoatCompanyId            int       `xorm:"default 0 comment('船公司') index INT(11)"`
	QuotationNo              string    `xorm:"default '0' comment('报价编号') VARCHAR(255)"`
	ValidStartAt             time.Time `xorm:"default '2019-05-26 01:22:52' comment('有效期') DATETIME"`
	ValidEndAt               time.Time `xorm:"default '2019-06-02 01:22:52' comment('有效期') DATETIME"`
	CodeOfTwoId              int       `xorm:"comment('航空公司') index INT(11)"`
	Remarks                  string    `xorm:"comment('备注') VARCHAR(255)"`
	AirFees                  string    `xorm:"default 0.0000 DECIMAL(15,4)"`
	LclFees                  string    `xorm:"default 0.0000 DECIMAL(15,4)"`
	UserCompanyId            int       `xorm:"default 0 index INT(11)"`
	Platform                 string    `xorm:"default 'offline' comment('来源') VARCHAR(64)"`
	CreatedAt                time.Time `xorm:"not null DATETIME"`
	UpdatedAt                time.Time `xorm:"not null DATETIME"`
	BargainMainId            int64     `xorm:"index BIGINT(20)"`
	BargainPlatformInquiryId int64     `xorm:"index BIGINT(20)"`
	SeaPolId                 int       `xorm:"INT(11)"`
	SeaPodId                 int       `xorm:"INT(11)"`
	Status                   int       `xorm:"default 0 comment('供应商报价是否提交') index INT(11)"`
	DeletedAt                time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
	IsDirectBargain          int       `xorm:"default 0 TINYINT(1)"`
	LockVersion              int       `xorm:"default 0 INT(11)"`
}
