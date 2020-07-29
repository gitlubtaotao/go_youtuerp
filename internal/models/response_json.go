/*
	response struct not table struct
	not create table
*/
package models

import "time"

// 查询订单对应的response
type ResponseOrderMaster struct {
	ID               uint       `json:"id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	SerialNumber     string     `json:"serial_number"`
	InstructionId    uint       `json:"instruction_id"`
	SalesmanId       uint       `json:"salesman_id"`
	OperationId      uint       `json:"operation_id" validate:"required"`
	TransportType    uint       `json:"transport_type" validate:"required" `
	Status           string     `json:"status" validate:"required"`
	CompanyId        uint       `json:"company_id"`
	ContactId        uint       `json:"contact_id"`
	MainTransport    uint       `json:"main_transport"`
	PayableStatus    string     ` json:"payable_status"`
	PaidStatus       string     `json:"paid_status"`
	ReceivableStatus string     `json:"receivable_status"`
	ReceivedStatus   string     `json:"received_status"`
	Remarks          string     `gorm:"size:522" json:"remarks"`
	OrderMasterId    uint       `sql:"index" json:"order_master_id"`
	Number           uint       `comment:"包装数量" json:"number"`
	PackageTypeId    uint       `comment:"包装类型" json:"package_type_id"`
	GrossWeight      string     `gorm:"size:64" comment:"毛重" json:"gross_weight"`
	Volume           string     `gorm:"size:64" comment:"体积" json:"volume"`
	MblSO            string     `gorm:"size:16;index:mbl_so" json:"mbl_so"`
	HblSO            string     `gorm:"size:16" json:"hbl_so"`
	SONo             string     `gorm:"size:16;index:so_no" json:"so_no"`
	CarrierId        uint       `sql:"index" json:"carrier_id"`
	POLId            uint       `gorm:"index:pol_pod_and_index" json:"pol_id"`
	PODId            uint       `gorm:"index:pol_pod_and_index" json:"pod_id"`
	POTId            uint       `json:"pot_id"`
	CutOffDay        *time.Time `sql:"index" json:"cut_off_day"`
	Departure        *time.Time `sql:"index" json:"departure"`
	Arrival          *time.Time `sql:"index" json:"arrival"`
	Vessel           string     `gorm:"size:64" json:"vessel"`
	Voyage           string     `gorm:"size:64" json:"voyage"`
	FlightNo         string     `gorm:"size:16;index:flight_no" json:"flight_no"`
	ChargedWeight    float64    `gorm:"default:0.0;PRECISION:4" comment:"计费重量" json:"charged_weight"`
	CourierCodeId    uint       `sql:"index" json:"courier_code_id"`
	CourierNo        string     `gorm:"size:16;index:courier_no" json:"courier_no"`
	ShipperId        uint       `json:"shipper_id"`
	SupplyAgentId    uint       `sql:"index" json:"supply_agent_id"`
	//Roles            []Role     `json:"roles"`
}

//系统流水好规则设置对应的response json
type ResponseNumberSetting struct {
	ID                    uint      `json:"id"`
	CreatedAt             time.Time `json:"created_at"`
	UserCompanyId         uint      `json:"user_company_id"`
	UserCompaniesNameNick string    `json:"user_companies_name_nick"`
	Prefix                string    `json:"prefix"`
	Length                uint      `json:"length"`
	DefaultRule           string    `json:"default_rule"`
	Special               string    `json:"special"`
	ApplicationNo         string    `json:"application_no"`
	DefaultNumber         uint      `json:"default_number"`
	CurrentNumber         uint      `json:"current_number"`
	ClearRule             string    `json:"clear_rule"`
}

type ResponseSetting struct {
	Key   string `json:"key"`
	Field string `json:"field"`
	Value string `json:"value"`
}

//查询部门信息对应的response
type ResponseDepartment struct {
	ID                    uint      `json:"id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	NameCn                string    `json:"name_cn"`
	NameEn                string    `json:"name_en"`
	UserCompanyId         int       `json:"user_company_id"`
	UserCompaniesNameNick string    `json:"user_companies_name_nick"`
}

//查询员工信息对应的response
type ResponseEmployee struct {
	ID                    uint      `json:"id"`
	CreatedAt             time.Time `json:"created_at"`
	Email                 string    `json:"email"`
	SignInCount           int       `json:"sign_in_count"`
	LastSignInAt          time.Time `json:"last_sign_in_at"`
	LastSignInIp          string    `json:"last_sign_in_ip"`
	UserCompanyId         int       `json:"user_company_id"`
	UserCompaniesNameNick string    `json:"user_companies_name_nick"`
	DepartmentId          int       `json:"department_id"`
	DepartmentsNameCN     string    `json:"departments_name_cn"`
	Name                  string    `json:"name"`
	UserNo                string    `json:"user_no"`
	Phone                 string    `json:"phone"`
	Remarks               string    `json:"remarks"`
	Sex                   uint      `json:"sex"`
	Address               string    `json:"address"`
}

//查询对账列表对应response struct
type ResponseFinanceFee struct {
	ID                  uint       `gorm:"primary_key"json:"id"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	DeletedAt           *time.Time `sql:"index"`
	Name                string     `gorm:"size:64;index:name" json:"name" validate:"required"`
	NameCn              string     `json:"name_cn"`
	NameEn              string     `json:"name_en"`
	PayOrReceive        string     `gorm:"size:16;index:pay_or_receive" comment:"类型" json:"pay_or_receive" validate:"required"`
	PayTypeId           uint       `gorm:"index:pay_type_id" comment:"结算方式" json:"pay_type_id"`
	FinanceCurrencyId   uint       `gorm:"index:finance_currency_id" comment:"币种类型" json:"finance_currency_id" validate:"required"`
	FinanceCurrencyRate float64    `gorm:"PRECISION: 4" json:"finance_currency_rate" validate:"required"`
	Quantity            float64    `gorm:"PRECISION: 4" json:"quantity" validate:"required"`
	UnitPrice           float64    `gorm:"PRECISION:4" json:"unit_price" validate:"required"`
	TaxRate             float64    `gorm:"PRECISION:4" json:"tax_rate"`
	TaxAmount           float64    `gorm:"PRECISION:4" json:"tax_amount" validate:"required"`
	NotTaxAmount        float64    `gorm:"PRECISION:4" json:"not_tax_amount"`
	ReceiveAmount       float64    `gorm:"PRECISION:4" json:"receive_amount"`
	Receivable          float64    `gorm:"PRECISION:4" json:"receivable"`
	PayAmount           float64    `gorm:"PRECISION:4" json:"pay_amount"`
	Payable             float64    `gorm:"PRECISION:4" json:"payable"`
	ClosingUnitId       uint       `gorm:"index:closing_unit_id" json:"closing_unit_id"`
	OrderMasterId       uint       `gorm:"index:order_master_id" json:"order_master_id"`
	SerialNumber        string     `json:"serial_number"`
	FinanceStatementId  uint       `json:"finance_statement_id"`
	FinanceStatementNo  string     `gorm:"size:32;index:finance_statement_no" json:"finance_statement_no"`
	Status              string     `gorm:"size:16;index:status;default:'init'" json:"status"`
	InvoiceAmount       float64    `gorm:"PRECISION:4" json:"invoice_amount"`
	InvoiceStatus       string     `gorm:"size:16;index:invoice_status;default:'init'" json:"invoice_status"`
	TypeId              uint       `gorm:"index:type_id;default:1" json:"type_id"`
	Remarks             string     `gorm:"size:128" json:"remarks"`
}

func (ResponseOrderMaster) TableName() string {
	return "order_masters"
}
func (ResponseNumberSetting) TableName() string {
	return "number_settings"
}

func (ResponseDepartment) TableName() string {
	return "departments"
}

func (ResponseEmployee) TableName() string {
	return "users"
}
