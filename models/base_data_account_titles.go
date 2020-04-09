package models

import (
	"time"
)

type BaseDataAccountTitles struct {
	Id        int64     `xorm:"pk autoincr BIGINT(20)"`
	Name      string    `xorm:"comment('会计科目名称') VARCHAR(50)"`
	CreatedAt time.Time `xorm:"not null DATETIME"`
	UpdatedAt time.Time `xorm:"not null DATETIME"`
	DeletedAt time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
