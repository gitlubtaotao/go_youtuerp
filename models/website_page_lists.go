package models

type WebsitePageLists struct {
	Id          int64  `xorm:"pk autoincr BIGINT(20)"`
	SourceType  string `xorm:"index(index_website_page_lists_on_source_type_and_source_id) VARCHAR(255)"`
	SourceId    int    `xorm:"index(index_website_page_lists_on_source_type_and_source_id) INT(11)"`
	Title       string `xorm:"index VARCHAR(64)"`
	Details     string `xorm:"comment('内容信息') TEXT"`
	UrlInfo     string `xorm:"comment('链接信息') VARCHAR(255)"`
	CatalogType string `xorm:"comment('分类信息') VARCHAR(255)"`
}
