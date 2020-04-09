package models

import (
	"time"
)

type Approvers struct {
	Id                    int64     `xorm:"pk autoincr BIGINT(20)"`
	ApprovalApplicationId int64     `xorm:"index BIGINT(20)"`
	UserId                int       `xorm:"not null comment('审批人id') index INT(11)"`
	UserName              string    `xorm:"not null comment('审批人姓名') VARCHAR(64)"`
	Sort                  int       `xorm:"not null default 1 comment('审批顺序') INT(11)"`
	Status                string    `xorm:"not null default 'created' comment('审批状态') index VARCHAR(32)"`
	CreatedAt             time.Time `xorm:"not null DATETIME"`
	UpdatedAt             time.Time `xorm:"not null DATETIME"`
	DeletedAt             time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	LockVersion           int       `xorm:"default 0 INT(11)"`
}
