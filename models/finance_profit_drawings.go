package models

import (
	"time"
)

type FinanceProfitDrawings struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt     time.Time `xorm:"not null DATETIME"`
	UpdatedAt     time.Time `xorm:"not null DATETIME"`
	Name          string    `xorm:"comment('提成名称') VARCHAR(255)"`
	UserCompanyId int       `xorm:"index INT(11)"`
	BaseAmount    string    `xorm:"DECIMAL(13,4)"`
	EndAmount     string    `xorm:"DECIMAL(13,4)"`
	Rate          string    `xorm:"DECIMAL(4,2)"`
	Method        int       `xorm:"INT(11)"`
	Remarks       string    `xorm:"VARCHAR(255)"`
	DeletedAt     time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
