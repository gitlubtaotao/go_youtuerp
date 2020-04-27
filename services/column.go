package services

import (
	"errors"
	"github.com/kataras/iris/v12/context"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"
	"youtuerp/tools"
)

var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

type IColumnService interface {
	DefaultColumn(model interface{}, args ...interface{}) (data []interface{}, err error)
	ColumnByBase(f reflect.StructField) (data []interface{})
	ColumnByOther(f reflect.StructField) (data []interface{})
	DefaultHiddenColumn(t reflect.Value, args ...interface{}) (data interface{})
	ToSnakeCase(str string) string
	//将struct转化成对应的map结构
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
	return c.structToMap(currentObject), nil
}

func (c *ColumnService) DefaultColumn(model interface{}, args ...interface{}) (dataArray []interface{}, err error) {
	t := reflect.TypeOf(model)
	v := reflect.ValueOf(model)
	hiddenColumn := c.DefaultHiddenColumn(v)
	if t.Kind() != reflect.Struct {
		err = errors.New("mode is not struct")
		return
	}
	c.sy.Lock()
	defer c.sy.Unlock()
	tableName := c.tableName(v)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Type.Name() == "Model" || f.Type.Name() == "Base" {
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
		if c.isHiddenColumn(hiddenColumn, data) {
			attr := map[string]interface{}{
				"data":  data,
				"type":  f.Type.Name(),
				"title": c.loader.GetMessage(tableName + "." + data),
			}
			dataArray = append(dataArray, attr)
		}
	}
	return
}

// 获取gorm.Model 默认的列
func (c *ColumnService) ColumnByBase(f reflect.StructField) []interface{} {
	dataArray := make([]interface{}, 0)
	stringArray := []string{"ID", "CreatedAt", "UpdatedAt"}
	for i := 0; i < len(stringArray); i++ {
		if field, ok := f.Type.FieldByName(stringArray[i]); ok {
			data := c.ToSnakeCase(field.Name)
			attr := map[string]interface{}{
				"data":   data,
				"type":   field.Type.Name(),
				"title":  c.loader.GetMessage("base." + data),
				"select": true,
			}
			dataArray = append(dataArray, attr)
		}
	}
	return dataArray
}

func (c *ColumnService) ColumnByOther(f reflect.StructField) (dataArray []interface{}) {
	t := f.Type
	value := reflect.New(t)
	tableName := c.tableName(value)
	hiddenColumn := c.DefaultHiddenColumn(value)
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
			"title":  c.loader.GetMessage(tableName + "." + data),
		}
		dataArray = append(dataArray, attr)
	}
	return
}

//
func (c *ColumnService) DefaultHiddenColumn(v reflect.Value, args ...interface{}) (data interface{}) {
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
func (c *ColumnService) ToSnakeCase(str string) string {
	snake := matchAllCap.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(strings.ToLower(snake))
}

//获取的model对应的table name
func (c *ColumnService) tableName(v reflect.Value) string {
	var data string
	methodName := v.MethodByName("TableName")
	if methodName.IsValid() {
		value := methodName.Call([]reflect.Value{})
		data = value[0].String()
	} else {
		data = c.ToSnakeCase(v.Kind().String())
	}
	return data
}

//strut value to map value
func (c *ColumnService) structToMap(currentObject interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	v := reflect.TypeOf(currentObject)
	utils := tools.TimeHelper{}
	reflectValue := reflect.ValueOf(currentObject)
	reflectValue = reflect.Indirect(reflectValue)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		temp := v.Field(i).Type
		kind := temp.Kind()
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()
		if kind == reflect.Struct {
			if temp.Name() == "Time" {
				res[tag] = utils.DefaultDate(field.(time.Time), c.loader.Language())
			} else {
				res[tag] = c.structToMap(field)
			}
		} else {
			if tag != "" {
				res[tag] = field
			}
		}
	}
	return res
}

func NewColumnService(loader context.Locale) IColumnService {
	return &ColumnService{sy: sync.Mutex{}, loader: loader}
}
