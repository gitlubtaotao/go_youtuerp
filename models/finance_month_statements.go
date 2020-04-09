package models

import (
	"time"
)

type FinanceMonthStatements struct {
	Id                  int64     `xorm:"pk autoincr BIGINT(20)"`
	ClosingUnitId       int       `xorm:"comment('结算单位,方便检索') INT(11)"`
	Status              string    `xorm:"default 'pending' comment('前台对账状态') VARCHAR(255)"`
	ConfirmBy           string    `xorm:"default 'client' comment('前台确认还是后台确认') VARCHAR(20)"`
	Remarks             string    `xorm:"comment('备注') VARCHAR(255)"`
	TemplateId          int       `xorm:"comment('月结账单模板的id') INT(11)"`
	FeeIdJson           string    `xorm:"comment('費用的id') TEXT"`
	CustomStatementHtml string    `xorm:"comment('如果操作在模板的基础上修改生成的账单，保存该html') TEXT"`
	CreatedAt           time.Time `xorm:"not null DATETIME"`
	UpdatedAt           time.Time `xorm:"not null DATETIME"`
}
