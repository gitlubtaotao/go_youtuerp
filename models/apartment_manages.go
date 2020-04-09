package models

import (
	"time"
)

type ApartmentManages struct {
	Id              int64     `xorm:"pk autoincr BIGINT(20)"`
	Name            string    `xorm:"comment('数据库名') unique VARCHAR(255)"`
	HostName        string    `xorm:"comment('主机名') VARCHAR(255)"`
	PortName        int       `xorm:"comment('端口') INT(11)"`
	AppId           string    `xorm:"VARCHAR(255)"`
	AppSecret       string    `xorm:"VARCHAR(255)"`
	DefaultPassword string    `xorm:"VARCHAR(255)"`
	CreatedAt       time.Time `xorm:"not null DATETIME"`
	UpdatedAt       time.Time `xorm:"not null DATETIME"`
	CompanyName     string    `xorm:"comment('公司名称') VARCHAR(255)"`
	ValidityDate    time.Time `xorm:"comment('有效期') DATETIME"`
	ContactInfo     string    `xorm:"comment('联系信息') TEXT"`
	TrailerLocation int       `xorm:"default 0 comment('拖车定位使用次数') INT(11)"`
}
