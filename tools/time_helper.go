//对时间格式进行转换
package tools

import (
	"time"
)

type ITimeHelper interface {
	DefaultDate(time time.Time, language string) string
	DefaultDateTime(time time.Time, language string) string
	TimeFormat(time time.Time, timeFormat string) string
}

type TimeHelper struct {
}

const (
	ZHDate        = "2006年01月02日"
	ShortDate     = "2006-01-02"
	ZHDateTime    = "2006年01月02日 15:04:05"
	ShortDateTime = "2006-01-02 15:04:05"
	ENDate        = "Jan 02, 2006"
	ENDateTime    = "Jan 02, 2006 15:04:05"
)

//转化时间
func (u TimeHelper) TimeFormat(time time.Time, timeFormat string) string {
	return time.Format(timeFormat)
}

func (u TimeHelper) DefaultDate(time time.Time, language string) string {
	if language == "zh-CN" {
		return u.TimeFormat(time, ZHDate)
	} else {
		return u.TimeFormat(time, ENDate)
	}
}

func (u TimeHelper) DefaultDateTime(time time.Time, language string) string {
	if language == "zh-CN" {
		return u.TimeFormat(time, ZHDateTime)
	} else {
		return u.TimeFormat(time, ENDateTime)
	}
}