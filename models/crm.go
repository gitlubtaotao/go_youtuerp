package models

import "time"

type CrmClue struct {
	ID            uint       `gorm:"primary_key"json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `sql:"index"`
	CompanyType   uint       `json:"company_type"`
	NameNick      string     `gorm:"size:64;index:name_nick" json:"name_nick" validate:"required"`
	NameCn        string     `gorm:"size:128" json:"name_cn" validate:"required"`
	NameEn        string     `gorm:"size:256" json:"name_en"`
	Tel           string     `gorm:"size:16;index:tel" json:"tel" validate:"required"`
	Source        string     `gorm:"size:16" json:"source"`
	Email         string     `gorm:"size:64" json:"email" validate:"required"`
	ZhAddress     string     `gorm:"size:512" json:"zh_address"`
	EnAddress     string     `gorm:"size:1024" json:"en_address"`
	Remarks       string     `gorm:"size:1024" json:"remarks"`
	UserName      string     `gorm:"size:64;index:user_name" json:"user_name"`
	UserTel       string     `gorm:"size:16;index:user_tel" json:"user_tel"`
	UserEmail     string     `gorm:"size:64" json:"user_email"`
	WechatId      string     `grom:"size:16" json:"wechat_id"`
	QQId          string     `gorm:"size:16" json:"qq_id"`
	UserRemarks   string     `gorm:"size:1024" json:"user_remarks"`
	CreateId      uint       `gorm:"index:create_id" json:"create_id"`
	CreateName    string     `gorm:"-" json:"create_name"`
	UserCompanyId uint       `gorm:"index:user_company_id"`
	City          string     `gorm:"size:64" json:"city"`
	Province      string     `gorm:"size:64" json:"province"`
	Distinct      string     `gorm:"size:64" json:"distinct"`
}

type CrmCooperator struct {
}
type CrmSupply struct {
}

type CrmUser struct {
}

func (CrmClue) TableName() string {
	return "crm_clues"
}

func (CrmCooperator) TableName() string {
	return "user_companies"
}

func (CrmSupply) TableName() string {
	return "user_companies"
}

func (CrmUser) TableName() string {
	return "users"
}
