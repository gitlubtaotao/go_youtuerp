package models

import (
	"time"
)

type FinanceCurrencies struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	RateRealtime   string    `xorm:"comment('实时汇率') DECIMAL(15,4)"`
	RateFix        string    `xorm:"comment('固定汇率') DECIMAL(15,4)"`
	RateResult     string    `xorm:"comment('实际汇率') DECIMAL(15,4)"`
	NameEn         string    `xorm:"comment('英文名') VARCHAR(255)"`
	NameCn         string    `xorm:"comment('中文名') VARCHAR(255)"`
	Symbo          string    `xorm:"comment('货币符号') index VARCHAR(255)"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedAt      time.Time `xorm:"not null DATETIME"`
	DeletedAt      time.Time `xorm:"index DATETIME"`
	IsLocalChanged int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
	ValidTime      time.Time `xorm:"comment('汇率生效时间') index DATETIME"`
	InvalidTime    time.Time `xorm:"comment('汇率失效时间') DATETIME"`
	LockVersion    int       `xorm:"default 0 INT(11)"`
}
