package models

import (
	"time"
)

type People struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	NameNick      string    `xorm:"comment('昵称') VARCHAR(255)"`
	NameCn        string    `xorm:"comment('中文名') VARCHAR(255)"`
	NameEn        string    `xorm:"comment('英文名') VARCHAR(255)"`
	Gender        int       `xorm:"default 0 comment('0：未知，1：男，2女') INT(11)"`
	Post          string    `xorm:"comment('职位') VARCHAR(255)"`
	Mobi          string    `xorm:"comment('手机号码') VARCHAR(255)"`
	Mobi2         string    `xorm:"comment('备用号码') VARCHAR(255)"`
	Telephone     string    `xorm:"comment('座机') VARCHAR(255)"`
	Telephone2    string    `xorm:"comment('备用座机') VARCHAR(255)"`
	Fax           string    `xorm:"comment('传真') VARCHAR(255)"`
	Email         string    `xorm:"VARCHAR(255)"`
	Qq            string    `xorm:"VARCHAR(255)"`
	AddrCompany   string    `xorm:"comment('公司地址') VARCHAR(255)"`
	AddrHome      string    `xorm:"comment('家庭地址') VARCHAR(255)"`
	Remark        string    `xorm:"comment('备注') VARCHAR(255)"`
	Remark2       string    `xorm:"comment('备用备注') VARCHAR(255)"`
	IsMainContact int       `xorm:"comment('是否主要联系人') TINYINT(1)"`
	CreatedAt     time.Time `xorm:"DATETIME"`
	UpdatedAt     time.Time `xorm:"DATETIME"`
	SourceType    string    `xorm:"index(index_person_on_source_type_and_source_id) VARCHAR(255)"`
	SourceId      int64     `xorm:"comment('数据来源') index(index_person_on_source_type_and_source_id) BIGINT(20)"`
	DeletedAt     time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
