package models

import (
	"time"
)

type YoutuCrmTags struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	Type          string    `xorm:"index index(index_youtu_crm_tags_on_type_and_transport_type) VARCHAR(255)"`
	TagId         int       `xorm:"index INT(11)"`
	TagName       string    `xorm:"VARCHAR(255)"`
	SourceType    string    `xorm:"index(index_youtu_crm_tags_on_source_type_and_source_id) VARCHAR(255)"`
	SourceId      int64     `xorm:"index(index_youtu_crm_tags_on_source_type_and_source_id) BIGINT(20)"`
	DeletedAt     time.Time `xorm:"default '1969-12-31 16:00:00' comment('默认软删除') DATETIME"`
	TransportType int       `xorm:"index index(index_youtu_crm_tags_on_type_and_transport_type) INT(11)"`
	Remarks       string    `xorm:"VARCHAR(255)"`
}
