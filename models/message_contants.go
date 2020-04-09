package models

import (
	"time"
)

type MessageContants struct {
	Id          int64     `xorm:"pk autoincr BIGINT(20)"`
	Category    string    `xorm:"comment('类型(group群组、personal个人、copy抄送)') VARCHAR(255)"`
	MailboxId   int       `xorm:"comment('对应信息id') index INT(11)"`
	UserId      int       `xorm:"comment('接收者id(group_id、person_id)') index INT(11)"`
	UserName    string    `xorm:"comment('接收者姓名') VARCHAR(255)"`
	CompanyId   int       `xorm:"comment('接收者公司id') index INT(11)"`
	CompanyName string    `xorm:"comment('接收者公司名称') VARCHAR(255)"`
	CreatedAt   time.Time `xorm:"not null DATETIME"`
	UpdatedAt   time.Time `xorm:"not null DATETIME"`
	Email       string    `xorm:"VARCHAR(255)"`
	DeletedAt   time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	Mark        string    `xorm:"index VARCHAR(8)"`
	LockVersion int       `xorm:"default 0 INT(11)"`
}
