package models

import (
	"gorm.io/gorm"
	"time"
)

type UserCompany struct {
	ID               uint         `gorm:"primary_key"json:"id"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
	DeletedAt        *time.Time   `sql:"index"`
	CompanyType      int          `form:"company_type" json:"company_type"`
	ParentId         int          `gorm:"-"`
	UserSalesmanId   int          `gorm:"-"`                                   // 所属的业务人员
	IsHeadOffice     bool         `gorm:"default:false" form:"is_head_office"` // 是否为总部
	AccountPeriod    string       `gorm:"-"`                                   // 公司结算类型
	Age              int          `gorm:"-"`                                   // 公司账龄
	Amount           float64      `gorm:"-"`                                   // 月结金额
	NameNick         string       `gorm:"unique;not null" form:"name_nick" json:"name_nick" validate:"required"`
	NameCn           string       `gorm:"unique;not null" form:"name_cn" json:"name_cn" validate:"required"`
	NameEn           string       `gorm:"unique;not null" form:"name_en" json:"name_en" validate:"required"`
	BusinessTypeName string       `gorm:"-"`
	Status           string       `grom:"-"`
	Telephone        string       `form:"telephone" json:"telephone" validate:"required"` // 座机
	Email            string       `json:"email" validate:"required"`
	Fax              string       `json:"fax"`
	ZhAddress        string       `gorm:"column:address;" form:"zh_address" json:"zh_address"`
	EnAddress        string       `gorm:"column:address2;" form:"en_address" json:"en_address"`
	Remark           string       `form:"remark" json:"remark"` // 公司备注
	Website          string       `form:"website" json:"website"`
	Employees        []Employee   `gorm:"foreignkey:user_company_id"`
	Departments      []Department `gorm:"foreignkey:user_company_id"`
	Accounts         []Account    `gorm:"foreignkey:user_company_id"`
}

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
	ID                    uint      `json:"id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	NameCn                string    `json:"name_cn"`
	NameEn                string    `json:"name_en"`
	UserCompanyId         int       `json:"user_company_id"`
	UserCompaniesNameNick string    `json:"user_companies_name_nick"`
}

type Employee struct {
	ID                  uint      `json:"id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	DeletedAt           *time.Time
	Email               string `json:"email" validate:"required,email"` // email
	EncryptedPassword   string
	ResetPasswordToken  string
	ResetPasswordSentAt time.Time
	RememberCreatedAt   time.Time
	SignInCount         int       `json:"sign_in_count"`
	CurrentSignInAt     time.Time `json:"current_sign_in_at"`
	LastSignInAt        time.Time ` json:"last_sign_in_at"`
	CurrentSignInIp     string    `json:"current_sign_in_ip"`
	LastSignInIp        string    `json:"last_sign_in_ip"`
	UserCompanyId       int       `form:"user_company_id" json:"user_company_id" validate:"required"`
	DepartmentId        int       `form:"department_id" json:"department_id"`
	Name                string    `form:"name" json:"name" validate:"required"` // 姓名
	AuthenticationToken string
	IsAdmin             bool        `gorm:"default:false" form:"is_admin" json:"is_admin"` // 是否为超级管理人员(系统默认只有一位)
	UserNo              string      `form:"user_no" json:"user_no"`                        // 工号
	Phone               string      `gorm:"UNIQUE_INDEX;size:64" form:"phone" json:"phone" validate:"required"`
	Address             string      `json:"address" form:"address"`
	Remarks             string      `gorm:"size:65535" json:"remarks" form:"remarks"`
	Sex                 uint        `gorm:"default:0" json:"sex" form:"sex"`
	UserCompany         UserCompany `gorm:"foreignkey:user_company_id" validate:"structonly"`
	Department          Department  `gorm:"foreignkey:department_id" validate:"structonly"`
	Avatar              string      `gorm:"size:255" json:"avatar" yaml:"avatar"`
	CompanyType         uint        `json:"company_type"`
}

type ResultEmployee struct {
	ID                    uint      `json:"id"`
	CreatedAt             time.Time `json:"created_at"`
	Email                 string    `json:"email"`
	SignInCount           int       `json:"sign_in_count"`
	LastSignInAt          time.Time `json:"last_sign_in_at"`
	LastSignInIp          string    `json:"last_sign_in_ip"`
	UserCompanyId         int       `json:"user_company_id"`
	UserCompaniesNameNick string    `json:"user_companies_name_nick"`
	DepartmentId          int       `json:"department_id"`
	DepartmentsNameCN     string    `json:"departments_name_cn"`
	Name                  string    `json:"name"`
	UserNo                string    `json:"user_no"`
	Phone                 string    `json:"phone"`
}

func (UserCompany) TableName() string {
	return "user_companies"
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

func (Employee) TableName() string {
	return "users"
}

func (ResultEmployee) TableName() string {
	return "users"
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	e.ResetPasswordSentAt = time.Now()
	e.RememberCreatedAt = time.Now()
	e.LastSignInAt = time.Now()
	e.CurrentSignInAt = time.Now()
	e.CompanyType = CompanyTypeB
	return
}
