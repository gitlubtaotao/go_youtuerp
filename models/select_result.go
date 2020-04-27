package models

//前端下拉选择框
type SelectResult struct {
	ID           uint   `json:"id"`
	NameNick     string `json:"name_nick"`
	NameCn       string `json:"name_cn"`
	NameEn       string `json:"name_en"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}
