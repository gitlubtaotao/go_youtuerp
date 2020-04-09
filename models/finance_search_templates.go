package models

import (
	"time"
)

type FinanceSearchTemplates struct {
	Id          int64     `xorm:"pk autoincr BIGINT(20)"`
	Name        string    `xorm:"not null comment('模板名称') VARCHAR(64)"`
	Enabled     int       `xorm:"not null default 1 comment('有效状态') TINYINT(1)"`
	UserId      int       `xorm:"not null comment('用户id') INT(11)"`
	Content     string    `xorm:"not null comment('搜索内容(json字符串)') TEXT"`
	Description string    `xorm:"comment('描述') TEXT"`
	CreatedAt   time.Time `xorm:"not null DATETIME"`
	UpdatedAt   time.Time `xorm:"not null DATETIME"`
}
