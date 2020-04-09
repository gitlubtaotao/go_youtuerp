package models

import (
	"time"
)

type Fielders struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt     time.Time `xorm:"not null DATETIME"`
	UpdatedAt     time.Time `xorm:"not null DATETIME"`
	Name          string    `xorm:"default '' comment('字段名称') VARCHAR(255)"`
	Label         string    `xorm:"default '' comment('字段备注') VARCHAR(255)"`
	FieldType     string    `xorm:"default '' comment('字段类型') VARCHAR(255)"`
	Comment       string    `xorm:"default '' comment('字段的描述') VARCHAR(255)"`
	Hint          string    `xorm:"default '' comment('字段的提示') VARCHAR(255)"`
	Accessibility int       `xorm:"not null comment('只读,可读可写,隐藏字段') INT(11)"`
	UserId        int64     `xorm:"index BIGINT(20)"`
	LockVersion   int       `xorm:"default 0 comment('行级乐观锁') INT(11)"`
	TableName     string    `xorm:"comment('表名') index VARCHAR(255)"`
}
