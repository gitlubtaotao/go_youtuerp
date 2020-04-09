package models

import (
	"time"
)

type FormerTrailsLoadInfos struct {
	Id                     int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt              time.Time `xorm:"not null DATETIME"`
	UpdatedAt              time.Time `xorm:"not null DATETIME"`
	Loader                 string    `xorm:"comment('装货联系人') VARCHAR(255)"`
	SupplierMobi           string    `xorm:"comment('联系人电话') VARCHAR(255)"`
	PlaceOfReceipt         string    `xorm:"comment('装货地信息') TEXT"`
	GrossWeight            string    `xorm:"comment('货物重量') VARCHAR(255)"`
	Size                   string    `xorm:"comment('货物体积') VARCHAR(255)"`
	EstimatedTimeOfLoading string    `xorm:"comment('预计装货时间') VARCHAR(255)"`
	LoadingDate            string    `xorm:"comment('预计装货时间') VARCHAR(255)"`
	SourceType             string    `xorm:"index(index_former_trails_load_infos_on_source_type_and_source_id) VARCHAR(255)"`
	SourceId               int64     `xorm:"index(index_former_trails_load_infos_on_source_type_and_source_id) BIGINT(20)"`
	Count                  int       `xorm:"comment('数量') INT(11)"`
	LoadOrUnload           int       `xorm:"comment('装货卸货,0为装货地，1为卸货地') TINYINT(1)"`
}
