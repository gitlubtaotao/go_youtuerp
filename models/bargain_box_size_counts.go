package models

import (
	"time"
)

type BargainBoxSizeCounts struct {
	Id              int64     `xorm:"pk autoincr BIGINT(20)"`
	BoxSizeId       int       `xorm:"default 0 index INT(11)"`
	BoxSizeCount    int       `xorm:"default 0 INT(11)"`
	TargetPrice     string    `xorm:"default 0.0000 DECIMAL(15,4)"`
	FinanceCurrency string    `xorm:"default 'USD' comment('币种') VARCHAR(64)"`
	FinishedPrice   string    `xorm:"default 0.0000 comment('成交价格') DECIMAL(15,4)"`
	CreatedAt       time.Time `xorm:"not null DATETIME"`
	UpdatedAt       time.Time `xorm:"not null DATETIME"`
	BargainMainId   int64     `xorm:"index BIGINT(20)"`
	DeletedAt       time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	LockVersion     int       `xorm:"default 0 INT(11)"`
}
