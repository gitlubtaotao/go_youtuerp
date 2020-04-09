package models

type WebsiteMenus struct {
	Id         int64  `xorm:"pk autoincr BIGINT(20)"`
	NameCn     string `xorm:"comment('菜单中文名') VARCHAR(255)"`
	NameEn     string `xorm:"comment('菜单英文名') VARCHAR(255)"`
	Enabled    int    `xorm:"default 1 comment('菜单是否显示') index TINYINT(1)"`
	Position   int    `xorm:"INT(11)"`
	HiddenName string `xorm:"comment('隐藏字段') index VARCHAR(64)"`
}
