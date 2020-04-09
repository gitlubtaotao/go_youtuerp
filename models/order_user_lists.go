package models

import (
	"time"
)

type OrderUserLists struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	OrderMasterId int64     `xorm:"index BIGINT(20)"`
	UserId        int64     `xorm:"index BIGINT(20)"`
	Role          string    `xorm:"VARCHAR(64)"`
	DeletedAt     time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
}
