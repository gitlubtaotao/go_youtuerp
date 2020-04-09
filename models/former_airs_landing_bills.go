package models

import (
	"time"
)

type FormerAirsLandingBills struct {
	Id                        int64     `xorm:"pk autoincr BIGINT(20)"`
	OrderMasterId             int64     `xorm:"index BIGINT(20)"`
	ShipperContent            string    `xorm:"comment('发货人信息') TEXT"`
	ConsigneeContent          string    `xorm:"comment('收货人信息') TEXT"`
	NotifyPartyContent        string    `xorm:"TEXT"`
	PodAgentContent           string    `xorm:"TEXT"`
	DepartureAirportId        int       `xorm:"index INT(11)"`
	DestinationAirportId      int       `xorm:"index INT(11)"`
	Flight                    string    `xorm:"comment('航班') VARCHAR(255)"`
	FlightDate                time.Time `xorm:"DATETIME"`
	CurrencyId                int64     `xorm:"comment('货币') index BIGINT(20)"`
	ChargableWeight           string    `xorm:"comment('计费重量') VARCHAR(255)"`
	ExecuteDate               time.Time `xorm:"comment('签发日期') DATETIME"`
	SerialNo                  string    `xorm:"comment('序列号') VARCHAR(255)"`
	SubBlNo                   string    `xorm:"comment('分提单号') VARCHAR(255)"`
	CreatedAt                 time.Time `xorm:"not null DATETIME"`
	UpdatedAt                 time.Time `xorm:"not null DATETIME"`
	ShipperType               string    `xorm:"index(shipper) VARCHAR(255)"`
	ShipperId                 int64     `xorm:"comment('发货人') index(shipper) BIGINT(20)"`
	ConsigneeType             string    `xorm:"index(consignee) VARCHAR(255)"`
	ConsigneeId               int64     `xorm:"comment('收货人') index(consignee) BIGINT(20)"`
	NotifyPartyType           string    `xorm:"index(notify_party) VARCHAR(255)"`
	NotifyPartyId             int64     `xorm:"comment('通知方') index(notify_party) BIGINT(20)"`
	PodAgentType              string    `xorm:"index(pod_agent) VARCHAR(255)"`
	PodAgentId                int64     `xorm:"comment('代理商') index(pod_agent) BIGINT(20)"`
	MainBlNo                  string    `xorm:"comment('主提单') VARCHAR(255)"`
	TransshipmentAirportId    int       `xorm:"index INT(11)"`
	CodeOfTwoId               int       `xorm:"index INT(11)"`
	OceanChangesPaytypeId     int       `xorm:"index INT(11)"`
	OtherChangesPaytypeId     int       `xorm:"index INT(11)"`
	Marks                     string    `xorm:"TEXT"`
	DescriptionOfGood         string    `xorm:"TEXT"`
	GrossWeight               string    `xorm:"TEXT"`
	Number                    string    `xorm:"TEXT"`
	Size                      string    `xorm:"TEXT"`
	Status                    string    `xorm:"VARCHAR(255)"`
	Remarks                   string    `xorm:"TEXT"`
	LandingBillType           string    `xorm:"VARCHAR(255)"`
	BaseDataPackageType       int       `xorm:"INT(11)"`
	CompanyInstructionId      int       `xorm:"INT(11)"`
	CompanyInstructionType    string    `xorm:"VARCHAR(255)"`
	CompanyInstructionContent string    `xorm:"TEXT"`
	DeliveryAddress           string    `xorm:"TEXT"`
	ArriveAddress             string    `xorm:"TEXT"`
	EstimatedTimeOfDeparture  time.Time `xorm:"DATETIME"`
	Ratio                     float32   `xorm:"comment('材积换算系数') FLOAT"`
	RatioWeight               float32   `xorm:"comment('材积重') FLOAT"`
	Bubble                    float32   `xorm:"comment('分泡%') FLOAT"`
	Dimension                 string    `xorm:"comment('体积') TEXT"`
	IsGoodsAttachment         int       `xorm:"default 0 comment('导出货物描述附件') TINYINT(1)"`
	BaseDataTradeTermId       int       `xorm:"INT(11)"`
	BaseDataItemId            int       `xorm:"INT(11)"`
	BookingRemarks            string    `xorm:"TEXT"`
	CostDescription           string    `xorm:"TEXT"`
	SupplierCompanyAgentId    int       `xorm:"INT(11)"`
	RandomFile                int       `xorm:"TINYINT(1)"`
	DeletedAt                 time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
