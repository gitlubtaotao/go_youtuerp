package models

import (
	"time"
)

type FreightSearchLogs struct {
	Id                   int64     `xorm:"pk autoincr BIGINT(20)"`
	DestinationAddress   string    `xorm:"comment('目的地') VARCHAR(50)"`
	DepartureCity        string    `xorm:"comment('城市') VARCHAR(50)"`
	DepartureDistrict    string    `xorm:"comment('地区') VARCHAR(50)"`
	DepartureAddress     string    `xorm:"comment('出发地') VARCHAR(255)"`
	CityName             string    `xorm:"comment('城市') VARCHAR(15)"`
	DepartureSeaPortName string    `xorm:"comment('起运港') VARCHAR(20)"`
	CityId               string    `xorm:"comment('城市') VARCHAR(50)"`
	DepartureSeaPortId   string    `xorm:"comment('起运港') VARCHAR(50)"`
	SeaLineId            string    `xorm:"comment('航线') VARCHAR(50)"`
	DestinationPortId    string    `xorm:"comment('目的港') VARCHAR(50)"`
	BoatCompanyId        string    `xorm:"comment('船公司') VARCHAR(50)"`
	QueryType            string    `xorm:"comment('查询类型【拖车，拼箱，整柜】') index VARCHAR(30)"`
	IpAddress            string    `xorm:"comment('客户登录ip') VARCHAR(30)"`
	UserId               int64     `xorm:"comment('系统用户id,可为空') BIGINT(20)"`
	CreatedAt            time.Time `xorm:"not null DATETIME"`
	UpdatedAt            time.Time `xorm:"not null DATETIME"`
	DeletedAt            time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	RawInfoJson          string    `xorm:"LONGTEXT"`
}
