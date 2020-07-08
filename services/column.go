package services

import (
	"errors"
	"github.com/kataras/iris/v12/context"
	"reflect"
	"sync"
	"youtuerp/tools"
)

type IColumnService interface {
	StructColumn(model interface{}, args ...interface{}) (data []interface{}, err error)
	StructToMap(currentObject interface{}) (map[string]interface{}, error)
}

type ColumnService struct {
	sy     sync.Mutex
	loader context.Locale
	BaseService
}

func (c *ColumnService) StructToMap(currentObject interface{}) (map[string]interface{}, error) {
	if currentObject == nil {
		return map[string]interface{}{}, errors.New(c.loader.GetMessage("error.params_error"))
	}
	return tools.OtherHelper{}.StructToMap(currentObject), nil
}

func (c *ColumnService) StructColumn(model interface{}, args ...interface{}) (dataArray []interface{}, err error) {
	t := reflect.TypeOf(model)
	v := reflect.ValueOf(model)
	hiddenColumn := c.defaultHiddenColumn(v)
	if t.Kind() != reflect.Struct {
		err = errors.New("model is not struct")
		return
	}
	c.sy.Lock()
	defer c.sy.Unlock()
	tableName := c.tableName(v)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Type.Kind() == reflect.Struct && f.Type.Name() != "Time" {
			continue
		}
		data := f.Tag.Get("json")
		if data == "" {
			continue
		}
		if c.isHiddenColumn(hiddenColumn, data) {
			title := c.loader.GetMessage(tableName + "." + data)
			if title == "" {
				continue
			}
			attr := map[string]interface{}{
				"data":  data,
				"title": title,
			}
			dataArray = append(dataArray, attr)
		}
	}
	dataArray = append(dataArray, c.structAddColumn(v, tableName)...)
	return
}

func (c *ColumnService) defaultHiddenColumn(v reflect.Value, args ...interface{}) (data interface{}) {
	methodName := v.MethodByName("DefaultHiddenColumn")
	if methodName.IsValid() {
		inputs := make([]reflect.Value, len(args))
		for i, _ := range args {
			inputs[i] = reflect.ValueOf(args[i])
		}
		returnValue := methodName.Call(inputs)
		data = returnValue[0].Interface()
	} else {
		data = []string{}
	}
	return
}
func (c *ColumnService) structAddColumn(v reflect.Value, tableName string) []interface{} {
	methodName := v.MethodByName("DefaultAddColumn")
	if !methodName.IsValid() {
		return nil
	}
	inputs := make([]reflect.Value, 0)
	returnValue := methodName.Call(inputs)
	dataColumn := returnValue[0].Interface()
	if reflect.ValueOf(dataColumn).Type().Kind() != reflect.Slice {
		return nil
	}
	stringColumn := dataColumn.([]string)
	dataArray := make([]interface{}, 0, len(stringColumn))
	for _, col := range stringColumn {
		title := c.loader.GetMessage(tableName + "." + col)
		if title == "" {
			continue
		}
		attr := map[string]interface{}{
			"data":  col,
			"title": title,
		}
		dataArray = append(dataArray, attr)
	}
	return dataArray
}

// 是否默认隐藏column
func (c *ColumnService) isHiddenColumn(hiddenColumns interface{}, column string) (isSelect bool) {
	if reflect.TypeOf(hiddenColumns).Kind() == reflect.Slice {
		s := reflect.ValueOf(hiddenColumns)
		for i := 0; i < s.Len(); i++ {
			if s.Index(i).String() == column {
				return false
			}
		}
	} else {
		return true
	}
	return true
}

//将蛇形字符转换成_风格
func (c *ColumnService) toSnakeCase(str string) string {
	return tools.OtherHelper{}.ToSnakeCase(str)
}

//获取的model对应的table name
func (c *ColumnService) tableName(v reflect.Value) string {
	var data string
	methodName := v.MethodByName("TableName")
	if methodName.IsValid() {
		value := methodName.Call([]reflect.Value{})
		data = value[0].String()
	} else {
		data = c.toSnakeCase(v.Kind().String())
	}
	return data
}

func NewColumnService(loader context.Locale) IColumnService {
	return &ColumnService{sy: sync.Mutex{}, loader: loader}
}
