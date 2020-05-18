package redis

import (
	"github.com/kataras/golog"
)

func (r Redis) GetCompany(id interface{}, field string) string {
	if field == "" {
		field = "name_nick"
	}
	if value := r.HGet("user_companies", id, field); value != "" {
		return value
	} else {
		if err := r.SetCompany(id); err != nil {
			golog.Errorf("set company redis is error %v", err)
			return ""
		}
		return r.HGet("user_companies", id, field)
	}
}

func (r Redis) GetCommon(table string, id interface{}, field string) string {
	if field == "" {
		field = "name"
	}
	if value := r.HGet(table, id, field); value != "" {
		return value
	} else {
		if err := r.SetCommon(table, id, []string{}); err != nil {
			golog.Errorf("set %v redis is error %v", table, err)
			return ""
		}
		return r.HGet(table, id, field)
	}
}


