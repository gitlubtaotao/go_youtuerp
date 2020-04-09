package models

import (
	"time"
)

type WebsiteArticles struct {
	Id        int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt time.Time `xorm:"not null DATETIME"`
	UpdatedAt time.Time `xorm:"not null DATETIME"`
	Title     string    `xorm:"comment('文章标题') index VARCHAR(64)"`
	Type      string    `xorm:"comment('分类') index VARCHAR(255)"`
	Keyword   string    `xorm:"comment('关键字') TEXT"`
	Details   string    `xorm:"comment('详情内容') LONGTEXT"`
	Enabled   int       `xorm:"default 0 comment('是否可见') TINYINT(1)"`
	Locale    string    `xorm:"default 'zh-CN' comment('语言类型') VARCHAR(255)"`
	ReadCount int       `xorm:"comment('已读次数') INT(11)"`
	DeletedAt time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
