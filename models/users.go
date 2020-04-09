package models

import (
	"time"
)

type Users struct {
	Id                  int64     `xorm:"pk autoincr BIGINT(20)"`
	Email               string    `xorm:"not null default '' comment('email') VARCHAR(255)"`
	EncryptedPassword   string    `xorm:"not null default '' VARCHAR(255)"`
	ResetPasswordToken  string    `xorm:"unique VARCHAR(255)"`
	ResetPasswordSentAt time.Time `xorm:"DATETIME"`
	RememberCreatedAt   time.Time `xorm:"DATETIME"`
	SignInCount         int       `xorm:"not null default 0 INT(11)"`
	CurrentSignInAt     time.Time `xorm:"DATETIME"`
	LastSignInAt        time.Time `xorm:"DATETIME"`
	CurrentSignInIp     string    `xorm:"VARCHAR(255)"`
	LastSignInIp        string    `xorm:"VARCHAR(255)"`
	CreatedAt           time.Time `xorm:"DATETIME"`
	UpdatedAt           time.Time `xorm:"DATETIME"`
	UserCompanyId       int       `xorm:"index INT(11)"`
	DepartmentId        int64     `xorm:"index BIGINT(20)"`
	DeletedAt           time.Time `xorm:"index DATETIME"`
	Name                string    `xorm:"comment('姓名') index VARCHAR(255)"`
	VerifiedAt          time.Time `xorm:"DATETIME"`
	IsSync              int       `xorm:"default 0 TINYINT(1)"`
	AuthenticationToken string    `xorm:"VARCHAR(255)"`
	InitialPassword     string    `xorm:"VARCHAR(255)"`
	DeviseUuid          string    `xorm:"comment('单点登录验证') VARCHAR(255)"`
	IsAdmin             int       `xorm:"default 0 comment('是否为超级管理人员(系统默认只有一位)') TINYINT(1)"`
	Provider            string    `xorm:"VARCHAR(255)"`
	Uid                 string    `xorm:"VARCHAR(255)"`
	UserNo              string    `xorm:"comment('工号') VARCHAR(255)"`
	Phone               string    `xorm:"VARCHAR(32)"`
}
