package services

import (
	"github.com/kataras/golog"
	"time"
	"youtuerp/models"
	"youtuerp/redis"
	"youtuerp/repositories"
)

type IFinanceFee interface {
	CopyFee(orderMasterId []uint, financeFeeIds []uint, companyId int) error
	//通过ids 查询费用信息
	FindFeesById(ids []uint, otherKeys ...string) ([]models.FinanceFee, error)
	//根据前端查询条件查询费用信息
	FindFees(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string) ([] models.FinanceFee, error)
	//更改费用状态
	ChangeStatusFees(ids []uint, status string) error
	//删除费用状态
	DeleteFees(ids []uint) error
	//插入费用信息
	BulkInsert([]models.FinanceFee) ([]models.FinanceFee, error)
	//获取订单对应的费用
	OrderFees(orderId uint, payOrReceive ...string) (map[string][]models.FinanceFee, error)
	//获取订单费用对应的select options
	OrderFeesOptions(companyId uint) map[string]interface{}
}

type FinanceFee struct {
	repo repositories.IFinanceFee
}

func (f FinanceFee) CopyFee(orderMasterId []uint, financeFeeIds []uint, companyId int) error {
	//var sy sync.WaitGroup
	if redis.SystemRateSetting() == models.SettingFeeRateNow {
		return f.realTimeRateCopyFee(orderMasterId, financeFeeIds, uint(companyId))
	} else {
		return f.monthRateCopyFee(orderMasterId, financeFeeIds, uint(companyId))
	}
}

func (f FinanceFee) FindFeesById(ids []uint, otherKeys ...string) ([]models.FinanceFee, error) {
	return f.repo.FindFeesById(ids, otherKeys...)
}

func (f FinanceFee) FindFees(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string) ([] models.FinanceFee, error) {
	return nil, nil
}

func (f FinanceFee) ChangeStatusFees(ids []uint, status string) error {
	return f.repo.ChangeStatusFees(ids, status)
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

//按照实时汇率复制费用
func (f FinanceFee) realTimeRateCopyFee(orderMasterIds []uint, financeFeeIds []uint, companyId uint) error {
	financeFees, err := f.FindFeesById(financeFeeIds)
	if err != nil {
		return err
	}
	orderMaster, err := f.getOrderMaster(orderMasterIds)
	if err != nil {
		return err
	}
	rates, err := NewFinanceBase().GetAllFeeRate(companyId)
	if err != nil {
		return err
	}
	for _, order := range orderMaster {
		go func(order models.ResultOrderMaster, financeFees []models.FinanceFee) {
			for i := 0; i < len(financeFees); i++ {
				financeFees[i] = f.commonHandlerByCopyFee(rates, financeFees[i], order)
			}
			_, _ = f.BulkInsert(financeFees)
		}(order, financeFees)
	}
	return nil
}

// 按照月结汇率复制费用
func (f FinanceFee) monthRateCopyFee(orderMasterIds []uint, financeFeeIds []uint, companyId uint) error {
	financeFees, err := f.FindFeesById(financeFeeIds)
	if err != nil {
		return err
	}
	orderMaster, err := f.getOrderMaster(orderMasterIds)
	if err != nil {
		return err
	}
	for _, order := range orderMaster {
		go func(order models.ResultOrderMaster, financeFees []models.FinanceFee, companyId uint) {
			rates, err := f.getRateByOrderCreate(companyId, order.CreatedAt)
			if err != nil {
				golog.Errorf("copy fee is error %v", err)
				return
			}
			for i := 0; i < len(financeFees); i++ {
				financeFees[i] = f.commonHandlerByCopyFee(rates, financeFees[i], order)
			}
			_, _ = f.BulkInsert(financeFees)
		}(order, financeFees, companyId)
	}
	return nil
}

//复制费用，查询对于的订单信息
func (f FinanceFee) getOrderMaster(orderMasterIds []uint) ([]models.ResultOrderMaster, error) {
	orderServer := NewOrderMasterService()
	keys := []string{
		"order_masters.id as id",
		"order_masters.instruction_id as instruction_id",
		"order_extend_infos.supply_agent_id as supply_agent_id",
		"order_masters.created_at as created_at",
	}
	return orderServer.FindMasterByIds(orderMasterIds, keys...)
}

//根据币种查询汇率
func (f FinanceFee) searchRate(financeRate []models.FinanceRate, financeCurrencyId uint) float64 {
	for _, rate := range financeRate {
		if rate.FinanceCurrencyId == financeCurrencyId {
			return rate.Rate
		}
	}
	return 0
}

//根据不同的订单查询对于的汇率信息
func (f FinanceFee) getRateByOrderCreate(companyId uint, createdAt time.Time) (rates []models.FinanceRate, err error) {
	filter := map[string]interface{}{
		"year":        createdAt.Year(),
		"start_month": int(createdAt.Month()),
		"end_month":   int(createdAt.Month()) + 1,
	}
	return NewFinanceBase().GetAllFeeRate(companyId, filter)
}

//复制费用实时汇率和按照月结汇率的共同操作
func (f FinanceFee) commonHandlerByCopyFee(rates []models.FinanceRate, financeFee models.FinanceFee, order models.ResultOrderMaster) models.FinanceFee {
	rate := f.searchRate(rates, financeFee.FinanceCurrencyId)
	if rate != 0 {
		financeFee.FinanceCurrencyRate = rate
	}
	financeFee.Status = models.FinanceFeeStatusInit
	financeFee.PayAmount = 0
	financeFee.ReceiveAmount = 0
	financeFee.InvoiceAmount = 0
	financeFee.InvoiceStatus = models.FinanceFeeInvoiceInit
	financeFee.OrderMasterId = order.ID
	financeFee.ID = 0
	if financeFee.PayOrReceive == "pay" && order.SupplyAgentId != 0 {
		financeFee.ClosingUnitId = order.SupplyAgentId
	} else if financeFee.PayOrReceive == "receive" && order.InstructionId != 0 {
		financeFee.ClosingUnitId = order.InstructionId
	}
	return financeFee
}

func NewFinanceFee() IFinanceFee {
	return FinanceFee{repo: repositories.NewFinanceFee()}
}
