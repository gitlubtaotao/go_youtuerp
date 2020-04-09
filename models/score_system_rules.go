package models

import (
	"time"
)

type ScoreSystemRules struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt  time.Time `xorm:"not null DATETIME"`
	UpdatedAt  time.Time `xorm:"not null DATETIME"`
	RuleType   string    `xorm:"comment('规则类型[获得积分,使用积分]') VARCHAR(255)"`
	Group      string    `xorm:"comment('分组[用户注册,在线登录,交易奖励,积分抵扣]') VARCHAR(255)"`
	Key        string    `xorm:"comment('类型') VARCHAR(255)"`
	Value      string    `xorm:"comment('积分值') VARCHAR(255)"`
	FilterType string    `xorm:"comment('类型') VARCHAR(255)"`
	Label      string    `xorm:"comment('标识') VARCHAR(255)"`
	CompanyId  int       `xorm:"default 0 index INT(11)"`
}
