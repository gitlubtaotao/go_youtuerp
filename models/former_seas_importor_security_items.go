package models

import (
	"time"
)

type FormerSeasImportorSecurityItems struct {
	Id                 int64     `xorm:"pk autoincr BIGINT(20)"`
	ImportorSecurityId int64     `xorm:"index BIGINT(20)"`
	PoNo               string    `xorm:"comment('PO号') VARCHAR(255)"`
	Item               string    `xorm:"comment('商品') VARCHAR(255)"`
	Description        string    `xorm:"comment('描述') VARCHAR(255)"`
	CreatedAt          time.Time `xorm:"not null DATETIME"`
	UpdatedAt          time.Time `xorm:"not null DATETIME"`
	HsCode             string    `xorm:"comment('HS商品编码(前六位)') VARCHAR(255)"`
}
