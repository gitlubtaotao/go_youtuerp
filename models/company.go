package models

import "time"

type Company struct {
	ID                uint       `gorm:"primary_key"json:"id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `sql:"index"`
	CompanyType       int        `form:"company_type" json:"company_type"`
	ParentId          int        `gorm:"index:parent_id"`
	UserSalesmanId    int        `gorm:"index:user_salesman_id"`              // 所属的业务人员
	IsHeadOffice      bool       `gorm:"default:false" form:"is_head_office"` // 是否为总部
	AccountPeriod     string     `gorm:"size:16;index:account_period"`        // 公司结算类型
	Age               int        // 公司账龄
	Amount            int        // 月结金额
	NameNick          string     `gorm:"unique;not null" form:"name_nick" json:"name_nick" validate:"required"`
	NameCn            string     `gorm:"unique;not null" form:"name_cn" json:"name_cn" validate:"required"`
	NameEn            string     `gorm:"unique;not null" form:"name_en" json:"name_en" validate:"required"`
	BusinessTypeName  string     `gorm:"size:64"`
	Status            string     `gorm:"size:64;index:status"`
	Telephone         string     `gorm:"size:16;index:telephone" json:"telephone" validate:"required"` // 座机
	Email             string     `gorm:"size:64;index:email" json:"email" validate:"required"`
	Fax               string     `json:"fax"`
	ZhAddress         string     `gorm:"column:address;" form:"zh_address" json:"zh_address"`
	EnAddress         string     `gorm:"column:address2;" form:"en_address" json:"en_address"`
	Remark            string     `form:"remark" json:"remark"` // 公司备注
	Website           string     `form:"website" json:"website"`
	City              string     `gorm:"size:64"`
	Province          string     `gorm:"size:64"`
	Distinct          string     `gorm:"size:64"`
	Code              string     `gorm:"size:64"`
	Source            string     `gorm:"size:64" json:"source"`
	FrequentlyUseInfo string     `gorm:"size:1024;column:frequently_use_info"`
}

//定义公司类型
const (
	CompanyTypeC = iota + 1
	CompanyTypeS
	CompanyTypeCS
	CompanyTypeB
)
const (
	CompanyStatusCancel    = "cancel"
	CompanyStatusRejected  = "rejected"
	CompanyStatusApproved  = "approved"
	CompanyStatusApproving = "approving"
)

func (Company) TableName() string {
	return "user_companies"
}
