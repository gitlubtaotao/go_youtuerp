package models

import "time"

type UserCompany struct {
	ID               uint       `gorm:"primary_key"json:"id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `sql:"index"`
	CompanyType      int        `form:"company_type" json:"company_type"`
	ParentId         int        `gorm:"-"`
	UserSalesmanId   int        `gorm:"-"`                                  // 所属的业务人员
	IsHeadOffice     bool       `gorm:"default:false" form:"is_head_office"` // 是否为总部
	AccountPeriod    string     `gorm:"-"`                                  // 公司结算类型
	Age              int        `gorm:"-"`                                  // 公司账龄
	Amount           float64    `gorm:"-"`                                  // 月结金额
	NameNick         string     `gorm:"unique;not null" form:"name_nick" json:"name_nick" validate:"required"`
	NameCn           string     `gorm:"unique;not null" form:"name_cn" json:"name_cn" validate:"required"`
	NameEn           string     `gorm:"unique;not null" form:"name_en" json:"name_en" validate:"required"`
	BusinessTypeName string     `gorm:"-"`
	Status           string     `grom:"-"`
	Telephone        string     `form:"telephone" json:"telephone" validate:"required"` // 座机
	Email            string     `json:"email" validate:"required"`
	Fax              string     `json:"fax"`
	ZhAddress        string     `gorm:"column:address;" form:"zh_address" json:"zh_address"`
	EnAddress        string     `gorm:"column:address2;" form:"en_address" json:"en_address"`
	Remark           string     `form:"remark" json:"remark"` // 公司备注
	Website          string     `form:"website" json:"website"`
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
