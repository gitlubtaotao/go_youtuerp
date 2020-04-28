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
	UserCompany   UserCompany `gorm:"foreignkey:user_company_id;auto_preload" table_name:"user_companies" json:"user_companies" validate:"-"`
}

//go index 查询结构
type ResultDepartment struct {
	ID                    uint  `json:"id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	NameCn                string `json:"name_cn"`
	NameEn                string `json:"name_en"`
	UserCompanyId         int `json:"user_company_id"`
	UserCompaniesNameNick string `json:"user_companies_name_nick"`
}

func (Department) TableName() string {
	return "departments"
}

func (ResultDepartment) TableName() string {
	return "departments"
}

func (Department) DefaultAddColumn() []string {
	return []string{}
}
