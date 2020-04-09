package models

import (
	"time"
)

type YoutuErpRoles struct {
	Id              int64     `xorm:"pk autoincr BIGINT(20)"`
	Name            string    `xorm:"index index(index_youtu_erp_roles_on_name_and_resource_type_and_resource_id) VARCHAR(255)"`
	ResourceType    string    `xorm:"index(index_youtu_erp_roles_on_name_and_resource_type_and_resource_id) index(index_youtu_erp_roles_on_resource_type_and_resource_id) VARCHAR(255)"`
	ResourceId      int64     `xorm:"index(index_youtu_erp_roles_on_name_and_resource_type_and_resource_id) index(index_youtu_erp_roles_on_resource_type_and_resource_id) BIGINT(20)"`
	CreatedAt       time.Time `xorm:"DATETIME"`
	UpdatedAt       time.Time `xorm:"DATETIME"`
	IsAdmin         int       `xorm:"default 0 TINYINT(1)"`
	UserCompanyId   int       `xorm:"default 0 comment('该角色所属的公司') index INT(11)"`
	ShareCompanyIds string    `xorm:"comment('可共享其他分公司的数据') VARCHAR(255)"`
	DeletedAt       time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
	LockVersion     int       `xorm:"default 0 INT(11)"`
}
