package models

import "time"

type User struct {
	ID                  uint       `gorm:"primary_key"json:"id"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	DeletedAt           *time.Time `sql:"index"`
	Email               string     `grom:"type:varchar(100);email;unique;not_null;" form:"email" json:"email" validate:"required,email"` // email
	EncryptedPassword   string
	ResetPasswordToken  string
	ResetPasswordSentAt time.Time
	RememberCreatedAt   time.Time
	SignInCount         int       `json:"sign_in_count"`
	CurrentSignInAt     time.Time `gorm:"default:current_time" json:"current_sign_in_at"`
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
	IsKeyContact        bool        `gorm:"default: false" json:"is_key_contact"`
	CompanyType         uint        `json:"company_type"`
}

func (User) TableName() string {
	return "users"
}

const (
	UserMan = iota + 1
	UserWom
)
