package models

import (
	"time"
)

type FormerWarehouseOrders struct {
	Id                        int64     `xorm:"pk autoincr BIGINT(20)"`
	OrderMasterId             int       `xorm:"comment('订单') INT(11)"`
	SerialNo                  string    `xorm:"comment('流水号') VARCHAR(255)"`
	CompanyInstructionType    string    `xorm:"comment('委托单位') VARCHAR(255)"`
	CompanyInstructionId      int       `xorm:"comment('委托单位') INT(11)"`
	CompanyInstructionContent string    `xorm:"TEXT"`
	WarehouseAddress          string    `xorm:"comment('仓库地址') TEXT"`
	DeliveryAddress           string    `xorm:"comment('送货说明') TEXT"`
	DistributionReason        string    `xorm:"comment('配货原因') TEXT"`
	Note                      string    `xorm:"comment('备注') TEXT"`
	ClientNote                string    `xorm:"comment('客户备注') TEXT"`
	WarehouseNo               string    `xorm:"comment('入仓单号') VARCHAR(128)"`
	WarehouseTime             time.Time `xorm:"comment('入仓时间') DATETIME"`
	DeletedAt                 time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
