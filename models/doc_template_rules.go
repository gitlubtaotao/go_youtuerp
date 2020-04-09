package models

import (
	"time"
)

type DocTemplateRules struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	Name          string    `xorm:"comment('模板名称') VARCHAR(255)"`
	TransportType int       `xorm:"default 1 comment('运输方式') INT(11)"`
	MainTransport int       `xorm:"default 1 comment('其他对应的主运输方式') INT(11)"`
	DeletedAt     time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	Content       string    `xorm:"comment('模板对应的列设置内容') MEDIUMTEXT"`
	RuleType      string    `xorm:"default 'sea_instruction' VARCHAR(32)"`
	CreatedAt     time.Time `xorm:"not null DATETIME"`
	UpdatedAt     time.Time `xorm:"not null DATETIME"`
}
