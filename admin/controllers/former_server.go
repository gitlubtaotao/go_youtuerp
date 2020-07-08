package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"sync"
	"youtuerp/models"
	"youtuerp/services"
)

type FormerServer struct {
	BaseController
	service services.IFormerServer
}

//根据订单类型获取不同的综合服务
func (f *FormerServer) GetOtherServer(ctx iris.Context) {
	var (
		id            uint
		err           error
		formerData    map[string]interface{}
		selectOptions map[string]interface{}
		crmOptions    map[string]interface{}
		sm            sync.Mutex
		sw            sync.WaitGroup
	)
	id, err = ctx.Params().GetUint("id")
	if err != nil {
		f.Render400(ctx, err, "")
		return
	}
	sw.Add(3)
	go func(id uint) {
		sm.Lock()
		defer sm.Unlock()
		formerData, _ = f.service.GetOtherServer(id)
		sw.Done()
	}(id)
	go func() {
		sm.Lock()
		defer sm.Unlock()
		selectOptions, _ = f.service.GetOtherServerOptions()
		sw.Done()
	}()
	go func() {
		sm.Lock()
		defer sm.Unlock()
		crmOptions, _ = f.getCrmOptions(ctx)
		sw.Done()
	}()
	sw.Wait()
	_, _ = ctx.JSON(iris.Map{
		"code":          http.StatusOK,
		"formerData":    formerData,
		"selectOptions": selectOptions,
		"crmOptions":    crmOptions})
}

//保存其他综合服务对应的内容,如果Id不存在则进行创建,Id存在则进行保存
func (f *FormerServer) SaveOtherServer(ctx iris.Context) {
	var (
		data       models.RenderFormerData
		id         uint
		err        error
		formerType string
	)
	if err := ctx.ReadJSON(&data); err != nil {
		f.Render400(ctx, err, "")
		return
	}
	if ok := ctx.URLParamExists("former_type"); !ok {
		f.Render400(ctx, nil, "")
		return
	}
	formerType = ctx.URLParam("former_type")
	id, err = f.service.SaveOtherServer(formerType, data)
	if err != nil {
		f.Render500(ctx, err, "")
	} else {
		_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "id": id})
	}
}

// 删除综合服务
func (f *FormerServer) DeleteOtherServer(ctx iris.Context) {
	var (
		id         uint
		err        error
		formerType string
	)
	if ok := ctx.URLParamExists("former_type"); !ok {
		f.Render400(ctx, nil, "")
		return
	}
	formerType = ctx.URLParam("former_type")
	if id, err = ctx.Params().GetUint("id"); err != nil {
		f.Render400(ctx, nil, "")
		return
	}
	if err = f.service.DeleteOtherServer(id, formerType); err != nil {
		f.Render500(ctx, err, "")
	} else {
		f.RenderSuccessJson(ctx, iris.Map{})
	}
}

//TODO-Tao 需要解决只更新修改的内容，未修改的内容无序进行更新
// 现有的解决方案是所有的字段进行更新
func (f *FormerServer) UpdateFormerData(ctx iris.Context) {
	var (
		formerType string
		id         uint
		err        error
		params     models.RenderFormerData
	)
	_ = ctx.ReadJSON(&params)
	formerType = ctx.URLParam("former_type")
	id, err = ctx.Params().GetUint("id")
	if err != nil {
		f.Render400(ctx, err, err.Error())
		return
	}
	if err = f.service.UpdateOperationInfo(id, formerType, params); err != nil {
		f.Render500(ctx, err, "")
	} else {
		f.RenderSuccessJson(ctx, iris.Map{})
	}
}

func (f *FormerServer) DeleteCargoInfo(ctx iris.Context) {
	formerType := ctx.URLParam("former_type")
	var deleteIds map[string][]int
	if err := ctx.ReadJSON(&deleteIds); err != nil {
		f.Render400(ctx, err, err.Error())
		return
	}
	err := f.service.DeleteCargoInfo(deleteIds["ids"], formerType)
	if err != nil {
		f.Render500(ctx, err, "")
	} else {
		f.RenderSuccessJson(ctx, iris.Map{})
	}
}

func (f *FormerServer) UpdateCargoInfo(ctx iris.Context) {
	var (
		id         uint
		err        error
		params     models.RenderFormerData
		formerType string
	)
	_ = ctx.ReadJSON(&params)
	formerType = ctx.URLParam("former_type")
	id, err = ctx.Params().GetUint("id")
	if err != nil {
		f.Render400(ctx, err, err.Error())
		return
	}
	data, _ := f.service.UpdateCargoInfo(id, formerType, params)
	_, _ = ctx.JSON(iris.Map{"code": http.StatusOK, "data": data})
}

//获取供应商和客户信息
func (f *FormerServer) getCrmOptions(ctx iris.Context) (map[string]interface{}, error) {
	var returnMap = make(map[string]interface{})
	selectServer := services.NewSelectService(ctx)
	if customerOptions, err := selectServer.GetCompanySelect("", map[string]interface{}{"company_type": []int{1, 3}}, []string{}); err != nil {
		return returnMap, err
	} else {
		returnMap["customerOptions"] = customerOptions
	}
	if supplyOptions, err := selectServer.GetCompanySelect("", map[string]interface{}{"company_type": []int{2, 3}}, []string{}); err != nil {
		return returnMap, err
	} else {
		returnMap["supplyOptions"] = supplyOptions
	}
	return returnMap, nil
}

//
func (f *FormerServer) Before(ctx iris.Context) {
	f.service = services.NewFormerServer()
	ctx.Next()
}
