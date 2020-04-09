package models

import (
	"time"
)

type ApprovalApplications struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	OperatorId    int       `xorm:"not null comment('申请人id') INT(11)"`
	OperatorName  string    `xorm:"not null comment('申请人姓名') VARCHAR(64)"`
	Note          string    `xorm:"comment('备注') TEXT"`
	Status        string    `xorm:"not null comment('状态') index VARCHAR(32)"`
	ApprovalTimes int       `xorm:"not null default 1 comment('当前审核次数') INT(11)"`
	Number        string    `xorm:"not null comment('申请单号') index VARCHAR(64)"`
	CreatedAt     time.Time `xorm:"not null DATETIME"`
	UpdatedAt     time.Time `xorm:"not null DATETIME"`
	UserCompanyId int64     `xorm:"index BIGINT(20)"`
	DeletedAt     time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	PayOrReceive  string    `xorm:"index VARCHAR(32)"`
	AuditorId     int       `xorm:"INT(11)"`
	AuditorName   string    `xorm:"VARCHAR(64)"`
	BatchNumber   string    `xorm:"default '0' comment('批量申请单号') VARCHAR(255)"`
	OrderMasterId int       `xorm:"index INT(11)"`
	ClosingUnitId int       `xorm:"index INT(11)"`
	LockVersion   int       `xorm:"default 0 INT(11)"`
}
