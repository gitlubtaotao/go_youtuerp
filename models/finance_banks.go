package models

import (
	"time"
)

type FinanceBanks struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	BankCn         string    `xorm:"index VARCHAR(255)"`
	BankEn         string    `xorm:"comment('英文名称') VARCHAR(255)"`
	BankAbbr       string    `xorm:"comment('英文简写') VARCHAR(255)"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedAt      time.Time `xorm:"not null DATETIME"`
	DeletedAt      time.Time `xorm:"index DATETIME"`
	IsLocalChanged int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
	LockVersion    int       `xorm:"default 0 INT(11)"`
}
