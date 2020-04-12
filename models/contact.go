package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Contact struct {
	gorm.Model
	Email               string `grom:"type:varchar(100);email;unique;not_null;" form:"email" json:"email" validate:"required,email,unique"` // email
	EncryptedPassword   string
	ResetPasswordToken  string
	ResetPasswordSentAt string
	RememberCreatedAt   string
	SignInCount         int       `json:"sign_in_count"`
	CurrentSignInAt     time.Time `json:"current_sign_in_at"`
	LastSignInAt        time.Time `json:"last_sign_in_at"`
	CurrentSignInIp     string    `json:"current_sign_in_ip" validate:"ip"`
	LastSignInIp        string    `json:"last_sign_in_ip" validate:"ip"`
	UserCompanyId       int       `form:"user_company_id" json:"user_company_id"`
	Name                string    `form:"name" json:"name" validate:"required,unique"` // 姓名
	AuthenticationToken string
	Phone               string `form:"phone" json:"phone" validate:"required"`
	Address             string `json:"address" form:"address"`
	Remarks             string `gorm:"size:65535" json:"remarks" form:"remarks"`
	Sex                 uint   `gorm:"default:0" json:"sex" form:"sex"`
	IsKeyContact        bool   `gorm:"default:false" json:"is_key_contact" form:"is_key_contact"`
}

const (
	Man = iota
	Woman
)

func (Contact) TableName() string {
	return "users"
}
