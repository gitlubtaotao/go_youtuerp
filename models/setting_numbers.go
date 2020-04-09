package models

import (
	"time"
)

type SettingNumbers struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt     time.Time `xorm:"not null DATETIME"`
	UpdatedAt     time.Time `xorm:"not null DATETIME"`
	Prefix        string    `xorm:"comment('前缀') VARCHAR(64)"`
	Special       string    `xorm:"comment('特殊字符') VARCHAR(64)"`
	NumberLength  int       `xorm:"comment('流水号长度') INT(11)"`
	YearRule      string    `xorm:"comment('年设置') VARCHAR(255)"`
	MonthRule     string    `xorm:"comment('月设置') VARCHAR(255)"`
	DayRule       string    `xorm:"comment('日设置') VARCHAR(255)"`
	UserNumber    string    `xorm:"default 'user' comment('用户工号') VARCHAR(255)"`
	ApplicationNo string    `xorm:"comment('应用') index VARCHAR(64)"`
	DefaultNumber int       `xorm:"default 0 comment('默认的流水号长度') INT(11)"`
	ClearMethod   string    `xorm:"comment('清空方式') VARCHAR(64)"`
	CurrentNumber string    `xorm:"comment('当前流水号') VARCHAR(255)"`
	Rule          string    `xorm:"comment('规则') VARCHAR(255)"`
	OldRule       string    `xorm:"comment('旧规则') VARCHAR(255)"`
	DeletedAt     time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
	IsChange      int       `xorm:"default 0 comment('是否修改') TINYINT(1)"`
	NumberRule    string    `xorm:"VARCHAR(255)"`
	UserCompanyId int64     `xorm:"BIGINT(20)"`
}
