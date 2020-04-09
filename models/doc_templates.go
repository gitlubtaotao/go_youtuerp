package models

import (
	"time"
)

type DocTemplates struct {
	Id                int64     `xorm:"pk autoincr BIGINT(20)"`
	Name              string    `xorm:"comment('模板名') VARCHAR(255)"`
	WidgetParams      string    `xorm:"MEDIUMTEXT"`
	Remark            string    `xorm:"comment('备注') VARCHAR(255)"`
	CreatedAt         time.Time `xorm:"not null DATETIME"`
	UpdatedAt         time.Time `xorm:"not null DATETIME"`
	UserCompanyId     int64     `xorm:"index BIGINT(20)"`
	TemplateType      string    `xorm:"comment('1为海运,2为陆运') index VARCHAR(255)"`
	TmpName           string    `xorm:"comment('上传后的文件名') VARCHAR(255)"`
	OriginalFilename  string    `xorm:"comment('文件原始名') VARCHAR(255)"`
	DeletedAt         time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	IsBootTemplate    int       `xorm:"default 0 comment('是否Bootstrap模板') TINYINT(1)"`
	TransportType     int       `xorm:"default 1 comment('运输类型') INT(11)"`
	DocTemplateRuleId int       `xorm:"index INT(11)"`
	YoutuDefault      int       `xorm:"default 0 comment('优途提供的模板默认为false') TINYINT(1)"`
	IsValid           int       `xorm:"default 1 comment('是否有效,默认为true有效') TINYINT(1)"`
}
