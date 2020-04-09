package models

import (
	"time"
)

type BaseDataInstructionTypes struct {
	Id        int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt time.Time `xorm:"not null DATETIME"`
	UpdatedAt time.Time `xorm:"not null DATETIME"`
	Name      string    `xorm:"VARCHAR(255)"`
}
