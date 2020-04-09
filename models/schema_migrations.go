package models

type SchemaMigrations struct {
	Version string `xorm:"not null pk VARCHAR(255)"`
}
