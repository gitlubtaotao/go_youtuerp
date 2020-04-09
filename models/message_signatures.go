package models

import (
	"time"
)

type MessageSignatures struct {
	Id        int64     `xorm:"pk autoincr BIGINT(20)"`
	Content   string    `xorm:"TEXT"`
	UserId    int       `xorm:"index INT(11)"`
	Name      string    `xorm:"VARCHAR(255)"`
	CreatedAt time.Time `xorm:"not null DATETIME"`
	UpdatedAt time.Time `xorm:"not null DATETIME"`
	IsDefault int       `xorm:"default 0 index TINYINT(1)"`
}
