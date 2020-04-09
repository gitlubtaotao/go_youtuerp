package models

import (
	"time"
)

type OrderWorks struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	OrderMasterId int       `xorm:"not null comment('订单id') index INT(11)"`
	Name          string    `xorm:"not null comment('任务名称') VARCHAR(64)"`
	Note          string    `xorm:"comment('备注') TEXT"`
	Enabled       int       `xorm:"not null default 1 comment('有效状态') TINYINT(1)"`
	Warning       int       `xorm:"not null default 0 comment('提醒标志') TINYINT(1)"`
	Status        string    `xorm:"not null default 'uncompleted' comment('状态') index VARCHAR(64)"`
	Rank          int       `xorm:"not null comment('排序权位') INT(11)"`
	CompletedTime time.Time `xorm:"comment('完成时间') DATETIME"`
	CreatedAt     time.Time `xorm:"not null DATETIME"`
	UpdatedAt     time.Time `xorm:"not null DATETIME"`
	NameCn        string    `xorm:"VARCHAR(255)"`
	LockVersion   int       `xorm:"default 0 INT(11)"`
}
