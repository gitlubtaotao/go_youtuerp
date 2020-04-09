package models

import (
	"time"
)

type BaseDataSeaPorts struct {
	Id                 int64     `xorm:"pk autoincr BIGINT(20)"`
	Name               string    `xorm:"comment('海港') VARCHAR(255)"`
	SeaLineName        string    `xorm:"comment('航线名') VARCHAR(255)"`
	CreatedAt          time.Time `xorm:"not null DATETIME"`
	UpdatedAt          time.Time `xorm:"not null DATETIME"`
	NameCn             string    `xorm:"comment('海港中文名') VARCHAR(255)"`
	NationCn           string    `xorm:"comment('国家中文名') VARCHAR(255)"`
	NationEn           string    `xorm:"comment('国家英文名') VARCHAR(255)"`
	SeaLineCn          string    `xorm:"comment('航线中文名') VARCHAR(255)"`
	SeaLineEn          string    `xorm:"comment('航线英文名') VARCHAR(255)"`
	NamePya            string    `xorm:"comment('海港拼音全写') VARCHAR(255)"`
	NamePyf            string    `xorm:"comment('海港拼音简写') VARCHAR(255)"`
	NationCode         string    `xorm:"comment('国家代码') VARCHAR(255)"`
	NationCode2        string    `xorm:"comment('国家代码') VARCHAR(255)"`
	CityId             int       `xorm:"comment('城市ID') INT(11)"`
	City               string    `xorm:"comment('码头所属城市，一般为内贸') VARCHAR(255)"`
	CityEn             string    `xorm:"comment('城市英文名') VARCHAR(255)"`
	CityPya            string    `xorm:"comment('城市拼音全写') VARCHAR(255)"`
	CityPyf            string    `xorm:"comment('城市拼音简写') VARCHAR(255)"`
	HasFreight         int       `xorm:"default 1 comment('附加费') TINYINT(1)"`
	Position           int       `xorm:"default 0 INT(11)"`
	CountryId          int       `xorm:"comment('所属国家ID') INT(11)"`
	SeaLineIdSet       string    `xorm:"comment('所属航线（多个）') VARCHAR(255)"`
	DeletedAt          time.Time `xorm:"index DATETIME"`
	IsLocalChanged     int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
	BaseDataSeaLinesId int64     `xorm:"index BIGINT(20)"`
}
