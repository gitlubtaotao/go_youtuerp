package models

import (
	"time"
)

type BargainMains struct {
	Id                      int64     `xorm:"pk autoincr BIGINT(20)"`
	MainNo                  string    `xorm:"default '0000' comment('价单编号') VARCHAR(255)"`
	BargainType             int       `xorm:"default 1 comment('价单类型') INT(11)"`
	TransportType           int       `xorm:"default 1 comment('运输类型') index INT(11)"`
	Status                  string    `xorm:"default 'init' comment('价单状态') index VARCHAR(64)"`
	CompanyInstructionId    int       `xorm:"default 0 comment('客户') index INT(11)"`
	CompanyContactName      string    `xorm:"comment('联系人') VARCHAR(255)"`
	CompanyContactPhone     string    `xorm:"comment('联系人电话') VARCHAR(64)"`
	UserOperationId         int       `xorm:"default 0 comment('业务人员') index INT(11)"`
	ExpectedShipmentTime    time.Time `xorm:"default '2019-05-26 01:22:51' comment('预计出货时间') DATETIME"`
	OrderMasterSerialNumber string    `xorm:"default '0' comment('订单编号') VARCHAR(128)"`
	OrderMasterId           int       `xorm:"default 0 comment('订单') index INT(11)"`
	Remarks                 string    `xorm:"default '0' comment('价单备注') VARCHAR(255)"`
	SeaLineId               int       `xorm:"comment('航线') index INT(11)"`
	CutOffDay               time.Time `xorm:"comment('截关/截港日期') DATETIME"`
	FlightDate              time.Time `xorm:"comment('航班日期') DATETIME"`
	BoatCompanyId           int       `xorm:"default 0 comment('船公司') index INT(11)"`
	SeaPolId                int       `xorm:"default 0 comment('起运港') index(sea_port_index) INT(11)"`
	SeaPodId                int       `xorm:"default 0 comment('目的港') index(sea_port_index) INT(11)"`
	GrossWeight             string    `xorm:"default '0' comment('毛重') VARCHAR(255)"`
	Size                    string    `xorm:"default '0' comment('体积') VARCHAR(255)"`
	AirPolId                int       `xorm:"default 0 comment('航空起运港') index(air_port_index) INT(11)"`
	AirPodId                int       `xorm:"default 0 comment('航空目的港') index(air_port_index) INT(11)"`
	CodeOfTwoId             int       `xorm:"default 0 comment('航空公司') index INT(11)"`
	SeaFreightId            int       `xorm:"default 0 comment('对应的海运费') INT(11)"`
	FinishedTime            time.Time `xorm:"comment('成交时间') DATETIME"`
	UserCompanyId           int       `xorm:"default 0 index INT(11)"`
	Platform                string    `xorm:"default 'offline' comment('线下') VARCHAR(64)"`
	CreatedAt               time.Time `xorm:"not null DATETIME"`
	UpdatedAt               time.Time `xorm:"not null DATETIME"`
	DeletedAt               time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	CityId                  int       `xorm:"INT(11)"`
	UserId                  int       `xorm:"comment('创建询价单用户id') INT(11)"`
	CompanyType             string    `xorm:"default 'cooperator' comment('客户类型:默认为合作单位') VARCHAR(32)"`
	CompanyName             string    `xorm:"comment('公司名称') VARCHAR(255)"`
	PlaceOfReceipt          string    `xorm:"comment('提货地址') TEXT"`
	NotifyPartyAddress      string    `xorm:"comment('交货地') TEXT"`
	BaseDataTradeTermsId    int       `xorm:"comment('贸易条款') INT(11)"`
	BaseDataItemId          int       `xorm:"comment('装运条款') INT(11)"`
	IsSaveClue              int       `xorm:"default 0 comment('客户信息是否保存为线索') TINYINT(1)"`
	LockVersion             int       `xorm:"default 0 INT(11)"`
	BargainRemarks          string    `xorm:"TEXT"`
}
