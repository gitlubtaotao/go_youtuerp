package models

import "time"

type Base struct {
	ID        uint       `gorm:"primary_key"json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
}

type ReadPassword struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type SelectResult struct {
	ID           uint   `json:"id"`
	NameNick     string `json:"name_nick"`
	NameCn       string `json:"name_cn"`
	NameEn       string `json:"name_en"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
	Code         string `json:"code"`
}

type RenderFormerData struct {
	OrderExtendInfo      OrderExtendInfo      `json:"order_extend_info"`
	FormerSeaInstruction FormerSeaInstruction `json:"former_sea_instruction"`
	FormerSeaBook        FormerSeaBook        `json:"former_sea_book"`
	FormerSeaSoNo        FormerSeaSoNo        `json:"former_sea_so_no"`
	SeaCargoInfo         []SeaCargoInfo       `json:"sea_cargo_info"`
	FormerTrailerOrder   FormerTrailerOrder   `json:"former_trailer_order"`
}
