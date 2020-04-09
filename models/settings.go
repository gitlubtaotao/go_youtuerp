package models

import (
	"time"
)

type Settings struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	SourceType string    `xorm:"index(index_settings_on_source_type_and_source_id) VARCHAR(255)"`
	SourceId   int64     `xorm:"index(index_settings_on_source_type_and_source_id) BIGINT(20)"`
	Key        string    `xorm:"index VARCHAR(255)"`
	Value      string    `xorm:"TEXT"`
	Group      string    `xorm:"index VARCHAR(255)"`
	FilterType string    `xorm:"VARCHAR(255)"`
	CreatedAt  time.Time `xorm:"not null DATETIME"`
	UpdatedAt  time.Time `xorm:"not null DATETIME"`
	Label      string    `xorm:"VARCHAR(255)"`
	CompanyId  int       `xorm:"default 0 index INT(11)"`
}
