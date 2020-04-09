package models

import (
	"time"
)

type ActiveStorageBlobs struct {
	Id          int64     `xorm:"pk autoincr BIGINT(20)"`
	Key         string    `xorm:"not null unique VARCHAR(255)"`
	Filename    string    `xorm:"not null VARCHAR(255)"`
	ContentType string    `xorm:"VARCHAR(255)"`
	Metadata    string    `xorm:"TEXT"`
	ByteSize    int64     `xorm:"not null BIGINT(20)"`
	Checksum    string    `xorm:"not null VARCHAR(255)"`
	CreatedAt   time.Time `xorm:"not null DATETIME"`
}
