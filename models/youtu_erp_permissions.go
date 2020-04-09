package models

import (
	"time"
)

type YoutuErpPermissions struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt      time.Time `xorm:"DATETIME"`
	UpdatedAt      time.Time `xorm:"DATETIME"`
	YoutuErpRoleId int       `xorm:"index INT(11)"`
	Name           string    `xorm:"VARCHAR(255)"`
	SubjectClass   string    `xorm:"index index(subject_class_on_action) VARCHAR(255)"`
	SubjectId      int       `xorm:"INT(11)"`
	Action         string    `xorm:"index index(subject_class_on_action) VARCHAR(255)"`
	Description    string    `xorm:"TEXT"`
	ActionType     string    `xorm:"VARCHAR(255)"`
	Klass          string    `xorm:"VARCHAR(255)"`
	DeletedAt      time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`
}
