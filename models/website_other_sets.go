package models

import (
	"time"
)

type WebsiteOtherSets struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	CompanyName      string    `xorm:"comment('公司名称') VARCHAR(255)"`
	CompanyNameEn    string    `xorm:"comment('公司英文名称') VARCHAR(255)"`
	CompanyAddress   string    `xorm:"comment('联系方式') TEXT"`
	CompanyAddressEn string    `xorm:"comment('英文联系方式') TEXT"`
	RecordNumber     string    `xorm:"comment('备案号') VARCHAR(255)"`
	DeletedAt        time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
