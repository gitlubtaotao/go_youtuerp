package models

import (
	"time"
)

type TrackModels struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt     time.Time `xorm:"not null DATETIME"`
	UpdatedAt     time.Time `xorm:"not null DATETIME"`
	OrderMasterId int       `xorm:"index INT(11)"`
	Content       string    `xorm:"TEXT"`
	Position      int       `xorm:"index INT(11)"`
	Status        string    `xorm:"VARCHAR(255)"`
	Process       string    `xorm:"VARCHAR(255)"`
}
