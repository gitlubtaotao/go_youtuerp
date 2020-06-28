package models

import "time"

type Setting struct {
	ID     uint   `gorm:"primary_key"json:"id"`
	Key    string `gorm:"size:64" json:"key"`
	Field  string `gorm:"size:64" json:"field"`
	Value  string `gorm:"type:varchar(1024)" json:"value"`
	UserId uint   `gorm:"index:user_id"`
}
type NumberSetting struct {
	ID            uint      `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	UserCompanyId uint      `gorm:"index:user_company_id" json:"user_company_id" validate:"required"`
	Prefix        string    `gorm:"size:64;comment:'前缀'" json:"prefix" validate:"required"`
	Length        uint      `gorm:"comment:'流水号长度'" json:"length" validate:"required,max=10,min=1"`
	YearRule      string    `gorm:"size:16;comment:'年规则'" json:"year_rule"`
	MonthRule     string    `gorm:"size:16;comment:'月规则'" json:"month_rule"`
	DayRule       string    `gorm:"size:16;comment:'日规则'" json:"day_rule"`
	DefaultRule   string    `gorm:"size:64" json:"default_rule"`
	Special       string    `gorm:"size:16" json:"special"`
	ApplicationNo string    `gorm:"size:64;index:application_no" json:"application_no" validate:"required"`
	DefaultNumber int       `gorm:"default:0" json:"default_number" validate:"required,min=1"`
	CurrentNumber int       `gorm:"default:0" json:"current_number"`
	ClearRule     string    `gorm:"size:16;comment:'清零规则'" json:"clear_rule"`
}

type NumberSettingHistory struct {
	ID              uint      `gorm:"primary_key" json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	NumberSettingId uint      `gorm:"index:number_setting_id" json:"number_setting_id"`
	Year            int       `gorm:"index:year_moth_day_index" json:"year"`
	Month           int       `gorm:"index:year_moth_day_index" json:"month"`
	Day             int       `gorm:"index:year_moth_day_index" json:"day"`
	CurrentNumber   int       `gorm:"default:0" json:"current_number"`
}

type ResultNumberSetting struct {
	ID                    uint      `json:"id"`
	CreatedAt             time.Time `json:"created_at"`
	UserCompanyId         uint      `json:"user_company_id"`
	UserCompaniesNameNick string    `json:"user_companies_name_nick"`
	Prefix                string    `json:"prefix"`
	Length                uint      `json:"length"`
	DefaultRule           string    `json:"default_rule"`
	Special               string    `json:"special"`
	ApplicationNo         string    `json:"application_no"`
	DefaultNumber         uint      `json:"default_number"`
	CurrentNumber         uint      `json:"current_number"`
	ClearRule             string    `json:"clear_rule"`
}

type ResultSetting struct {
	Key   string `json:"key"`
	Field string `json:"field"`
	Value string `json:"value"`
}

func (Setting) TableName() string {
	return "system_settings"
}

func (NumberSetting) TableName() string {
	return "number_settings"
}
func (ResultNumberSetting) TableName() string {
	return "number_settings"
}

const (
	NumberSettingOrderNumber = "order_serial_number"
)

const (
	NumberSettingYearClear  = "year"
	NumberSettingMonthClear = "month"
	NumberSettingDayClear   = "day"
	NumberSettingNonZero    = "non_zero"
)
const (
	SettingFeeRateNow   = "now"
	SettingFeeRateMonth = "month"
)
