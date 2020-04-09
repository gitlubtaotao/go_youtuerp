package models

import (
	"time"
)

type ActiveStorageAttachments struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	Name          string    `xorm:"not null unique(index_active_storage_attachments_uniqueness) VARCHAR(255)"`
	RecordType    string    `xorm:"not null unique(index_active_storage_attachments_uniqueness) VARCHAR(255)"`
	RecordId      int64     `xorm:"not null unique(index_active_storage_attachments_uniqueness) BIGINT(20)"`
	BlobId        int64     `xorm:"not null index unique(index_active_storage_attachments_uniqueness) BIGINT(20)"`
	CreatedAt     time.Time `xorm:"not null DATETIME"`
	Uploader      int       `xorm:"comment('上传者') INT(11)"`
	OrderMasterId int       `xorm:"default 0 index INT(11)"`
}
