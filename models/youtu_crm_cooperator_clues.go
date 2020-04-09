package models

import (
	"time"
)

type YoutuCrmCooperatorClues struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	NameCn           string    `xorm:"comment('公司中文名称') TEXT"`
	NameNick         string    `xorm:"comment('公司简称') VARCHAR(150)"`
	NameEn           string    `xorm:"comment('公司英文名称') VARCHAR(150)"`
	Website          string    `xorm:"comment('网址') TEXT"`
	CompanyType      int       `xorm:"comment('公司类型') SMALLINT(6)"`
	BusinessTypeName string    `xorm:"comment('公司业务类型') VARCHAR(50)"`
	CustomerSource   int       `xorm:"comment('客户来源: 线上注册,广告,陌拜') INT(11)"`
	Province         string    `xorm:"comment('客户所在省') VARCHAR(10)"`
	City             string    `xorm:"comment('客户所在市') VARCHAR(15)"`
	District         string    `xorm:"comment('客户所在区') VARCHAR(255)"`
	Address          string    `xorm:"comment('客户具体中文地址') TEXT"`
	Address2         string    `xorm:"comment('客户具体英文地址') TEXT"`
	Remark           string    `xorm:"comment('对公司的备注') TEXT"`
	Email            string    `xorm:"comment('公司邮箱') TEXT"`
	Phone            string    `xorm:"comment('公司座机') TEXT"`
	ContactName      string    `xorm:"comment('联系人姓名') TEXT"`
	ContactPost      string    `xorm:"comment('联系人职位') VARCHAR(255)"`
	ContactMobi      string    `xorm:"comment('联系人手机') TEXT"`
	ContactGender    int       `xorm:"comment('联系人性别') TINYINT(4)"`
	ContactEmail     string    `xorm:"comment('联系人邮箱') TEXT"`
	ContactTelephone string    `xorm:"comment('联系人座机') TEXT"`
	ContactWechat    string    `xorm:"comment('联系人微信') VARCHAR(255)"`
	ContactQq        string    `xorm:"comment('联系人QQ') VARCHAR(255)"`
	ContactFax       string    `xorm:"comment('联系人FAX') VARCHAR(80)"`
	ContactRemark    string    `xorm:"comment('对联系人的备注') TEXT"`
	AuditorId        int       `xorm:"comment('审核人的users.id') index INT(11)"`
	CreatorId        int       `xorm:"comment('线索创建人的users.id') INT(11)"`
	LiableUserId     int       `xorm:"comment('负责人的users.id') index INT(11)"`
	UserCompanyId    int       `xorm:"comment('线索转化为客户之后的公司user_companies.id') INT(11)"`
	DeletedAt        time.Time `xorm:"default '1969-12-31 16:00:00' comment('默认软删除') DATETIME"`
	CompanyId        int       `xorm:"index INT(11)"`
	CreatedAt        time.Time `xorm:"not null DATETIME"`
	UpdatedAt        time.Time `xorm:"not null DATETIME"`
	Status           string    `xorm:"comment('审核状态') index VARCHAR(255)"`
	Type             string    `xorm:"VARCHAR(255)"`
}
