package models

import (
	"time"
)

type FormerSeasCapVgms struct {
	Id                        int64     `xorm:"pk autoincr BIGINT(20)"`
	OrderMasterId             int64     `xorm:"index BIGINT(20)"`
	SerialNo                  string    `xorm:"comment('序列号') VARCHAR(255)"`
	VerifiedGrossMass         int       `xorm:"comment('称重公斤数') INT(11)"`
	ResponsibleParty          string    `xorm:"comment('责任方') VARCHAR(255)"`
	AuthorizedPerson          string    `xorm:"comment('负责人') VARCHAR(255)"`
	IncludeContainer          int       `xorm:"comment('称重方式两种，一种含柜一种不含柜，此项必填，0为不含柜，1为含柜') TINYINT(1)"`
	WeighingParty             string    `xorm:"comment('称重方，可选') VARCHAR(255)"`
	Status                    string    `xorm:"comment('状态') VARCHAR(255)"`
	CreatedAt                 time.Time `xorm:"not null DATETIME"`
	UpdatedAt                 time.Time `xorm:"not null DATETIME"`
	SoNo                      string    `xorm:"VARCHAR(255)"`
	ContainerNo               string    `xorm:"TEXT"`
	SealNo                    string    `xorm:"TEXT"`
	GrossUnit                 string    `xorm:"VARCHAR(255)"`
	CompanyInstructionId      int       `xorm:"INT(11)"`
	CompanyInstructionType    string    `xorm:"VARCHAR(255)"`
	CompanyInstructionContent string    `xorm:"TEXT"`
	DeletedAt                 time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
