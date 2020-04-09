package models

import (
	"time"
)

type ChangeLogs struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt     time.Time `xorm:"not null index DATETIME"`
	UpdatedAt     time.Time `xorm:"not null DATETIME"`
	SourceType    string    `xorm:"index(index_change_logs_on_source_type_and_source_id) VARCHAR(255)"`
	SourceId      int64     `xorm:"comment('数据来源') index(index_change_logs_on_source_type_and_source_id) BIGINT(20)"`
	Content       string    `xorm:"comment('操作内容') LONGTEXT"`
	UserId        int64     `xorm:"index BIGINT(20)"`
	IpAddress     string    `xorm:"comment('ip address') VARCHAR(255)"`
	Remark        string    `xorm:"comment('用户操作备注') index VARCHAR(255)"`
	UserCompanyId int       `xorm:"index INT(11)"`
	LogType       string    `xorm:"default 'log' VARCHAR(64)"`
}
