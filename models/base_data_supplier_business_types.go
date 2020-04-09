package models

import (
	"time"
)

type BaseDataSupplierBusinessTypes struct {
	Id        int64     `xorm:"pk autoincr BIGINT(20)"`
	Name      string    `xorm:"VARCHAR(255)"`
	CreatedAt time.Time `xorm:"not null DATETIME"`
	UpdatedAt time.Time `xorm:"not null DATETIME"`
	DeletedAt time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
