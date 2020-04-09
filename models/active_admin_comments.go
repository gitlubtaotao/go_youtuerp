package models

import (
	"time"
)

type ActiveAdminComments struct {
	Id           int64     `xorm:"pk autoincr BIGINT(20)"`
	Namespace    string    `xorm:"index VARCHAR(255)"`
	Body         string    `xorm:"TEXT"`
	ResourceType string    `xorm:"index(index_active_admin_comments_on_resource_type_and_resource_id) VARCHAR(255)"`
	ResourceId   int64     `xorm:"index(index_active_admin_comments_on_resource_type_and_resource_id) BIGINT(20)"`
	AuthorType   string    `xorm:"index(index_active_admin_comments_on_author_type_and_author_id) VARCHAR(255)"`
	AuthorId     int64     `xorm:"index(index_active_admin_comments_on_author_type_and_author_id) BIGINT(20)"`
	CreatedAt    time.Time `xorm:"not null DATETIME"`
	UpdatedAt    time.Time `xorm:"not null DATETIME"`
}
