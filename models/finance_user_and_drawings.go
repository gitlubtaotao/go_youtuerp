package models

import (
	"time"
)

type FinanceUserAndDrawings struct {
	Id                     int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt              time.Time `xorm:"not null DATETIME"`
	UpdatedAt              time.Time `xorm:"not null DATETIME"`
	UserId                 int64     `xorm:"index BIGINT(20)"`
	FinanceProfitDrawingId int64     `xorm:"index BIGINT(20)"`
	DeletedAt              time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
}
