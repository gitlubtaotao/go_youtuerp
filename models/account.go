package models

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID                uint        `gorm:"primary_key" json:"id"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	DeletedAt         gorm.DeletedAt  `sql:"index"`
	BankName          string      `gorm:"varchar(256);index:bank_name" json:"bank_name"`
	Name              string      `gorm:"size:64;index:name" json:"name" validate:"required"`
	BankNumber        string      `gorm:"type:varchar(256);index:bank_number" json:"bank_number" validate:"required"`
	BankAddress       string      `gorm:"type:varchar(1024)" json:"bank_address"`
	UserName          string      `gorm:"size:64;index:user_name" json:"user_name" validate:"required"`
	UserAddress       string      `gorm:"type:varchar(1024)" json:"user_address"`
	Location          string      `gorm:"type:varchar(1024)" json:"location"`
	SwiftCode         string      `gorm:"type:varchar(256);" json:"swift_code"`
	TaxRegisterNumber string      `gorm:"type:varchar(256);unique_index" json:"tax_register_number" validate:"required"`
	Category          uint        `gorm:"size:64;index:category" json:"category" validate:"required"`
	UserCompanyId     uint        `gorm:"index" json:"user_company_id" validate:"required"`
	UserCompany       UserCompany `gorm:"foreignkey:user_company_id" validate:"structonly"`
}

type Invoice struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `sql:"index"`
	Name          string     `gorm:"size:256;index:name" json:"name" validate:"required"`
	TaxNumber     string     `gorm:"size:64" json:"tax_number" validate:"required"`
	BankNumber    string     `gorm:"size:64" json:"bank_number" validate:"required"`
	BankName      string     `gorm:"size:64" json:"bank_name" validate:"required"`
	BankAddress   string     `gorm:"size:522" json:"bank_address"`
	BankTel       string     `gorm:"size:16" json:"bank_tel"`
	UserCompanyId uint       `gorm:"index:user_company_id" json:"user_company_id"`
	Category      uint       `gorm:"index:category" json:"category" validate:"required"`
}

type Address struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `sql:"index"`
	UserName      string     `gorm:"size:64;index:user_name" comment:"收件人姓名" json:"user_name" validate:"required"`
	UserTel       string     `gorm:"size:16" comment:"收件人电话" json:"user_tel"`
	UserAddress   string     `gorm:"size:522" comment:"邮件地址"  json:"user_address"`
	Code          string     `gorm:"size:16" comment:"邮编" json:"code"`
	Province      string     `gorm:"size:64" json:"province"`
	City          string     `gorm:"size:64" json:"city"`
	Distinct      string     `gorm:"size:64" json:"distinct"`
	UserCompanyId uint       `gorm:"index:user_company_id" json:"user_company_id"`
}

func (Account) TableName() string {
	return "accounts"
}
func (Invoice) TableName() string {
	return "invoices"
}

func (Address) TableName() string {
	return "address"
}
