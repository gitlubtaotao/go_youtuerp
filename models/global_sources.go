package models

import (
	"time"
)

type GlobalSources struct {
	Id                 int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt          time.Time `xorm:"not null DATETIME"`
	UpdatedAt          time.Time `xorm:"not null DATETIME"`
	SourceType         string    `xorm:"index(index_global_sources_on_source_type_and_source_id) VARCHAR(255)"`
	SourceId           int64     `xorm:"comment('数据来源') index(index_global_sources_on_source_type_and_source_id) BIGINT(20)"`
	SourcePlatform     string    `xorm:"comment('来源平台') index VARCHAR(255)"`
	SourceRecordPath   string    `xorm:"comment('来源路径') VARCHAR(255)"`
	SourceUuid         string    `xorm:"VARCHAR(255)"`
	SourceUyid         string    `xorm:"VARCHAR(255)"`
	SourceUuidRegister int       `xorm:"default 0 comment('唯一id注册') TINYINT(1)"`
	ParentSourceUuid   string    `xorm:"VARCHAR(255)"`
	DeletedAt          time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
