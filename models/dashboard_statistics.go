package models

import (
	"time"
)

type DashboardStatistics struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	ObjId          int       `xorm:"not null comment('对象id') INT(11)"`
	Category       string    `xorm:"not null comment('对象类别') VARCHAR(64)"`
	ObjName        string    `xorm:"comment('对象名称') VARCHAR(255)"`
	Description    string    `xorm:"comment('描述') TEXT"`
	OrganizationId int       `xorm:"comment('组织id') INT(11)"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedAt      time.Time `xorm:"not null DATETIME"`
}
