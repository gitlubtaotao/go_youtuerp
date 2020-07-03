package services

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"reflect"
	"sync"
	"youtuerp/models"
	"youtuerp/repositories"
)

type ISelectService interface {
	//获取公司信息下拉选择
	GetCompanySelect(searchName string, scope map[string]interface{}, selectKeys []string)([]map[string]interface{}, error)
	//获取操作盘下拉选择
	GetOperationSelect(formerType string) map[string]interface{}
	//根据不同的table查询数据
	FindTable(tableName string, name string, scope map[string]interface{}, selectKeys []string) (selectResult []map[string]interface{}, err error)
	
}

type SelectService struct {
	repo repositories.ISelectRepository
	ctx  iris.Context
	sy   sync.Mutex
}


func (s SelectService) GetCompanySelect(searchName string, scope map[string]interface{}, selectKeys []string)([]map[string]interface{}, error) {
	if len(scope) == 0 {
		scope = map[string]interface{}{"company_type": 4}
	}
	scope["status"] = models.CompanyStatusApproved
	if len(selectKeys) == 0 {
		selectKeys = []string{"id", "name_en", "name_nick", "name_cn", "frequently_use_info"}
	}
	return s.FindTable("user_companies",searchName,scope,selectKeys)
}


func (s SelectService) GetOperationSelect(formerType string) map[string]interface{} {
	returnAttr := make(map[string]interface{})
	crmOptions, _ := s.FindTable("user_companies", "", map[string]interface{}{"company_type": []int{1, 3}, "status": models.CompanyStatusApproved}, []string{"id", "name_cn","name_nick","name_en"})
	returnAttr["crmOptions"] = crmOptions
	stringArray := []string{models.CodePayType, models.CodeCapType, models.CodeInstructionType,
		models.CodeCustomType, models.CodeBillProduceType, models.CodeTransshipment,
		models.CodeTradeTerms, models.CodeShippingTerms, models.PackageType}
	codeService := NewBaseCode()
	for i := 0; i < len(stringArray); i++ {
		returnAttr[stringArray[i]] = codeService.FindCollect(stringArray[i])
	}
	carrierService := NewBaseCarrier()
	portService := NewBasePort()
	returnAttr["carrier"] = carrierService.FindCollect("1")
	returnAttr["port"] = portService.FindCollect("1")
	returnAttr["userInfo"] = NewEmployeeService().FindRedis()
	return returnAttr
}

func (s SelectService) FindTable(tableName string, name string, scope map[string]interface{}, selectKeys []string) ([]map[string]interface{}, error) {
	selectResult := make([]map[string]interface{}, 0)
	scopeResult := s.handleScope(scope)
	result, err := s.repo.FindTable(tableName, name, scopeResult, selectKeys)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	columnService := NewColumnService(s.ctx.GetLocale())
	for _, v := range result {
		src, _ := columnService.StructToMap(v)
		dst := s.handleResult(src)
		selectResult = append(selectResult, dst)
	}
	return selectResult, nil
}

func (s SelectService) handleScope(scope map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, 0)
	s.sy.Lock()
	defer s.sy.Unlock()
	for k, v := range scope {
		golog.Infof("k is %v,v is %v", k, v)
		var search string
		ty := reflect.TypeOf(v)
		switch ty.Kind() {
		case reflect.Slice, reflect.Array:
			search = k + "-in"
		default:
			search = k + "-eq"
		}
		dst[search] = v
	}
	return dst
}

func (s SelectService) handleResult(dest map[string]interface{}) (out map[string]interface{}) {
	out = make(map[string]interface{})
	out["value"] = dest["id"]
	if value := dest["name"]; value.(string) != "" {
		out["label"] = dest["name"]
		return out
	}
	if value := dest["serial_number"]; value.(string) != "" {
		out["label"] = dest["serial_number"]
		return out
	}
	if s.ctx.GetLocale().Language() == "en" {
		out["label"] = dest["name_en"]
		return
	}
	if value := dest["name_nick"]; value.(string) != "" {
		out["label"] = dest["name_nick"]
	} else {
		out["label"] = dest["name_cn"]
	}
	return
}

func NewSelectService(ctx iris.Context) ISelectService {
	return &SelectService{repo: repositories.NewSelectRepository(), ctx: ctx,}
}
