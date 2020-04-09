package models

import (
	"time"
)

type Addresses struct {
	Id              int64     `xorm:"pk autoincr BIGINT(20)"`
	UserCompaniesId int64     `xorm:"index BIGINT(20)"`
	UserName        string    `xorm:"not null comment('收件人姓名') VARCHAR(32)"`
	PhoneNumber     string    `xorm:"not null comment('收件人手机号') VARCHAR(32)"`
	Address         string    `xorm:"not null comment('收件人地址') TEXT"`
	PostCode        string    `xorm:"comment('邮编') VARCHAR(16)"`
	CreatedAt       time.Time `xorm:"not null DATETIME"`
	UpdatedAt       time.Time `xorm:"not null DATETIME"`
	DeletedAt       time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
}
