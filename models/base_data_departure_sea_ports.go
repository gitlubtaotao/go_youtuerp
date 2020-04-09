package models

import (
	"time"
)

type BaseDataDepartureSeaPorts struct {
	Id              int64     `xorm:"pk autoincr BIGINT(20)"`
	City            string    `xorm:"comment('起运港城市') VARCHAR(255)"`
	CreatedAt       time.Time `xorm:"not null DATETIME"`
	UpdatedAt       time.Time `xorm:"not null DATETIME"`
	Port            string    `xorm:"comment('起运港') VARCHAR(255)"`
	Position        int       `xorm:"default 90000000 comment('位置') index INT(11)"`
	CityId          string    `xorm:"comment('城市ID') index VARCHAR(255)"`
	PortEn          string    `xorm:"comment('港口英文名') VARCHAR(255)"`
	PortPya         string    `xorm:"comment('港口拼音全写') VARCHAR(255)"`
	PortPyf         string    `xorm:"comment('港口拼音简写') VARCHAR(255)"`
	CityEn          string    `xorm:"comment('城市英文名') VARCHAR(255)"`
	CityPya         string    `xorm:"comment('城市拼音全写') VARCHAR(255)"`
	CityPyf         string    `xorm:"comment('城市拼音简写') VARCHAR(255)"`
	HasFreight      int       `xorm:"default 1 comment('附加费') TINYINT(1)"`
	Longitude       string    `xorm:"comment('经度') VARCHAR(255)"`
	Latitude        string    `xorm:"comment('纬度') VARCHAR(255)"`
	LocationAddress string    `xorm:"comment('定位地址') VARCHAR(255)"`
	DeletedAt       time.Time `xorm:"index DATETIME"`
	IsLocalChanged  int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
}
