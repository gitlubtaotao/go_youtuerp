package models

import (
	"time"
)

type FinanceFeeTypes struct {
	Id                int64     `xorm:"pk autoincr BIGINT(20)"`
	Name              string    `xorm:"comment('费用名') index VARCHAR(255)"`
	NameCn            string    `xorm:"comment('中文名') VARCHAR(255)"`
	NameEn            string    `xorm:"comment('英文名') VARCHAR(255)"`
	ReceivableTagId   int       `xorm:"comment('默认应收标签') INT(11)"`
	PayableTagId      int       `xorm:"comment('默认应付标签') INT(11)"`
	DefaultValue      string    `xorm:"comment('默认金额') DECIMAL(15,2)"`
	Remarks           string    `xorm:"comment('备注') TEXT"`
	FinanceCurrencyId int       `xorm:"comment('外部关联货币') INT(11)"`
	CreatedAt         time.Time `xorm:"not null DATETIME"`
	UpdatedAt         time.Time `xorm:"not null DATETIME"`
	DeletedAt         time.Time `xorm:"index DATETIME"`
	IsLocalChanged    int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
	LockVersion       int       `xorm:"default 0 INT(11)"`
}
