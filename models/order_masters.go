package models

import (
	"time"
)

type OrderMasters struct {
	Platform               string    `xorm:"VARCHAR(64)"`
	Id                     int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt              time.Time `xorm:"not null DATETIME"`
	UpdatedAt              time.Time `xorm:"not null DATETIME"`
	SerialNumber           string    `xorm:"comment('流水号') index VARCHAR(255)"`
	UserCompanyId          int       `xorm:"INT(11)"`
	UserSalesmanId         int       `xorm:"index INT(11)"`
	UserOperatorId         int       `xorm:"index INT(11)"`
	Remark                 string    `xorm:"comment('描述') TEXT"`
	DeletedAt              time.Time `xorm:"index DATETIME"`
	RawInfoJson            string    `xorm:"comment('原始信息') LONGTEXT"`
	LockVersion            int       `xorm:"default 0 INT(11)"`
	AasmState              string    `xorm:"comment('aasm状态机') index VARCHAR(255)"`
	TransportType          int       `xorm:"default 1 index INT(11)"`
	CompanyId              int       `xorm:"default 0 index INT(11)"`
	PayStatus              string    `xorm:"index VARCHAR(255)"`
	ReceiveStatus          string    `xorm:"index VARCHAR(255)"`
	BookingDate            time.Time `xorm:"comment('和订舱日期保持一致，用于搜索的冗余字段') DATETIME"`
	ReceivableLockDate     time.Time `xorm:"comment('费用结清日期') DATETIME"`
	PayableStatus          string    `xorm:"VARCHAR(255)"`
	ReceiveableStatus      string    `xorm:"VARCHAR(255)"`
	SupplierCompanyAgentId int       `xorm:"INT(11)"`
	CreditNoteStatus       string    `xorm:"default 'unfinished' comment('应付对账状态') VARCHAR(15)"`
	DebitNoteStatus        string    `xorm:"default 'unfinished' comment('应收对账状态') VARCHAR(15)"`
	UserFileId             int       `xorm:"INT(11)"`
	UserMarketId           int       `xorm:"INT(11)"`
	UserCustomerId         int       `xorm:"INT(11)"`
	UserFeeId              int       `xorm:"INT(11)"`
	UserContactId          int       `xorm:"comment('订单委托单位的联系人') INT(11)"`
	MainTransport          int       `xorm:"default 1 comment('主要运输方式') INT(11)"`
}
