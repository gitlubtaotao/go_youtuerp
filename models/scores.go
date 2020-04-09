package models

import (
	"time"
)

type Scores struct {
	Id           int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt    time.Time `xorm:"not null DATETIME"`
	UpdatedAt    time.Time `xorm:"not null DATETIME"`
	SourceType   string    `xorm:"index(index_scores_on_source_type_and_source_id) VARCHAR(255)"`
	SourceId     int64     `xorm:"index(index_scores_on_source_type_and_source_id) BIGINT(20)"`
	ScoreGrowth  string    `xorm:"comment('积分成长值') DECIMAL(15,2)"`
	ScoreCurrent string    `xorm:"comment('可用积分') DECIMAL(15,2)"`
	Grade        string    `xorm:"comment('等级') VARCHAR(255)"`
}
