package models

import (
	"time"
)

type BaseDataCapTypeSizes struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	Name           string    `xorm:"comment('标题') VARCHAR(255)"`
	CapSize        string    `xorm:"comment('箱尺寸') VARCHAR(255)"`
	CapType        string    `xorm:"comment('箱类型') VARCHAR(255)"`
	Remark         string    `xorm:"comment('说明') VARCHAR(255)"`
	IsoCode        string    `xorm:"comment('ISO代码') VARCHAR(255)"`
	IsoRemark      string    `xorm:"comment('ISO代码说明') VARCHAR(255)"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedAt      time.Time `xorm:"not null DATETIME"`
	DeletedAt      time.Time `xorm:"index DATETIME"`
	IsLocalChanged int       `xorm:"default 1 comment('本地修改') TINYINT(1)"`
	ExtLength      string    `xorm:"comment('外长') VARCHAR(255)"`
	ExtWidth       string    `xorm:"comment('外宽') VARCHAR(255)"`
	ExtHeight      string    `xorm:"comment('外高') VARCHAR(255)"`
	IntLength      string    `xorm:"comment('内长') VARCHAR(255)"`
	IntWidth       string    `xorm:"comment('内宽') VARCHAR(255)"`
	IntHeight      string    `xorm:"comment('内高') VARCHAR(255)"`
	DoorWidth      string    `xorm:"comment('箱门宽') VARCHAR(255)"`
	DoorHeight     string    `xorm:"comment('箱门高') VARCHAR(255)"`
	GrossWeight    string    `xorm:"comment('总重') VARCHAR(255)"`
	Tare           string    `xorm:"comment('自重') VARCHAR(255)"`
	Payload        string    `xorm:"comment('载重') VARCHAR(255)"`
	CubicCapacity  string    `xorm:"comment('容积') VARCHAR(255)"`
	Ventilation    string    `xorm:"comment('通风量') VARCHAR(255)"`
	HatchDiameter  string    `xorm:"comment('舱口直径') VARCHAR(255)"`
	SelectCount    int       `xorm:"default 0 comment('选中的次数,用于排序常用') INT(11)"`
}
