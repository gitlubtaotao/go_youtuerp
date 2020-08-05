package enumerize

import (
	"github.com/kataras/iris/v12/context"
	"reflect"
	"strconv"
)

type Enumerize struct {
	Locale context.Locale
}

func NewEnum(locale context.Locale) Enumerize {
	return Enumerize{Locale: locale}
}

func (e Enumerize) DefaultText(key string, src interface{}) string {
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

func (e Enumerize) ClearRuleText(src interface{}) string {
	return e.DefaultText("clear_rule.", src)
}

func (e Enumerize) CompanyTypeText(src interface{}) string {
	return e.DefaultText("user_companies_company_type.", src)
}
