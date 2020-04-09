package models

import (
	"time"
)

type FinanceStatements struct {
	Id                      int64     `xorm:"pk autoincr BIGINT(20)"`
	ClosingUnitId           int       `xorm:"comment('结算单位,方便检索') index INT(11)"`
	StatementType           string    `xorm:"comment('账单类型 cash_now(票结) cash_month(月结)') VARCHAR(255)"`
	ConfirmBy               string    `xorm:"default 'client' comment('前台确认还是后台确认') VARCHAR(20)"`
	Status                  string    `xorm:"default 'pending' comment('前台对账状态') index VARCHAR(255)"`
	Remarks                 string    `xorm:"comment('备注') VARCHAR(255)"`
	TemplateId              int       `xorm:"comment('账单模板的id') INT(11)"`
	CustomStatementHtml     string    `xorm:"comment('如果操作在模板的基础上修改生成的账单，保存该html') TEXT"`
	CreatedAt               time.Time `xorm:"not null DATETIME"`
	UpdatedAt               time.Time `xorm:"not null DATETIME"`
	FeeIdJson               string    `xorm:"comment('費用的id') TEXT"`
	FinanceMonthStatementId int       `xorm:"comment('月结对账单id') INT(11)"`
	OrderMasterId           int       `xorm:"comment('订单号') INT(11)"`
	DebitNoteSn             string    `xorm:"comment('账单号') VARCHAR(255)"`
	BeginAt                 time.Time `xorm:"comment('月结开始时间') DATE"`
	EndAt                   time.Time `xorm:"comment('月结账单结束时间') DATE"`
	DeletedAt               time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
	UserCompanyId           int       `xorm:"comment('生成账单的公司,用于权限控制，满足子公司不能看总公司账单') index INT(11)"`
	UserId                  int       `xorm:"comment('导出账单的人员') INT(11)"`
	LockVersion             int       `xorm:"default 0 INT(11)"`
}
