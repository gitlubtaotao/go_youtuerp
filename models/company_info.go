package models

import "github.com/jinzhu/gorm"

type CompanyInfo struct {
	gorm.Model
	Telephone  string `gorm:"not null;" form:"telephone" json:"telephone" validate:"required"` // 座机
	Telephone2 string `form:"telephone2" json:"telephone2"`                                    // 备用座机
	Fax        string `form:"fax" json:"fax"`                                                  // 传真
	Fax2       string `form:"fax2" json:"fax2"`                                                // 备用传真
	ZhAddress  string `gorm:"column:address;" form:"address" json:"address"`                   // 地址
	EnAddress  string `gorm:"column:address2;" form:"address2" json:"address2"`                // 备用地址
	Website    string `form:"website" json:"website"`                                          // 网站
	SourceId   int    `gorm:"not null;" form:"source_id"`                     // 数据来源
	Code       string `form:"code" json:"code"`
	Email      string `form:"email" json:"email" validate:"email"`
	Province   string `form:"province" json:"province"` // 省份
	City       string `form:"city" json:"city"`         // 市
	District   string `form:"district" json:"district"` // 区
}

func (CompanyInfo) TableName() string {
	return "companies"
}
