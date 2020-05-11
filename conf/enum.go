package conf

import "github.com/kataras/iris/v12/context"

type Enum struct {
	Locale context.Locale
}

func (e Enum) DefaultText(src string) string {
	return e.Locale.GetMessage(src)
}

func (e Enum) ClearRule(src interface{}) string {
	return e.DefaultText("clear_rule." + src.(string))
}

func (e Enum) TransportType(src interface{}) string {
	return e.DefaultText("transport_type" + src.(string))
}