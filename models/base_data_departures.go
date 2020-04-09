package models

import (
	"time"
)

type BaseDataDepartures struct {
	Id              int64     `xorm:"pk autoincr BIGINT(20)"`
	City            string    `xorm:"index VARCHAR(255)"`
	Address         string    `xorm:"index VARCHAR(255)"`
	CreatedAt       time.Time `xorm:"not null DATETIME"`
	UpdatedAt       time.Time `xorm:"not null DATETIME"`
	Priority        int       `xorm:"default 0 INT(11)"`
	District        string    `xorm:"comment('区域') index VARCHAR(255)"`
	DestinationCity string    `xorm:"comment('目的地') index VARCHAR(255)"`
	CityPya         string    `xorm:"comment('城市拼音全写') VARCHAR(255)"`
	CityPyf         string    `xorm:"comment('城市拼音简写') VARCHAR(255)"`
	AddressPya      string    `xorm:"comment('地址拼音全写') VARCHAR(255)"`
	AddressPyf      string    `xorm:"comment('地址拼音简写') VARCHAR(255)"`
	DistrictPya     string    `xorm:"comment('目的地拼音全写') VARCHAR(255)"`
	DistrictPyf     string    `xorm:"comment('目的地拼音简写') VARCHAR(255)"`
	Longitude       string    `xorm:"comment('经度') VARCHAR(255)"`
	Latitude        string    `xorm:"comment('纬度') VARCHAR(255)"`
	LocationAddress string    `xorm:"comment('定位地址') VARCHAR(255)"`
	DeletedAt       time.Time `xorm:"index DATETIME"`
	IsLocalChanged  int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
}
