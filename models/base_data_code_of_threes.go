package models

import (
	"time"
)

type BaseDataCodeOfThrees struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	Code           string    `xorm:"comment('三字代码') VARCHAR(255)"`
	Name           string    `xorm:"index VARCHAR(255)"`
	EnName         string    `xorm:"comment('机场英文名') VARCHAR(255)"`
	Short          string    `xorm:"VARCHAR(255)"`
	Airport        string    `xorm:"comment('机场名称') index VARCHAR(255)"`
	EnAirport      string    `xorm:"comment('英文机场名称') VARCHAR(255)"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedAt      time.Time `xorm:"not null DATETIME"`
	City           string    `xorm:"comment('城市') VARCHAR(255)"`
	EnCity         string    `xorm:"comment('城市英文名') VARCHAR(255)"`
	FourCode       string    `xorm:"VARCHAR(255)"`
	NamePya        string    `xorm:"comment('名称拼音全写') VARCHAR(255)"`
	NamePyf        string    `xorm:"comment('名称拼音简写') VARCHAR(255)"`
	AirportPya     string    `xorm:"comment('机场拼音全写') VARCHAR(255)"`
	AirportPyf     string    `xorm:"comment('机场拼音简写') VARCHAR(255)"`
	CityPya        string    `xorm:"comment('城市拼音全写') VARCHAR(255)"`
	CityPyf        string    `xorm:"comment('城市拼音简写') VARCHAR(255)"`
	DeletedAt      time.Time `xorm:"index DATETIME"`
	IsLocalChanged int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
}
