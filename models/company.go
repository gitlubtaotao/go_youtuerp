package models

import "github.com/jinzhu/gorm"

type UserCompany struct {
	gorm.Model
	CompanyType       int         `form:"company_type" json:"company_type"`
	ParentId          int         `form:"parent_id" json:"parent_id"`               // 父级id,区分谁的客户
	UserSalesmanId    int         `form:"user_salesman_id" json:"user_salesman_id"` // 所属的业务人员
	IsHeadOffice      int         `form:"is_head_office" json:"is_head_office"`     // 是否为总部
	AccountPeriod     string      `form:"account_period" json:"account_period"`     // 公司结算类型
	Age               int         `form:"age" json:"age"`                           // 公司账龄
	Amount            float64     `form:"amount" json:"amount"`                     // 月结金额
	NameNick          string      `form:"name_nick" json:"name_nick" validate:"required,unique"`
	NameCn            string      `form:"name_cn" json:"name_cn" validate:"required,unique"`
	NameEn            string      `form:"name_en" json:"name_en" validate:"required,unique"`
	BusinessTypeName  string      `form:"business_type_name" json:"business_type_name"`
	ScaleGroupId      int         `form:"scale_group_id" json:"scale_group_id"`
	Status            string      `form:"status" json:"status"`
	Remark            string      `form:"remark" json:"remark"`         // 公司备注
	AuditorId         int         `form:"auditor_id" json:"auditor_id"` // 审核人的users.id
	CreatedId         int         `form:"created_id" json:"created_id"` // 创建者
	CustomerSource    int         `form:"customer_source" json:"customer_source"`
	TagGrade          int         `form:"tag_grade" json:"tag_grade"`
	IsVip             int         `form:"is_vip" json:"is_vip"`                           // 是否为VIP会员
	FrequentlyUseInfo string      `form:"frequently_use_info" json:"frequently_use_info"` // 公司常用信息
	Company           CompanyInfo `gorm:"foreignkey:source_id" table_name:"companies"`
}

//定义公司类型
const (
	Customer = iota
	Supplier
	CustomerSupplier
	Branch
)

func (UserCompany) TableName() string {
	return "user_companies"
}
