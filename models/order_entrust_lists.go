package models

import (
	"time"
)

type OrderEntrustLists struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedAt      time.Time `xorm:"not null DATETIME"`
	OrderMasterId  int64     `xorm:"index BIGINT(20)"`
	UserCompanyId  int64     `xorm:"index BIGINT(20)"`
	ClientUserId   int       `xorm:"comment('前端下单人') INT(11)"`
	PayStatus      string    `xorm:"default 'unfinished' VARCHAR(255)"`
	VerifyStatus   string    `xorm:"default 'unfinished' VARCHAR(255)"`
	LandingBillUrl string    `xorm:"comment('生成的提单URL') VARCHAR(255)"`
	InstructionUrl string    `xorm:"comment('委托的URL') VARCHAR(255)"`
	DeletedAt      time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
}
