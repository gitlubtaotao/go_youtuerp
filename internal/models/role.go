package models

import (
	"time"
)

type Role struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	UserId     uint      `gorm:"index:user_id" json:"user_id"`
	UserName   string    `gorm:"size:64" json:"user_name"`
	Name       string    `gorm:"size:64" json:"name"`
	SourceID   uint      `gorm:"index:idx_source_id_and_type" json:"source_id"`
	SourceType string    `gorm:"index:idx_source_id_and_type" json:"source_type"`
}

const (
	RoleNameSale     = "salesman"
	RoleNameAudit    = "audit"
	RoleNameCreate   = "created"
	RoleNameCustomer = "customer"
	RoleNameBusiness = "business"
	RoleNameFinance  = "finance"
	RoleNameFile     = "file"
)
