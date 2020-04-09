package models

import (
	"time"
)

type BaseDataItems struct {
	Id        int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt time.Time `xorm:"not null DATETIME"`
	UpdatedAt time.Time `xorm:"not null DATETIME"`
}
