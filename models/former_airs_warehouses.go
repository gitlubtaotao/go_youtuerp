package models

import (
	"time"
)

type FormerAirsWarehouses struct {
	Id                       int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt                time.Time `xorm:"not null DATETIME"`
	UpdatedAt                time.Time `xorm:"not null DATETIME"`
	SerialNo                 string    `xorm:"comment('序列号') VARCHAR(255)"`
	WarehouseNo              string    `xorm:"comment('入仓号') VARCHAR(255)"`
	Flight                   string    `xorm:"comment('航班') VARCHAR(255)"`
	FlightDate               time.Time `xorm:"comment('航班日期') DATETIME"`
	DepartureAirportId       int       `xorm:"comment('起运港') index INT(11)"`
	DestinationAirportId     int       `xorm:"comment('目的港') index INT(11)"`
	WarehouseAddress         string    `xorm:"comment('入仓地址') TEXT"`
	SupplierAgent            string    `xorm:"comment('联系人') VARCHAR(255)"`
	SupplierAgentMobi        string    `xorm:"comment('联系人电话') VARCHAR(255)"`
	Remarks                  string    `xorm:"comment('备注') TEXT"`
	Status                   string    `xorm:"VARCHAR(255)"`
	OrderMasterId            int       `xorm:"index INT(11)"`
	OceanChangesPaytypeId    int       `xorm:"comment('费用支付方式') INT(11)"`
	OtherChangesPaytypeId    int       `xorm:"comment('其他费用支付方式') INT(11)"`
	EstimatedTimeOfDeparture time.Time `xorm:"DATETIME"`
	CodeOfTwoId              int       `xorm:"index INT(11)"`
	TransshipmentAirportId   int       `xorm:"INT(11)"`
	DeletedAt                time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
