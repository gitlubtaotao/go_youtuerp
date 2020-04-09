package models

import (
	"time"
)

type CargoLists struct {
	Id                     int64     `xorm:"pk autoincr BIGINT(20)"`
	OrderMasterId          int64     `xorm:"index BIGINT(20)"`
	Enabled                int       `xorm:"default 1 comment('有效状态') TINYINT(1)"`
	CompanyInstructionId   int       `xorm:"comment('委托单位') INT(11)"`
	CompanyInstructionType string    `xorm:"comment('委托单位类型') VARCHAR(64)"`
	Name                   string    `xorm:"comment('货物名称') TEXT"`
	UnitCount              int       `xorm:"comment('数量') INT(11)"`
	PackageTypeId          int       `xorm:"comment('包装类型id') INT(11)"`
	PackageTypeName        string    `xorm:"comment('包装类型名称') VARCHAR(64)"`
	Length                 string    `xorm:"comment('长') DECIMAL(19,4)"`
	Width                  string    `xorm:"comment('宽') DECIMAL(19,4)"`
	Height                 string    `xorm:"comment('高') DECIMAL(19,4)"`
	Size                   string    `xorm:"comment('体积') DECIMAL(19,4)"`
	Count                  int       `xorm:"comment('总数') INT(11)"`
	UnitPrice              string    `xorm:"comment('单价') DECIMAL(19,4)"`
	Amount                 string    `xorm:"comment('总价') DECIMAL(19,4)"`
	Marks                  string    `xorm:"comment('唛头') VARCHAR(255)"`
	GrossWeight            string    `xorm:"comment('毛重') DECIMAL(19,4)"`
	UnitSize               string    `xorm:"comment('单件体积') DECIMAL(19,4)"`
	CreatedAt              time.Time `xorm:"not null DATETIME"`
	UpdatedAt              time.Time `xorm:"not null DATETIME"`
	DeletedAt              time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
	CustomsNumber          string    `xorm:"comment('报关序号') VARCHAR(255)"`
	DeclareElements        string    `xorm:"comment('申报要素') TEXT"`
	CustomsCode            string    `xorm:"comment('海关编码') VARCHAR(255)"`
	NetWeight              string    `xorm:"default 0.0000 DECIMAL(14,4)"`
	SourceType             string    `xorm:"index(index_cargo_lists_on_source_id_and_source_type) VARCHAR(255)"`
	SourceId               int       `xorm:"index(index_cargo_lists_on_source_id_and_source_type) INT(11)"`
	LockVersion            int       `xorm:"default 0 INT(11)"`
	FilterType             string    `xorm:"default 'order_master' VARCHAR(64)"`
	InstructionId          int       `xorm:"INT(11)"`
}
