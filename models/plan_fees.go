package models

import (
	"time"
)

type PlanFees struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	PlanMainId    int64     `xorm:"index BIGINT(20)"`
	CreatedAt     time.Time `xorm:"not null DATETIME"`
	UpdatedAt     time.Time `xorm:"not null DATETIME"`
	ClosingUnitId int       `xorm:"comment('结算单位') index INT(11)"`
	Name          string    `xorm:"comment('简称') VARCHAR(64)"`
	NameCn        string    `xorm:"comment('中文名称') VARCHAR(128)"`
	NameEn        string    `xorm:"comment('英文名称') VARCHAR(128)"`
	CurrencyId    int       `xorm:"comment('币种') INT(11)"`
	UnitPrice     string    `xorm:"comment('单价') DECIMAL(19,4)"`
	Count         string    `xorm:"comment('数量') DECIMAL(19,4)"`
	Amount        string    `xorm:"comment('总价') DECIMAL(19,4)"`
	PayOrReceive  string    `xorm:"comment('收支类型') index VARCHAR(32)"`
	Enabled       int       `xorm:"default 1 comment('有效状态') TINYINT(1)"`
	IsSync        int       `xorm:"default 0 TINYINT(1)"`
	DeletedAt     time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
}
