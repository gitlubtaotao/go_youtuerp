package models

import (
	"time"
)

type Schedules struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	Name       string    `xorm:"not null comment('待办事项名称') VARCHAR(255)"`
	Category   string    `xorm:"not null comment('类型') VARCHAR(64)"`
	Note       string    `xorm:"comment('备注') TEXT"`
	Status     string    `xorm:"not null default 'created' comment('状态') VARCHAR(64)"`
	UserId     int       `xorm:"not null comment('提醒人id') INT(11)"`
	UserName   string    `xorm:"not null comment('提醒人姓名') VARCHAR(32)"`
	ObjectId   int       `xorm:"not null comment('数据源id') INT(11)"`
	ObjectName string    `xorm:"not null comment('数据源名称') VARCHAR(64)"`
	ObjectUrl  string    `xorm:"not null comment('数据源url') VARCHAR(255)"`
	CreatedAt  time.Time `xorm:"not null DATETIME"`
	UpdatedAt  time.Time `xorm:"not null DATETIME"`
}
