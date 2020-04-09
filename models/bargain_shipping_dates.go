package models

import (
	"time"
)

type BargainShippingDates struct {
	Id                       int64     `xorm:"pk autoincr BIGINT(20)"`
	Days                     string    `xorm:"default '1' comment('航程/天数') VARCHAR(64)"`
	StartDate                int       `xorm:"default 1 comment('截关/截港/航班日期') INT(11)"`
	Remarks                  string    `xorm:"comment('备注') VARCHAR(255)"`
	BargainMainId            int64     `xorm:"index BIGINT(20)"`
	BargainSupplyQuotationId int64     `xorm:"index BIGINT(20)"`
	DeletedAt                time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	CutOffDay                int       `xorm:"default 1 INT(11)"`
	TransshipmentInfo        string    `xorm:"default '' VARCHAR(64)"`
	SeaTransshipmentId       int       `xorm:"INT(11)"`
	AirTransshipmentId       int       `xorm:"INT(11)"`
	LockVersion              int       `xorm:"default 0 INT(11)"`
}
