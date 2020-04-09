package models

import (
	"time"
)

type FormerSeasAms struct {
	Id                        int64     `xorm:"pk autoincr BIGINT(20)"`
	ShipperId                 int       `xorm:"comment('承运人，从cooperator中选择') index(index_shipper_type_and_shipper_id) INT(11)"`
	SoNo                      string    `xorm:"comment('so编号') VARCHAR(255)"`
	ConsigneeId               int       `xorm:"comment('收货人，从cooperator选择') index(index_consignee_type_and_consignee_id) INT(11)"`
	NotifyPartyId             int       `xorm:"index(index_notify_party_type_and_notify_party_id) INT(11)"`
	PrecarriageBy             string    `xorm:"comment('头程运输') VARCHAR(255)"`
	PlaceOfReceipt            string    `xorm:"comment('接货地') TEXT"`
	OceanVessel               string    `xorm:"comment('航线') VARCHAR(255)"`
	SeaPortPodId              int       `xorm:"index INT(11)"`
	SeaPortPolId              int       `xorm:"index INT(11)"`
	PlaceOfDelivery           string    `xorm:"comment('送货地') TEXT"`
	ShipmentTerm              string    `xorm:"comment('运输条款') VARCHAR(255)"`
	Marks                     string    `xorm:"comment('唛头') TEXT"`
	Number                    string    `xorm:"TEXT"`
	DescriptionOfGood         string    `xorm:"comment('货物描述') TEXT"`
	GrossWeight               string    `xorm:"comment('总重，单位kg') TEXT"`
	Size                      string    `xorm:"TEXT"`
	CreatedAt                 time.Time `xorm:"not null DATETIME"`
	UpdatedAt                 time.Time `xorm:"not null DATETIME"`
	OrderMasterId             int64     `xorm:"index BIGINT(20)"`
	Status                    string    `xorm:"VARCHAR(255)"`
	InvoiceNo                 string    `xorm:"comment('发票号') VARCHAR(255)"`
	SerialNo                  string    `xorm:"VARCHAR(255)"`
	ShipperType               string    `xorm:"index(index_shipper_type_and_shipper_id) VARCHAR(255)"`
	ShipperContent            string    `xorm:"TEXT"`
	ConsigneeType             string    `xorm:"index(index_consignee_type_and_consignee_id) VARCHAR(255)"`
	ConsigneeContent          string    `xorm:"TEXT"`
	NotifyPartyType           string    `xorm:"index(index_notify_party_type_and_notify_party_id) VARCHAR(255)"`
	NotifyPartyContent        string    `xorm:"TEXT"`
	BaseDataPackageType       int       `xorm:"INT(11)"`
	CompanyInstructionId      int       `xorm:"INT(11)"`
	CompanyInstructionType    string    `xorm:"VARCHAR(255)"`
	CompanyInstructionContent string    `xorm:"TEXT"`
	Vessel                    string    `xorm:"VARCHAR(255)"`
	Voyage                    string    `xorm:"VARCHAR(255)"`
	PodAgentContent           string    `xorm:"TEXT"`
	PodAgentId                string    `xorm:"TEXT"`
	PodAgentType              string    `xorm:"TEXT"`
	ShippedOnBoardDate        time.Time `xorm:"DATETIME"`
	Dimension                 string    `xorm:"TEXT"`
}
