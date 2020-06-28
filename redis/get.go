package redis

import (
	"github.com/kataras/golog"
	"strconv"
	"youtuerp/models"
	"youtuerp/tools/uploader"
)

func (r Redis) HGetCrm(id interface{}, field string) (value string) {
	var err error
	value, err = r.getUserCompany(models.CrmCompany{}.RedisKey(), id, field)
	if err != nil {
		golog.Errorf("get crm redis is error %v", err)
	}
	return value
}

func (r Redis) HGetCompany(id interface{}, field string) (value string) {
	var err error
	value, err = r.getUserCompany(models.Company{}.TableName(), id, field)
	if err != nil {
		golog.Errorf("get crm redis is error %v", err)
	}
	return value
}

func (r Redis) HGetRecord(table string, id interface{}, field string) (value string) {
	var err error
	if field == "" {
		field = "name"
	}
	key := r.CombineKey(table, id)
	
	if value, err = r.HGet(key, field); err != nil {
		golog.Errorf("HGetValue is error table_name is %v,id is %v,error is %v", table, id, err)
		return ""
	}
	if value != "" {
		return value
	}
	if err := r.HSetRecord(table, id, []string{}); err != nil {
		golog.Errorf("set %v redis is error %v", table, err)
		return ""
	}
	value, _ = r.HGet(key, field)
	return
}

func (r Redis) HGetValue(table string, id interface{}, field string) string {
	if field == "" {
		field = "name"
	}
	value, err := r.HGet(r.CombineKey(table, id), field)
	if err != nil {
		golog.Errorf("HGetValue is error %v", err)
		return ""
	}
	return value
}

func (r Redis) HGetAllValue(table string, id interface{}) (value map[string]string) {
	key := r.CombineKey(table, id)
	var err error
	if value, err = r.HGetAll(key); err != nil {
		golog.Errorf("HGetAllValue is error table is %v,id is %v,error is %v", table, id, err)
		return
	}
	return
}

//获取某一张表中所有的缓存集合
func (r Redis) HCollectOptions(table string) []map[string]string {
	var collect []map[string]string
	members := r.SMembers(table)
	r.sy.Lock()
	defer r.sy.Unlock()
	for _, i := range members {
		temp, err := r.HGetAll(r.CombineKey(table, i))
		if err != nil {
			golog.Warnf("HCollectOptions is error, id is %v,error is %v", i, err)
			continue
		}
		collect = append(collect, temp)
	}
	return collect
}

func (r Redis) getUserCompany(tableName string, id interface{}, field string) (value string, err error) {
	if field == "" {
		field = "name_nick"
	}
	key := r.CombineKey(tableName, id)
	if value, err = r.HGet(key, field); err != nil {
		return
	}
	if value != "" {
		return value, nil
	}
	if err := r.HSetCompany(tableName, id); err != nil {
		return "", err
	}
	return r.HGet(key, field)
}

func CompanyLogo() string {
	value, err := GetSystemSetting("base", "company_logo")
	if err == nil {
		golog.Errorf("company logo redis is error %v", err)
		return ""
	}
	upload := uploader.NewQiNiuUploaderDefault()
	return upload.PrivateReadURL(value)
}

//获取数据安全控制
func DataSecurityControl() bool {
	value, err := GetSystemSetting("base", "data_security_control")
	if err != nil {
		golog.Errorf("data security control is error %v", err)
		return false
	}
	if value == "false" {
		return false
	} else {
		return true
	}
}

//计费的计算方式
func ConversionMethod() (method string, remain int) {
	tempM, _ := GetSystemSetting("base", "conversion_method_calu")
	tempR, _ := GetSystemSetting("base", "conversion_method_remain")
	if tempM == "" || tempR == "" {
	}
	remain, _ = strconv.Atoi(tempR)
	return tempM, remain
}

//获取本位币计算方式
func SystemFinanceCurrency() string {
	value, err := GetSystemSetting("finance", "system_finance_currency")
	if err != nil {
		golog.Errorf("system finance currency is err %v", err)
		return ""
	}
	return value
}

//获取系统汇率计算方式
func SystemRateSetting() string {
	value, err := GetSystemSetting("finance", "system_finance_rate")
	if err != nil {
		golog.Errorf("system finance currency is err %v", err)
		return ""
	}
	return value
}

//获取是否跳过审批
func SystemFinanceApprove() string {
	value, err := GetSystemSetting("finance", "skip_approve")
	if err != nil {
		golog.Errorf("system finance approve is err %v", err)
		return ""
	}
	return value
}

//获取是否跳过复核
func SystemFinanceAudit() string {
	value, err := GetSystemSetting("finance", "skip_audit")
	if err != nil {
		golog.Errorf("system finance audit is err %v", err)
		return ""
	}
	return value
}


func GetSystemSetting(key string, field string) (value string, err error) {
	red := NewRedis()
	return red.HGet(red.CombineKey("system_settings", key), field)
}
