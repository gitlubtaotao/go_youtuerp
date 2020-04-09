package models

import (
	"time"
)

type PlanCustoms struct {
	Id                     int64     `xorm:"pk autoincr BIGINT(20)"`
	PlanMainId             int64     `xorm:"index BIGINT(20)"`
	CreatedAt              time.Time `xorm:"not null DATETIME"`
	UpdatedAt              time.Time `xorm:"not null DATETIME"`
	Enabled                int       `xorm:"default 1 comment('有效状态') TINYINT(1)"`
	Note                   string    `xorm:"comment('备注') TEXT"`
	CompanyInstructionId   int       `xorm:"comment('委托单位') index INT(11)"`
	CompanyInstructionType string    `xorm:"comment('委托单位类型') VARCHAR(255)"`
	BaseDataCiqtypeId      int       `xorm:"comment('报关类型') INT(11)"`
	CustomsAmount          string    `xorm:"comment('报关金额') DECIMAL(19,4)"`
	IsPay                  int       `xorm:"default 0 comment('是否买单') TINYINT(1)"`
	CustomBrokerId         int       `xorm:"comment('报关行') INT(11)"`
	Contact                string    `xorm:"comment('联系人') VARCHAR(64)"`
	ContactPhone           string    `xorm:"comment('联系人手机号') VARCHAR(32)"`
	ContractNo             string    `xorm:"comment('合同编号') VARCHAR(64)"`
	IsFumigation           int       `xorm:"default 0 comment('是否需要熏蒸') TINYINT(1)"`
	FumigationId           int       `xorm:"comment('熏蒸公司') INT(11)"`
	IsCertificate          int       `xorm:"default 0 comment('是否需要产地证') TINYINT(1)"`
	TraderId               int       `xorm:"comment('贸易商') INT(11)"`
	IsInsurance            int       `xorm:"default 0 comment('是否需要保险') TINYINT(1)"`
	InsuranceId            int       `xorm:"comment('保险公司') INT(11)"`
	Beneficiary            string    `xorm:"comment('受益人') VARCHAR(64)"`
	InsuranceAmount        string    `xorm:"comment('保险金额') DECIMAL(19,4)"`
	IsMagneticTest         int       `xorm:"default 0 comment('是否需要磁检') TINYINT(1)"`
	MagneticTestId         int       `xorm:"comment('磁检机构') INT(11)"`
	IsIdentification       int       `xorm:"default 0 comment('是否需要鉴定') TINYINT(1)"`
	IdentificationId       int       `xorm:"comment('鉴定机构') INT(11)"`
	IsWarehouse            int       `xorm:"default 0 comment('是否需要仓库/场装') TINYINT(1)"`
	WarehouseId            int       `xorm:"comment('仓库/场装') INT(11)"`
	DeletedAt              time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
	LockVersion            int       `xorm:"default 0 INT(11)"`
}
