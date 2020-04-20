package models

import (
	"time"
)

type Employee struct {
	ID                  uint       `gorm:"primary_key"json:"id"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	DeletedAt           *time.Time `sql:"index"`
	Email               string     `grom:"type:varchar(100);email;unique;not_null;" form:"email" json:"email" validate:"required,email,unique"` // email
	EncryptedPassword   string
	ResetPasswordToken  string
	ResetPasswordSentAt string
	RememberCreatedAt   string
	SignInCount         int       `json:"sign_in_count"`
	CurrentSignInAt     time.Time `gorm:"default:current_time" json:"current_sign_in_at"`
	LastSignInAt        time.Time ` json:"last_sign_in_at"`
	CurrentSignInIp     string    `json:"current_sign_in_ip" validate:"ip"`
	LastSignInIp        string    `json:"last_sign_in_ip" validate:"ip"`
	UserCompanyId       int       `form:"user_company_id" json:"user_company_id"`
	
	DepartmentId        int    `form:"department_id" json:"department_id"`
	Name                string `form:"name" json:"name" validate:"required,unique"` // 姓名
	AuthenticationToken string
	IsAdmin             bool        `gorm:"default:false" form:"is_admin" json:"is_admin"` // 是否为超级管理人员(系统默认只有一位)
	UserNo              string      `form:"user_no" json:"user_no"`                        // 工号
	Phone               string      `gorm:"UNIQUE_INDEX;size:64" form:"phone" json:"phone" validate:"required"`
	Address             string      `json:"address" form:"address"`
	Remarks             string      `gorm:"size:65535" json:"remarks" form:"remarks"`
	Sex                 uint        `gorm:"default:0" json:"sex" form:"sex"`
	UserCompany         UserCompany `gorm:"foreignkey:user_company_id"`
	Avatar              string      `gorm:"size:255" json:"avatar" yaml:"avatar"`
}

func (Employee) TableName() string {
	return "users"
}
