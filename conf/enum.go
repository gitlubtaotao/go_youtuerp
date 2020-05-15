package conf

import (
	"github.com/kataras/iris/v12/context"
	"reflect"
	"strconv"
)

type Enum struct {
	Locale context.Locale
}

func (e Enum) DefaultText(key string, src interface{}) string {
	ty := reflect.TypeOf(src)
	var dst string
	switch ty.Kind() {
	case reflect.String:
		dst = key + (src.(string))
	case reflect.Uint:
		dst = key + strconv.Itoa(int(src.(uint)))
	case reflect.Int:
		dst = key + strconv.Itoa(src.(int))
	case reflect.Int8:
		dst = key + strconv.Itoa(int(src.(int8)))
	default:
		dst = key + src.(string)
	}
	return e.Locale.GetMessage(dst)
}

func (e Enum) ClearRuleText(src interface{}) string {
	return e.DefaultText("clear_rule.", src)
}

func (e Enum) CompanyTypeText(src interface{}) string {
	return e.DefaultText("company_type.", src)
}
