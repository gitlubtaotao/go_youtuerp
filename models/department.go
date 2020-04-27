package models

import "time"

type Department struct {
	ID            uint        `gorm:"primary_key"json:"id"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	DeletedAt     *time.Time  `sql:"index"`
	NameCn        string      `form:"name_cn" json:"name_cn"` // 部门中文名
	NameEn        string      `form:"name_en" json:"name_en"` // 部门英文名
	UserCompanyId int         `form:"user_company_id" json:"user_company_id"`
	UserCompany   UserCompany `gorm:"foreignkey:user_company_id" json:"user_company" validate:"-"`
}

func (Department) TableName() string {
	return "departments"
}
