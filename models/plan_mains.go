package models

import (
	"time"
)

type PlanMains struct {
	Id                     int64     `xorm:"pk autoincr BIGINT(20)"`
	OrderMasterId          int64     `xorm:"index BIGINT(20)"`
	CreatedAt              time.Time `xorm:"not null DATETIME"`
	UpdatedAt              time.Time `xorm:"not null DATETIME"`
	Enabled                int       `xorm:"default 1 comment('有效状态') TINYINT(1)"`
	UserOperatorId         int       `xorm:"comment('操作员') INT(11)"`
	UserSalesmanId         int       `xorm:"comment('业务员') INT(11)"`
	UserFileId             int       `xorm:"comment('文件') INT(11)"`
	BusinessmanId          int       `xorm:"comment('商务') INT(11)"`
	UserFeeId              int       `xorm:"comment('财务') INT(11)"`
	UserCustomerId         int       `xorm:"comment('客服') INT(11)"`
	UserAuditId            int       `xorm:"comment('价格审核') INT(11)"`
	Profit                 string    `xorm:"comment('利润') DECIMAL(19,4)"`
	PolId                  int       `xorm:"comment('起运港') INT(11)"`
	PodId                  int       `xorm:"comment('目的港') INT(11)"`
	GoodCount              int       `xorm:"comment('包装件数') INT(11)"`
	PackageTypeId          int       `xorm:"comment('包装类型') INT(11)"`
	GrossWeight            string    `xorm:"comment('毛重') DECIMAL(19,4)"`
	GoodSize               string    `xorm:"comment('货物体积') DECIMAL(19,4)"`
	CompanyId              int       `xorm:"comment('船公司/航空公司/运输公司') INT(11)"`
	SupplierCompanyAgentId int       `xorm:"INT(11)"`
	MawbNo                 string    `xorm:"comment('mawb编号') VARCHAR(64)"`
	CutOffDate             time.Time `xorm:"comment('截关/开船日期') DATETIME"`
	DepartureDate          time.Time `xorm:"comment('离港日期/起飞时间') DATETIME"`
	ArrivalDate            time.Time `xorm:"comment('到达时间') DATETIME"`
	BillProduceId          int       `xorm:"comment('出单方式') INT(11)"`
	Note                   string    `xorm:"comment('备注') TEXT"`
	InlandDriverPhone      string    `xorm:"comment('内陆司机手机号') VARCHAR(32)"`
	InlandCarNumber        string    `xorm:"comment('内陆车牌') VARCHAR(32)"`
	HongkongDriverPhone    string    `xorm:"comment('香港司机手机号') VARCHAR(32)"`
	HongkongCarNumber      string    `xorm:"comment('香港车牌') VARCHAR(32)"`
	Status                 string    `xorm:"comment('状态') VARCHAR(32)"`
	PodAgentId             int       `xorm:"comment('目的港代理') INT(11)"`
	PodAgentType           string    `xorm:"comment('目的港代理类型') VARCHAR(64)"`
	PodContact             string    `xorm:"comment('目的港联系人') VARCHAR(32)"`
	PodEmail               string    `xorm:"comment('目的港联系人邮箱') VARCHAR(64)"`
	CourierNumber          string    `xorm:"comment('快递单号') VARCHAR(64)"`
	DeletedAt              time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
	LockVersion            int       `xorm:"default 0 INT(11)"`
	TransshipmentId        int       `xorm:"INT(11)"`
}
