package models

import (
	"time"
)

type FormerOtherServices struct {
	Remarks                   string    `xorm:"TEXT"`
	Id                        int64     `xorm:"pk autoincr BIGINT(20)"`
	OrderMasterId             int       `xorm:"comment('订单') index INT(11)"`
	CompanyInstructionId      int       `xorm:"comment('委托单位') INT(11)"`
	IsFumigation              int       `xorm:"default 0 TINYINT(1)"`
	FumigationId              int       `xorm:"INT(11)"`
	IsTraders                 int       `xorm:"default 0 TINYINT(1)"`
	TradersId                 int       `xorm:"INT(11)"`
	IsInsurance               int       `xorm:"default 0 TINYINT(1)"`
	InsuranceId               int       `xorm:"INT(11)"`
	Beneficiary               string    `xorm:"VARCHAR(255)"`
	IsMagneticTest            int       `xorm:"default 0 TINYINT(1)"`
	MagneticTestId            int       `xorm:"INT(11)"`
	IsIdentification          int       `xorm:"default 0 TINYINT(1)"`
	IdentificationId          int       `xorm:"INT(11)"`
	SerialNo                  string    `xorm:"VARCHAR(255)"`
	CompanyInstructionType    string    `xorm:"VARCHAR(255)"`
	CompanyInstructionContent string    `xorm:"TEXT"`
	DeletedAt                 time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	CommodityInspectionId     int       `xorm:"INT(11)"`
	IsCommodityInspection     int       `xorm:"default 0 TINYINT(1)"`
}
