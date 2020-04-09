package models

import (
	"time"
)

type ArInternalMetadata struct {
	Key       string    `xorm:"not null pk VARCHAR(255)"`
	Value     string    `xorm:"VARCHAR(255)"`
	CreatedAt time.Time `xorm:"not null DATETIME"`
	UpdatedAt time.Time `xorm:"not null DATETIME"`
}
