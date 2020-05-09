package models

type Setting struct {
	ID     uint   `gorm:"primary_key"json:"id"`
	Key    string `gorm:"size:64" json:"key"`
	Field  string `gorm:"size:64" json:"field"`
	Value  string `gorm:"type:varchar(1024)" json:"value"`
	UserId uint   `gorm:"index:user_id"`
}
type ResultSetting struct {
	Key   string `json:"key"`
	Field string `json:"field"`
	Value string `json:"value"`
}

func (Setting) TableName() string {
	return "system_settings"
}
