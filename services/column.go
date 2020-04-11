package services

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

type IColumnService interface {
	DefaultColumn(model interface{}) (data []interface{}, err error)
	ColumnByBase(f reflect.StructField) (data []interface{})
	ColumnByOther(f reflect.StructField) (data []interface{})
	DefaultHiddenColumn(t reflect.Type, args ...interface{}) (data interface{})
	ToSnakeCase(str string) string
}

type ColumnService struct {
	sy sync.Mutex
}

func (c *ColumnService) DefaultColumn(model interface{}) (dataArray []interface{}, err error) {
	t := reflect.TypeOf(model)
	hiddenColumn := c.DefaultHiddenColumn(t)
	if t.Kind() != reflect.Struct {
		err = errors.New("mode is not struct")
		return
	}
	c.sy.Lock()
	defer c.sy.Unlock()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Type.Name() == "Model" {
			dataArray = append(dataArray, c.ColumnByBase(f)...)
			continue
		}
		//必须补上table_name，表示需要关联对象
		if f.Tag.Get("table_name") != "" {
			//只查询两层结构
			if f.Type.Kind() == reflect.Struct {
				attr := make(map[string]interface{})
				attr[f.Tag.Get("table_name")] = c.ColumnByOther(f)
				dataArray = append(dataArray, attr)
				continue
			}
		}
		data := f.Tag.Get("json")
		if data == "" {
			continue
		}
		attr := map[string]interface{}{
			"data":   data,
			"type":   f.Type.Name(),
			"select": c.isHiddenColumn(hiddenColumn, data),
		}
		dataArray = append(dataArray, attr)
	}
	return
}

// 获取gorm.Model 默认的列
func (c *ColumnService) ColumnByBase(f reflect.StructField) []interface{} {
	dataArray := make([]interface{}, 0)
	stringArray := []string{"ID", "CreatedAt", "UpdatedAt"}
	for i := 0; i < len(stringArray); i++ {
		if field, ok := f.Type.FieldByName(stringArray[i]); ok {
			attr := map[string]interface{}{
				"data":   c.ToSnakeCase(field.Name),
				"type":   field.Type.Name(),
				"select": true,
			}
			dataArray = append(dataArray, attr)
		}
	}
	return dataArray
}

func (c *ColumnService) ColumnByOther(f reflect.StructField) (dataArray []interface{}) {
	t := f.Type
	hiddenColumn := c.DefaultHiddenColumn(t)
	if t.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Type.Name() == "Model" {
			continue
		}
		data := f.Tag.Get("json")
		if data == "" {
			continue
		}
		attr := map[string]interface{}{
			"data":   data,
			"select": c.isHiddenColumn(hiddenColumn, data),
			"type":   f.Type.Name(),
		}
		dataArray = append(dataArray, attr)
	}
	return
}

//
func (c *ColumnService) DefaultHiddenColumn(t reflect.Type, args ...interface{}) (data interface{}) {
	if methodName, ok := t.MethodByName("DefaultHiddenColumn"); ok {
		inputs := make([]reflect.Value, len(args))
		for i, _ := range args {
			inputs[i] = reflect.ValueOf(args[i])
		}
		returnValue := methodName.Func.Method(methodName.Index).Call(inputs)
		data = returnValue[0].Interface()
	} else {
		data = []string{}
	}
	return
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

func (c *ColumnService) ToSnakeCase(str string) string {
	snake := matchAllCap.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(strings.ToLower(snake))
}
func NewColumnService() IColumnService {
	return &ColumnService{}
}
