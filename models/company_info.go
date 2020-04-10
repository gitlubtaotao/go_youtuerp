package models

import "github.com/jinzhu/gorm"

type CompanyInfo struct {
	gorm.Model
	Telephone  string `form:"telephone" json:"telephone"`                               // 座机
	Telephone2 string `form:"telephone2" json:"telephone2"`                             // 备用座机
	Fax        string `form:"fax" json:"fax"`                                           // 传真
	Fax2       string `form:"fax2" json:"fax2"`                                         // 备用传真
	Address    string `form:"address" json:"address"`                                   // 地址
	Address2   string `form:"address2" json:"address2"`                                 // 备用地址
	Website    string `form:"website" json:"website"`                                   // 网站
	SourceId   int    `gorm:"UNIQUE_INDEX;NOT NULL;" form:"source_id" json:"source_id"` // 数据来源
	NameNick   string `form:"name_nick" json:"name_nick"`
	NameCn     string `form:"name_cn" json:"name_cn"`
	NameEn     string `form:"name_en" json:"name_en"`
	Code       string `form:"code" json:"code"`
	DeletedAt  string `form:"deleted_at" json:"deleted_at"`
	Email      string `form:"email" json:"email"`
	Province   string `form:"province" json:"province"` // 省份
	City       string `form:"city" json:"city"`         // 市
	District   string `form:"district" json:"district"` // 区
}

func (*CompanyInfo) TableName() string {
	return "companies"
}
