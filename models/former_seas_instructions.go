package models

import (
	"time"
)

type FormerSeasInstructions struct {
	Id                        int64     `xorm:"pk autoincr BIGINT(20)"`
	BaseDataInstructionTypeId int       `xorm:"index INT(11)"`
	SerialNo                  string    `xorm:"comment('序列号/流水号') VARCHAR(255)"`
	CompanyInstructionType    string    `xorm:"index(index_company_instruction) VARCHAR(255)"`
	CompanyInstructionId      int       `xorm:"index(index_company_instruction) INT(11)"`
	CompanyInstructionContent string    `xorm:"comment('委托公司详细要求') TEXT"`
	ContactInstructionType    string    `xorm:"index(index_contact_instruction) VARCHAR(255)"`
	ContactInstructionId      int       `xorm:"index(index_contact_instruction) INT(11)"`
	ContactInstructionContent string    `xorm:"comment('委托公司联系人详细要求内容') TEXT"`
	ShipperType               string    `xorm:"index(index_shipper) VARCHAR(255)"`
	ShipperId                 int       `xorm:"index(index_shipper) INT(11)"`
	ShipperContent            string    `xorm:"comment('发货人要求') TEXT"`
	ConsigneeType             string    `xorm:"index(index_consignee) VARCHAR(255)"`
	ConsigneeId               int       `xorm:"index(index_consignee) INT(11)"`
	ConsigneeContent          string    `xorm:"comment('收货人要求') TEXT"`
	NotifyPartyType           string    `xorm:"index(index_notify_party) VARCHAR(255)"`
	NotifyPartyId             int       `xorm:"index(index_notify_party) INT(11)"`
	NotifyPartyContent        string    `xorm:"comment('通知内容') TEXT"`
	PodAgentType              string    `xorm:"index(index_pod_agent) VARCHAR(255)"`
	PodAgentId                int       `xorm:"index(index_pod_agent) INT(11)"`
	PodAgentContent           string    `xorm:"comment('目的港代理详情') TEXT"`
	Marks                     string    `xorm:"comment('标记') TEXT"`
	Number                    string    `xorm:"comment('包装件数') TEXT"`
	DescriptionOfGood         string    `xorm:"comment('品名') TEXT"`
	GrossWeight               string    `xorm:"comment('毛重') TEXT"`
	Size                      string    `xorm:"comment('体积') TEXT"`
	InvoiceNo                 string    `xorm:"comment('发票号') TEXT"`
	SeaPortPolId              int       `xorm:"comment('装货港') index INT(11)"`
	SeaPortViaId              int       `xorm:"comment('中转港') index INT(11)"`
	SeaPortPodId              int       `xorm:"comment('目的港') index INT(11)"`
	TransshipmentId           int       `xorm:"index INT(11)"`
	ServiceContractNo         string    `xorm:"comment('合同编号') VARCHAR(255)"`
	BaseDataItemId            int       `xorm:"index INT(11)"`
	BaseDataTradeTermsId      int       `xorm:"index INT(11)"`
	WaysOfDeclarationId       int       `xorm:"index INT(11)"`
	WaysOfTransportationId    int       `xorm:"index INT(11)"`
	Remarks                   string    `xorm:"comment('备注') TEXT"`
	CreatedAt                 time.Time `xorm:"not null DATETIME"`
	UpdatedAt                 time.Time `xorm:"not null DATETIME"`
	OrderMasterId             int       `xorm:"index INT(11)"`
	LockVersion               int       `xorm:"default 0 INT(11)"`
	Status                    string    `xorm:"index VARCHAR(255)"`
	DeletedAt                 time.Time `xorm:"index DATETIME"`
	EstimatedTimeOfDeparture  time.Time `xorm:"comment('开船日期') DATETIME"`
	EstimatedTimeOfArrival    time.Time `xorm:"comment('到港日期') DATETIME"`
	CutOffDate                time.Time `xorm:"comment('截关日期') DATETIME"`
	BoatCompanyId             int       `xorm:"comment('船公司') index INT(11)"`
	Vessel                    string    `xorm:"comment('船名') VARCHAR(255)"`
	Voyage                    string    `xorm:"comment('航次') VARCHAR(255)"`
	CargoReceivedDate         time.Time `xorm:"DATETIME"`
	BaseDataPackageType       int       `xorm:"INT(11)"`
	BoxSizeCount              string    `xorm:"TEXT"`
	OceanChangesPaytypeId     int       `xorm:"comment('运费支付方式') INT(11)"`
	OtherChangesPaytypeId     int       `xorm:"comment('其他运费支付方式') INT(11)"`
	BaseDataMiscBillId        int       `xorm:"INT(11)"`
	Dimension                 string    `xorm:"comment('尺寸') TEXT"`
	Type                      string    `xorm:"default 'Former::Seas::Instruction' VARCHAR(255)"`
	BLNo                      string    `xorm:"comment('分单号') VARCHAR(255)"`
	HblReleaseDate            time.Time `xorm:"comment('分单释放时间') DATETIME"`
	IsCabinetAttachment       int       `xorm:"default 0 TINYINT(1)"`
	IsGoodsAttachment         int       `xorm:"default 0 TINYINT(1)"`
	HblRemark                 string    `xorm:"comment('HB/L备注') TEXT"`
	PlaceOfReceipt            string    `xorm:"TEXT"`
	PlaceOfDelivery           string    `xorm:"TEXT"`
	ShippedOnBoardDate        time.Time `xorm:"DATETIME"`
	VerifyDate                time.Time `xorm:"DATETIME"`
	UserVerifyId              int       `xorm:"INT(11)"`
	TlxNo                     string    `xorm:"VARCHAR(255)"`
	PlaceOfIssue              string    `xorm:"VARCHAR(255)"`
	DateOfIssue               time.Time `xorm:"DATETIME"`
}
