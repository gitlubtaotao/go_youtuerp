package models

import "time"

type FinanceRate struct {
	ID                uint      `gorm:"primary_key"json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Year              uint      `gorm:"index:year_start_and_end_month" json:"year"`
	StartMonth        uint      `gorm:"year_start_and_end_month" json:"start_month"`
	EndMonth          uint      `gorm:"year_start_and_end_month" json:"end_month"`
	Rate              float64   `gorm:"default:0.0;PRECISION:4" json:"rate"`
	FinanceCurrencyId uint      `gorm:"index:finance_currency_id" json:"finance_currency_id" validate:"required"`
	UserId            uint      `gorm:"index:user_id" json:"user_id" validate:"required"`
	CompanyId         uint      `gorm:"index:company_id" json:"company_id"`
}

type FinanceFeeType struct {
	ID                uint       `gorm:"primary_key"json:"id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `sql:"index"`
	Name              string     `gorm:"size:64;index:name" json:"name"`
	NameCn            string     `gorm:"index:name_cn" json:"name_cn"`
	NameEn            string     `json:"name_en"`
	FinanceCurrencyId uint       `gorm:"index:finance_currency_id" json:"finance_currency_id"`
	Remarks           string     `gorm:"size:256" json:"remarks"`
}

func (FinanceRate) TableName() string {
	return "finance_rates"
}

func (FinanceFeeType) TableName() string {
	return "finance_fee_types"
}
