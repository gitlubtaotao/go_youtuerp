package models

import (
	"time"
)

type ScoreSystemDetails struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedAt      time.Time `xorm:"not null DATETIME"`
	ScoreId        int       `xorm:"comment('对应的积分表') index INT(11)"`
	ScoreItem      string    `xorm:"comment('积分项目') VARCHAR(255)"`
	OperatorItem   string    `xorm:"comment('操作项目,[系统,人工]') VARCHAR(255)"`
	UserOperatorId int       `xorm:"comment('系统操作人员') index INT(11)"`
	UserAuditId    int       `xorm:"comment('系统审核人员') index INT(11)"`
	OperaScore     string    `xorm:"comment('操作的积分') DECIMAL(15,2)"`
	ScoreObject    string    `xorm:"comment('积分操作对象') VARCHAR(255)"`
	ScoreMultiple  string    `xorm:"comment('操作倍数') VARCHAR(10)"`
	Remark         string    `xorm:"comment('备注') TEXT"`
	OperatorDetail string    `xorm:"TEXT"`
	Status         string    `xorm:"default 'init' comment('操作状态') VARCHAR(255)"`
	SourceType     string    `xorm:"index(index_score_system_details_on_source_type_and_source_id) VARCHAR(255)"`
	SourceId       int64     `xorm:"comment('数据来源') index(index_score_system_details_on_source_type_and_source_id) BIGINT(20)"`
	DataObjectType string    `xorm:"index(index_and_data_object) VARCHAR(255)"`
	DataObjectId   int64     `xorm:"index(index_and_data_object) BIGINT(20)"`
	ScoreRuleId    int       `xorm:"index INT(11)"`
	CompanyId      int       `xorm:"default 0 index INT(11)"`
}
