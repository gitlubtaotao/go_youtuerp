package models

type WebsitePages struct {
	Id            int64  `xorm:"pk autoincr BIGINT(20)"`
	Title         string `xorm:"comment('标题') index VARCHAR(64)"`
	Keyword       string `xorm:"comment('关键字') TEXT"`
	Details       string `xorm:"comment('详情内容') LONGTEXT"`
	WebsiteMenuId int    `xorm:"default 0 comment('所属菜单') index INT(11)"`
	Locale        string `xorm:"default 'zh-CN' comment('语言选择') index VARCHAR(255)"`
}
