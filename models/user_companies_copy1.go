package models

import (
	"time"
)

type UserCompaniesCopy1 struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt        time.Time `xorm:"DATETIME"`
	UpdatedAt        time.Time `xorm:"DATETIME"`
	DeletedAt        time.Time `xorm:"index DATETIME"`
	CompanyType      int       `xorm:"default 1 INT(11)"`
	ParentId         int       `xorm:"default 0 comment('父级id,区分谁的客户') INT(11)"`
	UserSalesmanId   int       `xorm:"default 0 comment('所属的业务人员') index INT(11)"`
	IsHeadOffice     int       `xorm:"default 0 comment('是否为总部') TINYINT(1)"`
	AccountPeriod    string    `xorm:"comment('公司结算类型') VARCHAR(16)"`
	Age              int       `xorm:"comment('公司账龄') SMALLINT(6)"`
	Amount           string    `xorm:"comment('月结金额') DECIMAL(13,4)"`
	IsBlack          int       `xorm:"not null default 0 comment('是否加入黑名单') TINYINT(1)"`
	NameNick         string    `xorm:"VARCHAR(256)"`
	NameCn           string    `xorm:"VARCHAR(256)"`
	NameEn           string    `xorm:"VARCHAR(256)"`
	BusinessTypeName string    `xorm:"VARCHAR(255)"`
	ScaleGroupId     int       `xorm:"INT(11)"`
	Status           string    `xorm:"default 'approved' VARCHAR(15)"`
}
