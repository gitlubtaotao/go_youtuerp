package redis

//保存公司redis信息
func (r Redis) SetCompany(id interface{}) error {
	return r.HSetRecord("user_companies",
		map[string]interface{}{"id": id},
		[]string{"id", "name_nick", "code", "age", "amount", "account_period"})
}

func (r Redis) SetCommon(tableName string, id interface{}, selectKeys []string) error {
	if len(selectKeys) == 0 {
		selectKeys = []string{"id", "name"}
	}
	return r.HSetRecord(tableName, map[string]interface{}{"id": id}, selectKeys)
}


