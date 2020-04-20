package models

import "github.com/jinzhu/gorm"

type UserCompany struct {
	gorm.Model
	CompanyType      int     `form:"company_type" json:"company_type"`
	ParentId         int     `gorm:"-"`
	UserSalesmanId   int     `gorm:"-"`                                                        // 所属的业务人员
	IsHeadOffice     bool    `gorm:"default:true" form:"is_head_office" json:"is_head_office"` // 是否为总部
	AccountPeriod    string  `gorm:"-"`                                                        // 公司结算类型
	Age              int     `gorm:"-"`                                                        // 公司账龄
	Amount           float64 `gorm:"-"`                                                        // 月结金额
	NameNick         string  `form:"name_nick" json:"name_nick" validate:"required,unique"`
	NameCn           string  `form:"name_cn" json:"name_cn" validate:"required,unique"`
	NameEn           string  `form:"name_en" json:"name_en" validate:"required,unique"`
	BusinessTypeName string  `gorm:"-"`
	Status           string
	Remark           string      `form:"remark" json:"remark"` // 公司备注
	Company          CompanyInfo `gorm:"foreignkey:source_id" table_name:"companies"`
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
