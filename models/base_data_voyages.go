package models

import (
	"time"
)

type BaseDataVoyages struct {
	Id                    int64     `xorm:"pk autoincr BIGINT(20)"`
	BaseDataBoatCompanyId int64     `xorm:"index BIGINT(20)"`
	Name                  string    `xorm:"VARCHAR(128)"`
	Enabled               int       `xorm:"comment('是否有效') TINYINT(1)"`
	CreatedAt             time.Time `xorm:"not null DATETIME"`
	UpdatedAt             time.Time `xorm:"not null DATETIME"`
	DeletedAt             time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
