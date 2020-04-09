package models

import (
	"time"
)

type DocTemplateDefaults struct {
	Id                       int64     `xorm:"pk autoincr BIGINT(20)"`
	YoutuErpUserCompanyId    int64     `xorm:"index BIGINT(20)"`
	DocTemplateDocTemplateId int64     `xorm:"index BIGINT(20)"`
	DeletedAt                time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
	CreatedAt                time.Time `xorm:"not null DATETIME"`
	UpdatedAt                time.Time `xorm:"not null DATETIME"`
}
