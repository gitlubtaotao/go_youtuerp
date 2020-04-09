package models

import (
	"time"
)

type HelpFeedbacks struct {
	Id                int64     `xorm:"pk autoincr BIGINT(20)"`
	Content           string    `xorm:"comment('内容') TEXT"`
	FeedbackUserId    int       `xorm:"comment('反馈人') INT(11)"`
	Url               string    `xorm:"TEXT"`
	Cookie            string    `xorm:"TEXT"`
	Json              string    `xorm:"TEXT"`
	ScreenShot        []byte    `xorm:"comment('全屏截图') LONGBLOB"`
	ScreenCapture     []byte    `xorm:"comment('截图') LONGBLOB"`
	ScreenCaptureHtml string    `xorm:"TEXT"`
	CreatedAt         time.Time `xorm:"not null DATETIME"`
	UpdatedAt         time.Time `xorm:"not null DATETIME"`
	DeletedAt         time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
