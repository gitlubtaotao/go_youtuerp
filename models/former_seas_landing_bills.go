package models

import (
	"time"
)

type FormerSeasLandingBills struct {
	Id                        int64     `xorm:"pk autoincr BIGINT(20)"`
	BLNo                      string    `xorm:"comment('提单号') VARCHAR(255)"`
	VerifyDate                time.Time `xorm:"comment('校对日期') DATETIME"`
	UserVerifyId              int       `xorm:"index INT(11)"`
	ShipperType               string    `xorm:"index(shipper) VARCHAR(255)"`
	ShipperId                 int64     `xorm:"index(shipper) BIGINT(20)"`
	ConsigneeType             string    `xorm:"index(consignee) VARCHAR(255)"`
	ConsigneeId               int64     `xorm:"index(consignee) BIGINT(20)"`
	NotifyPartyType           string    `xorm:"index(notify_party) VARCHAR(255)"`
	NotifyPartyId             int64     `xorm:"index(notify_party) BIGINT(20)"`
	PodAgentType              string    `xorm:"index(pod_agent) VARCHAR(255)"`
	PodAgentId                int64     `xorm:"index(pod_agent) BIGINT(20)"`
	AmsActualShipperId        int64     `xorm:"index BIGINT(20)"`
	AmsActualConsigneeId      int64     `xorm:"index BIGINT(20)"`
	BoatCompanyId             int       `xorm:"index INT(11)"`
	Vessel                    string    `xorm:"comment('船名') VARCHAR(255)"`
	Voyage                    string    `xorm:"comment('航次') VARCHAR(255)"`
	SeaPortPodId              int       `xorm:"index INT(11)"`
	SeaPortPolId              int       `xorm:"index INT(11)"`
	SeaPortViaId              int       `xorm:"index INT(11)"`
	PlaceOfDelivery           string    `xorm:"comment('目的地') TEXT"`
	PlaceOfReceipt            string    `xorm:"comment('收货地址') TEXT"`
	PreCarriageById           int       `xorm:"index INT(11)"`
	EstimatedTimeOfDeparture  time.Time `xorm:"DATETIME"`
	EstimatedTimeOfArrival    time.Time `xorm:"DATETIME"`
	CutOffDate                time.Time `xorm:"comment('截关日期') DATETIME"`
	CargoReceivedDate         time.Time `xorm:"comment('收货人提货时间') DATETIME"`
	FreightPayableAt          string    `xorm:"comment('运费支付地') TEXT"`
	PlaceOfIssue              string    `xorm:"comment('签单地点') TEXT"`
	DateOfIssue               time.Time `xorm:"comment('签单日期') DATETIME"`
	BaseDataTradeTermsId      int       `xorm:"index INT(11)"`
	BaseDataMiscBillId        int       `xorm:"index INT(11)"`
	TlxNo                     string    `xorm:"comment('电放单号') VARCHAR(255)"`
	MblReleaseDate            time.Time `xorm:"DATETIME"`
	MblNo                     string    `xorm:"comment('船东提单号') VARCHAR(255)"`
	Number                    string    `xorm:"TEXT"`
	BaseDataPackageType       int       `xorm:"comment('包装类型') INT(11)"`
	GrossWeight               string    `xorm:"comment('毛重') TEXT"`
	Size                      string    `xorm:"TEXT"`
	DescriptionOfGood         string    `xorm:"comment('货物英文名称') TEXT"`
	Marks                     string    `xorm:"comment('标记唛头') TEXT"`
	Remarks                   string    `xorm:"comment('提单备注') TEXT"`
	ShippingInstruction       int       `xorm:"comment('补料,简称SI') TINYINT(1)"`
	CreatedAt                 time.Time `xorm:"not null DATETIME"`
	UpdatedAt                 time.Time `xorm:"not null DATETIME"`
	OrderMasterId             int64     `xorm:"index BIGINT(20)"`
	DeletedAt                 time.Time `xorm:"index DATETIME"`
	LockVersion               int       `xorm:"INT(11)"`
	SerialNo                  string    `xorm:"comment('序列号') VARCHAR(255)"`
	Status                    string    `xorm:"comment('主状态') VARCHAR(255)"`
	ShipperContent            string    `xorm:"comment('发货人详细') TEXT"`
	ConsigneeContent          string    `xorm:"comment('收货人详情') TEXT"`
	NotifyPartyContent        string    `xorm:"comment('通知人详情') TEXT"`
	PodAgentContent           string    `xorm:"TEXT"`
	LandingBillType           string    `xorm:"comment('分类') VARCHAR(255)"`
	BillOfLandingId           int       `xorm:"index INT(11)"`
	OceanChangesPaytypeId     int       `xorm:"index INT(11)"`
	OtherChangesPaytypeId     int       `xorm:"index INT(11)"`
	ShippedOnBoardDate        time.Time `xorm:"DATETIME"`
	PreCarriageBy             string    `xorm:"VARCHAR(255)"`
	CompanyInstructionId      int       `xorm:"INT(11)"`
	CompanyInstructionType    string    `xorm:"VARCHAR(255)"`
	CompanyInstructionContent string    `xorm:"TEXT"`
	PayPodId                  int       `xorm:"INT(11)"`
	Dimension                 string    `xorm:"comment('尺寸') TEXT"`
	IsGoodsAttachment         int       `xorm:"default 0 comment('导出货物描述附件') TINYINT(1)"`
	IsCabinetAttachment       int       `xorm:"default 0 comment('导出分柜附件') TINYINT(1)"`
	BaseDataItemId            int       `xorm:"INT(11)"`
	BookingRemark             string    `xorm:"comment('订舱备注') TEXT"`
	ChargeDescription         string    `xorm:"comment('费用描述') TEXT"`
	BoxSizeCount              string    `xorm:"comment('柜兴柜量') TEXT"`
	SupplierCompanyAgentId    int       `xorm:"INT(11)"`
}
