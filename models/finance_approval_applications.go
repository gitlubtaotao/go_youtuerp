package models

import (
	"time"
)

type FinanceApprovalApplications struct {
	Id                    int64     `xorm:"pk autoincr BIGINT(20)"`
	FinanceFeeId          int64     `xorm:"index BIGINT(20)"`
	ApprovalApplicationId int64     `xorm:"index BIGINT(20)"`
	CreatedAt             time.Time `xorm:"not null DATETIME"`
	UpdatedAt             time.Time `xorm:"not null DATETIME"`
	DeletedAt             time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
