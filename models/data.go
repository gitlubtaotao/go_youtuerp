package models

import "time"

type BaseDataLevel struct {
	ID     uint   `gorm:"primary_key"json:"id"`
	Code   string `gorm:"size:64" json:"code"`
	Name   string `gorm:"size:64" json:"name"`
	EnName string `gorm:"size:64" json:"en_name"`
}

type BaseDataCode struct {
	ID        uint       `gorm:"primary_key"json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
	CodeName  string     `gorm:"size:64;index:code_name" json:"code_name" validate:"required"`
	Name      string     `gorm:"size:128;index:name;" json:"name" validate:"required"`
	Remarks   string     `gorm:"size:522;" json:"remarks"`
}

type BaseDataPort struct {
	ID        uint       `gorm:"primary_key"json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
	Name      string     `gorm:"size:64;index:name" json:"name"`
	NameCn    string     `gorm:"size:128;" json:"name_cn"`
	NameEn    string     `gorm:"size:128;" json:"name_en"`
	Country   string     `gorm:"size:128" json:"country"`
	Region    string     `gorm:"size:128" json:"region"`
	City      string     `gorm:"size:128" json:"city"`
	Type      uint       `gorm:"index:type" json:"type"`
	Remarks   string     `gorm:"size:522;" json:"remarks"`
}

type BaseDataCarrier struct {
	ID        uint       `gorm:"primary_key"json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
	Name      string     `gorm:"size:64;index:name" json:"name" validate:"required"`
	NameCn    string     `gorm:"size:128;" json:"name_cn"`
	NameEn    string     `gorm:"size:128;" json:"name_en"`
	Url       string     `gorm:"size:128" json:"url"`
	Type      uint       `gorm:"index:type;" json:"type"`
	Remarks   string     `gorm:"size:522;" json:"remarks"`
}

type BaseWarehouse struct {
	ID             uint       `gorm:"primary_key"json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `sql:"index"`
	Name           string     `gorm:"size:64;index:name" json:"name"`
	Detail         string     `gorm:"size:256;" json:"detail"`
	ContactName    string     `gorm:"size:64;index:contact_name_and_contact_tel" json:"contact_name"`
	ContactTel     string     `gorm:"size:16;index:contact_name_and_contact_tel" json:"contact_tel"`
	ContactAddress string     `gorm:"size:522" json:"contact_address"`
	Remarks        string     `gorm:"size:522" json:"remarks"`
	Region         string     `gorm:"size:256" json:"region"`
}

const (
	BaseTypeSea = iota + 1
	BaseTypeAir
	BaseTypeCourier
)

const (
	CodeFinanceCurrency = "FinanceCurrency"
	CodePayType         = "PayType"
	CodeCapType         = "CapType"
	CodeInstructionType = "InstructionType"
	CodeCustomType      = "CustomType"
	CodeBillProduceType = "BillProduceType"
	CodeTransshipment   = "Transshipment"
	CodeTradeTerms      = "TradeTerms"
	CodeShippingTerms   = "ShippingTerms"
)


func (r BaseDataCode) TableName() string {
	return "base_data_codes"
}

func (BaseDataLevel) TableName() string {
	return "base_data_levels"
}

func (BaseDataCarrier) TableName() string {
	return "base_data_carriers"
}

func (BaseDataPort) TableName() string {
	return "base_data_ports"
}

func (BaseWarehouse) TableName() string {
	return "base_warehouses"
}
