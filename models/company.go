package models

import "time"

type UserCompany struct {
	ID               uint       `gorm:"primary_key"json:"id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `sql:"index"`
	CompanyType      int        `form:"company_type" json:"company_type"`
	ParentId         int        `gorm:"-"`
	UserSalesmanId   int        `gorm:"-"`                                                        // 所属的业务人员
	IsHeadOffice     bool       `gorm:"default:true" form:"is_head_office" json:"is_head_office"` // 是否为总部
	AccountPeriod    string     `gorm:"-"`                                                        // 公司结算类型
	Age              int        `gorm:"-"`                                                        // 公司账龄
	Amount           float64    `gorm:"-"`                                                        // 月结金额
	NameNick         string     `form:"name_nick" json:"name_nick" validate:"required,unique"`
	NameCn           string     `form:"name_cn" json:"name_cn" validate:"required,unique"`
	NameEn           string     `form:"name_en" json:"name_en" validate:"required,unique"`
	BusinessTypeName string     `gorm:"-"`
	Status           string     `grom:"-"`
	Telephone        string     `gorm:"not null;" form:"telephone" json:"telephone" validate:"required"` // 座机
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


