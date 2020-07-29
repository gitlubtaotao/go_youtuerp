/*
	client read json struct
	not create table
*/
package models

//读取前端传送的密码数据
type ReadPassword struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

//读取前端订单操作盘对应的数据
type RenderFormerData struct {
	OrderExtendInfo        OrderExtendInfo        `json:"order_extend_info"`
	FormerSeaInstruction   FormerSeaInstruction   `json:"former_sea_instruction"`
	FormerSeaBook          FormerSeaBook          `json:"former_sea_book"`
	FormerSeaSoNo          FormerSeaSoNo          `json:"former_sea_so_no"`
	SeaCargoInfo           []SeaCargoInfo         `json:"sea_cargo_info"`
	FormerTrailerOrder     FormerTrailerOrder     `json:"former_trailer_order"`
	FormerOtherService     FormerOtherService     `json:"former_other_service"`
	FormerWarehouseService FormerWarehouseService `json:"former_warehouse_service"`
	FormerCustomClearance  FormerCustomClearance  `json:"former_custom_clearance"`
}

//读取前端下拉框传送的数据
type ReadSelectResult struct {
	ID           uint   `json:"id"`
	NameNick     string `json:"name_nick"`
	NameCn       string `json:"name_cn"`
	NameEn       string `json:"name_en"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
	Code         string `json:"code"`
}
