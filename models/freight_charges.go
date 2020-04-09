package models

import (
	"time"
)

type FreightCharges struct {
	Id                 int64     `xorm:"pk autoincr BIGINT(20)"`
	Type               string    `xorm:"not null comment('类型【整箱,拼箱,拖车】') index VARCHAR(64)"`
	Name               string    `xorm:"not null comment('加价设置名称') index VARCHAR(128)"`
	Group              string    `xorm:"not null comment('加价组别【内部加价,客户加价,未登录用户】') index VARCHAR(64)"`
	DeparturePortIds   string    `xorm:"comment('起运港/城市') VARCHAR(255)"`
	BoatCompanyIds     string    `xorm:"comment('船公司') VARCHAR(255)"`
	SeaLineIds         string    `xorm:"comment('航线') VARCHAR(255)"`
	DestinationPortIds string    `xorm:"comment('目的港') VARCHAR(255)"`
	ScaleType          string    `xorm:"not null comment('加价类型【金额，百分比】') VARCHAR(64)"`
	ScaleGroupIds      string    `xorm:"comment('用户组ids') VARCHAR(255)"`
	ScaleValue         int       `xorm:"comment('加价数值') INT(11)"`
	Gp20               string    `xorm:"comment('20GP') DECIMAL(12,4)"`
	Gp40               string    `xorm:"comment('40GP') DECIMAL(12,4)"`
	Hq40               string    `xorm:"comment('40GP') DECIMAL(12,4)"`
	Tons               string    `xorm:"comment('tons') DECIMAL(12,4)"`
	Cbms               string    `xorm:"comment('cbms') DECIMAL(12,4)"`
	LclCharge          string    `xorm:"comment('附加费') DECIMAL(12,4)"`
	ChargeOne          string    `xorm:"comment('拖车价格一段') DECIMAL(12,4)"`
	ChargeTwo          string    `xorm:"comment('拖车价格二段') DECIMAL(12,4)"`
	ChargeThree        string    `xorm:"comment('拖车价格三段') DECIMAL(12,4)"`
	ChargeFour         string    `xorm:"comment('拖车价格四段') DECIMAL(12,4)"`
	ChargeFive         string    `xorm:"comment('拖车价格五段') DECIMAL(12,4)"`
	ChargeSix          string    `xorm:"comment('拖车价格六段') DECIMAL(12,4)"`
	Enabled            int       `xorm:"default 1 comment('有效状态') TINYINT(1)"`
	CreatedAt          time.Time `xorm:"not null DATETIME"`
	UpdatedAt          time.Time `xorm:"not null DATETIME"`
	Position           int64     `xorm:"BIGINT(20)"`
	UserCompanyId      int       `xorm:"index INT(11)"`
	DeletedAt          time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	LockVersion        int       `xorm:"default 0 INT(11)"`
}
