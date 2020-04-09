package models

import (
	"time"
)

type MessageContents struct {
	Id          int64     `xorm:"pk autoincr BIGINT(20)"`
	Title       string    `xorm:"not null comment('标题') VARCHAR(255)"`
	Context     string    `xorm:"comment('内容') TEXT"`
	Attachment  string    `xorm:"comment('附件') VARCHAR(255)"`
	MailboxId   int       `xorm:"comment('邮件id') index INT(11)"`
	CreatedAt   time.Time `xorm:"not null DATETIME"`
	UpdatedAt   time.Time `xorm:"not null DATETIME"`
	DeletedAt   time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	LockVersion int       `xorm:"default 0 INT(11)"`
}
