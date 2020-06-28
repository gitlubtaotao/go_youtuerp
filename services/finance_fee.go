package services

import (
	"youtuerp/models"
	"youtuerp/redis"
	"youtuerp/repositories"
)

type IFinanceFee interface {
	//更改费用状态
	ChangeStatusFees(ids []uint, status string) error
	//删除费用状态
	DeleteFees(ids []uint) error
	//插入费用信息
	BulkInsert([]models.FinanceFee) ([]models.FinanceFee, error)
	//获取订单对应的费用
	OrderFees(orderId uint, payOrReceive ...string) (map[string][]models.FinanceFee, error)
	OrderFeesOptions(companyId uint) map[string]interface{}
}

type FinanceFee struct {
	repo repositories.IFinanceFee
}

func (f FinanceFee) ChangeStatusFees(ids []uint, status string) error {
	return f.repo.ChangeStatusFees(ids,status)
}

func (f FinanceFee) DeleteFees(ids []uint) error {
	return f.repo.DeleteFees(ids)
}

func (f FinanceFee) BulkInsert(financeFess []models.FinanceFee) ([]models.FinanceFee, error) {
	return f.repo.BulkInsertOrUpdate(financeFess)
}

func (f FinanceFee) OrderFeesOptions(companyId uint) map[string]interface{} {
	var data = make(map[string]interface{})
	financeBaseService := NewFinanceBase()
	codeService := NewBaseCode()
	rateOptions, _ := financeBaseService.GetAllFeeRate(companyId)
	data["currency_rate_options"] = rateOptions
	data["fee_type_options"] = financeBaseService.FindFeeTypeRedis()
	data["finance_currency"] = codeService.FindCollect(models.CodeFinanceCurrency)
	data["pay_type_options"] = codeService.FindCollect(models.CIQType)
	data["finance_tag_options"] = codeService.FindCollect(models.FinanceTag)
	data["system_finance_approve"] = redis.SystemFinanceApprove()
	data["system_finance_audit"] = redis.SystemFinanceAudit()
	return data
}

func (f FinanceFee) OrderFees(orderId uint, payOrReceive ...string) (map[string][]models.FinanceFee, error) {
	attr := map[string]interface{}{
		"order_master_id": orderId,
	}
	return f.repo.OrderFees(attr, payOrReceive...)
}

func NewFinanceFee() IFinanceFee {
	return FinanceFee{repo: repositories.NewFinanceFee()}
}
