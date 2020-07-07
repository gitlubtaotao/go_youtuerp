package services

import (
	"github.com/kataras/golog"
	"strings"
	"sync"
	"youtuerp/models"
	"youtuerp/repositories"
)

type IFormerServer interface {
	//删除其他综合服务对应的表单
	DeleteOtherServer(id uint, formerType string) error
	//保存其他综合服务的对应的表单
	SaveOtherServer(formerType string, data models.RenderFormerData) (id uint, err error)
	//删除货物详情
	DeleteCargoInfo(ids []int, formerType string) error
	//保存海运货物详情
	UpdateCargoInfo(id uint, formerType string, readData models.RenderFormerData) (data interface{}, err error)
	//保存操作盘对应的信息
	UpdateOperationInfo(id uint, formerType string, readData models.RenderFormerData) error
	//获取其他服务的下来选择
	GetOtherServerOptions(orderMasterId uint, transportType string) (map[string]interface{}, error)
	//获取订单对应的拖车单,报告单,仓库场装单
	GetOtherServer(orderMasterId uint, transportType string) (map[string]interface{}, error)
	//获取so信息的下拉选择
	GetFormerSoNoOptions(orderMasterId uint, transportType string) ([]string, error)
}
type FormerServer struct {
	repo repositories.IFormerServer
}

func (f FormerServer) DeleteOtherServer(id uint, formerType string) error {
	return f.repo.DeleteOtherServer(id, formerType)
}

func (f FormerServer) SaveOtherServer(formerType string, data models.RenderFormerData) (id uint, err error) {
	return f.repo.SaveOtherServer(formerType, data)
}

func (f FormerServer) DeleteCargoInfo(ids []int, formerType string) error {
	return f.repo.DeleteCargoInfo(ids, formerType)
}

func (f FormerServer) UpdateCargoInfo(id uint, formerType string, readData models.RenderFormerData) (data interface{}, err error) {
	if formerType == "sea_cargo_info" {
		return f.repo.UpdateSeaCargoInfo(readData.SeaCargoInfo)
	}
	return nil, nil
}

func (f FormerServer) UpdateOperationInfo(id uint, formerType string, readData models.RenderFormerData) error {
	extendInfo := readData.OrderExtendInfo
	var err error
	if extendInfo.ID != 0 {
		err = f.repo.UpdateExtendInfo(extendInfo.ID, extendInfo)
	}
	if err != nil {
		return err
	}
	return f.repo.UpdateFormerData(formerType, readData)
}

func (f FormerServer) GetOtherServerOptions(orderMasterId uint, transportType string) (map[string]interface{}, error) {
	var resultMap = make(map[string]interface{})
	portServer := NewBasePort()
	resultMap["portOptions"] = portServer.FindCollect("1")
	baseCodeServer := NewBaseCode()
	resultMap["packageOptions"] = baseCodeServer.FindCollect(models.PackageType)
	resultMap["capTypeOptions"] = baseCodeServer.FindCollect(models.CodeCapType)
	resultMap["transTypeOptions"] = baseCodeServer.FindCollect(models.CodeTransType)
	resultMap["customTypeOptions"] = baseCodeServer.FindCollect(models.CodeCustomType)
	resultMap["currencyOptions"] = baseCodeServer.FindCollect(models.CodeFinanceCurrency)
	return resultMap, nil
}

func (f FormerServer) GetFormerSoNoOptions(orderMasterId uint, transportType string) ([]string, error) {
	server := NewOrderMasterService()
	var (
		data         interface{}
		err          error
		returnResult []string
	)
	if transportType == "1" {
		data, err = server.GetFormerData(orderMasterId, "former_sea_so_no", "")
		if err != nil {
			return []string{}, err
		}
		formerSoNo := data.(models.FormerSeaSoNo)
		if formerSoNo.SoNo == "" {
			returnResult = []string{}
		} else {
			returnResult = strings.Split(formerSoNo.SoNo, ",")
		}
	}
	return returnResult, err
}

func (f FormerServer) GetOtherServer(orderMasterId uint, transportType string) (map[string]interface{}, error) {
	var (
		formerOtherServers      []models.FormerOtherService
		formerTrailerOrders     []models.FormerTrailerOrder
		formerWarehouseServices []models.FormerWarehouseService
		formerCustomClearance   []models.FormerCustomClearance
		err                     error
		sy                      sync.Mutex
		sw                      sync.WaitGroup
	)
	sw.Add(4)
	go func() {
		sy.Lock()
		defer sy.Unlock()
		formerOtherServers, err = f.repo.GetFormerOtherService(orderMasterId)
		if err != nil {
			golog.Errorf("GetOtherServer is err %v ", err)
		}
		sw.Done()
	}()
	go func() {
		sy.Lock()
		defer sy.Unlock()
		formerTrailerOrders, err = f.repo.GetFormerTrailerOrder(orderMasterId)
		if err != nil {
			golog.Errorf("GetOtherServer is err %v ", err)
		}
		sw.Done()
	}()
	go func() {
		sy.Lock()
		defer sy.Unlock()
		formerWarehouseServices, err = f.repo.GetFormerWarehouseService(orderMasterId)
		if err != nil {
			golog.Errorf("GetOtherServer is err %v ", err)
		}
		sw.Done()
	}()
	go func() {
		sy.Lock()
		defer sy.Unlock()
		formerCustomClearance, err = f.repo.GetFormerCustomClearance(orderMasterId)
		if err != nil {
			golog.Errorf("GetOtherServer is err %v ", err)
		}
		sw.Done()
	}()
	sw.Wait()
	return map[string]interface{}{
		"formerTrailerOrders":     formerTrailerOrders,
		"formerOtherServers":      formerOtherServers,
		"formerWarehouseServices": formerWarehouseServices,
		"formerCustomClearances":  formerCustomClearance,
	}, nil
}

func NewFormerServer() IFormerServer {
	return FormerServer{repo: repositories.NewFormerServerRepository()}
}
