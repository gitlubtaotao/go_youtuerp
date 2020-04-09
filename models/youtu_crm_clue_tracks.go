package models

import (
	"time"
)

type YoutuCrmClueTracks struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	SourceType     string    `xorm:"index(index_youtu_crm_clue_tracks_on_source_type_and_source_id) VARCHAR(255)"`
	SourceId       int64     `xorm:"index(index_youtu_crm_clue_tracks_on_source_type_and_source_id) BIGINT(20)"`
	Description    string    `xorm:"comment('跟进描述') TEXT"`
	OfWay          string    `xorm:"comment('跟进方式') VARCHAR(255)"`
	TrackingUserId int       `xorm:"comment('跟进人') INT(11)"`
	NextTrackTime  time.Time `xorm:"comment('下次跟进时间') index DATETIME"`
	Nofity         int       `xorm:"default 0 comment('是否通知,可能有多种通知方式,定为integer预留扩展') SMALLINT(6)"`
	DeletedAt      time.Time `xorm:"default '1969-12-31 16:00:00' comment('默认软删除') DATETIME"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedAt      time.Time `xorm:"not null DATETIME"`
	UserCompanyId  int       `xorm:"default 0 comment('所属公司') index INT(11)"`
}
