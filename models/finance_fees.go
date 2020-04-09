package models

import (
	"time"
)

type FinanceFees struct {
	Id                           int64     `xorm:"pk autoincr BIGINT(20)"`
	PayOrReceive                 string    `xorm:"comment('填pay 或 receive') index VARCHAR(255)"`
	Name                         string    `xorm:"comment('费用名称') VARCHAR(255)"`
	NameCn                       string    `xorm:"comment('费用中文名') VARCHAR(255)"`
	NameEn                       string    `xorm:"comment('费用英文名') VARCHAR(255)"`
	PayTypeId                    int       `xorm:"comment('结算方式') INT(11)"`
	FinanceCurrencyId            int       `xorm:"INT(11)"`
	Quantity                     string    `xorm:"default 1.0000 comment('数量') DECIMAL(20,4)"`
	UnitPrice                    string    `xorm:"default 0.0000 comment('单价') DECIMAL(20,4)"`
	TaxRate                      string    `xorm:"default 0.0000 comment('税率') DECIMAL(20,4)"`
	TaxAmount                    string    `xorm:"default 0.0000 comment('含税金额') DECIMAL(20,4)"`
	StandardCurrencyExchangeRate string    `xorm:"comment('本位币汇率') DECIMAL(15,4)"`
	Remark                       string    `xorm:"comment('备注/附加说明') TEXT"`
	DebitNote                    string    `xorm:"comment('借方通知单') VARCHAR(255)"`
	ReceiveAmount                string    `xorm:"default 0.0000 comment('实收') DECIMAL(20,4)"`
	Receivable                   string    `xorm:"default 0.0000 comment('应收') DECIMAL(20,4)"`
	CreditNote                   string    `xorm:"comment('收款通知单') VARCHAR(255)"`
	PayAmount                    string    `xorm:"default 0.0000 comment('实付') DECIMAL(20,4)"`
	Payable                      string    `xorm:"default 0.0000 comment('应付') DECIMAL(20,4)"`
	ClosingUnitType              string    `xorm:"index(index_finance_fees_on_closing_unit_type_and_closing_unit_id) VARCHAR(255)"`
	ClosingUnitId                int64     `xorm:"comment('结算单位') index(index_finance_fees_on_closing_unit_type_and_closing_unit_id) BIGINT(20)"`
	CreatedAt                    time.Time `xorm:"not null DATETIME"`
	UpdatedAt                    time.Time `xorm:"not null DATETIME"`
	SourceType                   string    `xorm:"index(index_finance_fees_on_source_type_and_source_id) VARCHAR(255)"`
	SourceId                     int64     `xorm:"comment('数据来源') index(index_finance_fees_on_source_type_and_source_id) BIGINT(20)"`
	DeletedAt                    time.Time `xorm:"index DATETIME"`
	OrderMasterId                int       `xorm:"index INT(11)"`
	LockVersion                  int       `xorm:"INT(11)"`
	Status                       string    `xorm:"not null default 'created' comment('财务账单状态') VARCHAR(16)"`
	Locked                       int       `xorm:"not null default 0 comment('财务状态锁') TINYINT(1)"`
	InvoicesId                   int64     `xorm:"index BIGINT(20)"`
	AccountId                    int64     `xorm:"index BIGINT(20)"`
	FinanceCurrencyRate          string    `xorm:"DECIMAL(13,4)"`
	AddressId                    int64     `xorm:"index BIGINT(20)"`
	DebitNoteSn                  string    `xorm:"comment('账单流水号') VARCHAR(100)"`
	VerifyStatus                 string    `xorm:"default 'unfinished' comment('对账状态') VARCHAR(12)"`
	InvoiceAmount                string    `xorm:"default 0.0000 comment('开票金额') DECIMAL(20,4)"`
	InvoiceStatus                string    `xorm:"default 'uninvoiced' comment('开票状态') VARCHAR(64)"`
	AccountTitleId               int       `xorm:"default 1 comment('会计类目/费用种类') INT(11)"`
	FinanceStatementId           int       `xorm:"comment('对账单id') INT(11)"`
	PlanFeeId                    int       `xorm:"default 0 INT(11)"`
}
