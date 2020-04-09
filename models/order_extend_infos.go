package models

import (
	"time"
)

type OrderExtendInfos struct {
	Id                  int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt           time.Time `xorm:"not null DATETIME"`
	UpdatedAt           time.Time `xorm:"not null DATETIME"`
	Number              string    `xorm:"comment('件数') DECIMAL(13,4)"`
	BaseDataPackageType int       `xorm:"comment('类型') INT(11)"`
	GrossWeight         string    `xorm:"comment('毛重') DECIMAL(13,4)"`
	Size                string    `xorm:"comment('体积') DECIMAL(13,4)"`
	HblSo               string    `xorm:"comment('分单号') VARCHAR(128)"`
	MblSo               string    `xorm:"comment('主单号') VARCHAR(128)"`
	BoatCompanyId       int       `xorm:"comment('船公司') index INT(11)"`
	SeaPolId            int       `xorm:"comment('起运港') INT(11)"`
	SeaPodId            int       `xorm:"comment('目的港') INT(11)"`
	AirPolId            int       `xorm:"comment('起运港') INT(11)"`
	AirPodId            int       `xorm:"comment('目的港') INT(11)"`
	CutOffDay           time.Time `xorm:"comment('截关日期/开船日期') DATETIME"`
	FlightDate          time.Time `xorm:"comment('起飞日期') DATETIME"`
	EndDate             time.Time `xorm:"comment('到达时间') index DATETIME"`
	IsCutOff            int       `xorm:"comment('是否开船') INT(11)"`
	IsArrive            int       `xorm:"comment('是否到港') INT(11)"`
	OrderMasterId       int64     `xorm:"index BIGINT(20)"`
	Ratio               string    `xorm:"comment('换算系数') DECIMAL(13,4)"`
	RatioWeight         string    `xorm:"comment('材积重') DECIMAL(13,4)"`
	Bubble              string    `xorm:"comment('分泡') DECIMAL(13,4)"`
	SoNo                string    `xorm:"comment('so信息') TEXT"`
	CodeOfTwoId         int       `xorm:"index INT(11)"`
	CourierCodeId       int       `xorm:"INT(11)"`
	StartDate           time.Time `xorm:"index DATETIME"`
	IsHml               int       `xorm:"default 0 TINYINT(1)"`
	IsAms               int       `xorm:"default 0 TINYINT(1)"`
	IsIfs               int       `xorm:"default 0 TINYINT(1)"`
	ChargedWeight       string    `xorm:"DECIMAL(13,4)"`
	DeletedAt           time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
	MiscBill            int       `xorm:"INT(11)"`
	BoxSizeCount        string    `xorm:"VARCHAR(255)"`
	Vessel              string    `xorm:"VARCHAR(255)"`
	Voyage              string    `xorm:"VARCHAR(255)"`
	Flight              string    `xorm:"VARCHAR(255)"`
	LockVersion         int       `xorm:"default 0 INT(11)"`
}
