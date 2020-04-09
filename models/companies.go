package models

import (
	"time"
)

type Companies struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	Telephone  string    `xorm:"comment('座机') VARCHAR(255)"`
	Telephone2 string    `xorm:"comment('备用座机') VARCHAR(255)"`
	Fax        string    `xorm:"comment('传真') VARCHAR(255)"`
	Fax2       string    `xorm:"comment('备用传真') VARCHAR(255)"`
	Address    string    `xorm:"comment('地址') TEXT"`
	Address2   string    `xorm:"comment('备用地址') TEXT"`
	Website    string    `xorm:"comment('网站') VARCHAR(255)"`
	CreatedAt  time.Time `xorm:"DATETIME"`
	UpdatedAt  time.Time `xorm:"DATETIME"`
	SourceType string    `xorm:"index(index_companies_on_source_type_and_source_id) VARCHAR(255)"`
	SourceId   int64     `xorm:"comment('数据来源') index(index_companies_on_source_type_and_source_id) BIGINT(20)"`
	NameNick   string    `xorm:"VARCHAR(255)"`
	NameCn     string    `xorm:"VARCHAR(255)"`
	NameEn     string    `xorm:"VARCHAR(255)"`
	Code       string    `xorm:"VARCHAR(255)"`
	DeletedAt  time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	Email      string    `xorm:"VARCHAR(255)"`
	Province   string    `xorm:"comment('省份') VARCHAR(10)"`
	City       string    `xorm:"comment('市') VARCHAR(15)"`
	District   string    `xorm:"comment('区') VARCHAR(255)"`
}
