package models

import (
	"time"
)

type MessageMailboxes struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	FromId     int64     `xorm:"not null comment('发件人id') BIGINT(20)"`
	FromType   string    `xorm:"comment('发件人类型') index VARCHAR(24)"`
	FromName   string    `xorm:"comment('发件人姓名') VARCHAR(255)"`
	FromEmail  string    `xorm:"comment('发送者邮件') VARCHAR(255)"`
	Status     string    `xorm:"comment('状态') VARCHAR(24)"`
	IsReplay   string    `xorm:"not null comment('是否可回复') VARCHAR(24)"`
	SendTime   time.Time `xorm:"not null comment('发送时间') DATETIME"`
	SourceId   int       `xorm:"comment('信息来源id') unique(index_message_mailboxes_on_source_type_and_source_id) INT(11)"`
	SourceType string    `xorm:"comment('信息来源') unique(index_message_mailboxes_on_source_type_and_source_id) VARCHAR(255)"`
	Category   string    `xorm:"comment('类别') VARCHAR(24)"`
	Ways       string    `xorm:"comment('渠道') VARCHAR(255)"`
	CreatedAt  time.Time `xorm:"not null DATETIME"`
	UpdatedAt  time.Time `xorm:"not null DATETIME"`
	DeletedAt  time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	Enabled    int       `xorm:"TINYINT(4)"`
}
