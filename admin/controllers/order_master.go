package controllers

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"net/http"
	"strconv"
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
	o.RenderModuleColumn(ctx, models.ResultOrderMaster{})
}

func (o *OrderMaster) Get(ctx iris.Context) {
	records, total, err := o.service.FindMaster(
		o.GetPer(ctx),
		o.GetPage(ctx),
		o.handlerParams(),
		[]string{},
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
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": order})
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

//修改的订单的状态
func (o *OrderMaster) ChangeStatus(ctx iris.Context) {
	var (
		id     uint
		status string
		err    error
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		o.Render400(ctx, err, "")
		return
	}
	if !ctx.URLParamExists("status") {
		o.Render400(ctx, nil, "")
		return
	}
	status = ctx.URLParam("status")
	if err = o.service.ChangeStatus(id, status); err != nil {
		o.Render400(ctx, err, err.Error())
	} else {
		o.RenderSuccessJson(ctx, iris.Map{})
	}
}

func (o *OrderMaster) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetUint("id")
	if err != nil {
		o.Render400(ctx, err, "")
		return
	}
	if err = o.service.DeleteMaster(id); err != nil {
		o.Render500(ctx, err, "")
	} else {
		o.RenderSuccessJson(ctx, iris.Map{})
	}
}

func (o *OrderMaster) Operation(ctx iris.Context) {
	var (
		id            uint
		err           error
		order         models.OrderMaster
		data          interface{}
		sy            sync.WaitGroup
		mx            sync.Mutex
		selectOptions map[string]interface{}
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		o.Render400(ctx, err, "")
		return
	}
	order, err = o.service.FirstMaster(id, "OrderExtendInfo", "Roles")
	if err != nil {
		o.Render400(ctx, err, "")
		return
	}
	sy.Add(2)
	go func(id uint) {
		mx.Lock()
		defer mx.Unlock()
		data, err = o.getOperationFormerData(order)
		sy.Done()
	}(id)
	go func() {
		mx.Lock()
		defer mx.Unlock()
		selectService := services.NewSelectService(ctx)
		selectOptions = selectService.GetOperationSelect(order.TransportType)
		sy.Done()
	}()
	sy.Wait()
	if err != nil {
		o.Render500(ctx, err, "")
		return
	}
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "formerData": data, "order": o.handlerOrderInfo(order), "selectOptions": selectOptions})
}

// 获取表单的数据
func (o *OrderMaster) GetFormerData(ctx iris.Context) {
	var (
		formerType    string
		orderMasterId uint
		err           error
		formerData    interface{}
	)
	if orderMasterId, err = ctx.Params().GetUint("id"); err != nil {
		o.Render400(ctx, err, err.Error())
		return
	}
	golog.Infof("order master id is %v", orderMasterId)
	formerType = ctx.URLParam("formerType")
	formerData, err = o.service.GetFormerData(orderMasterId, formerType)
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "formerData": formerData})
}

//获取SoNo 下拉信息
func (o *OrderMaster) GetSoNoOptions(ctx iris.Context) {
	var (
		id     uint
		err    error
		result []string
	)
	if id, err = ctx.Params().GetUint("id"); err != nil {
		o.Render400(ctx, err, err.Error())
		return
	}
	server := services.NewFormerServer()
	result, err = server.GetFormerSoNoOptions(id, "1")
	if err != nil {
		o.Render500(ctx, err, "")
	} else {
		o.RenderSuccessJson(ctx, iris.Map{"so_no": result})
	}
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
	data["cut_off_day"] = toolTime.InterfaceFormat(data["cut_off_day"], "zh-CN")
	data["departure"] = toolTime.InterfaceFormat(data["departure"], "zh-CN")
	data["arrival"] = toolTime.InterfaceFormat(data["arrival"], "zh-CN")
	data["instruction_id_value"] = data["instruction_id"]
	data["instruction_id"] = red.HGetCrm(data["instruction_id"], "")
	data["supply_agent_id"] = red.HGetCrm(data["supply_agent_id"], "")
	data["salesman_id_value"] = data["salesman_id"]
	data["salesman_id"] = red.HGetRecord("users", data["salesman_id"], "")
	data["operation_id_value"] = data["operation_id"]
	data["operation_id"] = red.HGetRecord("users", data["operation_id"], "")
	data["pol_id"] = o.showPort(data["transport_type"], data["pol_id"])
	data["pod_id"] = o.showPort(data["transport_type"], data["pod_id"])
	data["carrier_id"] = o.showCarrier(data["transport_type"], data["carrier_id"])
	data["transport_type_value"] = data["transport_type"]
	data["transport_type"] = o.service.ShowTransport(o.enum, order)
	data["status_value"] = data["status"]
	data["status"] = o.service.ShowStatus(o.enum, order.Status)
	data["company_id"] = red.HGetCompany(data["company_id"], "")
	for _, v := range []string{"paid_status", "received_status", "payable_status", "receivable_status"} {
		data[v+"_value"] = data[v]
		data[v] = o.service.ShowFinanceStatus(o.enum, v, data[v])
	}
	return data
}

// 处理前端查询字段
func (o *OrderMaster) handlerParams() map[string]interface{} {
	var readMap map[string]interface{}
	err := o.ctx.ReadJSON(&readMap)
	if err != nil {
		return readMap
	}
	for _, item := range []string{"created_at", "cut_off_day", "departure", "arrival"} {
		o.HandlerFilterDate(readMap, item)
	}
	if status := readMap["status-eq"]; status == "" {
		readMap["status-notEq"] = models.OrderStatusCancel
		delete(readMap,"status-eq")
	}
	golog.Infof("map is %v", readMap)
	return readMap
}

// 处理查询时间问题
func (o *OrderMaster) HandlerFilterDate(filters map[string]interface{}, field string) {
	timeField, ok := filters[field]
	if !ok {
		return
	}
	stringTime := timeField.(string)
	timeArray := strings.Split(stringTime, ",")
	if len(timeArray) == 2 {
		filters[field+"-gtEq"] = o.stringToDate(timeArray[0])
		filters[field+"-ltEq"] = o.stringToDate(timeArray[1])
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

//处理订单操作信息
func (o *OrderMaster) handlerOrderInfo(order models.OrderMaster) map[string]interface{} {
	data, _ := o.StructToMap(order, o.ctx)
	data["transport_type_text"] = o.service.ShowTransport(o.enum, order)
	data["status_text"] = o.service.ShowStatus(o.enum, order.Status)
	data["instruction_name"] = red.HGetCrm(order.InstructionId, "")
	data["operation_name"] = red.HGetValue("users", order.OperationId, "")
	data["salesman_name"] = red.HGetValue("users", order.SalesmanId, "")
	return data
}

//操作盘处理表单信息
func (o *OrderMaster) getOperationFormerData(order models.OrderMaster) (interface{}, error) {
	transportType := order.TransportType
	switch transportType {
	case 1, 4:
		return o.service.GetSeaFormerInstruction(order, models.InstructionMaster)
	case 3:
		return nil, nil
	}
	return nil, nil
}

func (o *OrderMaster) showPort(transportType interface{}, portId interface{}) string {
	tableName := models.BaseDataPort{}.TableName()
	uTransportType := transportType.(uint)
	var key string
	switch uTransportType {
	case 1, 4:
		key = strconv.Itoa(models.BaseTypeSea)
	case 2:
		key = strconv.Itoa(models.BaseTypeAir)
	}
	return red.HGetValue(tableName+key, portId, "name")
}

func (o *OrderMaster) showCarrier(transportType interface{}, carrierId interface{}) string {
	tableName := models.BaseDataCarrier{}.TableName()
	uTransportType := transportType.(uint)
	var key string
	switch uTransportType {
	case 1, 4:
		key = strconv.Itoa(models.BaseTypeSea)
	case 2:
		key = strconv.Itoa(models.BaseTypeAir)
	}
	return red.HGetValue(tableName+key, carrierId, "name")
}
