package models

import "github.com/jinzhu/gorm"

type UserCompany struct {
	gorm.Model
	CompanyType       int     `gorm:"company_type" json:"company_type"`
	ParentId          int     `gorm:"parent_id" json:"parent_id"`               // 父级id,区分谁的客户
	UserSalesmanId    int     `gorm:"user_salesman_id" json:"user_salesman_id"` // 所属的业务人员
	IsHeadOffice       int     `gorm:"is_head_office" json:"is_head_office"`     // 是否为总部
	AccountPeriod     string  `gorm:"account_period" json:"account_period"`     // 公司结算类型
	Age               int     `gorm:"age" json:"age"`                           // 公司账龄
	Amount            float64 `gorm:"amount" json:"amount"`                     // 月结金额
	NameNick          string  `gorm:"name_nick" json:"name_nick"`
	NameCn            string  `gorm:"name_cn" json:"name_cn"`
	NameEn            string  `gorm:"name_en" json:"name_en"`
	BusinessTypeName  string  `gorm:"business_type_name" json:"business_type_name"`
	ScaleGroupId      int     `gorm:"scale_group_id" json:"scale_group_id"`
	Status            string  `gorm:"status" json:"status"`
	Remark            string  `gorm:"remark" json:"remark"`         // 公司备注
	AuditorId         int     `gorm:"auditor_id" json:"auditor_id"` // 审核人的users.id
	CreatedId         int     `gorm:"created_id" json:"created_id"` // 创建者
	CustomerSource    int     `gorm:"customer_source" json:"customer_source"`
	TagGrade          int     `gorm:"tag_grade" json:"tag_grade"`
	IsVip             int     `gorm:"is_vip" json:"is_vip"`                           // 是否为VIP会员
	FrequentlyUseInfo string  `gorm:"frequently_use_info" json:"frequently_use_info"` // 公司常用信息
}

func (*UserCompany) TableName() string {
	return "user_companies"
}
