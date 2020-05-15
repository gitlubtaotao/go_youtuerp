package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type CrmClue struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `sql:"index"`
	CompanyType   int        `json:"company_type"`
	NameNick      string     `gorm:"size:64;index:name_nick" json:"name_nick" validate:"required"`
	NameCn        string     `gorm:"size:128" json:"name_cn" validate:"required"`
	NameEn        string     `gorm:"size:256" json:"name_en"`
	Tel           string     `gorm:"size:16;index:tel" json:"tel" validate:"required"`
	Source        string     `gorm:"size:256" json:"source"`
	Email         string     `gorm:"size:64" json:"email" validate:"required"`
	ZhAddress     string     `gorm:"size:512" json:"zh_address"`
	EnAddress     string     `gorm:"size:1024" json:"en_address"`
	Remarks       string     `gorm:"size:1024" json:"remarks"`
	UserName      string     `gorm:"size:64;index:user_name" json:"user_name"`
	UserTel       string     `gorm:"size:16;index:user_tel" json:"user_tel"`
	UserEmail     string     `gorm:"size:64" json:"user_email"`
	WechatId      string     `grom:"size:16" json:"wechat_id"`
	QQId          string     `gorm:"size:16" json:"qq_id"`
	UserRemarks   string     `gorm:"size:1024" json:"user_remarks"`
	CreateId      uint       `gorm:"index:create_id" json:"create_id"`
	CreateName    string     `gorm:"-" json:"create_name"`
	UserCompanyId uint       `gorm:"index:user_company_id"`
	City          string     `gorm:"size:64" json:"city"`
	Province      string     `gorm:"size:64" json:"province"`
	Distinct      string     `gorm:"size:64" json:"distinct"`
	CrmTracks     []CrmTrack `gorm:"polymorphic:Source;"`
	UserCreate    Employee   `gorm:"ForeignKey:CreateId" validate:"-"`
	Status        uint       `gorm:"index:status;default:0" json:"status"`
}

type CrmCompany struct {
	ID               uint       `gorm:"primary_key"json:"id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `sql:"index"`
	CompanyType      int        `form:"company_type" json:"company_type"`
	ParentId         int        `gorm:"index:parent_id" json:"parent_id"`
	ParentNameNick   string     `gorm:"-"`
	UserSalesmanId   int        `gorm:"index:user_salesman_id"` // 所属的业务人员
	UserAuditId      int        `gorm:"-" json:"user_audit_id"`
	UserCreateId     int        `gorm:"-" json:"user_create_id"`
	AccountPeriod    string     `gorm:"size:16" json:"account_period"` // 公司结算类型
	Age              int        `json:"age"`
	Amount           float64    `json:"amount"`
	NameNick         string     `gorm:"unique;not null" form:"name_nick" json:"name_nick" validate:"required"`
	NameCn           string     `gorm:"unique;not null" form:"name_cn" json:"name_cn" validate:"required"`
	NameEn           string     `gorm:"unique;not null" form:"name_en" json:"name_en" validate:"required"`
	BusinessTypeName string     `gorm:"size:256" json:"business_type_name"`
	Status           string     `gorm:"size:64" json:"status"`
	Telephone        string     `form:"telephone" json:"telephone" validate:"required"` // 座机
	Email            string     `json:"email" validate:"required"`
	Fax              string     `json:"fax"`
	ZhAddress        string     `gorm:"column:address;" form:"zh_address" json:"zh_address"`
	EnAddress        string     `gorm:"column:address2;" form:"en_address" json:"en_address"`
	Remark           string     `form:"remark" json:"remark"` // 公司备注
	Website          string     `form:"website" json:"website"`
	City             string     `gorm:"size:64"`
	Province         string     `gorm:"size:64"`
	Distinct         string     `gorm:"size:64"`
	Code             string     `json:"code"`
	Roles            []Role     `gorm:"polymorphic:Source;" json:"roles"`
}

type CrmUser struct {
}

type CrmTrack struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `sql:"index"`
	OfWay         string     `gorm:"size:64" json:"of_way" validate:"required"`
	TrackUserId   uint       `gorm:"index:track_user_id" json:"track_user_id" validate:"required"`
	NextTrackTime time.Time  `gorm:"comment:'下次跟进时间'" json:"next_track_time" validate:"required"`
	Description   string     `gorm:"size:1024;comment:'跟进内容'"  json:"description" validate:"required"`
	ContactName   string     `gorm:"size:64;comment: '联系人名字'" json:"contact_name"`
	ContactTel    string     `gorm:"size:16" json:"contact_tel"`
	ContactEmail  string     `gorm:"size:64" json:"contact_email"`
	SourceId      uint       `gorm:"index:idx_source_id_and_type" json:"source_id"`
	SourceType    string     `gorm:"index:idx_source_id_and_type" json:"source_type"`
}

const (
	Unconverted = iota
	Converted
)

func (CrmClue) TableName() string {
	return "crm_clues"
}

func (CrmCompany) TableName() string {
	return "user_companies"
}

func (CrmUser) TableName() string {
	return "users"
}

func (CrmTrack) TableName() string {
	return "crm_tracks"
}

func (c *CrmClue) BeforeCreate(scope *gorm.Scope) (err error) {
	c.Status = 0
	return
}

func (c *CrmCompany) BeforeCreate(scope *gorm.Scope) (err error) {
	c.Status = "approving"
	return
}
