package services

import (
	"github.com/kataras/iris/v12"
	"youtuerp/repositories"
)

type ISelectService interface {
	FindModel(model interface{}, scope map[string]interface{}, selectKey []string) (selectResult []map[string]interface{}, err error)
}

type SelectService struct {
	repo repositories.ISelectRepository
	ctx  iris.Context
}

func (s *SelectService) FindModel(model interface{}, scope map[string]interface{}, selectKey []string) (selectResult []map[string]interface{}, err error) {
	result, err := s.repo.FindModel(model, scope, selectKey)
	if err != nil {
		return nil, err
	}
	columnService := NewColumnService(s.ctx.GetLocale())
	for _, v := range result {
		temp, _ := columnService.StructToMap(v)
		temp = s.afterHandler(temp)
		selectResult = append(selectResult, temp)
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
