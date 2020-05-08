package models

import "time"

type Account struct {
	ID                uint        `gorm:"primary_key" json:"id"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	DeletedAt         *time.Time  `sql:"index"`
	BankName          string      `gorm:"varchar(256);index:bank_name" json:"bank_name"`
	Name              string      `gorm:"size:64;index:name" json:"name" validate:"required"`
	BankNumber        string      `gorm:"type:varchar(256);index:bank_number" json:"bank_number" validate:"required"`
	BankAddress       string      `gorm:"type:varchar(1024)" json:"bank_address"`
	UserName          string      `gorm:"size:64;index:user_name" json:"user_name" validate:"required"`
	UserAddress       string      `gorm:"type:varchar(1024)" json:"user_address"`
	Location          string      `gorm:"type:varchar(1024)" json:"location"`
	SwiftCode         string      `gorm:"type:varchar(256);" json:"swift_code"`
	TaxRegisterNumber string      `gorm:"type:varchar(256);unique_index" json:"tax_register_number" validate:"required"`
	Category          uint      `gorm:"size:64;index:category" json:"category" validate:"required"`
	UserCompanyId     uint        `gorm:"index" json:"user_company_id" validate:"required"`
	UserCompany       UserCompany `gorm:"foreignkey:user_company_id" validate:"structonly"`
}


type ResultAccount struct {
	ID                  uint      `json:"id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	BankName            string    `json:"bank_name"`
	Name                string    `json:"name"`
	BankNumber          string    `json:"bank_number"`
	BankAddress         string    `json:"bank_address"`
	UserName            string    `json:"user_name"`
	UserAddress         string    `json:"user_address"`
	Location            string    `json:"location"`
	SwiftCode           string    `json:"swift_code"`
	TaxRegisterNumber   string    `json:"tax_register_number"`
	Category            uint    `json:"category"`
	UserCompanyId       uint      `json:"user_company_id"`
	UserCompaniesNameNick string    `json:"user_companies_name_nick"`
}

func (Account) TableName() string {
	return "accounts"
}
func (ResultAccount) TableName() string {
	return "accounts"
}

