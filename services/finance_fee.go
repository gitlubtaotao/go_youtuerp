package services

import (
	"errors"
	"github.com/kataras/golog"
	"time"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/redis"
	"youtuerp/repositories"
	"youtuerp/tools"
)

type IFinanceFee interface {
	//对返回前端的费用进行预处理
	HandleFeesShow(fee interface{}, enum conf.Enum) map[string]interface{}
	FindFinanceFees(per, page uint, filter map[string]interface{},
		selectKeys []string, orders []string) ([]models.ResultFinanceFee, uint, error)
	/*将查询到的历史费用复制到对应的订单中
	orderMasterId: 指定订单
	feeIds: 对应费用ids
	companyId: 当前操作所属公司
	closingUnitId: 结算单位
	payOrReceive: 收入/支出
	*/
	BulkHistoryFee(orderMasterId uint, feeIds []uint, companyId uint, closingUnitId uint, payOrReceive string) ([]models.FinanceFee, error)
	//根据结算单位，时间范围查询该结算单位的历史费用，通过name 进行分组
	GetHistoryFee(filter map[string]interface{}) ([]map[string]interface{}, error)
	/*将费用复制不同的订单
	orderMasterIds: 指定的订单Ids
	financeFeeIds: 复制费用的Ids
	companyId：当前公司
	*/
	CopyFee(orderMasterIds []uint, financeFeeIds []uint, companyId int) error
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
	BaseService
	repo repositories.IFinanceFee
}

func (f FinanceFee) HandleFeesShow(fee interface{}, enum conf.Enum) map[string]interface{} {
	data := toolOther.StructToMap(fee)
	golog.Infof("data is %v",data)
	data["pay_or_receive"] = enum.DefaultText("finance_fees_pay_or_receive.", data["pay_or_receive"].(string))
	data["pay_type_id"] = f.baseDataFindFast(models.CIQType, data["pay_type_id"])
	data["status_value"] = data["status"]
	data["invoice_status_value"] = data["invoice_status"]
	data["invoice_status"] = enum.DefaultText("finance_fees_invoice_status.", data["invoice_status"])
	data["type_id"] = f.baseDataFindFast(models.FinanceTag, data["type_id"])
	data["finance_currency_id"] = f.baseDataFindFast(models.CodeFinanceCurrency, data["finance_currency_id"])
	data["closing_unit_id_value"] = data["closing_unit_id"]
	data["closing_unit_id"] = red.HGetCrm(data["closing_unit_id"], "")
	data["order_master_id_value"] = data["order_master_id"]
	data["order_master_id"] = data["serial_number"]
	data["status"] = enum.DefaultText("finance_fees_status.", data["status"])
	return data
}

func (f FinanceFee) FindFinanceFees(per, page uint, filter map[string]interface{},
	selectKeys []string, orders []string) ([]models.ResultFinanceFee, uint, error) {
	return f.repo.FindFinanceFees(per, page, filter, selectKeys, orders)
}

func (f FinanceFee) GetHistoryFee(filter map[string]interface{}) ([]map[string]interface{}, error) {
	var returnResult []map[string]interface{}
	financeFees, err := f.repo.GetHistoryFee(filter, 50, []string{})
	if err != nil {
		return returnResult, err
	}
	tool := tools.OtherHelper{}
	codeServer := NewBaseCode()
	for _, fee := range financeFees {
		result := tool.StructToMap(fee)
		result["finance_currency_id"] = codeServer.HGetValue(models.CodeFinanceCurrency, result["finance_currency_id"], "")
		result["type_id"] = codeServer.HGetValue(models.FinanceTag, result["type_id"], "")
		result["pay_type_id"] = codeServer.HGetValue(models.CIQType, result["pay_type_id"], "")
		returnResult = append(returnResult, result)
	}
	return returnResult, nil
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
	return data
}

func (f FinanceFee) OrderFees(orderId uint, payOrReceive ...string) (map[string][]models.FinanceFee, error) {
	attr := map[string]interface{}{
		"order_master_id": orderId,
	}
	return f.repo.OrderFees(attr, payOrReceive...)
}

func (f FinanceFee) BulkHistoryFee(orderMasterId uint, feeIds []uint, companyId uint, closingUnitId uint, payOrReceive string) ([]models.FinanceFee, error) {
	var (
		financeFees  []models.FinanceFee
		rates        []models.FinanceRate
		orderMasters []models.ResultOrderMaster
		err          error
	)
	financeFees, err = f.FindFeesById(feeIds)
	if err != nil {
		return nil, err
	}
	orderMasters, err = f.getOrderMaster([]uint{orderMasterId})
	if err != nil {
		return nil, err
	}
	if len(orderMasters) == 0 {
		return nil, errors.New("订单不存在")
	}
	order := orderMasters[0]
	financeBase := NewFinanceBase()
	if redis.SystemRateSetting() == models.SettingFeeRateNow {
		rates, err = financeBase.GetAllFeeRate(companyId)
	} else {
		rates, err = f.getRateByOrderCreate(companyId, order.CreatedAt)
	}
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(financeFees); i++ {
		financeFees[i] = f.commonHandlerByCopyFee(rates, financeFees[i], order)
		financeFees[i].ClosingUnitId = closingUnitId
		financeFees[i].PayOrReceive = payOrReceive
		if payOrReceive == "pay" {
			financeFees[i].Payable = financeFees[i].TaxAmount
			financeFees[i].Receivable = 0
		} else {
			financeFees[i].Receivable = financeFees[i].TaxAmount
			financeFees[i].Payable = 0
		}
	}
	return f.BulkInsert(financeFees)
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

func (f FinanceFee) baseDataFindFast(key string, value interface{}) string {
	return red.HGetValue(models.BaseDataCode{}.TableName()+key, value, "name")
}

func NewFinanceFee() IFinanceFee {
	return FinanceFee{repo: repositories.NewFinanceFee()}
}
