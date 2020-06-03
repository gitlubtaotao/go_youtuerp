package controllers

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"net/http"
	"strings"
	"sync"
	"time"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/redis"
	"youtuerp/services"
	"youtuerp/tools"
)

type OrderMaster struct {
	BaseController
	ctx             iris.Context
	service         services.IOrderMasterService
	companyService  services.ICompanyService
	employeeService services.IEmployeeService
	mu              sync.Mutex
	enum            conf.Enum
}

var (
	tool     = tools.OtherHelper{}
	toolTime = tools.TimeHelper{}
	red      = redis.Redis{}
)

func (o *OrderMaster) GetColumn(ctx iris.Context) {
	o.RenderModuleColumn(ctx, models.OrderMaster{})
}

func (o *OrderMaster) Get(ctx iris.Context) {
	o.handlerParams()
	records, total, err := o.service.FindMaster(o.GetPer(ctx), o.GetPage(ctx), o.handlerParams(), []string{},
		[]string{})
	if err != nil {
		o.Render500(ctx, err, "")
		return
	}
	dataArray := make([]map[string]interface{}, 0, len(records))
	o.mu.Lock()
	defer o.mu.Unlock()
	for _, record := range records {
		dataArray = append(dataArray, o.handlerData(record))
	}
	_, _ = ctx.JSON(iris.Map{
		"code":  http.StatusOK,
		"data":  dataArray,
		"total": total,
	})
}

func (o *OrderMaster) Create(ctx iris.Context) {
	var (
		order models.OrderMaster
		err   error
	)
	if err = ctx.ReadJSON(&order); err != nil {
		o.Render400(ctx, err, err.Error())
		return
	}
	currentUser, _ := o.CurrentUser(ctx)
	if order.OperationId == 0 {
		order.OperationId = currentUser.ID
	}
	order.CompanyId = uint(currentUser.UserCompanyId)
	order, err = o.service.CreateMaster(order, ctx.GetLocale().Language())
	if err != nil {
		o.Render400(ctx, err, err.Error())
		return
	}
	o.RenderSuccessJson(ctx, iris.Map{})
}

func (o *OrderMaster) Edit(ctx iris.Context) {
	var (
		id    uint
		err   error
		order interface{}
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		o.Render400(ctx, err, "")
		return
	}
	if order, err = o.service.FirstMaster(id, "Roles"); err != nil {
		o.Render500(ctx, err, "")
		return
	}
	o.RenderSuccessJson(ctx, order)
}

func (o *OrderMaster) Update(ctx iris.Context) {
	var (
		id    uint
		err   error
		order models.OrderMaster
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		o.Render400(ctx, err, "")
		return
	}
	if err = ctx.ReadJSON(&order); err != nil {
		o.Render400(ctx, err, "")
		return
	}
	if err = o.service.UpdateMaster(id, order, ctx.GetLocale().Language()); err != nil {
		o.Render400(ctx, err, "")
		return
	}
	order.ID = id
	o.RenderSuccessJson(ctx, order)
}

func (o *OrderMaster) Before(ctx iris.Context) {
	o.ctx = ctx
	o.service = services.NewOrderMasterService()
	o.companyService = services.NewCompanyService()
	o.employeeService = services.NewEmployeeService()
	o.enum = conf.NewEnum(ctx.GetLocale())
	ctx.Next()
}

// 对输出的字段进行处理
func (o *OrderMaster) handlerData(order models.ResultOrderMaster) map[string]interface{} {
	data := tool.StructToMap(order)
	var sy sync.WaitGroup
	var mx sync.Mutex
	sy.Add(3)
	go func() {
		mx.Lock()
		defer mx.Unlock()
		data["order_extend_infos_cut_off_day"] = toolTime.InterfaceFormat(data["cut_off_day"], "zh-CN")
		data["order_extend_infos_departure"] = toolTime.InterfaceFormat(data["departure"], "zh-CN")
		data["order_extend_infos_arrival"] = toolTime.InterfaceFormat(data["arrival"], "zh-CN")
		sy.Done()
	}()
	go func() {
		mx.Lock()
		defer mx.Unlock()
		data["instruction_id_value"] = data["instruction_id"]
		data["instruction_id"] = red.HGetCrm(data["instruction_id"], "")
		data["salesman_id_value"] = data["salesman_id"]
		data["salesman_id"] = red.HGetRecord("users", data["salesman_id"], "")
		data["operation_id_value"] = data["operation_id"]
		data["operation_id"] = red.HGetRecord("users", data["operation_id"], "")
		sy.Done()
	}()
	go func() {
		mx.Lock()
		defer mx.Unlock()
		data["transport_type_value"] = data["transport_type"]
		data["transport_type"] = o.service.ShowTransport(o.enum, order)
		data["status_value"] = data["status"]
		data["status"] = o.service.ShowStatus(o.enum, order)
		data["company_id"] = red.HGetCompany(data["company_id"], "")
		for _, v := range []string{"paid_status", "received_status", "payable_status", "receivable_status"} {
			data[v+"_value"] = data[v]
			data[v] = o.service.ShowFinanceStatus(o.enum, v, data[v])
		}
		sy.Done()
	}()
	sy.Wait()
	return data
}

// 处理前端查询字段
func (o *OrderMaster) handlerParams() map[string]interface{} {
	filters := make(map[string]interface{}, 10)
	params := o.ctx
	var sy sync.WaitGroup
	var mu sync.Mutex
	sy.Add(4)
	go func() {
		intArray := []string{"transport_type", "instruction_id", "supply_agent_id",
			"company_id", "salesman_id", "operation_id", "carrier_id", "pol_id",
			"pod_id", "courier_code_id"}
		mu.Lock()
		defer mu.Unlock()
		for i := 0; i < len(intArray); i++ {
			filters[intArray[i]+"-eq"] = params.URLParamIntDefault(intArray[i], 0)
		}
		sy.Done()
	}()
	go func() {
		stringArray := []string{"serial_number", "mbl_so", "so_no", "flight_no", "courier_no"}
		mu.Lock()
		defer mu.Unlock()
		for i := 0; i < len(stringArray); i++ {
			filters[stringArray[i]+"-rCount"] = params.URLParamDefault(stringArray[i], "")
		}
		sy.Done()
	}()
	//
	go func() {
		eqArray := []string{"paid_status", "payable_status",
			"receivable_status", "received_status"}
		mu.Lock()
		defer mu.Unlock()
		for i := 0; i < len(eqArray); i++ {
			filters[eqArray[i]+"-eq"] = params.URLParamDefault(eqArray[i], "")
		}
		sy.Done()
	}()
	go func() {
		mu.Lock()
		defer mu.Unlock()
		for _, item := range []string{"created_at", "cut_off_day", "departure", "arrival"} {
			o.HandlerFilterDate(filters, item)
		}
		sy.Done()
	}()
	sy.Wait()
	status := params.URLParamDefault("status", "")
	if status != "" {
		filters["status-eq"] = status
	} else {
		filters["status-in"] = []string{models.OrderStatusPro, models.OrderStatusLocked, models.OrderStatusFinished}
	}
	return filters
}

// 处理查询时间问题
func (o *OrderMaster) HandlerFilterDate(filters map[string]interface{}, field string) {
	timeField := strings.Split(o.ctx.URLParamDefault(field, ""), ",")
	if len(timeField) == 2 {
		filters[field+"-gtEq"] = o.stringToDate(timeField[0])
		filters[field+"-ltEq"] = o.stringToDate(timeField[1])
	}
}

// 将string转化成日期格式
func (o *OrderMaster) stringToDate(strTime string) time.Time {
	result, err := tools.TimeHelper{}.StringToTime(strTime)
	if err != nil {
		golog.Errorf("string to date is error %v", err)
	}
	return result
}
