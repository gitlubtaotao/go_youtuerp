package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"strings"
	"sync"
	"youtuerp/conf"
	"youtuerp/models"
	"youtuerp/services"
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

func (o *OrderMaster) GetColumn(ctx iris.Context) {
	o.RenderModuleColumn(ctx, models.ResponseOrderMaster{})
}

//前端获取订单数据
func (o *OrderMaster) Get(ctx iris.Context) {
	data, total, err := o.service.FindMasterByIndex(
		o.GetPer(ctx),
		o.GetPage(ctx),
		o.handlerParams(),
		[]string{},
		[]string{},
		o.enum)
	if err != nil {
		o.Render500(ctx, err, "")
		return
	}
	_, _ = ctx.JSON(iris.Map{
		"code":  http.StatusOK,
		"data":  data,
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
		delete(readMap, "status-eq")
	}
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
