//其他一些帮助方法
package tools

import (
	"net"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"
)

var (
	MaxDepth    = 32
	matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
)

type IOtherHelper interface {
	GetIPAddress() (string, error)
	MapMerge(dst, src map[string]interface{}) map[string]interface{}
	StructToMap(currentObject interface{}) map[string]interface{}
}

type OtherHelper struct {
}

//获取访问的iP真实地址
func (o OtherHelper) GetIPAddress(r *http.Request) (string, error) {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" {
			return ip, nil
		}
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip, nil
	}
	
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip, nil
	}
	return "", nil
}

//map 进行合并
func (o OtherHelper) MapMerge(dst, src map[string]interface{}) map[string]interface{} {
	return o.merge(dst, src, 0)
}

func (o OtherHelper) StructToMap(currentObject interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	v := reflect.TypeOf(currentObject)
	utils := TimeHelper{}
	reflectValue := reflect.ValueOf(currentObject)
	reflectValue = reflect.Indirect(reflectValue)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		temp := v.Field(i).Type
		kind := temp.Kind()
		tag := v.Field(i).Tag.Get("json")
		if tag == "" {
			tag = o.ToSnakeCase(v.Field(i).Name)
		}
		field := reflectValue.Field(i).Interface()
		if kind == reflect.Struct {
			if temp.Name() == "Time" {
				res[tag] = utils.DefaultDate(field.(time.Time), "zh-CN")
			} else {
				res[tag] = o.StructToMap(field)
			}
		} else {
			if tag != "" {
				res[tag] = field
			}
		}
	}
	return res
}

//字符串转化
func (o OtherHelper) ToSnakeCase(str string) string {
	snake := matchAllCap.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(strings.ToLower(snake))
}

//slice 包含某元素
func ContainsSlice(src []interface{}, val string) (int, bool) {
	for i, item := range src {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func StructToChange(src interface{}) map[string]interface{} {
	t := reflect.TypeOf(src)
	v := reflect.ValueOf(src)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Tag.Get("json")
		if name != "" {
			data[name] = v.Field(i).Interface()
		}
	}
	return data
}

func (o OtherHelper) merge(dst, src map[string]interface{}, depth int) map[string]interface{} {
	if depth > MaxDepth {
		panic("too deep!")
	}
	for key, srcVal := range src {
		if dstVal, ok := dst[key]; ok {
			srcMap, srcMapOk := o.mapify(srcVal)
			dstMap, dstMapOk := o.mapify(dstVal)
			if srcMapOk && dstMapOk {
				srcVal = o.merge(dstMap, srcMap, depth+1)
			}
		}
		dst[key] = srcVal
	}
	return dst
}

func (o OtherHelper) mapify(i interface{}) (map[string]interface{}, bool) {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		m := map[string]interface{}{}
		for _, k := range value.MapKeys() {
			m[k.String()] = value.MapIndex(k).Interface()
		}
		return m, true
	}
	return map[string]interface{}{}, false
}
