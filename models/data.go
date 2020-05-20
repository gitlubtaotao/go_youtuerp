package models

import "time"

type BaseDataLevel struct {
	ID     uint   `gorm:"primary_key"json:"id"`
	Code   string `gorm:"size:64" json:"code"`
	Name   string `gorm:"size:64" json:"name"`
	EnName string `gorm:"size:64" json:"en_name"`
}

type BaseDataCode struct {
	ID        uint       `gorm:"primary_key"json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
	CodeName  string     `gorm:"size:64;index:code_name" json:"code_name"`
	Name      string     `gorm:"size:256;index:name;" json:"name"`
	Remarks   string     `gorm:"size:522;" json:"remarks"`
}

func (BaseDataCode) TableName() string {
	return "base_data_codes"
}

func (BaseDataLevel) TableName() string {
	return "base_data_levels"
}
