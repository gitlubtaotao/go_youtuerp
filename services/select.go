package services

import (
	"github.com/kataras/iris/v12"
	"youtuerp/repositories"
)

type ISelectService interface {
	FindTable(tableName string, name string, scope map[string]interface{}, selectKeys []string) (selectResult []map[string]interface{}, err error)
	FindModel(model interface{}, scope map[string]interface{}, selectKey []string) (selectResult []map[string]interface{}, err error)
}

type SelectService struct {
	repo repositories.ISelectRepository
	ctx  iris.Context
}


func (s *SelectService) FindTable(tableName string, name string, scope map[string]interface{}, selectKeys []string) ([]map[string]interface{}, error) {
	selectResult := make([]map[string]interface{},0)
	result, err := s.repo.FindTable(tableName, name, scope, selectKeys)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	columnService := NewColumnService(s.ctx.GetLocale())
	for _, v := range result {
		src, _ := columnService.StructToMap(v)
		dst := s.afterHandler(src)
		selectResult = append(selectResult, dst)
	}
	return selectResult, nil
}

func (s *SelectService) FindModel(model interface{}, scope map[string]interface{}, selectKey []string) (selectResult []map[string]interface{}, err error) {
	result, err := s.repo.FindModel(model, scope, selectKey)
	if err != nil {
		return nil, err
	}
	columnService := NewColumnService(s.ctx.GetLocale())
	for _, v := range result {
		src, _ := columnService.StructToMap(v)
		dst := s.afterHandler(src)
		selectResult = append(selectResult, dst)
	}
	return selectResult, nil
}

func (s *SelectService) afterHandler(dest map[string]interface{}) (out map[string]interface{}) {
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
	return &SelectService{
		repo: repositories.NewSelectRepository(),
		ctx:  ctx,
	}
}
