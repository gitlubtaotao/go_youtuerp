package models

import (
	"time"
)

type FormerSeasCargoLists struct {
	Id                      int64     `xorm:"pk autoincr BIGINT(20)"`
	SoNo                    string    `xorm:"index VARCHAR(255)"`
	ContainerNo             string    `xorm:"VARCHAR(255)"`
	SealNo                  string    `xorm:"VARCHAR(255)"`
	CreatedAt               time.Time `xorm:"not null DATETIME"`
	UpdatedAt               time.Time `xorm:"not null DATETIME"`
	OrderMasterId           int       `xorm:"index INT(11)"`
	LockVersion             int       `xorm:"INT(11)"`
	VerifiedGrossMass       string    `xorm:"DECIMAL(15,4)"`
	IncludeContainer        int       `xorm:"default 0 comment('称重方式') INT(11)"`
	GrossUnit               int       `xorm:"default 0 comment('重量单位') INT(11)"`
	GrossWeight             string    `xorm:"comment('重量') DECIMAL(15,4)"`
	Measurement             string    `xorm:"comment('体积') DECIMAL(15,4)"`
	Count                   int       `xorm:"comment('件数') INT(11)"`
	CapTypeSizeId           int64     `xorm:"comment('柜型') index BIGINT(20)"`
	Marks                   string    `xorm:"comment('唛头') VARCHAR(255)"`
	DescriptionOfGood       string    `xorm:"comment('货物描述') VARCHAR(255)"`
	PackageTypeId           int64     `xorm:"comment('包裝類型') index BIGINT(20)"`
	SerialNo                string    `xorm:"comment('流水号') VARCHAR(255)"`
	Status                  string    `xorm:"VARCHAR(255)"`
	CompanyInstructionId    int       `xorm:"INT(11)"`
	ContainerWeight         string    `xorm:"DECIMAL(15,4)"`
	VgmWeight               string    `xorm:"DECIMAL(15,4)"`
	FormerSeasInstructionId int       `xorm:"INT(11)"`
	FilterType              string    `xorm:"default 'order_master' VARCHAR(64)"`
}
