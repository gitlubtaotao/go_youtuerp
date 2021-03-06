package controllers

import (
	"errors"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"net/http"
	"sync"
	"time"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/services"
)

type FinanceFee struct {
	BaseController
	service services.IFinanceFee
	ctx     iris.Context
	enum    conf.Enum
}

//复制费用
type CopyFeeReceive struct {
	FeeIds         []uint `json:"fee_ids"`
	OrderMasterIds []uint `json:"order_master_ids"`
	PayOrReceive   string `json:"pay_or_receive"`
	ClosingUnitId  uint   `json:"closing_unit_id"`
}

func (f *FinanceFee) Create(ctx iris.Context) {
	var (
		financeFees []models.FinanceFee
		err         error
	)
	if err = ctx.ReadJSON(&financeFees); err != nil {
		f.Render400(ctx, err, err.Error())
		return
	}
	financeFees, err = f.service.BulkInsert(financeFees)
	if err != nil {
		f.Render500(ctx, err, err.Error())
		return
	}
	golog.Infof("finance fee is %v", financeFees)
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": financeFees})
}

//获取订单对应的费用
func (f *FinanceFee) OrderFees(ctx iris.Context) {
	var (
		id                uint
		err               error
		options           map[string]interface{}
		closingUnitOption []map[string]interface{}
		data              map[string][]models.FinanceFee
		sy                sync.WaitGroup
		sm                sync.Mutex
	)
	currentUser, _ := f.CurrentUser(ctx)
	sy.Add(3)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		f.Render400(ctx, err, err.Error())
		return
	}
	go func(id uint) {
		sm.Lock()
		defer sm.Unlock()
		data, err = f.service.OrderFees(id, "pay", "receive")
		sy.Done()
	}(id)
	go func(companyId uint) {
		sm.Lock()
		defer sm.Unlock()
		options = f.service.OrderFeesOptions(companyId)
		sy.Done()
	}(uint(currentUser.UserCompanyId))
	go func() {
		sm.Lock()
		defer sm.Unlock()
		scope := map[string]interface{}{"company_type": []int{1, 2, 3, 4}}
		closingUnitOption, _ = services.NewSelectService(ctx).GetCompanySelect("", scope, []string{})
		sy.Done()
	}()
	sy.Wait()
	if err != nil {
		f.Render500(ctx, err, "")
		return
	}
	_, _ = ctx.JSON(iris.Map{
		"code":                 http.StatusOK,
		"data":                 data,
		"options":              options,
		"closing_unit_options": closingUnitOption})
}

func (f *FinanceFee) DeleteFee(ctx iris.Context) {
	var (
		deleteIds map[string][]uint
		err       error
	)
	if err = ctx.ReadJSON(&deleteIds); err != nil {
		f.Render400(ctx, err, "")
		return
	}
	if ids, ok := deleteIds["ids"]; ok {
		if err = f.service.DeleteFees(ids); err != nil {
			f.Render500(ctx, err, "")
		} else {
			f.RenderSuccessJson(ctx, iris.Map{})
		}
	}
}

func (f *FinanceFee) ChangeStatus(ctx iris.Context) {
	var (
		temp   map[string][]uint
		err    error
		ids    []uint
		status string
	)
	if err = ctx.ReadJSON(&temp); err != nil {
		f.Render400(ctx, err, "")
		return
	}
	if _, ok := temp["ids"]; !ok {
		err = errors.New(ctx.GetLocale().GetMessage("err.params_error"))
		f.Render400(ctx, err, err.Error())
		return
	}
	ids = temp["ids"]
	status = ctx.URLParam("status")
	if err = f.service.ChangeStatusFees(ids, status); err != nil {
		f.Render400(ctx, err, err.Error())
	} else {
		f.RenderSuccessJson(ctx, iris.Map{})
	}
}

//根据选择的订单，将费用复制不同的订单中
func (f *FinanceFee) CopyFee(ctx iris.Context) {
	var (
		receiveIds CopyFeeReceive
		err        error
	)
	if err = ctx.ReadJSON(&receiveIds); err != nil {
		f.Render400(ctx, err, "")
		return
	}
	currentUser, _ := f.CurrentUser(ctx)
	err = f.service.CopyFee(receiveIds.OrderMasterIds, receiveIds.FeeIds, currentUser.UserCompanyId)
	if err != nil {
		f.Render500(ctx, err, "")
	} else {
		f.RenderSuccessJson(ctx, iris.Map{})
	}
}

//根据不同的结算单位,查询该结算单位下不重复的费用
//同时也可以根据订单号进行查询
func (f *FinanceFee) GetHistoryFee(ctx iris.Context) {
	var (
		searchParams = make(map[string]interface{})
		dateRange    int
		err          error
		cusErr       = errors.New(ctx.GetLocale().GetMessage("error.params_error"))
		financeFees  []map[string]interface{}
	)
	if ok := ctx.URLParamExists("closing_unit_id"); !ok {
		f.Render400(ctx, cusErr, cusErr.Error())
		return
	}
	searchParams["closing_unit_id-eq"] = ctx.URLParam("closing_unit_id")
	if ok := ctx.URLParamExists("date_range"); ok {
		dateRange, err = ctx.URLParamInt("date_range")
		if err != nil {
			f.Render400(ctx, err, "")
			return
		}
	}
	if ok := ctx.URLParamExists("pay_or_receive"); !ok {
		f.Render400(ctx, cusErr, cusErr.Error())
		return
	}
	searchParams["pay_or_receive-eq"] = ctx.URLParam("pay_or_receive")
	searchParams = f.handlerCreatedAt(dateRange, searchParams)
	financeFees, err = f.service.GetHistoryFee(searchParams)
	if err != nil {
		f.Render500(ctx, err, "")
	} else {
		_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": financeFees})
	}
}

//根据不同的订单插入历史费用
func (f *FinanceFee) BulkHistoryFee(ctx iris.Context) {
	var (
		orderMasterId uint
		err           error
		result        CopyFeeReceive
	)
	orderMasterId, err = ctx.Params().GetUint("orderMasterId")
	if err != nil {
		f.Render400(ctx, err, "")
		return
	}
	if err = ctx.ReadJSON(&result); err != nil {
		f.Render400(ctx, err, "")
		return
	}
	currentUser, _ := f.CurrentUser(ctx)
	if financeFees, err := f.service.BulkHistoryFee(
		orderMasterId,
		result.FeeIds,
		uint(currentUser.UserCompanyId),
		result.ClosingUnitId,
		result.PayOrReceive); err != nil {
		f.Render500(ctx, err, "")
	} else {
		_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": financeFees})
	}
}

//获取费用对应的列设置
func (f *FinanceFee) GetColumn(ctx iris.Context) {
	f.RenderModuleColumn(ctx, models.FinanceFee{})
}

//获取对账列表对应的费用信息
func (f *FinanceFee) GetConfirmBillList(ctx iris.Context) {
	var (
		readMap     map[string]interface{}
		err         error
		total       uint
		financeFees []models.ResultFinanceFee
	)
	if err = ctx.ReadJSON(&readMap); err != nil {
		f.Render400(ctx, err, "")
		return
	}
	for _, item := range []string{
		"finance_fees.created_at",
		"order_extend_infos.departure",
		"order_extend_infos.arrival",
		"order_masters.created_at",} {
		f.HandlerFilterDate(readMap, item)
	}
	financeFees, total, err = f.service.FindFinanceFees(f.GetPer(ctx), f.GetPage(ctx), readMap, []string{}, []string{"finance_fees.id asc"})
	if err != nil {
		f.Render400(ctx, err, "")
		return
	}
	dataArray := make([]map[string]interface{}, 0)
	for _, v := range financeFees {
		dataArray = append(dataArray, f.service.HandleFeesShow(v, f.enum))
	}
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": dataArray, "total": total})
}

//处理历史费用订单创建时间
func (f *FinanceFee) handlerCreatedAt(dateRange int, searchParams map[string]interface{}) map[string]interface{} {
	beg := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
	searchParams["created_at-ltEq"] = beg
	searchParams["created_at-gtEq"] = beg.AddDate(0, -dateRange, 0)
	return searchParams
}
func (f *FinanceFee) Before(ctx iris.Context) {
	f.service = services.NewFinanceFee()
	f.enum = conf.Enum{Locale: ctx.GetLocale()}
	f.ctx = ctx
	ctx.Next()
}
