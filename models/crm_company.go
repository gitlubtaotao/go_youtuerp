package models

import "github.com/jinzhu/gorm"

type CrmCompany struct {
	gorm.Model
	CompanyType      int         `form:"company_type" json:"company_type"`
	ParentId         int         `json:"parent_id" form:"parent_id"`
	UserSalesmanId   int         `form:"user_salesman_id" json:"user_salesman_id"` // 所属的业务人员
	AccountPeriod    string      `form:"account_period" json:"account_period"`     // 公司结算类型
	Age              int         `form:"age" json:"age"`                           // 公司账龄
	Amount           float64     `form:"amount" json:"amount"`                     // 月结金额
	NameNick         string      `form:"name_nick" json:"name_nick" validate:"required,unique"`
	NameCn           string      `form:"name_cn" json:"name_cn" validate:"required,unique"`
	NameEn           string      `form:"name_en" json:"name_en" validate:"required,unique"`
	BusinessTypeName string      `form:"business_type_name" json:"business_type_name"`
	Status           string      `form:"status" json:"status"`
	Remark           string      `form:"remark" json:"remark"`         // 公司备注
	AuditorId        int         `form:"auditor_id" json:"auditor_id"` // 审核人的users.id
	CreatedId        int         `form:"created_id" json:"created_id"` // 创建者
	CustomerSource   int         `form:"customer_source" json:"customer_source"`
	TagGrade         int         `form:"tag_grade" json:"tag_grade"`
	IsVip            int         `form:"is_vip" json:"is_vip"` // 是否为VIP会员
	Company          CompanyInfo `gorm:"foreignkey:source_id" table_name:"companies"`
}

func (CrmCompany) TableName() string {
	return "user_companies"
}
