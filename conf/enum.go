package conf

import (
	"github.com/kataras/iris/v12/context"
	"strconv"
)

type Enum struct {
	Locale context.Locale
}

func (e Enum) DefaultText(src string) string {
	return e.Locale.GetMessage(src)
}

func (e Enum) ClearRuleText(src interface{}) string {
	return e.DefaultText("clear_rule." + src.(string))
}

func (e Enum) ClearRuleOptions() []interface{} {
	e.Locale.GetMessage("clear_rule")
	return []interface{}{}
}



func (e Enum) TransportTypeText(src interface{}) string {
	value := strconv.Itoa(int(src.(uint)))
	return e.DefaultText("transport_type." + value)
}


