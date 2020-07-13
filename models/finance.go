package models

import (
	"gorm.io/gorm"
	"time"
)

//订单费用信息
type FinanceFee struct {
	ID                  uint       `gorm:"primary_key"json:"id"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `sql:"index"`
	Name                string     `gorm:"size:64;index:name" json:"name" validate:"required"`
	NameCn              string     `json:"name_cn"`
	NameEn              string     `json:"name_en"`
	PayOrReceive        string     `gorm:"size:16;index:pay_or_receive" comment:"类型" json:"pay_or_receive" validate:"required"`
	PayTypeId           uint       `gorm:"index:pay_type_id" comment:"结算方式" json:"pay_type_id"`
	FinanceCurrencyId   uint       `gorm:"index:finance_currency_id" comment:"币种类型" json:"finance_currency_id" validate:"required"`
	FinanceCurrencyRate float64    `gorm:"PRECISION: 4" json:"finance_currency_rate" validate:"required"`
	Quantity            float64    `gorm:"PRECISION: 4" json:"quantity" validate:"required"`
	UnitPrice           float64    `gorm:"PRECISION:4" json:"unit_price" validate:"required"`
	TaxRate             float64    `gorm:"PRECISION:4" json:"tax_rate"`
	TaxAmount           float64    `gorm:"PRECISION:4" json:"tax_amount" validate:"required"`
	NotTaxAmount        float64    `gorm:"PRECISION:4" json:"not_tax_amount"`
	ReceiveAmount       float64    `gorm:"PRECISION:4" json:"receive_amount"`
	Receivable          float64    `gorm:"PRECISION:4" json:"receivable"`
	PayAmount           float64    `gorm:"PRECISION:4" json:"pay_amount"`
	Payable             float64    `gorm:"PRECISION:4" json:"payable"`
	ClosingUnitId       uint       `gorm:"index:closing_unit_id" json:"closing_unit_id"`
	OrderMasterId       uint       `gorm:"index:order_master_id" json:"order_master_id"`
	FinanceStatementId  uint       `json:"finance_statement_id"`
	FinanceStatementNo  string     `gorm:"size:32;index:finance_statement_no" json:"finance_statement_no"`
	Status              string     `gorm:"size:16;index:status;default:'init'" json:"status"`
	InvoiceAmount       float64    `gorm:"PRECISION:4" json:"invoice_amount"`
	InvoiceStatus       string     `gorm:"size:16;index:invoice_status;default:'init'" json:"invoice_status"`
	TypeId              uint       `gorm:"index:type_id;default:1" json:"type_id"`
	Remarks             string     `gorm:"size:128" json:"remarks"`
}

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
	DeletedAt         gorm.DeletedAt `sql:"index"`
	Name              string     `gorm:"size:64;index:name" json:"name"`
	NameCn            string     `gorm:"index:name_cn" json:"name_cn"`
	NameEn            string     `json:"name_en"`
	FinanceCurrencyId uint       `gorm:"index:finance_currency_id" json:"finance_currency_id"`
	Remarks           string     `gorm:"size:256" json:"remarks"`
}

type ResultFinanceFee struct {
	ID                  uint       `gorm:"primary_key"json:"id"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	DeletedAt           *time.Time `sql:"index"`
	Name                string     `gorm:"size:64;index:name" json:"name" validate:"required"`
	NameCn              string     `json:"name_cn"`
	NameEn              string     `json:"name_en"`
	PayOrReceive        string     `gorm:"size:16;index:pay_or_receive" comment:"类型" json:"pay_or_receive" validate:"required"`
	PayTypeId           uint       `gorm:"index:pay_type_id" comment:"结算方式" json:"pay_type_id"`
	FinanceCurrencyId   uint       `gorm:"index:finance_currency_id" comment:"币种类型" json:"finance_currency_id" validate:"required"`
	FinanceCurrencyRate float64    `gorm:"PRECISION: 4" json:"finance_currency_rate" validate:"required"`
	Quantity            float64    `gorm:"PRECISION: 4" json:"quantity" validate:"required"`
	UnitPrice           float64    `gorm:"PRECISION:4" json:"unit_price" validate:"required"`
	TaxRate             float64    `gorm:"PRECISION:4" json:"tax_rate"`
	TaxAmount           float64    `gorm:"PRECISION:4" json:"tax_amount" validate:"required"`
	NotTaxAmount        float64    `gorm:"PRECISION:4" json:"not_tax_amount"`
	ReceiveAmount       float64    `gorm:"PRECISION:4" json:"receive_amount"`
	Receivable          float64    `gorm:"PRECISION:4" json:"receivable"`
	PayAmount           float64    `gorm:"PRECISION:4" json:"pay_amount"`
	Payable             float64    `gorm:"PRECISION:4" json:"payable"`
	ClosingUnitId       uint       `gorm:"index:closing_unit_id" json:"closing_unit_id"`
	OrderMasterId       uint       `gorm:"index:order_master_id" json:"order_master_id"`
	SerialNumber        string     `gorm:"-" json:"serial_number"`
	FinanceStatementId  uint       `json:"finance_statement_id"`
	FinanceStatementNo  string     `gorm:"size:32;index:finance_statement_no" json:"finance_statement_no"`
	Status              string     `gorm:"size:16;index:status;default:'init'" json:"status"`
	InvoiceAmount       float64    `gorm:"PRECISION:4" json:"invoice_amount"`
	InvoiceStatus       string     `gorm:"size:16;index:invoice_status;default:'init'" json:"invoice_status"`
	TypeId              uint       `gorm:"index:type_id;default:1" json:"type_id"`
	Remarks             string     `gorm:"size:128" json:"remarks"`
}

func (FinanceRate) TableName() string {
	return "finance_rates"
}

func (FinanceFeeType) TableName() string {
	return "finance_fee_types"
}

func (FinanceFee) TableName() string {
	return "finance_fees"
}

const (
	FinanceFeeStatusInit        = "init"
	FinanceFeeStatusDismiss     = "dismiss"
	FinanceFeeStatusVerify      = "verify"
	FinanceFeeStatusPending     = "pending"
	FinanceFeeStatusApproval    = "approval"
	FinanceFeeStatusReview      = "review"
	FinanceFeeStatusUnapplied   = "unapplied"
	FinanceFeeStatusPartApplied = "part_applied"
	FinanceFeeStatusApplied     = "applied"
)
const (
	FinanceFeeInvoiceInit = "init"
)
