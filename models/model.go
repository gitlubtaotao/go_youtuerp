package models

import "time"

type Base struct {
	ID        uint       `gorm:"primary_key"json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
}

//
//type Finance_banks struct {
//	Id int `gorm:"id" json:"id"`
//	BankCn string `gorm:"bank_cn" json:"bank_cn"`
//	BankEn string `gorm:"bank_en" json:"bank_en"` // 英文名称
//	BankAbbr string `gorm:"bank_abbr" json:"bank_abbr"` // 英文简写
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Finance_banks) TableName() string {
//	return "finance_banks"
//}
//
//type Former_sea_deliveries struct {
//	Id int `gorm:"id" json:"id"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	PolId int `gorm:"pol_id" json:"pol_id"` // 起运港
//	PodId int `gorm:"pod_id" json:"pod_id"` // 目的港
//	EstimatedTimeOfDeparture string `gorm:"estimated_time_of_departure" json:"estimated_time_of_departure"` // 离港日期
//	EstimatedTimeOfArrival string `gorm:"estimated_time_of_arrival" json:"estimated_time_of_arrival"` // 到港日期
//	CourierCompanyId int `gorm:"courier_company_id" json:"courier_company_id"` // 快递公司
//	CourierNumber string `gorm:"courier_number" json:"courier_number"` // 快递单号
//	BoatCompanyId int `gorm:"boat_company_id" json:"boat_company_id"` // 船公司
//	Vessel string `gorm:"vessel" json:"vessel"` // 船名
//	Voyage string `gorm:"voyage" json:"voyage"` // 航次
//	WarehousingDate string `gorm:"warehousing_date" json:"warehousing_date"` // 入仓时间
//	DeliverDate string `gorm:"deliver_date" json:"deliver_date"` // 派送日期
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"` // 委托单位
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"` // 委托单位类型
//	CompanyInstructionContent string `gorm:"company_instruction_content" json:"company_instruction_content"` // 委托单位详情
//	ShipperId int `gorm:"shipper_id" json:"shipper_id"` // 发货人
//	ShipperType string `gorm:"shipper_type" json:"shipper_type"`
//	ShipperContent string `gorm:"shipper_content" json:"shipper_content"` // 发货人详情
//	ConsigneeId int `gorm:"consignee_id" json:"consignee_id"` // 收货人
//	ConsigneeType string `gorm:"consignee_type" json:"consignee_type"`
//	ConsigneeContent string `gorm:"consignee_content" json:"consignee_content"` // 收货人详情
//	NotifyPartyId int `gorm:"notify_party_id" json:"notify_party_id"` // 通知人
//	NotifyPartyType string `gorm:"notify_party_type" json:"notify_party_type"`
//	NotifyPartyContent string `gorm:"notify_party_content" json:"notify_party_content"` // 通知人详情
//	OceanChangesPaytypeId int `gorm:"ocean_changes_paytype_id" json:"ocean_changes_paytype_id"` // 运费付款方式
//	OtherChangesPaytypeId int `gorm:"other_changes_paytype_id" json:"other_changes_paytype_id"` // 其他付款方式
//	InstructionTypeId int `gorm:"instruction_type_id" json:"instruction_type_id"` // 委托类型
//	WaysOfDeclarationId int `gorm:"ways_of_declaration_id" json:"ways_of_declaration_id"` // 报关方式
//	BaseDataTradeTermsId int `gorm:"base_data_trade_terms_id" json:"base_data_trade_terms_id"` // 贸易条款
//	TransshipmentId int `gorm:"transshipment_id" json:"transshipment_id"` // 转运
//	BaseDataItemId int `gorm:"base_data_item_id" json:"base_data_item_id"` // 装运条款
//	HsCode string `gorm:"hs_code" json:"hs_code"` // 商品编码
//	ServiceContractNo string `gorm:"service_contract_no" json:"service_contract_no"` // 合同编码
//	InvoiceNo string `gorm:"invoice_no" json:"invoice_no"` // 发票号
//	Insurance int `gorm:"insurance" json:"insurance"` // 是否需要发票
//	InsuranceAmount float64 `gorm:"insurance_amount" json:"insurance_amount"` // 保险金额
//	ViaId int `gorm:"via_id" json:"via_id"` // 中转港
//	ReceiverAddress string `gorm:"receiver_address" json:"receiver_address"` // 收货地址
//	ConsignorAddress string `gorm:"consignor_address" json:"consignor_address"` // 发货地址
//	ReceiverCityId int `gorm:"receiver_city_id" json:"receiver_city_id"` // 收货城市
//	ConsignorCityId int `gorm:"consignor_city_id" json:"consignor_city_id"` // 发货城市
//	BargeDate string `gorm:"barge_date" json:"barge_date"` // 驳船日期
//	BargeId int `gorm:"barge_id" json:"barge_id"` // 驳船名
//	Enabled int `gorm:"enabled" json:"enabled"` // 有效状态
//	Status string `gorm:"status" json:"status"` // 状态
//	Note string `gorm:"note" json:"note"` // 备注
//	ClientNote string `gorm:"client_note" json:"client_note"` // 客户备注
//	SerialNumber string `gorm:"serial_number" json:"serial_number"` // 流水号
//	Marks string `gorm:"marks" json:"marks"` // 唛头
//	Number string `gorm:"number" json:"number"` // 数量
//	DescriptionOfGood string `gorm:"description_of_good" json:"description_of_good"` // 货物描述
//	GrossWeight string `gorm:"gross_weight" json:"gross_weight"` // 毛重
//	Size string `gorm:"size" json:"size"` // 体积
//	SerialNo string `gorm:"serial_no" json:"serial_no"` // 序列号
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Ratio float64 `gorm:"ratio" json:"ratio"` // 材积换算系数
//	RatioWeight float64 `gorm:"ratio_weight" json:"ratio_weight"` // 材积重
//	Bubble float64 `gorm:"bubble" json:"bubble"` // 分泡%
//	Dimension string `gorm:"dimension" json:"dimension"` // 体积
//	ChargedWeight float64 `gorm:"charged_weight" json:"charged_weight"` // 计费重
//	BaseDataPackageType int `gorm:"base_data_package_type" json:"base_data_package_type"` // 包装类型
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Former_sea_deliveries) TableName() string {
//	return "former_sea_deliveries"
//}
//
//type Order_masters struct {
//	Platform string `gorm:"platform" json:"platform"`
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	SerialNumber string `gorm:"serial_number" json:"serial_number"` // 流水号
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	UserSalesmanId int `gorm:"user_salesman_id" json:"user_salesman_id"`
//	UserOperatorId int `gorm:"user_operator_id" json:"user_operator_id"`
//	Remark string `gorm:"remark" json:"remark"` // 描述
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	RawInfoJson string `gorm:"raw_info_json" json:"raw_info_json"` // 原始信息
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//	AasmState string `gorm:"aasm_state" json:"aasm_state"` // aasm状态机
//	TransportType int `gorm:"transport_type" json:"transport_type"`
//	CompanyId int `gorm:"company_id" json:"company_id"`
//	PayStatus string `gorm:"pay_status" json:"pay_status"`
//	ReceiveStatus string `gorm:"receive_status" json:"receive_status"`
//	BookingDate string `gorm:"booking_date" json:"booking_date"` // 和订舱日期保持一致，用于搜索的冗余字段
//	ReceivableLockDate string `gorm:"receivable_lock_date" json:"receivable_lock_date"` // 费用结清日期
//	PayableStatus string `gorm:"payable_status" json:"payable_status"`
//	ReceiveableStatus string `gorm:"receiveable_status" json:"receiveable_status"`
//	SupplierCompanyAgentId int `gorm:"supplier_company_agent_id" json:"supplier_company_agent_id"`
//	CreditNoteStatus string `gorm:"credit_note_status" json:"credit_note_status"` // 应付对账状态
//	DebitNoteStatus string `gorm:"debit_note_status" json:"debit_note_status"` // 应收对账状态
//	UserFileId int `gorm:"user_file_id" json:"user_file_id"`
//	UserMarketId int `gorm:"user_market_id" json:"user_market_id"`
//	UserCustomerId int `gorm:"user_customer_id" json:"user_customer_id"`
//	UserFeeId int `gorm:"user_fee_id" json:"user_fee_id"`
//	UserContactId int `gorm:"user_contact_id" json:"user_contact_id"` // 订单委托单位的联系人
//	MainTransport int `gorm:"main_transport" json:"main_transport"` // 主要运输方式
//}
//
//func (*Order_masters) TableName() string {
//	return "order_masters"
//}
//
//type Base_data_cap_type_sizes struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"` // 标题
//	CapSize string `gorm:"cap_size" json:"cap_size"` // 箱尺寸
//	CapType string `gorm:"cap_type" json:"cap_type"` // 箱类型
//	Remark string `gorm:"remark" json:"remark"` // 说明
//	IsoCode string `gorm:"iso_code" json:"iso_code"` // ISO代码
//	IsoRemark string `gorm:"iso_remark" json:"iso_remark"` // ISO代码说明
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//	ExtLength string `gorm:"ext_length" json:"ext_length"` // 外长
//	ExtWidth string `gorm:"ext_width" json:"ext_width"` // 外宽
//	ExtHeight string `gorm:"ext_height" json:"ext_height"` // 外高
//	IntLength string `gorm:"int_length" json:"int_length"` // 内长
//	IntWidth string `gorm:"int_width" json:"int_width"` // 内宽
//	IntHeight string `gorm:"int_height" json:"int_height"` // 内高
//	DoorWidth string `gorm:"door_width" json:"door_width"` // 箱门宽
//	DoorHeight string `gorm:"door_height" json:"door_height"` // 箱门高
//	GrossWeight string `gorm:"gross_weight" json:"gross_weight"` // 总重
//	Tare string `gorm:"tare" json:"tare"` // 自重
//	Payload string `gorm:"payload" json:"payload"` // 载重
//	CubicCapacity string `gorm:"cubic_capacity" json:"cubic_capacity"` // 容积
//	Ventilation string `gorm:"ventilation" json:"ventilation"` // 通风量
//	HatchDiameter string `gorm:"hatch_diameter" json:"hatch_diameter"` // 舱口直径
//	SelectCount int `gorm:"select_count" json:"select_count"` // 选中的次数,用于排序常用
//}
//
//func (*Base_data_cap_type_sizes) TableName() string {
//	return "base_data_cap_type_sizes"
//}
//
//type Base_data_courier_codes struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Name string `gorm:"name" json:"name"`
//	Code string `gorm:"code" json:"code"`
//	CourierType int `gorm:"courier_type" json:"courier_type"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Base_data_courier_codes) TableName() string {
//	return "base_data_courier_codes"
//}
//
//type Base_data_trade_terms struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//}
//
//func (*Base_data_trade_terms) TableName() string {
//	return "base_data_trade_terms"
//}
//
//type Finance_approval_applications struct {
//	Id int `gorm:"id" json:"id"`
//	FinanceFeeId int `gorm:"finance_fee_id" json:"finance_fee_id"`
//	ApprovalApplicationId int `gorm:"approval_application_id" json:"approval_application_id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Finance_approval_applications) TableName() string {
//	return "finance_approval_applications"
//}
//
//type Finance_fees struct {
//	Id int `gorm:"id" json:"id"`
//	PayOrReceive string `gorm:"pay_or_receive" json:"pay_or_receive"` // 填pay 或 receive
//	Name string `gorm:"name" json:"name"` // 费用名称
//	NameCn string `gorm:"name_cn" json:"name_cn"` // 费用中文名
//	NameEn string `gorm:"name_en" json:"name_en"` // 费用英文名
//	PayTypeId int `gorm:"pay_type_id" json:"pay_type_id"` // 结算方式
//	FinanceCurrencyId int `gorm:"finance_currency_id" json:"finance_currency_id"`
//	Quantity float64 `gorm:"quantity" json:"quantity"` // 数量
//	UnitPrice float64 `gorm:"unit_price" json:"unit_price"` // 单价
//	TaxRate float64 `gorm:"tax_rate" json:"tax_rate"` // 税率
//	TaxAmount float64 `gorm:"tax_amount" json:"tax_amount"` // 含税金额
//	StandardCurrencyExchangeRate float64 `gorm:"standard_currency_exchange_rate" json:"standard_currency_exchange_rate"` // 本位币汇率
//	Remark string `gorm:"remark" json:"remark"` // 备注/附加说明
//	DebitNote string `gorm:"debit_note" json:"debit_note"` // 借方通知单
//	ReceiveAmount float64 `gorm:"receive_amount" json:"receive_amount"` // 实收
//	Receivable float64 `gorm:"receivable" json:"receivable"` // 应收
//	CreditNote string `gorm:"credit_note" json:"credit_note"` // 收款通知单
//	PayAmount float64 `gorm:"pay_amount" json:"pay_amount"` // 实付
//	Payable float64 `gorm:"payable" json:"payable"` // 应付
//	ClosingUnitType string `gorm:"closing_unit_type" json:"closing_unit_type"`
//	ClosingUnitId int `gorm:"closing_unit_id" json:"closing_unit_id"` // 结算单位
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	SourceType string `gorm:"source_type" json:"source_type"`
//	SourceId int `gorm:"source_id" json:"source_id"` // 数据来源
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//	Status string `gorm:"status" json:"status"` // 财务账单状态
//	Locked int `gorm:"locked" json:"locked"` // 财务状态锁
//	InvoicesId int `gorm:"invoices_id" json:"invoices_id"`
//	AccountId int `gorm:"account_id" json:"account_id"`
//	FinanceCurrencyRate float64 `gorm:"finance_currency_rate" json:"finance_currency_rate"`
//	AddressId int `gorm:"address_id" json:"address_id"`
//	DebitNoteSn string `gorm:"debit_note_sn" json:"debit_note_sn"` // 账单流水号
//	VerifyStatus string `gorm:"verify_status" json:"verify_status"` // 对账状态
//	InvoiceAmount float64 `gorm:"invoice_amount" json:"invoice_amount"` // 开票金额
//	InvoiceStatus string `gorm:"invoice_status" json:"invoice_status"` // 开票状态
//	AccountTitleId int `gorm:"account_title_id" json:"account_title_id"` // 会计类目/费用种类
//	FinanceStatementId int `gorm:"finance_statement_id" json:"finance_statement_id"` // 对账单id
//	PlanFeeId int `gorm:"plan_fee_id" json:"plan_fee_id"`
//}
//
//func (*Finance_fees) TableName() string {
//	return "finance_fees"
//}
//
//type Former_seas_instructions struct {
//	Id int `gorm:"id" json:"id"`
//	BaseDataInstructionTypeId int `gorm:"base_data_instruction_type_id" json:"base_data_instruction_type_id"`
//	SerialNo string `gorm:"serial_no" json:"serial_no"` // 序列号/流水号
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"`
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"`
//	CompanyInstructionContent string `gorm:"company_instruction_content" json:"company_instruction_content"` // 委托公司详细要求
//	ContactInstructionType string `gorm:"contact_instruction_type" json:"contact_instruction_type"`
//	ContactInstructionId int `gorm:"contact_instruction_id" json:"contact_instruction_id"`
//	ContactInstructionContent string `gorm:"contact_instruction_content" json:"contact_instruction_content"` // 委托公司联系人详细要求内容
//	ShipperType string `gorm:"shipper_type" json:"shipper_type"`
//	ShipperId int `gorm:"shipper_id" json:"shipper_id"`
//	ShipperContent string `gorm:"shipper_content" json:"shipper_content"` // 发货人要求
//	ConsigneeType string `gorm:"consignee_type" json:"consignee_type"`
//	ConsigneeId int `gorm:"consignee_id" json:"consignee_id"`
//	ConsigneeContent string `gorm:"consignee_content" json:"consignee_content"` // 收货人要求
//	NotifyPartyType string `gorm:"notify_party_type" json:"notify_party_type"`
//	NotifyPartyId int `gorm:"notify_party_id" json:"notify_party_id"`
//	NotifyPartyContent string `gorm:"notify_party_content" json:"notify_party_content"` // 通知内容
//	PodAgentType string `gorm:"pod_agent_type" json:"pod_agent_type"`
//	PodAgentId int `gorm:"pod_agent_id" json:"pod_agent_id"`
//	PodAgentContent string `gorm:"pod_agent_content" json:"pod_agent_content"` // 目的港代理详情
//	Marks string `gorm:"marks" json:"marks"` // 标记
//	Number string `gorm:"number" json:"number"` // 包装件数
//	DescriptionOfGood string `gorm:"description_of_good" json:"description_of_good"` // 品名
//	GrossWeight string `gorm:"gross_weight" json:"gross_weight"` // 毛重
//	Size string `gorm:"size" json:"size"` // 体积
//	InvoiceNo string `gorm:"invoice_no" json:"invoice_no"` // 发票号
//	SeaPortPolId int `gorm:"sea_port_pol_id" json:"sea_port_pol_id"` // 装货港
//	SeaPortViaId int `gorm:"sea_port_via_id" json:"sea_port_via_id"` // 中转港
//	SeaPortPodId int `gorm:"sea_port_pod_id" json:"sea_port_pod_id"` // 目的港
//	TransshipmentId int `gorm:"transshipment_id" json:"transshipment_id"`
//	ServiceContractNo string `gorm:"service_contract_no" json:"service_contract_no"` // 合同编号
//	BaseDataItemId int `gorm:"base_data_item_id" json:"base_data_item_id"`
//	BaseDataTradeTermsId int `gorm:"base_data_trade_terms_id" json:"base_data_trade_terms_id"`
//	WaysOfDeclarationId int `gorm:"ways_of_declaration_id" json:"ways_of_declaration_id"`
//	WaysOfTransportationId int `gorm:"ways_of_transportation_id" json:"ways_of_transportation_id"`
//	Remarks string `gorm:"remarks" json:"remarks"` // 备注
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//	Status string `gorm:"status" json:"status"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	EstimatedTimeOfDeparture string `gorm:"estimated_time_of_departure" json:"estimated_time_of_departure"` // 开船日期
//	EstimatedTimeOfArrival string `gorm:"estimated_time_of_arrival" json:"estimated_time_of_arrival"` // 到港日期
//	CutOffDate string `gorm:"cut_off_date" json:"cut_off_date"` // 截关日期
//	BoatCompanyId int `gorm:"boat_company_id" json:"boat_company_id"` // 船公司
//	Vessel string `gorm:"vessel" json:"vessel"` // 船名
//	Voyage string `gorm:"voyage" json:"voyage"` // 航次
//	CargoReceivedDate string `gorm:"cargo_received_date" json:"cargo_received_date"`
//	BaseDataPackageType int `gorm:"base_data_package_type" json:"base_data_package_type"`
//	BoxSizeCount string `gorm:"box_size_count" json:"box_size_count"`
//	OceanChangesPaytypeId int `gorm:"ocean_changes_paytype_id" json:"ocean_changes_paytype_id"` // 运费支付方式
//	OtherChangesPaytypeId int `gorm:"other_changes_paytype_id" json:"other_changes_paytype_id"` // 其他运费支付方式
//	BaseDataMiscBillId int `gorm:"base_data_misc_bill_id" json:"base_data_misc_bill_id"`
//	Dimension string `gorm:"dimension" json:"dimension"` // 尺寸
//	Type string `gorm:"type" json:"type"`
//	BLNo string `gorm:"b_l_no" json:"b_l_no"` // 分单号
//	HblReleaseDate string `gorm:"hbl_release_date" json:"hbl_release_date"` // 分单释放时间
//	IsCabinetAttachment int `gorm:"is_cabinet_attachment" json:"is_cabinet_attachment"`
//	IsGoodsAttachment int `gorm:"is_goods_attachment" json:"is_goods_attachment"`
//	HblRemark string `gorm:"hbl_remark" json:"hbl_remark"` // HB/L备注
//	PlaceOfReceipt string `gorm:"place_of_receipt" json:"place_of_receipt"`
//	PlaceOfDelivery string `gorm:"place_of_delivery" json:"place_of_delivery"`
//	ShippedOnBoardDate string `gorm:"shipped_on_board_date" json:"shipped_on_board_date"`
//	VerifyDate string `gorm:"verify_date" json:"verify_date"`
//	UserVerifyId int `gorm:"user_verify_id" json:"user_verify_id"`
//	TlxNo string `gorm:"tlx_no" json:"tlx_no"`
//	PlaceOfIssue string `gorm:"place_of_issue" json:"place_of_issue"`
//	DateOfIssue string `gorm:"date_of_issue" json:"date_of_issue"`
//}
//
//func (*Former_seas_instructions) TableName() string {
//	return "former_seas_instructions"
//}
//
//type Former_seas_so_nos struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Vessel string `gorm:"vessel" json:"vessel"`
//	Voyage string `gorm:"voyage" json:"voyage"` // 航次
//	SeaPortPodId int `gorm:"sea_port_pod_id" json:"sea_port_pod_id"`
//	SeaPortPolId int `gorm:"sea_port_pol_id" json:"sea_port_pol_id"`
//	SeaPortViaId int `gorm:"sea_port_via_id" json:"sea_port_via_id"`
//	SoNo string `gorm:"so_no" json:"so_no"`
//	PodAgentType string `gorm:"pod_agent_type" json:"pod_agent_type"`
//	PodAgentId int `gorm:"pod_agent_id" json:"pod_agent_id"`
//	PodAgentContent string `gorm:"pod_agent_content" json:"pod_agent_content"`
//	ShippedOnBoardDate string `gorm:"shipped_on_board_date" json:"shipped_on_board_date"` // 装船日期
//	OceanChangesPaytypeId int `gorm:"ocean_changes_paytype_id" json:"ocean_changes_paytype_id"` // 运费付款方式
//	OtherChangesPaytypeId int `gorm:"other_changes_paytype_id" json:"other_changes_paytype_id"` // 其他运费付款方式
//	CyOpenDate string `gorm:"cy_open_date" json:"cy_open_date"` // 开舱时间
//	VoucherCutOff string `gorm:"voucher_cut_off" json:"voucher_cut_off"` // 截放行条时间
//	VgmSubmissionDeadline string `gorm:"vgm_submission_deadline" json:"vgm_submission_deadline"` // VGM截止时间
//	SiCutOff string `gorm:"si_cut_off" json:"si_cut_off"` // 截补料时间
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	SerialNo string `gorm:"serial_no" json:"serial_no"`
//	Status string `gorm:"status" json:"status"`
//	BoatCompanyId int `gorm:"boat_company_id" json:"boat_company_id"`
//	CutOffDate string `gorm:"cut_off_date" json:"cut_off_date"`
//	EstimatedTimeOfDeparture string `gorm:"estimated_time_of_departure" json:"estimated_time_of_departure"`
//	EstimatedTimeOfArrival string `gorm:"estimated_time_of_arrival" json:"estimated_time_of_arrival"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Former_seas_so_nos) TableName() string {
//	return "former_seas_so_nos"
//}
//
//type Plan_fees struct {
//	Id int `gorm:"id" json:"id"`
//	PlanMainId int `gorm:"plan_main_id" json:"plan_main_id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	ClosingUnitId int `gorm:"closing_unit_id" json:"closing_unit_id"` // 结算单位
//	Name string `gorm:"name" json:"name"` // 简称
//	NameCn string `gorm:"name_cn" json:"name_cn"` // 中文名称
//	NameEn string `gorm:"name_en" json:"name_en"` // 英文名称
//	CurrencyId int `gorm:"currency_id" json:"currency_id"` // 币种
//	UnitPrice float64 `gorm:"unit_price" json:"unit_price"` // 单价
//	Count float64 `gorm:"count" json:"count"` // 数量
//	Amount float64 `gorm:"amount" json:"amount"` // 总价
//	PayOrReceive string `gorm:"pay_or_receive" json:"pay_or_receive"` // 收支类型
//	Enabled int `gorm:"enabled" json:"enabled"` // 有效状态
//	IsSync int `gorm:"is_sync" json:"is_sync"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Plan_fees) TableName() string {
//	return "plan_fees"
//}
//
//type Base_data_sea_lines_ports struct {
//	Id int `gorm:"id" json:"id"`
//	SeaLineId int `gorm:"sea_line_id" json:"sea_line_id"` // 航线
//	SeaPortId int `gorm:"sea_port_id" json:"sea_port_id"` // 港口
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//}
//
//func (*Base_data_sea_lines_ports) TableName() string {
//	return "base_data_sea_lines_ports"
//}
//
//type Former_seas_importor_security_items struct {
//	Id int `gorm:"id" json:"id"`
//	ImportorSecurityId int `gorm:"importor_security_id" json:"importor_security_id"`
//	PoNo string `gorm:"po_no" json:"po_no"` // PO号
//	Item string `gorm:"item" json:"item"` // 商品
//	Description string `gorm:"description" json:"description"` // 描述
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	HsCode string `gorm:"hs_code" json:"hs_code"` // HS商品编码(前六位)
//}
//
//func (*Former_seas_importor_security_items) TableName() string {
//	return "former_seas_importor_security_items"
//}
//
//type Invoice_apply_lists struct {
//	Id int `gorm:"id" json:"id"`
//	InvoiceId int `gorm:"invoice_id" json:"invoice_id"`
//	AddressId int `gorm:"address_id" json:"address_id"`
//	Category string `gorm:"category" json:"category"`
//	Number string `gorm:"number" json:"number"` // 流水号
//	Status string `gorm:"status" json:"status"`
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	ClosingUnitId int `gorm:"closing_unit_id" json:"closing_unit_id"`
//	ApplyUserId int `gorm:"apply_user_id" json:"apply_user_id"`
//	FinanceCurrencyId int `gorm:"finance_currency_id" json:"finance_currency_id"`
//	Amount float64 `gorm:"amount" json:"amount"`
//	FinanceFeeIds string `gorm:"finance_fee_ids" json:"finance_fee_ids"`
//	Note string `gorm:"note" json:"note"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Invoice_apply_lists) TableName() string {
//	return "invoice_apply_lists"
//}
//
//type Bargain_box_size_counts struct {
//	Id int `gorm:"id" json:"id"`
//	BoxSizeId int `gorm:"box_size_id" json:"box_size_id"`
//	BoxSizeCount int `gorm:"box_size_count" json:"box_size_count"`
//	TargetPrice float64 `gorm:"target_price" json:"target_price"`
//	FinanceCurrency string `gorm:"finance_currency" json:"finance_currency"` // 币种
//	FinishedPrice float64 `gorm:"finished_price" json:"finished_price"` // 成交价格
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	BargainMainId int `gorm:"bargain_main_id" json:"bargain_main_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Bargain_box_size_counts) TableName() string {
//	return "bargain_box_size_counts"
//}
//
//type Former_seas_ams struct {
//	Id int `gorm:"id" json:"id"`
//	ShipperId int `gorm:"shipper_id" json:"shipper_id"` // 承运人，从cooperator中选择
//	SoNo string `gorm:"so_no" json:"so_no"` // so编号
//	ConsigneeId int `gorm:"consignee_id" json:"consignee_id"` // 收货人，从cooperator选择
//	NotifyPartyId int `gorm:"notify_party_id" json:"notify_party_id"`
//	PrecarriageBy string `gorm:"precarriage_by" json:"precarriage_by"` // 头程运输
//	PlaceOfReceipt string `gorm:"place_of_receipt" json:"place_of_receipt"` // 接货地
//	OceanVessel string `gorm:"ocean_vessel" json:"ocean_vessel"` // 航线
//	SeaPortPodId int `gorm:"sea_port_pod_id" json:"sea_port_pod_id"`
//	SeaPortPolId int `gorm:"sea_port_pol_id" json:"sea_port_pol_id"`
//	PlaceOfDelivery string `gorm:"place_of_delivery" json:"place_of_delivery"` // 送货地
//	ShipmentTerm string `gorm:"shipment_term" json:"shipment_term"` // 运输条款
//	Marks string `gorm:"marks" json:"marks"` // 唛头
//	Number string `gorm:"number" json:"number"`
//	DescriptionOfGood string `gorm:"description_of_good" json:"description_of_good"` // 货物描述
//	GrossWeight string `gorm:"gross_weight" json:"gross_weight"` // 总重，单位kg
//	Size string `gorm:"size" json:"size"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	Status string `gorm:"status" json:"status"`
//	InvoiceNo string `gorm:"invoice_no" json:"invoice_no"` // 发票号
//	SerialNo string `gorm:"serial_no" json:"serial_no"`
//	ShipperType string `gorm:"shipper_type" json:"shipper_type"`
//	ShipperContent string `gorm:"shipper_content" json:"shipper_content"`
//	ConsigneeType string `gorm:"consignee_type" json:"consignee_type"`
//	ConsigneeContent string `gorm:"consignee_content" json:"consignee_content"`
//	NotifyPartyType string `gorm:"notify_party_type" json:"notify_party_type"`
//	NotifyPartyContent string `gorm:"notify_party_content" json:"notify_party_content"`
//	BaseDataPackageType int `gorm:"base_data_package_type" json:"base_data_package_type"`
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"`
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"`
//	CompanyInstructionContent string `gorm:"company_instruction_content" json:"company_instruction_content"`
//	Vessel string `gorm:"vessel" json:"vessel"`
//	Voyage string `gorm:"voyage" json:"voyage"`
//	PodAgentContent string `gorm:"pod_agent_content" json:"pod_agent_content"`
//	PodAgentId string `gorm:"pod_agent_id" json:"pod_agent_id"`
//	PodAgentType string `gorm:"pod_agent_type" json:"pod_agent_type"`
//	ShippedOnBoardDate string `gorm:"shipped_on_board_date" json:"shipped_on_board_date"`
//	Dimension string `gorm:"dimension" json:"dimension"`
//}
//
//func (*Former_seas_ams) TableName() string {
//	return "former_seas_ams"
//}
//
//type User_companies_copy1 struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	CompanyType int `gorm:"company_type" json:"company_type"`
//	ParentId int `gorm:"parent_id" json:"parent_id"` // 父级id,区分谁的客户
//	UserSalesmanId int `gorm:"user_salesman_id" json:"user_salesman_id"` // 所属的业务人员
//	IsHeadOffice int `gorm:"is_head_office" json:"is_head_office"` // 是否为总部
//	AccountPeriod string `gorm:"account_period" json:"account_period"` // 公司结算类型
//	Age int `gorm:"age" json:"age"` // 公司账龄
//	Amount float64 `gorm:"amount" json:"amount"` // 月结金额
//	IsBlack int `gorm:"is_black" json:"is_black"` // 是否加入黑名单
//	NameNick string `gorm:"name_nick" json:"name_nick"`
//	NameCn string `gorm:"name_cn" json:"name_cn"`
//	NameEn string `gorm:"name_en" json:"name_en"`
//	BusinessTypeName string `gorm:"business_type_name" json:"business_type_name"`
//	ScaleGroupId int `gorm:"scale_group_id" json:"scale_group_id"`
//	Status string `gorm:"status" json:"status"`
//}
//
//func (*User_companies_copy1) TableName() string {
//	return "user_companies_copy1"
//}
//
//type User_companies_freight_charges struct {
//	Id int `gorm:"id" json:"id"`
//	FreightChargeId int `gorm:"freight_charge_id" json:"freight_charge_id"`
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	Enabled int `gorm:"enabled" json:"enabled"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//}
//
//func (*User_companies_freight_charges) TableName() string {
//	return "user_companies_freight_charges"
//}
//
//type Website_page_lists struct {
//	Id int `gorm:"id" json:"id"`
//	SourceType string `gorm:"source_type" json:"source_type"`
//	SourceId int `gorm:"source_id" json:"source_id"`
//	Title string `gorm:"title" json:"title"`
//	Details string `gorm:"details" json:"details"` // 内容信息
//	UrlInfo string `gorm:"url_info" json:"url_info"` // 链接信息
//	CatalogType string `gorm:"catalog_type" json:"catalog_type"` // 分类信息
//}
//
//func (*Website_page_lists) TableName() string {
//	return "website_page_lists"
//}
//
//type Youtu_erp_roles struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"`
//	ResourceType string `gorm:"resource_type" json:"resource_type"`
//	ResourceId int `gorm:"resource_id" json:"resource_id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	IsAdmin int `gorm:"is_admin" json:"is_admin"`
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"` // 该角色所属的公司
//	ShareCompanyIds string `gorm:"share_company_ids" json:"share_company_ids"` // 可共享其他分公司的数据
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Youtu_erp_roles) TableName() string {
//	return "youtu_erp_roles"
//}
//
//type Youtu_erp_role_controls struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"` // 角色名
//	UserId int `gorm:"user_id" json:"user_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Youtu_erp_role_controls) TableName() string {
//	return "youtu_erp_role_controls"
//}
//
//type Active_storage_blobs struct {
//	Id int `gorm:"id" json:"id"`
//	Key string `gorm:"key" json:"key"`
//	Filename string `gorm:"filename" json:"filename"`
//	ContentType string `gorm:"content_type" json:"content_type"`
//	Metadata string `gorm:"metadata" json:"metadata"`
//	ByteSize int `gorm:"byte_size" json:"byte_size"`
//	Checksum string `gorm:"checksum" json:"checksum"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//}
//
//func (*Active_storage_blobs) TableName() string {
//	return "active_storage_blobs"
//}
//
//type Base_data_code_of_threes struct {
//	Id int `gorm:"id" json:"id"`
//	Code string `gorm:"code" json:"code"` // 三字代码
//	Name string `gorm:"name" json:"name"`
//	EnName string `gorm:"en_name" json:"en_name"` // 机场英文名
//	Short string `gorm:"short" json:"short"`
//	Airport string `gorm:"airport" json:"airport"` // 机场名称
//	EnAirport string `gorm:"en_airport" json:"en_airport"` // 英文机场名称
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	City string `gorm:"city" json:"city"` // 城市
//	EnCity string `gorm:"en_city" json:"en_city"` // 城市英文名
//	FourCode string `gorm:"four_code" json:"four_code"`
//	NamePya string `gorm:"name_pya" json:"name_pya"` // 名称拼音全写
//	NamePyf string `gorm:"name_pyf" json:"name_pyf"` // 名称拼音简写
//	AirportPya string `gorm:"airport_pya" json:"airport_pya"` // 机场拼音全写
//	AirportPyf string `gorm:"airport_pyf" json:"airport_pyf"` // 机场拼音简写
//	CityPya string `gorm:"city_pya" json:"city_pya"` // 城市拼音全写
//	CityPyf string `gorm:"city_pyf" json:"city_pyf"` // 城市拼音简写
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//}
//
//func (*Base_data_code_of_threes) TableName() string {
//	return "base_data_code_of_threes"
//}
//
//type Ar_internal_metadata struct {
//	Key string `gorm:"key" json:"key"`
//	Value string `gorm:"value" json:"value"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//}
//
//func (*Ar_internal_metadata) TableName() string {
//	return "ar_internal_metadata"
//}
//
//type Base_data_items struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//}
//
//func (*Base_data_items) TableName() string {
//	return "base_data_items"
//}
//
//type Former_seas_bookings struct {
//	Id int `gorm:"id" json:"id"`
//	SerialNo string `gorm:"serial_no" json:"serial_no"` // 序列号
//	SoNo string `gorm:"so_no" json:"so_no"` // SO编号
//	BookingDate string `gorm:"booking_date" json:"booking_date"` // 订舱日期
//	ConfirmDate string `gorm:"confirm_date" json:"confirm_date"` // 确认日期
//	SupplierCompanyAgentId int `gorm:"supplier_company_agent_id" json:"supplier_company_agent_id"` // 供应商选择
//	SupplierAgentId int `gorm:"supplier_agent_id" json:"supplier_agent_id"` // 该供应商联系人
//	CarrierContact string `gorm:"carrier_contact" json:"carrier_contact"` // 船公司联系人
//	CarrierTel string `gorm:"carrier_tel" json:"carrier_tel"` // 船公司联系电话
//	EstimatedTimeOfDeparture string `gorm:"estimated_time_of_departure" json:"estimated_time_of_departure"` // 开船日期
//	EstimatedTimeOfArrival string `gorm:"estimated_time_of_arrival" json:"estimated_time_of_arrival"` // 到港日期
//	CutOffDate string `gorm:"cut_off_date" json:"cut_off_date"` // 截关日期
//	CutDocDate string `gorm:"cut_doc_date" json:"cut_doc_date"` // 截文件日期
//	CutCargoDate string `gorm:"cut_cargo_date" json:"cut_cargo_date"` // 截货日期
//	CargoReceivedDate string `gorm:"cargo_received_date" json:"cargo_received_date"` // 收货人提货时间
//	Vessel string `gorm:"vessel" json:"vessel"` // 船名
//	Voyage string `gorm:"voyage" json:"voyage"` // 航次
//	BaseDataSeaPortId int `gorm:"base_data_sea_port_id" json:"base_data_sea_port_id"`
//	ChargeDescription string `gorm:"charge_description" json:"charge_description"` // 费用描述
//	BookingNote string `gorm:"booking_note" json:"booking_note"` // 订舱备注
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Status string `gorm:"status" json:"status"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	InvoiceStatus string `gorm:"invoice_status" json:"invoice_status"`
//	AssignStatus string `gorm:"assign_status" json:"assign_status"`
//	BoatCompanyId int `gorm:"boat_company_id" json:"boat_company_id"`
//	BoatCompanyName string `gorm:"boat_company_name" json:"boat_company_name"` // 船名
//	SupplierPhone string `gorm:"supplier_phone" json:"supplier_phone"`
//	Remarks string `gorm:"remarks" json:"remarks"`
//	InvoiceNo string `gorm:"invoice_no" json:"invoice_no"`
//	ShipperType string `gorm:"shipper_type" json:"shipper_type"`
//	ShipperId int `gorm:"shipper_id" json:"shipper_id"`
//	ShipperContent string `gorm:"shipper_content" json:"shipper_content"` // 发货人
//	ConsigneeType string `gorm:"consignee_type" json:"consignee_type"`
//	ConsigneeId int `gorm:"consignee_id" json:"consignee_id"`
//	ConsigneeContent string `gorm:"consignee_content" json:"consignee_content"` // 收货人
//	NotifyPartyType string `gorm:"notify_party_type" json:"notify_party_type"`
//	NotifyPartyId int `gorm:"notify_party_id" json:"notify_party_id"`
//	NotifyPartyContent string `gorm:"notify_party_content" json:"notify_party_content"` // 通知人
//	CargoEn string `gorm:"cargo_en" json:"cargo_en"` // 货物英文名
//	CargoCn string `gorm:"cargo_cn" json:"cargo_cn"` // 货物中文名
//	Marks string `gorm:"marks" json:"marks"` // 标记
//	Number string `gorm:"number" json:"number"` // 包装件数
//	DescriptionOfGood string `gorm:"description_of_good" json:"description_of_good"` // 品名
//	GrossWeight string `gorm:"gross_weight" json:"gross_weight"` // 毛重
//	Size string `gorm:"size" json:"size"` // 体积
//	SeaPortPolId int `gorm:"sea_port_pol_id" json:"sea_port_pol_id"`
//	SeaPortViaId int `gorm:"sea_port_via_id" json:"sea_port_via_id"`
//	SeaPortPodId int `gorm:"sea_port_pod_id" json:"sea_port_pod_id"`
//	PlaceOfDelivery string `gorm:"place_of_delivery" json:"place_of_delivery"` // 目的地
//	PlaceOfReceipt string `gorm:"place_of_receipt" json:"place_of_receipt"` // 接货地
//	PodAgentType string `gorm:"pod_agent_type" json:"pod_agent_type"`
//	PodAgentId int `gorm:"pod_agent_id" json:"pod_agent_id"`
//	PodAgentContent string `gorm:"pod_agent_content" json:"pod_agent_content"`
//	SupplierCompanyAgentContent string `gorm:"supplier_company_agent_content" json:"supplier_company_agent_content"`
//	BaseDataPackageType int `gorm:"base_data_package_type" json:"base_data_package_type"`
//	SupplierAgentContent string `gorm:"supplier_agent_content" json:"supplier_agent_content"`
//	BoxSizeCount string `gorm:"box_size_count" json:"box_size_count"`
//	OceanChangesPaytypeId int `gorm:"ocean_changes_paytype_id" json:"ocean_changes_paytype_id"` // 运费支付方式
//	OtherChangesPaytypeId int `gorm:"other_changes_paytype_id" json:"other_changes_paytype_id"` // 其他运费支付方式
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	Dimension string `gorm:"dimension" json:"dimension"` // 尺寸
//}
//
//func (*Former_seas_bookings) TableName() string {
//	return "former_seas_bookings"
//}
//
//type Former_trails_orders struct {
//	Id int `gorm:"id" json:"id"`
//	CustomerNo string `gorm:"customer_no" json:"customer_no"` // 工厂客户号
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	Remarks string `gorm:"remarks" json:"remarks"` // 备注
//	InvoiceNo string `gorm:"invoice_no" json:"invoice_no"` // 发票号
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	Status string `gorm:"status" json:"status"`
//	SerialNo string `gorm:"serial_no" json:"serial_no"`
//	InvoiceStatus string `gorm:"invoice_status" json:"invoice_status"`
//	AssignStatus string `gorm:"assign_status" json:"assign_status"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	AssociatedFormers string `gorm:"associated_formers" json:"associated_formers"`
//	SoNo string `gorm:"so_no" json:"so_no"`
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"`
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"`
//	CompanyInstructionContent string `gorm:"company_instruction_content" json:"company_instruction_content"`
//	DriverMobi string `gorm:"driver_mobi" json:"driver_mobi"` // 司机联系电话
//	LoadingDate string `gorm:"loading_date" json:"loading_date"`
//	LicensePlateNumber string `gorm:"license_plate_number" json:"license_plate_number"`
//	Departure string `gorm:"departure" json:"departure"`
//	Destination string `gorm:"destination" json:"destination"`
//	BaseDataInstructionTypeId int `gorm:"base_data_instruction_type_id" json:"base_data_instruction_type_id"` // 委托类型
//	WaysOfDeclarationId int `gorm:"ways_of_declaration_id" json:"ways_of_declaration_id"` // 报关方式
//	BaseDataTradeTermsId int `gorm:"base_data_trade_terms_id" json:"base_data_trade_terms_id"` // 贸易条款
//	TransshipmentId int `gorm:"transshipment_id" json:"transshipment_id"` // 转运
//	BaseDataItemId int `gorm:"base_data_item_id" json:"base_data_item_id"` // 装运条款
//	ServiceContractNo string `gorm:"service_contract_no" json:"service_contract_no"` // 合同编号
//	OfWay int `gorm:"of_way" json:"of_way"` // 运输方式
//	Number string `gorm:"number" json:"number"` // 拖车货物总件数
//	GrossWeight string `gorm:"gross_weight" json:"gross_weight"` // 拖车总重量
//	Size string `gorm:"size" json:"size"` // 货物体积
//	DescriptionOfGood string `gorm:"description_of_good" json:"description_of_good"` // 货物描述
//	Marks string `gorm:"marks" json:"marks"`
//	IsWeighing int `gorm:"is_weighing" json:"is_weighing"` // 过磅
//	IsLockers int `gorm:"is_lockers" json:"is_lockers"` // 小柜摆尾
//	IsDeclare int `gorm:"is_declare" json:"is_declare"` // 报关单随车
//	IsDrivingLicense int `gorm:"is_driving_license" json:"is_driving_license"`
//	ContainerNo string `gorm:"container_no" json:"container_no"`
//	SealNo string `gorm:"seal_no" json:"seal_no"`
//	BoxSizeCount string `gorm:"box_size_count" json:"box_size_count"`
//	Dimension string `gorm:"dimension" json:"dimension"` // 尺寸
//	IsGoodsAttachment int `gorm:"is_goods_attachment" json:"is_goods_attachment"` // 导出货描列表
//	BaseDataPackageType int `gorm:"base_data_package_type" json:"base_data_package_type"`
//}
//
//func (*Former_trails_orders) TableName() string {
//	return "former_trails_orders"
//}
//
//type Score_system_rules struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	RuleType string `gorm:"rule_type" json:"rule_type"` // 规则类型[获得积分,使用积分]
//	Group string `gorm:"group" json:"group"` // 分组[用户注册,在线登录,交易奖励,积分抵扣]
//	Key string `gorm:"key" json:"key"` // 类型
//	Value string `gorm:"value" json:"value"` // 积分值
//	FilterType string `gorm:"filter_type" json:"filter_type"` // 类型
//	Label string `gorm:"label" json:"label"` // 标识
//	CompanyId int `gorm:"company_id" json:"company_id"`
//}
//
//func (*Score_system_rules) TableName() string {
//	return "score_system_rules"
//}
//
//type Bargain_surcharge_fees struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"` // 费用简称
//	NameCn string `gorm:"name_cn" json:"name_cn"` // 费用英文名
//	NameEn string `gorm:"name_en" json:"name_en"` // 费用中文名
//	UnitPrice float64 `gorm:"unit_price" json:"unit_price"`
//	Account int `gorm:"account" json:"account"` // 数量
//	FinanceCurrency string `gorm:"finance_currency" json:"finance_currency"` // 币种
//	Remarks string `gorm:"remarks" json:"remarks"`
//	Parcel float64 `gorm:"parcel" json:"parcel"`
//	SupplyId int `gorm:"supply_id" json:"supply_id"` // 供应商
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	SourceType string `gorm:"source_type" json:"source_type"`
//	SourceId int `gorm:"source_id" json:"source_id"`
//	BargainMainId int `gorm:"bargain_main_id" json:"bargain_main_id"`
//	20FR float64 `gorm:"20FR" json:"20FR"`
//	20GP float64 `gorm:"20GP" json:"20GP"`
//	20OT float64 `gorm:"20OT" json:"20OT"`
//	20RF float64 `gorm:"20RF" json:"20RF"`
//	40GP float64 `gorm:"40GP" json:"40GP"`
//	40FR float64 `gorm:"40FR" json:"40FR"`
//	40OT float64 `gorm:"40OT" json:"40OT"`
//	40HQ float64 `gorm:"40HQ" json:"40HQ"`
//	40RF float64 `gorm:"40RF" json:"40RF"`
//	45GP float64 `gorm:"45GP" json:"45GP"`
//	45HQ float64 `gorm:"45HQ" json:"45HQ"`
//	20RH float64 `gorm:"20RH" json:"20RH"`
//	20HQ float64 `gorm:"20HQ" json:"20HQ"`
//	20HT float64 `gorm:"20HT" json:"20HT"`
//	40HT float64 `gorm:"40HT" json:"40HT"`
//	40RH float64 `gorm:"40RH" json:"40RH"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Bargain_surcharge_fees) TableName() string {
//	return "bargain_surcharge_fees"
//}
//
//type Base_data_sea_lines struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"` // 航线名
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	EnName string `gorm:"en_name" json:"en_name"` // 航线英文名
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//}
//
//func (*Base_data_sea_lines) TableName() string {
//	return "base_data_sea_lines"
//}
//
//type Scale_groups struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"` // 群组名
//	Enabled int `gorm:"enabled" json:"enabled"` // 有效状态
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//}
//
//func (*Scale_groups) TableName() string {
//	return "scale_groups"
//}
//
//type Accounts struct {
//	Id int `gorm:"id" json:"id"`
//	FinanceBanksId int `gorm:"finance_banks_id" json:"finance_banks_id"`
//	Name string `gorm:"name" json:"name"` // 账户名称
//	UserName string `gorm:"user_name" json:"user_name"` // 开户人姓名
//	BankName string `gorm:"bank_name" json:"bank_name"` // 开户行
//	BankNumber string `gorm:"bank_number" json:"bank_number"` // 银行账号
//	Category string `gorm:"category" json:"category"` // 账户类型
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	BankAddress string `gorm:"bank_address" json:"bank_address"` // 开户行地址
//	SwiftCode string `gorm:"swift_code" json:"swift_code"` // swift code
//	TaxRegisterNumber string `gorm:"tax_register_number" json:"tax_register_number"`
//	BeneficiaryAddress string `gorm:"beneficiary_address" json:"beneficiary_address"` // 开户人地址
//	BeneficiaryLocation string `gorm:"beneficiary_location" json:"beneficiary_location"` // 开户人位置
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Accounts) TableName() string {
//	return "accounts"
//}
//
//type Doc_template_rules struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"` // 模板名称
//	TransportType int `gorm:"transport_type" json:"transport_type"` // 运输方式
//	MainTransport int `gorm:"main_transport" json:"main_transport"` // 其他对应的主运输方式
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	Content string `gorm:"content" json:"content"` // 模板对应的列设置内容
//	RuleType string `gorm:"rule_type" json:"rule_type"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//}
//
//func (*Doc_template_rules) TableName() string {
//	return "doc_template_rules"
//}
//
//type Former_other_services struct {
//	Remarks string `gorm:"remarks" json:"remarks"`
//	Id int `gorm:"id" json:"id"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"` // 订单
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"` // 委托单位
//	IsFumigation int `gorm:"is_fumigation" json:"is_fumigation"`
//	FumigationId int `gorm:"fumigation_id" json:"fumigation_id"`
//	IsTraders int `gorm:"is_traders" json:"is_traders"`
//	TradersId int `gorm:"traders_id" json:"traders_id"`
//	IsInsurance int `gorm:"is_insurance" json:"is_insurance"`
//	InsuranceId int `gorm:"insurance_id" json:"insurance_id"`
//	Beneficiary string `gorm:"beneficiary" json:"beneficiary"`
//	IsMagneticTest int `gorm:"is_magnetic_test" json:"is_magnetic_test"`
//	MagneticTestId int `gorm:"magnetic_test_id" json:"magnetic_test_id"`
//	IsIdentification int `gorm:"is_identification" json:"is_identification"`
//	IdentificationId int `gorm:"identification_id" json:"identification_id"`
//	SerialNo string `gorm:"serial_no" json:"serial_no"`
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"`
//	CompanyInstructionContent string `gorm:"company_instruction_content" json:"company_instruction_content"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	CommodityInspectionId int `gorm:"commodity_inspection_id" json:"commodity_inspection_id"`
//	IsCommodityInspection int `gorm:"is_commodity_inspection" json:"is_commodity_inspection"`
//}
//
//func (*Former_other_services) TableName() string {
//	return "former_other_services"
//}
//
//type Former_seas_landing_bills struct {
//	Id int `gorm:"id" json:"id"`
//	BLNo string `gorm:"b_l_no" json:"b_l_no"` // 提单号
//	VerifyDate string `gorm:"verify_date" json:"verify_date"` // 校对日期
//	UserVerifyId int `gorm:"user_verify_id" json:"user_verify_id"`
//	ShipperType string `gorm:"shipper_type" json:"shipper_type"`
//	ShipperId int `gorm:"shipper_id" json:"shipper_id"`
//	ConsigneeType string `gorm:"consignee_type" json:"consignee_type"`
//	ConsigneeId int `gorm:"consignee_id" json:"consignee_id"`
//	NotifyPartyType string `gorm:"notify_party_type" json:"notify_party_type"`
//	NotifyPartyId int `gorm:"notify_party_id" json:"notify_party_id"`
//	PodAgentType string `gorm:"pod_agent_type" json:"pod_agent_type"`
//	PodAgentId int `gorm:"pod_agent_id" json:"pod_agent_id"`
//	AmsActualShipperId int `gorm:"ams_actual_shipper_id" json:"ams_actual_shipper_id"`
//	AmsActualConsigneeId int `gorm:"ams_actual_consignee_id" json:"ams_actual_consignee_id"`
//	BoatCompanyId int `gorm:"boat_company_id" json:"boat_company_id"`
//	Vessel string `gorm:"vessel" json:"vessel"` // 船名
//	Voyage string `gorm:"voyage" json:"voyage"` // 航次
//	SeaPortPodId int `gorm:"sea_port_pod_id" json:"sea_port_pod_id"`
//	SeaPortPolId int `gorm:"sea_port_pol_id" json:"sea_port_pol_id"`
//	SeaPortViaId int `gorm:"sea_port_via_id" json:"sea_port_via_id"`
//	PlaceOfDelivery string `gorm:"place_of_delivery" json:"place_of_delivery"` // 目的地
//	PlaceOfReceipt string `gorm:"place_of_receipt" json:"place_of_receipt"` // 收货地址
//	PreCarriageById int `gorm:"pre_carriage_by_id" json:"pre_carriage_by_id"`
//	EstimatedTimeOfDeparture string `gorm:"estimated_time_of_departure" json:"estimated_time_of_departure"`
//	EstimatedTimeOfArrival string `gorm:"estimated_time_of_arrival" json:"estimated_time_of_arrival"`
//	CutOffDate string `gorm:"cut_off_date" json:"cut_off_date"` // 截关日期
//	CargoReceivedDate string `gorm:"cargo_received_date" json:"cargo_received_date"` // 收货人提货时间
//	FreightPayableAt string `gorm:"freight_payable_at" json:"freight_payable_at"` // 运费支付地
//	PlaceOfIssue string `gorm:"place_of_issue" json:"place_of_issue"` // 签单地点
//	DateOfIssue string `gorm:"date_of_issue" json:"date_of_issue"` // 签单日期
//	BaseDataTradeTermsId int `gorm:"base_data_trade_terms_id" json:"base_data_trade_terms_id"`
//	BaseDataMiscBillId int `gorm:"base_data_misc_bill_id" json:"base_data_misc_bill_id"`
//	TlxNo string `gorm:"tlx_no" json:"tlx_no"` // 电放单号
//	MblReleaseDate string `gorm:"mbl_release_date" json:"mbl_release_date"`
//	MblNo string `gorm:"mbl_no" json:"mbl_no"` // 船东提单号
//	Number string `gorm:"number" json:"number"`
//	BaseDataPackageType int `gorm:"base_data_package_type" json:"base_data_package_type"` // 包装类型
//	GrossWeight string `gorm:"gross_weight" json:"gross_weight"` // 毛重
//	Size string `gorm:"size" json:"size"`
//	DescriptionOfGood string `gorm:"description_of_good" json:"description_of_good"` // 货物英文名称
//	Marks string `gorm:"marks" json:"marks"` // 标记唛头
//	Remarks string `gorm:"remarks" json:"remarks"` // 提单备注
//	ShippingInstruction int `gorm:"shipping_instruction" json:"shipping_instruction"` // 补料,简称SI
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//	SerialNo string `gorm:"serial_no" json:"serial_no"` // 序列号
//	Status string `gorm:"status" json:"status"` // 主状态
//	ShipperContent string `gorm:"shipper_content" json:"shipper_content"` // 发货人详细
//	ConsigneeContent string `gorm:"consignee_content" json:"consignee_content"` // 收货人详情
//	NotifyPartyContent string `gorm:"notify_party_content" json:"notify_party_content"` // 通知人详情
//	PodAgentContent string `gorm:"pod_agent_content" json:"pod_agent_content"`
//	LandingBillType string `gorm:"landing_bill_type" json:"landing_bill_type"` // 分类
//	BillOfLandingId int `gorm:"bill_of_landing_id" json:"bill_of_landing_id"`
//	OceanChangesPaytypeId int `gorm:"ocean_changes_paytype_id" json:"ocean_changes_paytype_id"`
//	OtherChangesPaytypeId int `gorm:"other_changes_paytype_id" json:"other_changes_paytype_id"`
//	ShippedOnBoardDate string `gorm:"shipped_on_board_date" json:"shipped_on_board_date"`
//	PreCarriageBy string `gorm:"pre_carriage_by" json:"pre_carriage_by"`
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"`
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"`
//	CompanyInstructionContent string `gorm:"company_instruction_content" json:"company_instruction_content"`
//	PayPodId int `gorm:"pay_pod_id" json:"pay_pod_id"`
//	Dimension string `gorm:"dimension" json:"dimension"` // 尺寸
//	IsGoodsAttachment int `gorm:"is_goods_attachment" json:"is_goods_attachment"` // 导出货物描述附件
//	IsCabinetAttachment int `gorm:"is_cabinet_attachment" json:"is_cabinet_attachment"` // 导出分柜附件
//	BaseDataItemId int `gorm:"base_data_item_id" json:"base_data_item_id"`
//	BookingRemark string `gorm:"booking_remark" json:"booking_remark"` // 订舱备注
//	ChargeDescription string `gorm:"charge_description" json:"charge_description"` // 费用描述
//	BoxSizeCount string `gorm:"box_size_count" json:"box_size_count"` // 柜兴柜量
//	SupplierCompanyAgentId int `gorm:"supplier_company_agent_id" json:"supplier_company_agent_id"`
//}
//
//func (*Former_seas_landing_bills) TableName() string {
//	return "former_seas_landing_bills"
//}
//
//type Order_works struct {
//	Id int `gorm:"id" json:"id"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"` // 订单id
//	Name string `gorm:"name" json:"name"` // 任务名称
//	Note string `gorm:"note" json:"note"` // 备注
//	Enabled int `gorm:"enabled" json:"enabled"` // 有效状态
//	Warning int `gorm:"warning" json:"warning"` // 提醒标志
//	Status string `gorm:"status" json:"status"` // 状态
//	Rank int `gorm:"rank" json:"rank"` // 排序权位
//	CompletedTime string `gorm:"completed_time" json:"completed_time"` // 完成时间
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	NameCn string `gorm:"name_cn" json:"name_cn"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Order_works) TableName() string {
//	return "order_works"
//}
//
//type Base_data_boat_companies_copy1 struct {
//	Id int `gorm:"id" json:"id"`
//	ChinaName string `gorm:"china_name" json:"china_name"` // 中文名
//	Name string `gorm:"name" json:"name"` // 船运公司名称
//	Url string `gorm:"url" json:"url"` // 官网地址
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	ChinaNamePya string `gorm:"china_name_pya" json:"china_name_pya"` // 船运公司拼音全写
//	ChinaNamePyf string `gorm:"china_name_pyf" json:"china_name_pyf"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//}
//
//func (*Base_data_boat_companies_copy1) TableName() string {
//	return "base_data_boat_companies_copy1"
//}
//
//type Change_logs struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	SourceType string `gorm:"source_type" json:"source_type"`
//	SourceId int `gorm:"source_id" json:"source_id"` // 数据来源
//	Content string `gorm:"content" json:"content"` // 操作内容
//	UserId int `gorm:"user_id" json:"user_id"`
//	IpAddress string `gorm:"ip_address" json:"ip_address"` // ip address
//	Remark string `gorm:"remark" json:"remark"` // 用户操作备注
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	LogType string `gorm:"log_type" json:"log_type"`
//}
//
//func (*Change_logs) TableName() string {
//	return "change_logs"
//}
//
//type Finance_user_and_drawings struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	UserId int `gorm:"user_id" json:"user_id"`
//	FinanceProfitDrawingId int `gorm:"finance_profit_drawing_id" json:"finance_profit_drawing_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Finance_user_and_drawings) TableName() string {
//	return "finance_user_and_drawings"
//}
//
//type Former_airs_landing_bills struct {
//	Id int `gorm:"id" json:"id"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	ShipperContent string `gorm:"shipper_content" json:"shipper_content"` // 发货人信息
//	ConsigneeContent string `gorm:"consignee_content" json:"consignee_content"` // 收货人信息
//	NotifyPartyContent string `gorm:"notify_party_content" json:"notify_party_content"`
//	PodAgentContent string `gorm:"pod_agent_content" json:"pod_agent_content"`
//	DepartureAirportId int `gorm:"departure_airport_id" json:"departure_airport_id"`
//	DestinationAirportId int `gorm:"destination_airport_id" json:"destination_airport_id"`
//	Flight string `gorm:"flight" json:"flight"` // 航班
//	FlightDate string `gorm:"flight_date" json:"flight_date"`
//	CurrencyId int `gorm:"currency_id" json:"currency_id"` // 货币
//	ChargableWeight string `gorm:"chargable_weight" json:"chargable_weight"` // 计费重量
//	ExecuteDate string `gorm:"execute_date" json:"execute_date"` // 签发日期
//	SerialNo string `gorm:"serial_no" json:"serial_no"` // 序列号
//	SubBlNo string `gorm:"sub_bl_no" json:"sub_bl_no"` // 分提单号
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	ShipperType string `gorm:"shipper_type" json:"shipper_type"`
//	ShipperId int `gorm:"shipper_id" json:"shipper_id"` // 发货人
//	ConsigneeType string `gorm:"consignee_type" json:"consignee_type"`
//	ConsigneeId int `gorm:"consignee_id" json:"consignee_id"` // 收货人
//	NotifyPartyType string `gorm:"notify_party_type" json:"notify_party_type"`
//	NotifyPartyId int `gorm:"notify_party_id" json:"notify_party_id"` // 通知方
//	PodAgentType string `gorm:"pod_agent_type" json:"pod_agent_type"`
//	PodAgentId int `gorm:"pod_agent_id" json:"pod_agent_id"` // 代理商
//	MainBlNo string `gorm:"main_bl_no" json:"main_bl_no"` // 主提单
//	TransshipmentAirportId int `gorm:"transshipment_airport_id" json:"transshipment_airport_id"`
//	CodeOfTwoId int `gorm:"code_of_two_id" json:"code_of_two_id"`
//	OceanChangesPaytypeId int `gorm:"ocean_changes_paytype_id" json:"ocean_changes_paytype_id"`
//	OtherChangesPaytypeId int `gorm:"other_changes_paytype_id" json:"other_changes_paytype_id"`
//	Marks string `gorm:"marks" json:"marks"`
//	DescriptionOfGood string `gorm:"description_of_good" json:"description_of_good"`
//	GrossWeight string `gorm:"gross_weight" json:"gross_weight"`
//	Number string `gorm:"number" json:"number"`
//	Size string `gorm:"size" json:"size"`
//	Status string `gorm:"status" json:"status"`
//	Remarks string `gorm:"remarks" json:"remarks"`
//	LandingBillType string `gorm:"landing_bill_type" json:"landing_bill_type"`
//	BaseDataPackageType int `gorm:"base_data_package_type" json:"base_data_package_type"`
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"`
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"`
//	CompanyInstructionContent string `gorm:"company_instruction_content" json:"company_instruction_content"`
//	DeliveryAddress string `gorm:"delivery_address" json:"delivery_address"`
//	ArriveAddress string `gorm:"arrive_address" json:"arrive_address"`
//	EstimatedTimeOfDeparture string `gorm:"estimated_time_of_departure" json:"estimated_time_of_departure"`
//	Ratio float64 `gorm:"ratio" json:"ratio"` // 材积换算系数
//	RatioWeight float64 `gorm:"ratio_weight" json:"ratio_weight"` // 材积重
//	Bubble float64 `gorm:"bubble" json:"bubble"` // 分泡%
//	Dimension string `gorm:"dimension" json:"dimension"` // 体积
//	IsGoodsAttachment int `gorm:"is_goods_attachment" json:"is_goods_attachment"` // 导出货物描述附件
//	BaseDataTradeTermId int `gorm:"base_data_trade_term_id" json:"base_data_trade_term_id"`
//	BaseDataItemId int `gorm:"base_data_item_id" json:"base_data_item_id"`
//	BookingRemarks string `gorm:"booking_remarks" json:"booking_remarks"`
//	CostDescription string `gorm:"cost_description" json:"cost_description"`
//	SupplierCompanyAgentId int `gorm:"supplier_company_agent_id" json:"supplier_company_agent_id"`
//	RandomFile int `gorm:"random_file" json:"random_file"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Former_airs_landing_bills) TableName() string {
//	return "former_airs_landing_bills"
//}
//
//type Message_signatures struct {
//	Id int `gorm:"id" json:"id"`
//	Content string `gorm:"content" json:"content"`
//	UserId int `gorm:"user_id" json:"user_id"`
//	Name string `gorm:"name" json:"name"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	IsDefault int `gorm:"is_default" json:"is_default"`
//}
//
//func (*Message_signatures) TableName() string {
//	return "message_signatures"
//}
//
//type Plan_customs struct {
//	Id int `gorm:"id" json:"id"`
//	PlanMainId int `gorm:"plan_main_id" json:"plan_main_id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Enabled int `gorm:"enabled" json:"enabled"` // 有效状态
//	Note string `gorm:"note" json:"note"` // 备注
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"` // 委托单位
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"` // 委托单位类型
//	BaseDataCiqtypeId int `gorm:"base_data_ciqtype_id" json:"base_data_ciqtype_id"` // 报关类型
//	CustomsAmount float64 `gorm:"customs_amount" json:"customs_amount"` // 报关金额
//	IsPay int `gorm:"is_pay" json:"is_pay"` // 是否买单
//	CustomBrokerId int `gorm:"custom_broker_id" json:"custom_broker_id"` // 报关行
//	Contact string `gorm:"contact" json:"contact"` // 联系人
//	ContactPhone string `gorm:"contact_phone" json:"contact_phone"` // 联系人手机号
//	ContractNo string `gorm:"contract_no" json:"contract_no"` // 合同编号
//	IsFumigation int `gorm:"is_fumigation" json:"is_fumigation"` // 是否需要熏蒸
//	FumigationId int `gorm:"fumigation_id" json:"fumigation_id"` // 熏蒸公司
//	IsCertificate int `gorm:"is_certificate" json:"is_certificate"` // 是否需要产地证
//	TraderId int `gorm:"trader_id" json:"trader_id"` // 贸易商
//	IsInsurance int `gorm:"is_insurance" json:"is_insurance"` // 是否需要保险
//	InsuranceId int `gorm:"insurance_id" json:"insurance_id"` // 保险公司
//	Beneficiary string `gorm:"beneficiary" json:"beneficiary"` // 受益人
//	InsuranceAmount float64 `gorm:"insurance_amount" json:"insurance_amount"` // 保险金额
//	IsMagneticTest int `gorm:"is_magnetic_test" json:"is_magnetic_test"` // 是否需要磁检
//	MagneticTestId int `gorm:"magnetic_test_id" json:"magnetic_test_id"` // 磁检机构
//	IsIdentification int `gorm:"is_identification" json:"is_identification"` // 是否需要鉴定
//	IdentificationId int `gorm:"identification_id" json:"identification_id"` // 鉴定机构
//	IsWarehouse int `gorm:"is_warehouse" json:"is_warehouse"` // 是否需要仓库/场装
//	WarehouseId int `gorm:"warehouse_id" json:"warehouse_id"` // 仓库/场装
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Plan_customs) TableName() string {
//	return "plan_customs"
//}
//
//type Plan_integrates struct {
//	Id int `gorm:"id" json:"id"`
//	PlanMainId int `gorm:"plan_main_id" json:"plan_main_id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Enabled int `gorm:"enabled" json:"enabled"` // 有效状态
//	Category string `gorm:"category" json:"category"` // 类型[拖车，中港，内陆]
//	TransportType string `gorm:"transport_type" json:"transport_type"` // 运输方式
//	TrailerCompanyId int `gorm:"trailer_company_id" json:"trailer_company_id"` // 拖车行
//	TrailerContact string `gorm:"trailer_contact" json:"trailer_contact"` // 拖车联系人
//	TrailerPhone string `gorm:"trailer_phone" json:"trailer_phone"` // 拖车联系人电话
//	LoadingDate string `gorm:"loading_date" json:"loading_date"` // 装货时间
//	Departure string `gorm:"departure" json:"departure"` // 装卸货地
//	CarNumber string `gorm:"car_number" json:"car_number"` // 车牌
//	IsFumigation int `gorm:"is_fumigation" json:"is_fumigation"` // 是否需要熏蒸
//	FumigationId int `gorm:"fumigation_id" json:"fumigation_id"` // 熏蒸公司
//	IsCertificate int `gorm:"is_certificate" json:"is_certificate"` // 是否需要产地证
//	TraderId int `gorm:"trader_id" json:"trader_id"` // 贸易商
//	IsInsurance int `gorm:"is_insurance" json:"is_insurance"` // 是否需要保险
//	InsuranceId int `gorm:"insurance_id" json:"insurance_id"` // 保险公司
//	Beneficiary string `gorm:"beneficiary" json:"beneficiary"` // 受益人
//	InsuranceAmount float64 `gorm:"insurance_amount" json:"insurance_amount"` // 保险金额
//	IsMagneticTest int `gorm:"is_magnetic_test" json:"is_magnetic_test"` // 是否需要磁检
//	MagneticTestId int `gorm:"magnetic_test_id" json:"magnetic_test_id"` // 磁检机构
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"` // 委托单位
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"` // 委托单位类型
//	IsIdentification int `gorm:"is_identification" json:"is_identification"` // 是否需要鉴定
//	IdentificationId int `gorm:"identification_id" json:"identification_id"` // 鉴定机构
//	IsDrivingLicense int `gorm:"is_driving_license" json:"is_driving_license"` // 转关带司机本
//	IsDeclarationLicence int `gorm:"is_declaration_licence" json:"is_declaration_licence"` // 报关证单随车
//	IsWeighing int `gorm:"is_weighing" json:"is_weighing"` // 是否需要过磅
//	IsLockers int `gorm:"is_lockers" json:"is_lockers"` // 是否需要小柜摆尾
//	Note string `gorm:"note" json:"note"` // 备注
//	IsWarehouse int `gorm:"is_warehouse" json:"is_warehouse"` // 是否需要仓库/场装
//	WarehouseId int `gorm:"warehouse_id" json:"warehouse_id"` // 仓库/场装
//	Pol string `gorm:"pol" json:"pol"` // 起运港
//	PlaceOfDeparture string `gorm:"place_of_departure" json:"place_of_departure"` // 出发地
//	Destination string `gorm:"destination" json:"destination"` // 目的地
//	CommodityInspectionId int `gorm:"commodity_inspection_id" json:"commodity_inspection_id"`
//	IsCommodityInspection int `gorm:"is_commodity_inspection" json:"is_commodity_inspection"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Plan_integrates) TableName() string {
//	return "plan_integrates"
//}
//
//type Youtu_erp_user_subscribes struct {
//	Id int `gorm:"id" json:"id"`
//	UserId int `gorm:"user_id" json:"user_id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	OpenId string `gorm:"open_id" json:"open_id"`
//	Nickname string `gorm:"nickname" json:"nickname"`
//	Sex int `gorm:"sex" json:"sex"`
//	Province string `gorm:"province" json:"province"`
//	Country string `gorm:"country" json:"country"`
//	Headimgurl string `gorm:"headimgurl" json:"headimgurl"`
//	Unionid string `gorm:"unionid" json:"unionid"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Youtu_erp_user_subscribes) TableName() string {
//	return "youtu_erp_user_subscribes"
//}
//
//type Bargain_sea_fees struct {
//	Id int `gorm:"id" json:"id"`
//	UnitPrice float64 `gorm:"unit_price" json:"unit_price"`
//	Count int `gorm:"count" json:"count"` // 数量
//	FinanceCurrency string `gorm:"finance_currency" json:"finance_currency"` // 币种
//	SupplyId int `gorm:"supply_id" json:"supply_id"` // 供应商
//	BargainSupplyQuotationId int `gorm:"bargain_supply_quotation_id" json:"bargain_supply_quotation_id"` // 供应商报价
//	BargainCustomerQuotationId int `gorm:"bargain_customer_quotation_id" json:"bargain_customer_quotation_id"` // 客户报价单
//	BoxSizeId int `gorm:"box_size_id" json:"box_size_id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	BargainMainId int `gorm:"bargain_main_id" json:"bargain_main_id"`
//	BargainBoxSizeCountId int `gorm:"bargain_box_size_count_id" json:"bargain_box_size_count_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Bargain_sea_fees) TableName() string {
//	return "bargain_sea_fees"
//}
//
//type Companies struct {
//	Id int `gorm:"id" json:"id"`
//	Telephone string `gorm:"telephone" json:"telephone"` // 座机
//	Telephone2 string `gorm:"telephone2" json:"telephone2"` // 备用座机
//	Fax string `gorm:"fax" json:"fax"` // 传真
//	Fax2 string `gorm:"fax2" json:"fax2"` // 备用传真
//	Address string `gorm:"address" json:"address"` // 地址
//	Address2 string `gorm:"address2" json:"address2"` // 备用地址
//	Website string `gorm:"website" json:"website"` // 网站
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	SourceType string `gorm:"source_type" json:"source_type"`
//	SourceId int `gorm:"source_id" json:"source_id"` // 数据来源
//	NameNick string `gorm:"name_nick" json:"name_nick"`
//	NameCn string `gorm:"name_cn" json:"name_cn"`
//	NameEn string `gorm:"name_en" json:"name_en"`
//	Code string `gorm:"code" json:"code"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	Email string `gorm:"email" json:"email"`
//	Province string `gorm:"province" json:"province"` // 省份
//	City string `gorm:"city" json:"city"` // 市
//	District string `gorm:"district" json:"district"` // 区
//}
//
//func (*Companies) TableName() string {
//	return "companies"
//}
//
//type Invoices struct {
//	Id int `gorm:"id" json:"id"`
//	UserCompaniesId int `gorm:"user_companies_id" json:"user_companies_id"`
//	FinanceBanksId int `gorm:"finance_banks_id" json:"finance_banks_id"`
//	Name string `gorm:"name" json:"name"` // 发票名称
//	TaxpayerNumber string `gorm:"taxpayer_number" json:"taxpayer_number"` // 纳税人识别号
//	Address string `gorm:"address" json:"address"` // 地址
//	PhoneNumber string `gorm:"phone_number" json:"phone_number"` // 电话
//	BankName string `gorm:"bank_name" json:"bank_name"`
//	BankNumber string `gorm:"bank_number" json:"bank_number"` // 银行账号
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	FinanceCurrencyId int `gorm:"finance_currency_id" json:"finance_currency_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Invoices) TableName() string {
//	return "invoices"
//}
//
//type Invoice_info_finance_fees struct {
//	Id int `gorm:"id" json:"id"`
//	InvoiceInfoId int `gorm:"invoice_info_id" json:"invoice_info_id"`
//	FinanceFeeId int `gorm:"finance_fee_id" json:"finance_fee_id"`
//	Enabled int `gorm:"enabled" json:"enabled"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Amount float64 `gorm:"amount" json:"amount"` // 开票金额
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Invoice_info_finance_fees) TableName() string {
//	return "invoice_info_finance_fees"
//}
//
//type Order_user_lists struct {
//	Id int `gorm:"id" json:"id"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	UserId int `gorm:"user_id" json:"user_id"`
//	Role string `gorm:"role" json:"role"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Order_user_lists) TableName() string {
//	return "order_user_lists"
//}
//
//
//
//type Website_menus struct {
//	Id int `gorm:"id" json:"id"`
//	NameCn string `gorm:"name_cn" json:"name_cn"` // 菜单中文名
//	NameEn string `gorm:"name_en" json:"name_en"` // 菜单英文名
//	Enabled int `gorm:"enabled" json:"enabled"` // 菜单是否显示
//	Position int `gorm:"position" json:"position"`
//	HiddenName string `gorm:"hidden_name" json:"hidden_name"` // 隐藏字段
//}
//
//func (*Website_menus) TableName() string {
//	return "website_menus"
//}
//
//type Doc_templates struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"` // 模板名
//	WidgetParams string `gorm:"widget_params" json:"widget_params"`
//	Remark string `gorm:"remark" json:"remark"` // 备注
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	TemplateType string `gorm:"template_type" json:"template_type"` // 1为海运,2为陆运
//	TmpName string `gorm:"tmp_name" json:"tmp_name"` // 上传后的文件名
//	OriginalFilename string `gorm:"original_filename" json:"original_filename"` // 文件原始名
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsBootTemplate int `gorm:"is_boot_template" json:"is_boot_template"` // 是否Bootstrap模板
//	TransportType int `gorm:"transport_type" json:"transport_type"` // 运输类型
//	DocTemplateRuleId int `gorm:"doc_template_rule_id" json:"doc_template_rule_id"`
//	YoutuDefault int `gorm:"youtu_default" json:"youtu_default"` // 优途提供的模板默认为false
//	IsValid int `gorm:"is_valid" json:"is_valid"` // 是否有效,默认为true有效
//}
//
//func (*Doc_templates) TableName() string {
//	return "doc_templates"
//}
//
//type Former_airs_warehouses struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	SerialNo string `gorm:"serial_no" json:"serial_no"` // 序列号
//	WarehouseNo string `gorm:"warehouse_no" json:"warehouse_no"` // 入仓号
//	Flight string `gorm:"flight" json:"flight"` // 航班
//	FlightDate string `gorm:"flight_date" json:"flight_date"` // 航班日期
//	DepartureAirportId int `gorm:"departure_airport_id" json:"departure_airport_id"` // 起运港
//	DestinationAirportId int `gorm:"destination_airport_id" json:"destination_airport_id"` // 目的港
//	WarehouseAddress string `gorm:"warehouse_address" json:"warehouse_address"` // 入仓地址
//	SupplierAgent string `gorm:"supplier_agent" json:"supplier_agent"` // 联系人
//	SupplierAgentMobi string `gorm:"supplier_agent_mobi" json:"supplier_agent_mobi"` // 联系人电话
//	Remarks string `gorm:"remarks" json:"remarks"` // 备注
//	Status string `gorm:"status" json:"status"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	OceanChangesPaytypeId int `gorm:"ocean_changes_paytype_id" json:"ocean_changes_paytype_id"` // 费用支付方式
//	OtherChangesPaytypeId int `gorm:"other_changes_paytype_id" json:"other_changes_paytype_id"` // 其他费用支付方式
//	EstimatedTimeOfDeparture string `gorm:"estimated_time_of_departure" json:"estimated_time_of_departure"`
//	CodeOfTwoId int `gorm:"code_of_two_id" json:"code_of_two_id"`
//	TransshipmentAirportId int `gorm:"transshipment_airport_id" json:"transshipment_airport_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Former_airs_warehouses) TableName() string {
//	return "former_airs_warehouses"
//}
//
//type Base_data_codes struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	CodeName string `gorm:"code_name" json:"code_name"`
//	Name string `gorm:"name" json:"name"`
//	Remark string `gorm:"remark" json:"remark"` // 备注
//	SourceType string `gorm:"source_type" json:"source_type"` // 来源
//	CodeLevelId int `gorm:"code_level_id" json:"code_level_id"`
//	Status int `gorm:"status" json:"status"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"` // 行级乐观锁
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//}
//
//func (*Base_data_codes) TableName() string {
//	return "base_data_codes"
//}
//
//type Base_data_vessels struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"`
//	BoatCompanyId string `gorm:"boat_company_id" json:"boat_company_id"`
//	Enabled int `gorm:"enabled" json:"enabled"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Base_data_vessels) TableName() string {
//	return "base_data_vessels"
//}
//
//type Former_airs_bookings struct {
//	Id int `gorm:"id" json:"id"`
//	ShipperContent string `gorm:"shipper_content" json:"shipper_content"` // 发货人信息
//	ConsigneeContent string `gorm:"consignee_content" json:"consignee_content"` // 收货人信息
//	NotifyPartyContent string `gorm:"notify_party_content" json:"notify_party_content"`
//	MawbNo string `gorm:"mawb_no" json:"mawb_no"` // 航空提单号码
//	CarrierName string `gorm:"carrier_name" json:"carrier_name"` // 航空公司名称
//	HawbNo string `gorm:"hawb_no" json:"hawb_no"` // 公司提单号
//	DepartureAirportId int `gorm:"departure_airport_id" json:"departure_airport_id"`
//	CodeOfTwoId int `gorm:"code_of_two_id" json:"code_of_two_id"`
//	Destination string `gorm:"destination" json:"destination"` // 目的地地址
//	OceanChangesPaytypeId string `gorm:"ocean_changes_paytype_id" json:"ocean_changes_paytype_id"`
//	OtherChangesPaytypeId string `gorm:"other_changes_paytype_id" json:"other_changes_paytype_id"`
//	Marks string `gorm:"marks" json:"marks"` // 唛头
//	PackageNo string `gorm:"package_no" json:"package_no"` // 包装数量
//	DescriptionOfGood string `gorm:"description_of_good" json:"description_of_good"`
//	GrossWeight string `gorm:"gross_weight" json:"gross_weight"` // 毛重
//	Measurement string `gorm:"measurement" json:"measurement"` // 尺码
//	ShipperDeclaration string `gorm:"shipper_declaration" json:"shipper_declaration"` // 发货人声明
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	SerialNo string `gorm:"serial_no" json:"serial_no"` // 序列号
//	BookingDate string `gorm:"booking_date" json:"booking_date"` // 订舱日期
//	ConfirmDate string `gorm:"confirm_date" json:"confirm_date"` // 确认日期
//	ShipperType string `gorm:"shipper_type" json:"shipper_type"`
//	ShipperId int `gorm:"shipper_id" json:"shipper_id"`
//	ConsigneeType string `gorm:"consignee_type" json:"consignee_type"`
//	ConsigneeId int `gorm:"consignee_id" json:"consignee_id"`
//	NotifyPartyType string `gorm:"notify_party_type" json:"notify_party_type"`
//	NotifyPartyId int `gorm:"notify_party_id" json:"notify_party_id"`
//	PodAgentType string `gorm:"pod_agent_type" json:"pod_agent_type"`
//	PodAgentId int `gorm:"pod_agent_id" json:"pod_agent_id"`
//	PodAgentContent string `gorm:"pod_agent_content" json:"pod_agent_content"`
//	DestinationAirportId int `gorm:"destination_airport_id" json:"destination_airport_id"`
//	TransshipmentAirportId int `gorm:"transshipment_airport_id" json:"transshipment_airport_id"`
//	Status string `gorm:"status" json:"status"`
//	InvoiceStatus string `gorm:"invoice_status" json:"invoice_status"`
//	InvoiceNo string `gorm:"invoice_no" json:"invoice_no"`
//	AssignStatus string `gorm:"assign_status" json:"assign_status"`
//	SupplierAgentType string `gorm:"supplier_agent_type" json:"supplier_agent_type"`
//	SupplierAgentId int `gorm:"supplier_agent_id" json:"supplier_agent_id"`
//	SupplierAgentContent string `gorm:"supplier_agent_content" json:"supplier_agent_content"`
//	FlightDate string `gorm:"flight_date" json:"flight_date"`
//	CostDescription string `gorm:"cost_description" json:"cost_description"`
//	Remarks string `gorm:"remarks" json:"remarks"`
//	UserMarketId int `gorm:"user_market_id" json:"user_market_id"`
//	RandomFile int `gorm:"random_file" json:"random_file"`
//	AssociatedFormers string `gorm:"associated_formers" json:"associated_formers"`
//	SupplierCompanyAgentId int `gorm:"supplier_company_agent_id" json:"supplier_company_agent_id"`
//	SupplierCompanyAgentContent string `gorm:"supplier_company_agent_content" json:"supplier_company_agent_content"`
//	BaseDataPackageType int `gorm:"base_data_package_type" json:"base_data_package_type"`
//	Flight string `gorm:"flight" json:"flight"`
//	DeliveryAddress string `gorm:"delivery_address" json:"delivery_address"`
//	ArriveAddress string `gorm:"arrive_address" json:"arrive_address"`
//	EstimatedTimeOfDeparture string `gorm:"estimated_time_of_departure" json:"estimated_time_of_departure"`
//	Dimension string `gorm:"dimension" json:"dimension"` // 尺寸
//}
//
//func (*Former_airs_bookings) TableName() string {
//	return "former_airs_bookings"
//}
//
//type Former_trails_transports struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Type string `gorm:"type" json:"type"`
//	SerialNo string `gorm:"serial_no" json:"serial_no"` // 流水号
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"`
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"`
//	CompanyInstructionContent string `gorm:"company_instruction_content" json:"company_instruction_content"` // 委托单位详情
//	OfWay int `gorm:"of_way" json:"of_way"` // 运输方式
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"` // 运输公司
//	UserId int `gorm:"user_id" json:"user_id"` // 联系人
//	InvoiceNo string `gorm:"invoice_no" json:"invoice_no"` // 发票
//	InvoiceStatus string `gorm:"invoice_status" json:"invoice_status"`
//	AssignStatus string `gorm:"assign_status" json:"assign_status"`
//	Status string `gorm:"status" json:"status"` // 状态
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//	PlaceOfDeparture string `gorm:"place_of_departure" json:"place_of_departure"` // 起运地
//	Destination string `gorm:"destination" json:"destination"` // 目的地
//	AssociatedFormers string `gorm:"associated_formers" json:"associated_formers"` // 关联表单
//	TruckerRepTel string `gorm:"trucker_rep_tel" json:"trucker_rep_tel"` // 联系电话
//	LicensePlateNumber string `gorm:"license_plate_number" json:"license_plate_number"` // 车牌号
//	HkLicensePlateNumber string `gorm:"hk_license_plate_number" json:"hk_license_plate_number"` // 香港车牌号
//	DriverMobi string `gorm:"driver_mobi" json:"driver_mobi"` // 司机联系电话
//	HkDriverMobi string `gorm:"hk_driver_mobi" json:"hk_driver_mobi"` // 香港司机联系电话
//	CustomsOrderInfo string `gorm:"customs_order_info" json:"customs_order_info"` // 报关信息
//	Remarks string `gorm:"remarks" json:"remarks"` // 备注
//	DescriptionOfGood string `gorm:"description_of_good" json:"description_of_good"` // 品名
//	Number string `gorm:"number" json:"number"` // 件数
//	GrossWeight string `gorm:"gross_weight" json:"gross_weight"` // 重量
//	Size string `gorm:"size" json:"size"` // 体积
//	LoadingDate string `gorm:"loading_date" json:"loading_date"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	Receiver string `gorm:"receiver" json:"receiver"`
//	UserOperatorId int `gorm:"user_operator_id" json:"user_operator_id"` // 操作
//	UserSalesmanId int `gorm:"user_salesman_id" json:"user_salesman_id"` // 业务
//	UserMarketId int `gorm:"user_market_id" json:"user_market_id"` // 市场
//	UserCustomerId int `gorm:"user_customer_id" json:"user_customer_id"` // 客服
//	UserFileId int `gorm:"user_file_id" json:"user_file_id"` // 文件
//	BaseDataInstructionTypeId int `gorm:"base_data_instruction_type_id" json:"base_data_instruction_type_id"` // 委托类型
//	WaysOfDeclarationId int `gorm:"ways_of_declaration_id" json:"ways_of_declaration_id"` // 报关方式
//	TransshipmentId int `gorm:"transshipment_id" json:"transshipment_id"` // 转运
//	BaseDataItemId int `gorm:"base_data_item_id" json:"base_data_item_id"` // 装运条款
//	ServiceContractNo string `gorm:"service_contract_no" json:"service_contract_no"` // 合同编号
//	Marks string `gorm:"marks" json:"marks"`
//	IsWeighing int `gorm:"is_weighing" json:"is_weighing"` // 过磅
//	IsDrivingLicense int `gorm:"is_driving_license" json:"is_driving_license"` // 转关司机带本
//	IsDeclare int `gorm:"is_declare" json:"is_declare"`
//	Dimension string `gorm:"dimension" json:"dimension"` // 尺寸
//	IsGoodsAttachment int `gorm:"is_goods_attachment" json:"is_goods_attachment"` // 导出货描列表
//	BaseDataPackageType int `gorm:"base_data_package_type" json:"base_data_package_type"`
//}
//
//func (*Former_trails_transports) TableName() string {
//	return "former_trails_transports"
//}
//
//type Freight_charges struct {
//	Id int `gorm:"id" json:"id"`
//	Type string `gorm:"type" json:"type"` // 类型【整箱,拼箱,拖车】
//	Name string `gorm:"name" json:"name"` // 加价设置名称
//	Group string `gorm:"group" json:"group"` // 加价组别【内部加价,客户加价,未登录用户】
//	DeparturePortIds string `gorm:"departure_port_ids" json:"departure_port_ids"` // 起运港/城市
//	BoatCompanyIds string `gorm:"boat_company_ids" json:"boat_company_ids"` // 船公司
//	SeaLineIds string `gorm:"sea_line_ids" json:"sea_line_ids"` // 航线
//	DestinationPortIds string `gorm:"destination_port_ids" json:"destination_port_ids"` // 目的港
//	ScaleType string `gorm:"scale_type" json:"scale_type"` // 加价类型【金额，百分比】
//	ScaleGroupIds string `gorm:"scale_group_ids" json:"scale_group_ids"` // 用户组ids
//	ScaleValue int `gorm:"scale_value" json:"scale_value"` // 加价数值
//	Gp20 float64 `gorm:"gp20" json:"gp20"` // 20GP
//	Gp40 float64 `gorm:"gp40" json:"gp40"` // 40GP
//	Hq40 float64 `gorm:"hq40" json:"hq40"` // 40GP
//	Tons float64 `gorm:"tons" json:"tons"` // tons
//	Cbms float64 `gorm:"cbms" json:"cbms"` // cbms
//	LclCharge float64 `gorm:"lcl_charge" json:"lcl_charge"` // 附加费
//	ChargeOne float64 `gorm:"charge_one" json:"charge_one"` // 拖车价格一段
//	ChargeTwo float64 `gorm:"charge_two" json:"charge_two"` // 拖车价格二段
//	ChargeThree float64 `gorm:"charge_three" json:"charge_three"` // 拖车价格三段
//	ChargeFour float64 `gorm:"charge_four" json:"charge_four"` // 拖车价格四段
//	ChargeFive float64 `gorm:"charge_five" json:"charge_five"` // 拖车价格五段
//	ChargeSix float64 `gorm:"charge_six" json:"charge_six"` // 拖车价格六段
//	Enabled int `gorm:"enabled" json:"enabled"` // 有效状态
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Position int `gorm:"position" json:"position"`
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Freight_charges) TableName() string {
//	return "freight_charges"
//}
//
//type Plan_mains struct {
//	Id int `gorm:"id" json:"id"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Enabled int `gorm:"enabled" json:"enabled"` // 有效状态
//	UserOperatorId int `gorm:"user_operator_id" json:"user_operator_id"` // 操作员
//	UserSalesmanId int `gorm:"user_salesman_id" json:"user_salesman_id"` // 业务员
//	UserFileId int `gorm:"user_file_id" json:"user_file_id"` // 文件
//	BusinessmanId int `gorm:"businessman_id" json:"businessman_id"` // 商务
//	UserFeeId int `gorm:"user_fee_id" json:"user_fee_id"` // 财务
//	UserCustomerId int `gorm:"user_customer_id" json:"user_customer_id"` // 客服
//	UserAuditId int `gorm:"user_audit_id" json:"user_audit_id"` // 价格审核
//	Profit float64 `gorm:"profit" json:"profit"` // 利润
//	PolId int `gorm:"pol_id" json:"pol_id"` // 起运港
//	PodId int `gorm:"pod_id" json:"pod_id"` // 目的港
//	GoodCount int `gorm:"good_count" json:"good_count"` // 包装件数
//	PackageTypeId int `gorm:"package_type_id" json:"package_type_id"` // 包装类型
//	GrossWeight float64 `gorm:"gross_weight" json:"gross_weight"` // 毛重
//	GoodSize float64 `gorm:"good_size" json:"good_size"` // 货物体积
//	CompanyId int `gorm:"company_id" json:"company_id"` // 船公司/航空公司/运输公司
//	SupplierCompanyAgentId int `gorm:"supplier_company_agent_id" json:"supplier_company_agent_id"`
//	MawbNo string `gorm:"mawb_no" json:"mawb_no"` // mawb编号
//	CutOffDate string `gorm:"cut_off_date" json:"cut_off_date"` // 截关/开船日期
//	DepartureDate string `gorm:"departure_date" json:"departure_date"` // 离港日期/起飞时间
//	ArrivalDate string `gorm:"arrival_date" json:"arrival_date"` // 到达时间
//	BillProduceId int `gorm:"bill_produce_id" json:"bill_produce_id"` // 出单方式
//	Note string `gorm:"note" json:"note"` // 备注
//	InlandDriverPhone string `gorm:"inland_driver_phone" json:"inland_driver_phone"` // 内陆司机手机号
//	InlandCarNumber string `gorm:"inland_car_number" json:"inland_car_number"` // 内陆车牌
//	HongkongDriverPhone string `gorm:"hongkong_driver_phone" json:"hongkong_driver_phone"` // 香港司机手机号
//	HongkongCarNumber string `gorm:"hongkong_car_number" json:"hongkong_car_number"` // 香港车牌
//	Status string `gorm:"status" json:"status"` // 状态
//	PodAgentId int `gorm:"pod_agent_id" json:"pod_agent_id"` // 目的港代理
//	PodAgentType string `gorm:"pod_agent_type" json:"pod_agent_type"` // 目的港代理类型
//	PodContact string `gorm:"pod_contact" json:"pod_contact"` // 目的港联系人
//	PodEmail string `gorm:"pod_email" json:"pod_email"` // 目的港联系人邮箱
//	CourierNumber string `gorm:"courier_number" json:"courier_number"` // 快递单号
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//	TransshipmentId int `gorm:"transshipment_id" json:"transshipment_id"`
//}
//
//func (*Plan_mains) TableName() string {
//	return "plan_mains"
//}
//
//type Base_data_departures struct {
//	Id int `gorm:"id" json:"id"`
//	City string `gorm:"city" json:"city"`
//	Address string `gorm:"address" json:"address"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Priority int `gorm:"priority" json:"priority"`
//	District string `gorm:"district" json:"district"` // 区域
//	DestinationCity string `gorm:"destination_city" json:"destination_city"` // 目的地
//	CityPya string `gorm:"city_pya" json:"city_pya"` // 城市拼音全写
//	CityPyf string `gorm:"city_pyf" json:"city_pyf"` // 城市拼音简写
//	AddressPya string `gorm:"address_pya" json:"address_pya"` // 地址拼音全写
//	AddressPyf string `gorm:"address_pyf" json:"address_pyf"` // 地址拼音简写
//	DistrictPya string `gorm:"district_pya" json:"district_pya"` // 目的地拼音全写
//	DistrictPyf string `gorm:"district_pyf" json:"district_pyf"` // 目的地拼音简写
//	Longitude string `gorm:"longitude" json:"longitude"` // 经度
//	Latitude string `gorm:"latitude" json:"latitude"` // 纬度
//	LocationAddress string `gorm:"location_address" json:"location_address"` // 定位地址
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//}
//
//func (*Base_data_departures) TableName() string {
//	return "base_data_departures"
//}
//
//type Finance_fee_types struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"` // 费用名
//	NameCn string `gorm:"name_cn" json:"name_cn"` // 中文名
//	NameEn string `gorm:"name_en" json:"name_en"` // 英文名
//	ReceivableTagId int `gorm:"receivable_tag_id" json:"receivable_tag_id"` // 默认应收标签
//	PayableTagId int `gorm:"payable_tag_id" json:"payable_tag_id"` // 默认应付标签
//	DefaultValue float64 `gorm:"default_value" json:"default_value"` // 默认金额
//	Remarks string `gorm:"remarks" json:"remarks"` // 备注
//	FinanceCurrencyId int `gorm:"finance_currency_id" json:"finance_currency_id"` // 外部关联货币
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Finance_fee_types) TableName() string {
//	return "finance_fee_types"
//}
//
//type Finance_search_templates struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"` // 模板名称
//	Enabled int `gorm:"enabled" json:"enabled"` // 有效状态
//	UserId int `gorm:"user_id" json:"user_id"` // 用户id
//	Content string `gorm:"content" json:"content"` // 搜索内容(json字符串)
//	Description string `gorm:"description" json:"description"` // 描述
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//}
//
//func (*Finance_search_templates) TableName() string {
//	return "finance_search_templates"
//}
//
//type Message_mailboxes struct {
//	Id int `gorm:"id" json:"id"`
//	FromId int `gorm:"from_id" json:"from_id"` // 发件人id
//	FromType string `gorm:"from_type" json:"from_type"` // 发件人类型
//	FromName string `gorm:"from_name" json:"from_name"` // 发件人姓名
//	FromEmail string `gorm:"from_email" json:"from_email"` // 发送者邮件
//	Status string `gorm:"status" json:"status"` // 状态
//	IsReplay string `gorm:"is_replay" json:"is_replay"` // 是否可回复
//	SendTime string `gorm:"send_time" json:"send_time"` // 发送时间
//	SourceId int `gorm:"source_id" json:"source_id"` // 信息来源id
//	SourceType string `gorm:"source_type" json:"source_type"` // 信息来源
//	Category string `gorm:"category" json:"category"` // 类别
//	Ways string `gorm:"ways" json:"ways"` // 渠道
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	Enabled int `gorm:"enabled" json:"enabled"`
//}
//
//func (*Message_mailboxes) TableName() string {
//	return "message_mailboxes"
//}
//
//type Order_entrust_lists struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	ClientUserId int `gorm:"client_user_id" json:"client_user_id"` // 前端下单人
//	PayStatus string `gorm:"pay_status" json:"pay_status"`
//	VerifyStatus string `gorm:"verify_status" json:"verify_status"`
//	LandingBillUrl string `gorm:"landing_bill_url" json:"landing_bill_url"` // 生成的提单URL
//	InstructionUrl string `gorm:"instruction_url" json:"instruction_url"` // 委托的URL
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Order_entrust_lists) TableName() string {
//	return "order_entrust_lists"
//}
//
//type Order_extend_infos struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Number float64 `gorm:"number" json:"number"` // 件数
//	BaseDataPackageType int `gorm:"base_data_package_type" json:"base_data_package_type"` // 类型
//	GrossWeight float64 `gorm:"gross_weight" json:"gross_weight"` // 毛重
//	Size float64 `gorm:"size" json:"size"` // 体积
//	HblSo string `gorm:"hbl_so" json:"hbl_so"` // 分单号
//	MblSo string `gorm:"mbl_so" json:"mbl_so"` // 主单号
//	BoatCompanyId int `gorm:"boat_company_id" json:"boat_company_id"` // 船公司
//	SeaPolId int `gorm:"sea_pol_id" json:"sea_pol_id"` // 起运港
//	SeaPodId int `gorm:"sea_pod_id" json:"sea_pod_id"` // 目的港
//	AirPolId int `gorm:"air_pol_id" json:"air_pol_id"` // 起运港
//	AirPodId int `gorm:"air_pod_id" json:"air_pod_id"` // 目的港
//	CutOffDay string `gorm:"cut_off_day" json:"cut_off_day"` // 截关日期/开船日期
//	FlightDate string `gorm:"flight_date" json:"flight_date"` // 起飞日期
//	EndDate string `gorm:"end_date" json:"end_date"` // 到达时间
//	IsCutOff int `gorm:"is_cut_off" json:"is_cut_off"` // 是否开船
//	IsArrive int `gorm:"is_arrive" json:"is_arrive"` // 是否到港
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	Ratio float64 `gorm:"ratio" json:"ratio"` // 换算系数
//	RatioWeight float64 `gorm:"ratio_weight" json:"ratio_weight"` // 材积重
//	Bubble float64 `gorm:"bubble" json:"bubble"` // 分泡
//	SoNo string `gorm:"so_no" json:"so_no"` // so信息
//	CodeOfTwoId int `gorm:"code_of_two_id" json:"code_of_two_id"`
//	CourierCodeId int `gorm:"courier_code_id" json:"courier_code_id"`
//	StartDate string `gorm:"start_date" json:"start_date"`
//	IsHml int `gorm:"is_hml" json:"is_hml"`
//	IsAms int `gorm:"is_ams" json:"is_ams"`
//	IsIfs int `gorm:"is_ifs" json:"is_ifs"`
//	ChargedWeight float64 `gorm:"charged_weight" json:"charged_weight"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	MiscBill int `gorm:"misc_bill" json:"misc_bill"`
//	BoxSizeCount string `gorm:"box_size_count" json:"box_size_count"`
//	Vessel string `gorm:"vessel" json:"vessel"`
//	Voyage string `gorm:"voyage" json:"voyage"`
//	Flight string `gorm:"flight" json:"flight"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Order_extend_infos) TableName() string {
//	return "order_extend_infos"
//}
//
//type Verifications struct {
//	Id int `gorm:"id" json:"id"`
//	Number string `gorm:"number" json:"number"` // 核销单号
//	ClosingUnitName string `gorm:"closing_unit_name" json:"closing_unit_name"` // 结算单位名称
//	Currency string `gorm:"currency" json:"currency"` // 币种
//	Amount float64 `gorm:"amount" json:"amount"` // 不含税金额
//	TaxAmount float64 `gorm:"tax_amount" json:"tax_amount"` // 税金
//	Category string `gorm:"category" json:"category"` // 类别[冲销，核销]
//	ClosingUnitType string `gorm:"closing_unit_type" json:"closing_unit_type"`
//	ClosingUnitId int `gorm:"closing_unit_id" json:"closing_unit_id"` // 结算单位
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	UserId int `gorm:"user_id" json:"user_id"`
//	Remark string `gorm:"remark" json:"remark"` // 备注
//	PayTypeId int `gorm:"pay_type_id" json:"pay_type_id"` // 付款类型
//	CompanyAccountId int `gorm:"company_account_id" json:"company_account_id"`
//	ClosingUnitAccountId int `gorm:"closing_unit_account_id" json:"closing_unit_account_id"`
//	AmountCost float64 `gorm:"amount_cost" json:"amount_cost"` // 总的核销金额
//	FinanceCurrencyId int `gorm:"finance_currency_id" json:"finance_currency_id"`
//	PayOrReceive string `gorm:"pay_or_receive" json:"pay_or_receive"`
//	PayOrReceiveTime string `gorm:"pay_or_receive_time" json:"pay_or_receive_time"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Verifications) TableName() string {
//	return "verifications"
//}
//
//type Base_data_customer_business_types struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Base_data_customer_business_types) TableName() string {
//	return "base_data_customer_business_types"
//}
//
//type Score_system_details struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	ScoreId int `gorm:"score_id" json:"score_id"` // 对应的积分表
//	ScoreItem string `gorm:"score_item" json:"score_item"` // 积分项目
//	OperatorItem string `gorm:"operator_item" json:"operator_item"` // 操作项目,[系统,人工]
//	UserOperatorId int `gorm:"user_operator_id" json:"user_operator_id"` // 系统操作人员
//	UserAuditId int `gorm:"user_audit_id" json:"user_audit_id"` // 系统审核人员
//	OperaScore float64 `gorm:"opera_score" json:"opera_score"` // 操作的积分
//	ScoreObject string `gorm:"score_object" json:"score_object"` // 积分操作对象
//	ScoreMultiple string `gorm:"score_multiple" json:"score_multiple"` // 操作倍数
//	Remark string `gorm:"remark" json:"remark"` // 备注
//	OperatorDetail string `gorm:"operator_detail" json:"operator_detail"`
//	Status string `gorm:"status" json:"status"` // 操作状态
//	SourceType string `gorm:"source_type" json:"source_type"`
//	SourceId int `gorm:"source_id" json:"source_id"` // 数据来源
//	DataObjectType string `gorm:"data_object_type" json:"data_object_type"`
//	DataObjectId int `gorm:"data_object_id" json:"data_object_id"`
//	ScoreRuleId int `gorm:"score_rule_id" json:"score_rule_id"`
//	CompanyId int `gorm:"company_id" json:"company_id"`
//}
//
//func (*Score_system_details) TableName() string {
//	return "score_system_details"
//}
//
//type Bargain_supply_quotations struct {
//	Id int `gorm:"id" json:"id"`
//	SupplyId int `gorm:"supply_id" json:"supply_id"` // 供应商
//	BoatCompanyId int `gorm:"boat_company_id" json:"boat_company_id"` // 船公司
//	QuotationNo string `gorm:"quotation_no" json:"quotation_no"` // 报价编号
//	ValidStartAt string `gorm:"valid_start_at" json:"valid_start_at"` // 有效期
//	ValidEndAt string `gorm:"valid_end_at" json:"valid_end_at"` // 有效期
//	CodeOfTwoId int `gorm:"code_of_two_id" json:"code_of_two_id"` // 航空公司
//	Remarks string `gorm:"remarks" json:"remarks"` // 备注
//	AirFees float64 `gorm:"air_fees" json:"air_fees"`
//	LclFees float64 `gorm:"lcl_fees" json:"lcl_fees"`
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	Platform string `gorm:"platform" json:"platform"` // 来源
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	BargainMainId int `gorm:"bargain_main_id" json:"bargain_main_id"`
//	BargainPlatformInquiryId int `gorm:"bargain_platform_inquiry_id" json:"bargain_platform_inquiry_id"`
//	SeaPolId int `gorm:"sea_pol_id" json:"sea_pol_id"`
//	SeaPodId int `gorm:"sea_pod_id" json:"sea_pod_id"`
//	Status int `gorm:"status" json:"status"` // 供应商报价是否提交
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsDirectBargain int `gorm:"is_direct_bargain" json:"is_direct_bargain"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Bargain_supply_quotations) TableName() string {
//	return "bargain_supply_quotations"
//}
//
//type Base_data_package_types struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"` // 包装单位简称
//	NameCn string `gorm:"name_cn" json:"name_cn"` // 中文名称
//	NameEn string `gorm:"name_en" json:"name_en"` // 英文名称
//	EdiCode string `gorm:"edi_code" json:"edi_code"` // EDI代码
//	Remark string `gorm:"remark" json:"remark"` // 备注
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//}
//
//func (*Base_data_package_types) TableName() string {
//	return "base_data_package_types"
//}
//
//type Former_seas_cap_vgms struct {
//	Id int `gorm:"id" json:"id"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	SerialNo string `gorm:"serial_no" json:"serial_no"` // 序列号
//	VerifiedGrossMass int `gorm:"verified_gross_mass" json:"verified_gross_mass"` // 称重公斤数
//	ResponsibleParty string `gorm:"responsible_party" json:"responsible_party"` // 责任方
//	AuthorizedPerson string `gorm:"authorized_person" json:"authorized_person"` // 负责人
//	IncludeContainer int `gorm:"include_container" json:"include_container"` // 称重方式两种，一种含柜一种不含柜，此项必填，0为不含柜，1为含柜
//	WeighingParty string `gorm:"weighing_party" json:"weighing_party"` // 称重方，可选
//	Status string `gorm:"status" json:"status"` // 状态
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	SoNo string `gorm:"so_no" json:"so_no"`
//	ContainerNo string `gorm:"container_no" json:"container_no"`
//	SealNo string `gorm:"seal_no" json:"seal_no"`
//	GrossUnit string `gorm:"gross_unit" json:"gross_unit"`
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"`
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"`
//	CompanyInstructionContent string `gorm:"company_instruction_content" json:"company_instruction_content"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Former_seas_cap_vgms) TableName() string {
//	return "former_seas_cap_vgms"
//}
//
//type Youtu_crm_cooperator_clues struct {
//	Id int `gorm:"id" json:"id"`
//	NameCn string `gorm:"name_cn" json:"name_cn"` // 公司中文名称
//	NameNick string `gorm:"name_nick" json:"name_nick"` // 公司简称
//	NameEn string `gorm:"name_en" json:"name_en"` // 公司英文名称
//	Website string `gorm:"website" json:"website"` // 网址
//	CompanyType int `gorm:"company_type" json:"company_type"` // 公司类型
//	BusinessTypeName string `gorm:"business_type_name" json:"business_type_name"` // 公司业务类型
//	CustomerSource int `gorm:"customer_source" json:"customer_source"` // 客户来源: 线上注册,广告,陌拜
//	Province string `gorm:"province" json:"province"` // 客户所在省
//	City string `gorm:"city" json:"city"` // 客户所在市
//	District string `gorm:"district" json:"district"` // 客户所在区
//	Address string `gorm:"address" json:"address"` // 客户具体中文地址
//	Address2 string `gorm:"address2" json:"address2"` // 客户具体英文地址
//	Remark string `gorm:"remark" json:"remark"` // 对公司的备注
//	Email string `gorm:"email" json:"email"` // 公司邮箱
//	Phone string `gorm:"phone" json:"phone"` // 公司座机
//	ContactName string `gorm:"contact_name" json:"contact_name"` // 联系人姓名
//	ContactPost string `gorm:"contact_post" json:"contact_post"` // 联系人职位
//	ContactMobi string `gorm:"contact_mobi" json:"contact_mobi"` // 联系人手机
//	ContactGender int `gorm:"contact_gender" json:"contact_gender"` // 联系人性别
//	ContactEmail string `gorm:"contact_email" json:"contact_email"` // 联系人邮箱
//	ContactTelephone string `gorm:"contact_telephone" json:"contact_telephone"` // 联系人座机
//	ContactWechat string `gorm:"contact_wechat" json:"contact_wechat"` // 联系人微信
//	ContactQq string `gorm:"contact_qq" json:"contact_qq"` // 联系人QQ
//	ContactFax string `gorm:"contact_fax" json:"contact_fax"` // 联系人FAX
//	ContactRemark string `gorm:"contact_remark" json:"contact_remark"` // 对联系人的备注
//	AuditorId int `gorm:"auditor_id" json:"auditor_id"` // 审核人的users.id
//	CreatorId int `gorm:"creator_id" json:"creator_id"` // 线索创建人的users.id
//	LiableUserId int `gorm:"liable_user_id" json:"liable_user_id"` // 负责人的users.id
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"` // 线索转化为客户之后的公司user_companies.id
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"` // 默认软删除
//	CompanyId int `gorm:"company_id" json:"company_id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Status string `gorm:"status" json:"status"` // 审核状态
//	Type string `gorm:"type" json:"type"`
//}
//
//func (*Youtu_crm_cooperator_clues) TableName() string {
//	return "youtu_crm_cooperator_clues"
//}
//
//type Apartment_manages struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"` // 数据库名
//	HostName string `gorm:"host_name" json:"host_name"` // 主机名
//	PortName int `gorm:"port_name" json:"port_name"` // 端口
//	AppId string `gorm:"app_id" json:"app_id"`
//	AppSecret string `gorm:"app_secret" json:"app_secret"`
//	DefaultPassword string `gorm:"default_password" json:"default_password"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	CompanyName string `gorm:"company_name" json:"company_name"` // 公司名称
//	ValidityDate string `gorm:"validity_date" json:"validity_date"` // 有效期
//	ContactInfo string `gorm:"contact_info" json:"contact_info"` // 联系信息
//	TrailerLocation int `gorm:"trailer_location" json:"trailer_location"` // 拖车定位使用次数
//}
//
//func (*Apartment_manages) TableName() string {
//	return "apartment_manages"
//}
//
//type Dashboard_statistics struct {
//	Id int `gorm:"id" json:"id"`
//	ObjId int `gorm:"obj_id" json:"obj_id"` // 对象id
//	Category string `gorm:"category" json:"category"` // 对象类别
//	ObjName string `gorm:"obj_name" json:"obj_name"` // 对象名称
//	Description string `gorm:"description" json:"description"` // 描述
//	OrganizationId int `gorm:"organization_id" json:"organization_id"` // 组织id
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//}
//
//func (*Dashboard_statistics) TableName() string {
//	return "dashboard_statistics"
//}
//
//type Fielders struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Name string `gorm:"name" json:"name"` // 字段名称
//	Label string `gorm:"label" json:"label"` // 字段备注
//	FieldType string `gorm:"field_type" json:"field_type"` // 字段类型
//	Comment string `gorm:"comment" json:"comment"` // 字段的描述
//	Hint string `gorm:"hint" json:"hint"` // 字段的提示
//	Accessibility int `gorm:"accessibility" json:"accessibility"` // 只读,可读可写,隐藏字段
//	UserId int `gorm:"user_id" json:"user_id"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"` // 行级乐观锁
//	TableName string `gorm:"table_name" json:"table_name"` // 表名
//}
//
//func (*Fielders) TableName() string {
//	return "fielders"
//}
//
//type Help_feedbacks struct {
//	Id int `gorm:"id" json:"id"`
//	Content string `gorm:"content" json:"content"` // 内容
//	FeedbackUserId int `gorm:"feedback_user_id" json:"feedback_user_id"` // 反馈人
//	Url string `gorm:"url" json:"url"`
//	Cookie string `gorm:"cookie" json:"cookie"`
//	Json string `gorm:"json" json:"json"`
//	ScreenShot string `gorm:"screen_shot" json:"screen_shot"` // 全屏截图
//	ScreenCapture string `gorm:"screen_capture" json:"screen_capture"` // 截图
//	ScreenCaptureHtml string `gorm:"screen_capture_html" json:"screen_capture_html"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Help_feedbacks) TableName() string {
//	return "help_feedbacks"
//}
//
//type Scores struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	SourceType string `gorm:"source_type" json:"source_type"`
//	SourceId int `gorm:"source_id" json:"source_id"`
//	ScoreGrowth float64 `gorm:"score_growth" json:"score_growth"` // 积分成长值
//	ScoreCurrent float64 `gorm:"score_current" json:"score_current"` // 可用积分
//	Grade string `gorm:"grade" json:"grade"` // 等级
//}
//
//func (*Scores) TableName() string {
//	return "scores"
//}
//
//type Approvers struct {
//	Id int `gorm:"id" json:"id"`
//	ApprovalApplicationId int `gorm:"approval_application_id" json:"approval_application_id"`
//	UserId int `gorm:"user_id" json:"user_id"` // 审批人id
//	UserName string `gorm:"user_name" json:"user_name"` // 审批人姓名
//	Sort int `gorm:"sort" json:"sort"` // 审批顺序
//	Status string `gorm:"status" json:"status"` // 审批状态
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Approvers) TableName() string {
//	return "approvers"
//}
//
//type Bargain_customer_quotations struct {
//	Id int `gorm:"id" json:"id"`
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"` // 客户
//	CustomerQuotationNo string `gorm:"customer_quotation_no" json:"customer_quotation_no"` // 编号
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"` // 订单
//	FinishedTime string `gorm:"finished_time" json:"finished_time"` // 完成时间
//	Status string `gorm:"status" json:"status"` // 状态
//	Remarks string `gorm:"remarks" json:"remarks"` // 状态
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	AirFees float64 `gorm:"air_fees" json:"air_fees"`
//	LclFees float64 `gorm:"lcl_fees" json:"lcl_fees"`
//	BargainSupplyQuotationId int `gorm:"bargain_supply_quotation_id" json:"bargain_supply_quotation_id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	BargainMainId int `gorm:"bargain_main_id" json:"bargain_main_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsDirectBargain int `gorm:"is_direct_bargain" json:"is_direct_bargain"` // 直接生成客户报价
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Bargain_customer_quotations) TableName() string {
//	return "bargain_customer_quotations"
//}
//
//type Base_data_sea_ports struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"` // 海港
//	SeaLineName string `gorm:"sea_line_name" json:"sea_line_name"` // 航线名
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	NameCn string `gorm:"name_cn" json:"name_cn"` // 海港中文名
//	NationCn string `gorm:"nation_cn" json:"nation_cn"` // 国家中文名
//	NationEn string `gorm:"nation_en" json:"nation_en"` // 国家英文名
//	SeaLineCn string `gorm:"sea_line_cn" json:"sea_line_cn"` // 航线中文名
//	SeaLineEn string `gorm:"sea_line_en" json:"sea_line_en"` // 航线英文名
//	NamePya string `gorm:"name_pya" json:"name_pya"` // 海港拼音全写
//	NamePyf string `gorm:"name_pyf" json:"name_pyf"` // 海港拼音简写
//	NationCode string `gorm:"nation_code" json:"nation_code"` // 国家代码
//	NationCode2 string `gorm:"nation_code2" json:"nation_code2"` // 国家代码
//	CityId int `gorm:"city_id" json:"city_id"` // 城市ID
//	City string `gorm:"city" json:"city"` // 码头所属城市，一般为内贸
//	CityEn string `gorm:"city_en" json:"city_en"` // 城市英文名
//	CityPya string `gorm:"city_pya" json:"city_pya"` // 城市拼音全写
//	CityPyf string `gorm:"city_pyf" json:"city_pyf"` // 城市拼音简写
//	HasFreight int `gorm:"has_freight" json:"has_freight"` // 附加费
//	Position int `gorm:"position" json:"position"`
//	CountryId int `gorm:"country_id" json:"country_id"` // 所属国家ID
//	SeaLineIdSet string `gorm:"sea_line_id_set" json:"sea_line_id_set"` // 所属航线（多个）
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//	BaseDataSeaLinesId int `gorm:"base_data_sea_lines_id" json:"base_data_sea_lines_id"`
//}
//
//func (*Base_data_sea_ports) TableName() string {
//	return "base_data_sea_ports"
//}
//
//type Users struct {
//	Id int `gorm:"id" json:"id"`
//	Email string `gorm:"email" json:"email"` // email
//	EncryptedPassword string `gorm:"encrypted_password" json:"encrypted_password"`
//	ResetPasswordToken string `gorm:"reset_password_token" json:"reset_password_token"`
//	ResetPasswordSentAt string `gorm:"reset_password_sent_at" json:"reset_password_sent_at"`
//	RememberCreatedAt string `gorm:"remember_created_at" json:"remember_created_at"`
//	SignInCount int `gorm:"sign_in_count" json:"sign_in_count"`
//	CurrentSignInAt string `gorm:"current_sign_in_at" json:"current_sign_in_at"`
//	LastSignInAt string `gorm:"last_sign_in_at" json:"last_sign_in_at"`
//	CurrentSignInIp string `gorm:"current_sign_in_ip" json:"current_sign_in_ip"`
//	LastSignInIp string `gorm:"last_sign_in_ip" json:"last_sign_in_ip"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	DepartmentId int `gorm:"department_id" json:"department_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	Name string `gorm:"name" json:"name"` // 姓名
//	VerifiedAt string `gorm:"verified_at" json:"verified_at"`
//	IsSync int `gorm:"is_sync" json:"is_sync"`
//	AuthenticationToken string `gorm:"authentication_token" json:"authentication_token"`
//	InitialPassword string `gorm:"initial_password" json:"initial_password"`
//	DeviseUuid string `gorm:"devise_uuid" json:"devise_uuid"` // 单点登录验证
//	IsAdmin int `gorm:"is_admin" json:"is_admin"` // 是否为超级管理人员(系统默认只有一位)
//	Provider string `gorm:"provider" json:"provider"`
//	Uid string `gorm:"uid" json:"uid"`
//	UserNo string `gorm:"user_no" json:"user_no"` // 工号
//	Phone string `gorm:"phone" json:"phone"`
//}
//
//func (*Users) TableName() string {
//	return "users"
//}
//
//type Youtu_erp_permissions struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	YoutuErpRoleId int `gorm:"youtu_erp_role_id" json:"youtu_erp_role_id"`
//	Name string `gorm:"name" json:"name"`
//	SubjectClass string `gorm:"subject_class" json:"subject_class"`
//	SubjectId int `gorm:"subject_id" json:"subject_id"`
//	Action string `gorm:"action" json:"action"`
//	Description string `gorm:"description" json:"description"`
//	ActionType string `gorm:"action_type" json:"action_type"`
//	Klass string `gorm:"klass" json:"klass"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Youtu_erp_permissions) TableName() string {
//	return "youtu_erp_permissions"
//}
//
//type Base_data_boat_companies struct {
//	Id int `gorm:"id" json:"id"`
//	ChinaName string `gorm:"china_name" json:"china_name"` // 中文名
//	Name string `gorm:"name" json:"name"` // 船运公司名称
//	Url string `gorm:"url" json:"url"` // 官网地址
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	ChinaNamePya string `gorm:"china_name_pya" json:"china_name_pya"` // 船运公司拼音全写
//	ChinaNamePyf string `gorm:"china_name_pyf" json:"china_name_pyf"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//}
//
//func (*Base_data_boat_companies) TableName() string {
//	return "base_data_boat_companies"
//}
//
//type Departments struct {
//	Id int `gorm:"id" json:"id"`
//	NameCn string `gorm:"name_cn" json:"name_cn"` // 部门中文名
//	NameEn string `gorm:"name_en" json:"name_en"` // 部门英文名
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Departments) TableName() string {
//	return "departments"
//}
//
//type Doc_template_defaults struct {
//	Id int `gorm:"id" json:"id"`
//	YoutuErpUserCompanyId int `gorm:"youtu_erp_user_company_id" json:"youtu_erp_user_company_id"`
//	DocTemplateDocTemplateId int `gorm:"doc_template_doc_template_id" json:"doc_template_doc_template_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//}
//
//func (*Doc_template_defaults) TableName() string {
//	return "doc_template_defaults"
//}
//
//type Bargain_mains struct {
//	Id int `gorm:"id" json:"id"`
//	MainNo string `gorm:"main_no" json:"main_no"` // 价单编号
//	BargainType int `gorm:"bargain_type" json:"bargain_type"` // 价单类型
//	TransportType int `gorm:"transport_type" json:"transport_type"` // 运输类型
//	Status string `gorm:"status" json:"status"` // 价单状态
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"` // 客户
//	CompanyContactName string `gorm:"company_contact_name" json:"company_contact_name"` // 联系人
//	CompanyContactPhone string `gorm:"company_contact_phone" json:"company_contact_phone"` // 联系人电话
//	UserOperationId int `gorm:"user_operation_id" json:"user_operation_id"` // 业务人员
//	ExpectedShipmentTime string `gorm:"expected_shipment_time" json:"expected_shipment_time"` // 预计出货时间
//	OrderMasterSerialNumber string `gorm:"order_master_serial_number" json:"order_master_serial_number"` // 订单编号
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"` // 订单
//	Remarks string `gorm:"remarks" json:"remarks"` // 价单备注
//	SeaLineId int `gorm:"sea_line_id" json:"sea_line_id"` // 航线
//	CutOffDay string `gorm:"cut_off_day" json:"cut_off_day"` // 截关/截港日期
//	FlightDate string `gorm:"flight_date" json:"flight_date"` // 航班日期
//	BoatCompanyId int `gorm:"boat_company_id" json:"boat_company_id"` // 船公司
//	SeaPolId int `gorm:"sea_pol_id" json:"sea_pol_id"` // 起运港
//	SeaPodId int `gorm:"sea_pod_id" json:"sea_pod_id"` // 目的港
//	GrossWeight string `gorm:"gross_weight" json:"gross_weight"` // 毛重
//	Size string `gorm:"size" json:"size"` // 体积
//	AirPolId int `gorm:"air_pol_id" json:"air_pol_id"` // 航空起运港
//	AirPodId int `gorm:"air_pod_id" json:"air_pod_id"` // 航空目的港
//	CodeOfTwoId int `gorm:"code_of_two_id" json:"code_of_two_id"` // 航空公司
//	SeaFreightId int `gorm:"sea_freight_id" json:"sea_freight_id"` // 对应的海运费
//	FinishedTime string `gorm:"finished_time" json:"finished_time"` // 成交时间
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	Platform string `gorm:"platform" json:"platform"` // 线下
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	CityId int `gorm:"city_id" json:"city_id"`
//	UserId int `gorm:"user_id" json:"user_id"` // 创建询价单用户id
//	CompanyType string `gorm:"company_type" json:"company_type"` // 客户类型:默认为合作单位
//	CompanyName string `gorm:"company_name" json:"company_name"` // 公司名称
//	PlaceOfReceipt string `gorm:"place_of_receipt" json:"place_of_receipt"` // 提货地址
//	NotifyPartyAddress string `gorm:"notify_party_address" json:"notify_party_address"` // 交货地
//	BaseDataTradeTermsId int `gorm:"base_data_trade_terms_id" json:"base_data_trade_terms_id"` // 贸易条款
//	BaseDataItemId int `gorm:"base_data_item_id" json:"base_data_item_id"` // 装运条款
//	IsSaveClue int `gorm:"is_save_clue" json:"is_save_clue"` // 客户信息是否保存为线索
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//	BargainRemarks string `gorm:"bargain_remarks" json:"bargain_remarks"`
//}
//
//func (*Bargain_mains) TableName() string {
//	return "bargain_mains"
//}
//
//type Finance_statements struct {
//	Id int `gorm:"id" json:"id"`
//	ClosingUnitId int `gorm:"closing_unit_id" json:"closing_unit_id"` // 结算单位,方便检索
//	StatementType string `gorm:"statement_type" json:"statement_type"` // 账单类型 cash_now(票结) cash_month(月结)
//	ConfirmBy string `gorm:"confirm_by" json:"confirm_by"` // 前台确认还是后台确认
//	Status string `gorm:"status" json:"status"` // 前台对账状态
//	Remarks string `gorm:"remarks" json:"remarks"` // 备注
//	TemplateId int `gorm:"template_id" json:"template_id"` // 账单模板的id
//	CustomStatementHtml string `gorm:"custom_statement_html" json:"custom_statement_html"` // 如果操作在模板的基础上修改生成的账单，保存该html
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	FeeIdJson string `gorm:"fee_id_json" json:"fee_id_json"` // 費用的id
//	FinanceMonthStatementId int `gorm:"finance_month_statement_id" json:"finance_month_statement_id"` // 月结对账单id
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"` // 订单号
//	DebitNoteSn string `gorm:"debit_note_sn" json:"debit_note_sn"` // 账单号
//	BeginAt string `gorm:"begin_at" json:"begin_at"` // 月结开始时间
//	EndAt string `gorm:"end_at" json:"end_at"` // 月结账单结束时间
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"` // 生成账单的公司,用于权限控制，满足子公司不能看总公司账单
//	UserId int `gorm:"user_id" json:"user_id"` // 导出账单的人员
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Finance_statements) TableName() string {
//	return "finance_statements"
//}
//
//type Former_warehouse_orders struct {
//	Id int `gorm:"id" json:"id"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"` // 订单
//	SerialNo string `gorm:"serial_no" json:"serial_no"` // 流水号
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"` // 委托单位
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"` // 委托单位
//	CompanyInstructionContent string `gorm:"company_instruction_content" json:"company_instruction_content"`
//	WarehouseAddress string `gorm:"warehouse_address" json:"warehouse_address"` // 仓库地址
//	DeliveryAddress string `gorm:"delivery_address" json:"delivery_address"` // 送货说明
//	DistributionReason string `gorm:"distribution_reason" json:"distribution_reason"` // 配货原因
//	Note string `gorm:"note" json:"note"` // 备注
//	ClientNote string `gorm:"client_note" json:"client_note"` // 客户备注
//	WarehouseNo string `gorm:"warehouse_no" json:"warehouse_no"` // 入仓单号
//	WarehouseTime string `gorm:"warehouse_time" json:"warehouse_time"` // 入仓时间
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Former_warehouse_orders) TableName() string {
//	return "former_warehouse_orders"
//}
//
//type Track_models struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	Content string `gorm:"content" json:"content"`
//	Position int `gorm:"position" json:"position"`
//	Status string `gorm:"status" json:"status"`
//	Process string `gorm:"process" json:"process"`
//}
//
//func (*Track_models) TableName() string {
//	return "track_models"
//}
//
//type Youtu_crm_clue_tracks struct {
//	Id int `gorm:"id" json:"id"`
//	SourceType string `gorm:"source_type" json:"source_type"`
//	SourceId int `gorm:"source_id" json:"source_id"`
//	Description string `gorm:"description" json:"description"` // 跟进描述
//	OfWay string `gorm:"of_way" json:"of_way"` // 跟进方式
//	TrackingUserId int `gorm:"tracking_user_id" json:"tracking_user_id"` // 跟进人
//	NextTrackTime string `gorm:"next_track_time" json:"next_track_time"` // 下次跟进时间
//	Nofity int `gorm:"nofity" json:"nofity"` // 是否通知,可能有多种通知方式,定为integer预留扩展
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"` // 默认软删除
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"` // 所属公司
//}
//
//func (*Youtu_crm_clue_tracks) TableName() string {
//	return "youtu_crm_clue_tracks"
//}
//
//type Base_data_paytypes struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//}
//
//func (*Base_data_paytypes) TableName() string {
//	return "base_data_paytypes"
//}
//
//type Message_contents struct {
//	Id int `gorm:"id" json:"id"`
//	Title string `gorm:"title" json:"title"` // 标题
//	Context string `gorm:"context" json:"context"` // 内容
//	Attachment string `gorm:"attachment" json:"attachment"` // 附件
//	MailboxId int `gorm:"mailbox_id" json:"mailbox_id"` // 邮件id
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Message_contents) TableName() string {
//	return "message_contents"
//}
//
//type Settings struct {
//	Id int `gorm:"id" json:"id"`
//	SourceType string `gorm:"source_type" json:"source_type"`
//	SourceId int `gorm:"source_id" json:"source_id"`
//	Key string `gorm:"key" json:"key"`
//	Value string `gorm:"value" json:"value"`
//	Group string `gorm:"group" json:"group"`
//	FilterType string `gorm:"filter_type" json:"filter_type"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Label string `gorm:"label" json:"label"`
//	CompanyId int `gorm:"company_id" json:"company_id"`
//}
//
//func (*Settings) TableName() string {
//	return "settings"
//}
//
//type Users_youtu_erp_roles struct {
//	UserId int `gorm:"user_id" json:"user_id"`
//	RoleId int `gorm:"role_id" json:"role_id"`
//}
//
//func (*Users_youtu_erp_roles) TableName() string {
//	return "users_youtu_erp_roles"
//}
//
//type Website_pages struct {
//	Id int `gorm:"id" json:"id"`
//	Title string `gorm:"title" json:"title"` // 标题
//	Keyword string `gorm:"keyword" json:"keyword"` // 关键字
//	Details string `gorm:"details" json:"details"` // 详情内容
//	WebsiteMenuId int `gorm:"website_menu_id" json:"website_menu_id"` // 所属菜单
//	Locale string `gorm:"locale" json:"locale"` // 语言选择
//}
//
//func (*Website_pages) TableName() string {
//	return "website_pages"
//}
//
//type Addresses struct {
//	Id int `gorm:"id" json:"id"`
//	UserCompaniesId int `gorm:"user_companies_id" json:"user_companies_id"`
//	UserName string `gorm:"user_name" json:"user_name"` // 收件人姓名
//	PhoneNumber string `gorm:"phone_number" json:"phone_number"` // 收件人手机号
//	Address string `gorm:"address" json:"address"` // 收件人地址
//	PostCode string `gorm:"post_code" json:"post_code"` // 邮编
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Addresses) TableName() string {
//	return "addresses"
//}
//
//type Base_data_instruction_types struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Name string `gorm:"name" json:"name"`
//}
//
//func (*Base_data_instruction_types) TableName() string {
//	return "base_data_instruction_types"
//}
//
//type Cargo_lists struct {
//	Id int `gorm:"id" json:"id"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	Enabled int `gorm:"enabled" json:"enabled"` // 有效状态
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"` // 委托单位
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"` // 委托单位类型
//	Name string `gorm:"name" json:"name"` // 货物名称
//	UnitCount int `gorm:"unit_count" json:"unit_count"` // 数量
//	PackageTypeId int `gorm:"package_type_id" json:"package_type_id"` // 包装类型id
//	PackageTypeName string `gorm:"package_type_name" json:"package_type_name"` // 包装类型名称
//	Length float64 `gorm:"length" json:"length"` // 长
//	Width float64 `gorm:"width" json:"width"` // 宽
//	Height float64 `gorm:"height" json:"height"` // 高
//	Size float64 `gorm:"size" json:"size"` // 体积
//	Count int `gorm:"count" json:"count"` // 总数
//	UnitPrice float64 `gorm:"unit_price" json:"unit_price"` // 单价
//	Amount float64 `gorm:"amount" json:"amount"` // 总价
//	Marks string `gorm:"marks" json:"marks"` // 唛头
//	GrossWeight float64 `gorm:"gross_weight" json:"gross_weight"` // 毛重
//	UnitSize float64 `gorm:"unit_size" json:"unit_size"` // 单件体积
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	CustomsNumber string `gorm:"customs_number" json:"customs_number"` // 报关序号
//	DeclareElements string `gorm:"declare_elements" json:"declare_elements"` // 申报要素
//	CustomsCode string `gorm:"customs_code" json:"customs_code"` // 海关编码
//	NetWeight float64 `gorm:"net_weight" json:"net_weight"`
//	SourceType string `gorm:"source_type" json:"source_type"`
//	SourceId int `gorm:"source_id" json:"source_id"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//	FilterType string `gorm:"filter_type" json:"filter_type"`
//	InstructionId int `gorm:"instruction_id" json:"instruction_id"`
//}
//
//func (*Cargo_lists) TableName() string {
//	return "cargo_lists"
//}
//
//type Former_airs_instructions struct {
//	Id int `gorm:"id" json:"id"`
//	ShipperContent string `gorm:"shipper_content" json:"shipper_content"` // 发货人信息
//	ConsigneeContent string `gorm:"consignee_content" json:"consignee_content"` // 收货人信息
//	NotifyPartyContent string `gorm:"notify_party_content" json:"notify_party_content"`
//	DepartureAirportId int `gorm:"departure_airport_id" json:"departure_airport_id"`
//	CodeOfTwoId int `gorm:"code_of_two_id" json:"code_of_two_id"`
//	OceanChangesPaytypeId string `gorm:"ocean_changes_paytype_id" json:"ocean_changes_paytype_id"`
//	OtherChangesPaytypeId string `gorm:"other_changes_paytype_id" json:"other_changes_paytype_id"`
//	Marks string `gorm:"marks" json:"marks"` // 唛头
//	Number string `gorm:"number" json:"number"`
//	GrossWeight string `gorm:"gross_weight" json:"gross_weight"` // 毛重
//	Size string `gorm:"size" json:"size"`
//	ShipperDeclaration string `gorm:"shipper_declaration" json:"shipper_declaration"` // 发货人声明
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	SerialNo string `gorm:"serial_no" json:"serial_no"`
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"`
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"`
//	ContactInstructionType string `gorm:"contact_instruction_type" json:"contact_instruction_type"`
//	ContactInstructionId int `gorm:"contact_instruction_id" json:"contact_instruction_id"`
//	CompanyInstructionContent string `gorm:"company_instruction_content" json:"company_instruction_content"`
//	ContactInstructionContent string `gorm:"contact_instruction_content" json:"contact_instruction_content"`
//	ShipperType string `gorm:"shipper_type" json:"shipper_type"`
//	ShipperId int `gorm:"shipper_id" json:"shipper_id"`
//	ConsigneeType string `gorm:"consignee_type" json:"consignee_type"`
//	ConsigneeId int `gorm:"consignee_id" json:"consignee_id"`
//	NotifyPartyType string `gorm:"notify_party_type" json:"notify_party_type"`
//	NotifyPartyId int `gorm:"notify_party_id" json:"notify_party_id"`
//	PodAgentType string `gorm:"pod_agent_type" json:"pod_agent_type"`
//	PodAgentId int `gorm:"pod_agent_id" json:"pod_agent_id"`
//	PodAgentContent string `gorm:"pod_agent_content" json:"pod_agent_content"`
//	DestinationAirportId int `gorm:"destination_airport_id" json:"destination_airport_id"`
//	TransshipmentAirportId int `gorm:"transshipment_airport_id" json:"transshipment_airport_id"`
//	Status string `gorm:"status" json:"status"`
//	InvoiceNo string `gorm:"invoice_no" json:"invoice_no"`
//	RandomFile int `gorm:"random_file" json:"random_file"`
//	FlightDate string `gorm:"flight_date" json:"flight_date"`
//	Remarks string `gorm:"remarks" json:"remarks"`
//	WaysOfDeclarationId int `gorm:"ways_of_declaration_id" json:"ways_of_declaration_id"`
//	InstructionTypeId int `gorm:"instruction_type_id" json:"instruction_type_id"`
//	ServiceContractNo string `gorm:"service_contract_no" json:"service_contract_no"`
//	BaseDataItemId int `gorm:"base_data_item_id" json:"base_data_item_id"`
//	BaseDataTradeTermId int `gorm:"base_data_trade_term_id" json:"base_data_trade_term_id"`
//	TransshipmentId int `gorm:"transshipment_id" json:"transshipment_id"`
//	DescriptionOfGood string `gorm:"description_of_good" json:"description_of_good"`
//	BaseDataPackageType int `gorm:"base_data_package_type" json:"base_data_package_type"`
//	Flight string `gorm:"flight" json:"flight"`
//	DeliveryAddress string `gorm:"delivery_address" json:"delivery_address"`
//	ArriveAddress string `gorm:"arrive_address" json:"arrive_address"`
//	EstimatedTimeOfDeparture string `gorm:"estimated_time_of_departure" json:"estimated_time_of_departure"`
//	BaseDataBillProduceId int `gorm:"base_data_bill_produce_id" json:"base_data_bill_produce_id"`
//	Dimension string `gorm:"dimension" json:"dimension"` // 尺寸
//	PayPodId int `gorm:"pay_pod_id" json:"pay_pod_id"`
//	HawbRemarks string `gorm:"hawb_remarks" json:"hawb_remarks"`
//	IsGoodsAttachment int `gorm:"is_goods_attachment" json:"is_goods_attachment"`
//	SubBlNo string `gorm:"sub_bl_no" json:"sub_bl_no"`
//	CurrencyId int `gorm:"currency_id" json:"currency_id"`
//	ExecuteDate string `gorm:"execute_date" json:"execute_date"`
//	HblRemarks string `gorm:"hbl_remarks" json:"hbl_remarks"`
//	Ratio float64 `gorm:"ratio" json:"ratio"` // 材积换算系数
//	RatioWeight float64 `gorm:"ratio_weight" json:"ratio_weight"` // 材积重
//	Bubble float64 `gorm:"bubble" json:"bubble"` // 分泡%
//	ChargableWeight float64 `gorm:"chargable_weight" json:"chargable_weight"` // 计费重
//	Type string `gorm:"type" json:"type"`
//}
//
//func (*Former_airs_instructions) TableName() string {
//	return "former_airs_instructions"
//}
//
//type Global_sources struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	SourceType string `gorm:"source_type" json:"source_type"`
//	SourceId int `gorm:"source_id" json:"source_id"` // 数据来源
//	SourcePlatform string `gorm:"source_platform" json:"source_platform"` // 来源平台
//	SourceRecordPath string `gorm:"source_record_path" json:"source_record_path"` // 来源路径
//	SourceUuid string `gorm:"source_uuid" json:"source_uuid"`
//	SourceUyid string `gorm:"source_uyid" json:"source_uyid"`
//	SourceUuidRegister int `gorm:"source_uuid_register" json:"source_uuid_register"` // 唯一id注册
//	ParentSourceUuid string `gorm:"parent_source_uuid" json:"parent_source_uuid"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Global_sources) TableName() string {
//	return "global_sources"
//}
//
//type User_freight_charges struct {
//	Id int `gorm:"id" json:"id"`
//	UserId int `gorm:"user_id" json:"user_id"`
//	FreightChargeId int `gorm:"freight_charge_id" json:"freight_charge_id"`
//	Enabled int `gorm:"enabled" json:"enabled"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//}
//
//func (*User_freight_charges) TableName() string {
//	return "user_freight_charges"
//}
//
//type Youtu_crm_tags struct {
//	Id int `gorm:"id" json:"id"`
//	Type string `gorm:"type" json:"type"`
//	TagId int `gorm:"tag_id" json:"tag_id"`
//	TagName string `gorm:"tag_name" json:"tag_name"`
//	SourceType string `gorm:"source_type" json:"source_type"`
//	SourceId int `gorm:"source_id" json:"source_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"` // 默认软删除
//	TransportType int `gorm:"transport_type" json:"transport_type"`
//	Remarks string `gorm:"remarks" json:"remarks"`
//}
//
//func (*Youtu_crm_tags) TableName() string {
//	return "youtu_crm_tags"
//}
//
//type Finance_currencies struct {
//	Id int `gorm:"id" json:"id"`
//	RateRealtime float64 `gorm:"rate_realtime" json:"rate_realtime"` // 实时汇率
//	RateFix float64 `gorm:"rate_fix" json:"rate_fix"` // 固定汇率
//	RateResult float64 `gorm:"rate_result" json:"rate_result"` // 实际汇率
//	NameEn string `gorm:"name_en" json:"name_en"` // 英文名
//	NameCn string `gorm:"name_cn" json:"name_cn"` // 中文名
//	Symbo string `gorm:"symbo" json:"symbo"` // 货币符号
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//	ValidTime string `gorm:"valid_time" json:"valid_time"` // 汇率生效时间
//	InvalidTime string `gorm:"invalid_time" json:"invalid_time"` // 汇率失效时间
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Finance_currencies) TableName() string {
//	return "finance_currencies"
//}
//
//type Former_air_deliveries struct {
//	Id int `gorm:"id" json:"id"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	PolId int `gorm:"pol_id" json:"pol_id"` // 起运港
//	PodId int `gorm:"pod_id" json:"pod_id"` // 目的港
//	FlightDate string `gorm:"flight_date" json:"flight_date"` // 航班日期
//	ArrivalDate string `gorm:"arrival_date" json:"arrival_date"` // 到达日期
//	CourierCompanyId int `gorm:"courier_company_id" json:"courier_company_id"` // 快递公司
//	CourierNumber string `gorm:"courier_number" json:"courier_number"` // 快递单号
//	FlightCompanyId int `gorm:"flight_company_id" json:"flight_company_id"` // 航空公司
//	FlightNumber string `gorm:"flight_number" json:"flight_number"` // 航班
//	WarehousingDate string `gorm:"warehousing_date" json:"warehousing_date"` // 入仓日期
//	DeliverDate string `gorm:"deliver_date" json:"deliver_date"` // 派送日期
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"` // 委托单位
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"` // 委托单位类型
//	CompanyInstructionContent string `gorm:"company_instruction_content" json:"company_instruction_content"` // 委托单位详情
//	ShipperId int `gorm:"shipper_id" json:"shipper_id"` // 发货人
//	ShipperType string `gorm:"shipper_type" json:"shipper_type"`
//	ShipperContent string `gorm:"shipper_content" json:"shipper_content"` // 发货人详情
//	ConsigneeId int `gorm:"consignee_id" json:"consignee_id"` // 收货人
//	ConsigneeType string `gorm:"consignee_type" json:"consignee_type"`
//	ConsigneeContent string `gorm:"consignee_content" json:"consignee_content"` // 收货人详情
//	NotifyPartyId int `gorm:"notify_party_id" json:"notify_party_id"` // 通知人
//	NotifyPartyType string `gorm:"notify_party_type" json:"notify_party_type"`
//	NotifyPartyContent string `gorm:"notify_party_content" json:"notify_party_content"` // 通知人详情
//	OceanChangesPaytypeId int `gorm:"ocean_changes_paytype_id" json:"ocean_changes_paytype_id"` // 运费付款方式
//	OtherChangesPaytypeId int `gorm:"other_changes_paytype_id" json:"other_changes_paytype_id"` // 其他付款方式
//	WaysOfDeclarationId int `gorm:"ways_of_declaration_id" json:"ways_of_declaration_id"` // 报关方式
//	InstructionTypeId int `gorm:"instruction_type_id" json:"instruction_type_id"` // 委托类型
//	BaseDataTradeTermsId int `gorm:"base_data_trade_terms_id" json:"base_data_trade_terms_id"` // 贸易条款
//	TransshipmentId int `gorm:"transshipment_id" json:"transshipment_id"` // 转运
//	BaseDataItemId int `gorm:"base_data_item_id" json:"base_data_item_id"` // 装运条款
//	HsCode string `gorm:"hs_code" json:"hs_code"` // 商品编码
//	ServiceContractNo string `gorm:"service_contract_no" json:"service_contract_no"` // 合同编码
//	InvoiceNo string `gorm:"invoice_no" json:"invoice_no"` // 发票号
//	Insurance int `gorm:"insurance" json:"insurance"` // 是否需要发票
//	InsuranceAmount float64 `gorm:"insurance_amount" json:"insurance_amount"` // 保险金额
//	ReceiverAddress string `gorm:"receiver_address" json:"receiver_address"` // 收货地址
//	ConsignorAddress string `gorm:"consignor_address" json:"consignor_address"` // 发货地址
//	ReceiverCityId int `gorm:"receiver_city_id" json:"receiver_city_id"` // 收货城市
//	ConsignorCityId int `gorm:"consignor_city_id" json:"consignor_city_id"` // 发货城市
//	ViaId int `gorm:"via_id" json:"via_id"` // 中转港
//	TransshipmentFlight string `gorm:"transshipment_flight" json:"transshipment_flight"` // 转运航班
//	TransshipmentEtd string `gorm:"transshipment_etd" json:"transshipment_etd"` // 转运ETD
//	Enabled int `gorm:"enabled" json:"enabled"` // 有效状态
//	Status string `gorm:"status" json:"status"` // 状态
//	Note string `gorm:"note" json:"note"` // 备注
//	ClientNote string `gorm:"client_note" json:"client_note"` // 客户备注
//	SerialNumber string `gorm:"serial_number" json:"serial_number"` // 流水号
//	Marks string `gorm:"marks" json:"marks"` // 唛头
//	Number string `gorm:"number" json:"number"` // 数量
//	DescriptionOfGood string `gorm:"description_of_good" json:"description_of_good"` // 货物描述
//	GrossWeight string `gorm:"gross_weight" json:"gross_weight"` // 毛重
//	Size string `gorm:"size" json:"size"` // 体积
//	SerialNo string `gorm:"serial_no" json:"serial_no"` // 序列号
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Ratio float64 `gorm:"ratio" json:"ratio"` // 材积换算系数
//	RatioWeight float64 `gorm:"ratio_weight" json:"ratio_weight"` // 材积重
//	Bubble float64 `gorm:"bubble" json:"bubble"` // 分泡%
//	Dimension string `gorm:"dimension" json:"dimension"` // 体积
//	ChargedWeight float64 `gorm:"charged_weight" json:"charged_weight"` // 计费重
//	BaseDataPackageType int `gorm:"base_data_package_type" json:"base_data_package_type"` // 包装类型
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Former_air_deliveries) TableName() string {
//	return "former_air_deliveries"
//}
//
//type People struct {
//	Id int `gorm:"id" json:"id"`
//	NameNick string `gorm:"name_nick" json:"name_nick"` // 昵称
//	NameCn string `gorm:"name_cn" json:"name_cn"` // 中文名
//	NameEn string `gorm:"name_en" json:"name_en"` // 英文名
//	Gender int `gorm:"gender" json:"gender"` // 0：未知，1：男，2女
//	Post string `gorm:"post" json:"post"` // 职位
//	Mobi string `gorm:"mobi" json:"mobi"` // 手机号码
//	Mobi2 string `gorm:"mobi2" json:"mobi2"` // 备用号码
//	Telephone string `gorm:"telephone" json:"telephone"` // 座机
//	Telephone2 string `gorm:"telephone2" json:"telephone2"` // 备用座机
//	Fax string `gorm:"fax" json:"fax"` // 传真
//	Email string `gorm:"email" json:"email"`
//	Qq string `gorm:"qq" json:"qq"`
//	AddrCompany string `gorm:"addr_company" json:"addr_company"` // 公司地址
//	AddrHome string `gorm:"addr_home" json:"addr_home"` // 家庭地址
//	Remark string `gorm:"remark" json:"remark"` // 备注
//	Remark2 string `gorm:"remark2" json:"remark2"` // 备用备注
//	IsMainContact int `gorm:"is_main_contact" json:"is_main_contact"` // 是否主要联系人
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	SourceType string `gorm:"source_type" json:"source_type"`
//	SourceId int `gorm:"source_id" json:"source_id"` // 数据来源
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*People) TableName() string {
//	return "people"
//}
//
//type Bargain_shipping_dates struct {
//	Id int `gorm:"id" json:"id"`
//	Days string `gorm:"days" json:"days"` // 航程/天数
//	StartDate int `gorm:"start_date" json:"start_date"` // 截关/截港/航班日期
//	Remarks string `gorm:"remarks" json:"remarks"` // 备注
//	BargainMainId int `gorm:"bargain_main_id" json:"bargain_main_id"`
//	BargainSupplyQuotationId int `gorm:"bargain_supply_quotation_id" json:"bargain_supply_quotation_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	CutOffDay int `gorm:"cut_off_day" json:"cut_off_day"`
//	TransshipmentInfo string `gorm:"transshipment_info" json:"transshipment_info"`
//	SeaTransshipmentId int `gorm:"sea_transshipment_id" json:"sea_transshipment_id"`
//	AirTransshipmentId int `gorm:"air_transshipment_id" json:"air_transshipment_id"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Bargain_shipping_dates) TableName() string {
//	return "bargain_shipping_dates"
//}
//
//type Base_data_account_titles struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"` // 会计科目名称
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Base_data_account_titles) TableName() string {
//	return "base_data_account_titles"
//}
//
//type Former_seas_cargo_lists struct {
//	Id int `gorm:"id" json:"id"`
//	SoNo string `gorm:"so_no" json:"so_no"`
//	ContainerNo string `gorm:"container_no" json:"container_no"`
//	SealNo string `gorm:"seal_no" json:"seal_no"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//	VerifiedGrossMass float64 `gorm:"verified_gross_mass" json:"verified_gross_mass"`
//	IncludeContainer int `gorm:"include_container" json:"include_container"` // 称重方式
//	GrossUnit int `gorm:"gross_unit" json:"gross_unit"` // 重量单位
//	GrossWeight float64 `gorm:"gross_weight" json:"gross_weight"` // 重量
//	Measurement float64 `gorm:"measurement" json:"measurement"` // 体积
//	Count int `gorm:"count" json:"count"` // 件数
//	CapTypeSizeId int `gorm:"cap_type_size_id" json:"cap_type_size_id"` // 柜型
//	Marks string `gorm:"marks" json:"marks"` // 唛头
//	DescriptionOfGood string `gorm:"description_of_good" json:"description_of_good"` // 货物描述
//	PackageTypeId int `gorm:"package_type_id" json:"package_type_id"` // 包裝類型
//	SerialNo string `gorm:"serial_no" json:"serial_no"` // 流水号
//	Status string `gorm:"status" json:"status"`
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"`
//	ContainerWeight float64 `gorm:"container_weight" json:"container_weight"`
//	VgmWeight float64 `gorm:"vgm_weight" json:"vgm_weight"`
//	FormerSeasInstructionId int `gorm:"former_seas_instruction_id" json:"former_seas_instruction_id"`
//	FilterType string `gorm:"filter_type" json:"filter_type"`
//}
//
//func (*Former_seas_cargo_lists) TableName() string {
//	return "former_seas_cargo_lists"
//}
//
//type Former_seas_vgms struct {
//	Id int `gorm:"id" json:"id"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	SerialNo string `gorm:"serial_no" json:"serial_no"` // 序列号
//	VerifiedGrossMass int `gorm:"verified_gross_mass" json:"verified_gross_mass"` // 称重公斤数
//	ResponsibleParty string `gorm:"responsible_party" json:"responsible_party"` // 责任方
//	AuthorizedPerson string `gorm:"authorized_person" json:"authorized_person"` // 负责人
//	IncludeContainer int `gorm:"include_container" json:"include_container"` // 称重方式两种，一种含柜一种不含柜，此项必填，0为不含柜，1为含柜
//	WeighingParty string `gorm:"weighing_party" json:"weighing_party"` // 称重方，可选
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Status string `gorm:"status" json:"status"`
//}
//
//func (*Former_seas_vgms) TableName() string {
//	return "former_seas_vgms"
//}
//
//type Former_trails_load_infos struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Loader string `gorm:"loader" json:"loader"` // 装货联系人
//	SupplierMobi string `gorm:"supplier_mobi" json:"supplier_mobi"` // 联系人电话
//	PlaceOfReceipt string `gorm:"place_of_receipt" json:"place_of_receipt"` // 装货地信息
//	GrossWeight string `gorm:"gross_weight" json:"gross_weight"` // 货物重量
//	Size string `gorm:"size" json:"size"` // 货物体积
//	EstimatedTimeOfLoading string `gorm:"estimated_time_of_loading" json:"estimated_time_of_loading"` // 预计装货时间
//	LoadingDate string `gorm:"loading_date" json:"loading_date"` // 预计装货时间
//	SourceType string `gorm:"source_type" json:"source_type"`
//	SourceId int `gorm:"source_id" json:"source_id"`
//	Count int `gorm:"count" json:"count"` // 数量
//	LoadOrUnload int `gorm:"load_or_unload" json:"load_or_unload"` // 装货卸货,0为装货地，1为卸货地
//}
//
//func (*Former_trails_load_infos) TableName() string {
//	return "former_trails_load_infos"
//}
//
//type Invoice_info struct {
//	Id int `gorm:"id" json:"id"`
//	InvoiceId int `gorm:"invoice_id" json:"invoice_id"`
//	AddressId int `gorm:"address_id" json:"address_id"`
//	Category string `gorm:"category" json:"category"` // 发票类型
//	Number string `gorm:"number" json:"number"` // 发票号
//	CourierCompanyId int `gorm:"courier_company_id" json:"courier_company_id"` // 快递公司
//	CourierCompanyName string `gorm:"courier_company_name" json:"courier_company_name"` // 快递公司名称
//	CourierNumber string `gorm:"courier_number" json:"courier_number"` // 快递单号
//	Amount float64 `gorm:"amount" json:"amount"` // 发票金额
//	Status string `gorm:"status" json:"status"` // 发票状态
//	Enabled int `gorm:"enabled" json:"enabled"` // 有效状态
//	Note string `gorm:"note" json:"note"` // 备注
//	Date string `gorm:"date" json:"date"` // 开票时间
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	ClosingUnitId int `gorm:"closing_unit_id" json:"closing_unit_id"` // 结算单位
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"` // 所属公司
//	UserId int `gorm:"user_id" json:"user_id"`
//	ApplyUserId int `gorm:"apply_user_id" json:"apply_user_id"`
//	FinanceCurrencyId int `gorm:"finance_currency_id" json:"finance_currency_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Invoice_info) TableName() string {
//	return "invoice_info"
//}
//
//type Setting_number_histories struct {
//	Id int `gorm:"id" json:"id"`
//	Year int `gorm:"year" json:"year"`
//	Month int `gorm:"month" json:"month"`
//	Day int `gorm:"day" json:"day"`
//	SettingNumberId int `gorm:"setting_number_id" json:"setting_number_id"`
//	CurrentNumber int `gorm:"current_number" json:"current_number"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//}
//
//func (*Setting_number_histories) TableName() string {
//	return "setting_number_histories"
//}
//
//type Base_data_code_levels struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Name string `gorm:"name" json:"name"` // 层级名称
//	CodeName string `gorm:"code_name" json:"code_name"`
//	Status int `gorm:"status" json:"status"` // 是否有效
//	SourceType string `gorm:"source_type" json:"source_type"` // 来源
//	Remark string `gorm:"remark" json:"remark"` // 说明
//	LockVersion int `gorm:"lock_version" json:"lock_version"` // 行级乐观锁
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//}
//
//func (*Base_data_code_levels) TableName() string {
//	return "base_data_code_levels"
//}
//
//type Finance_profit_drawings struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Name string `gorm:"name" json:"name"` // 提成名称
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	BaseAmount float64 `gorm:"base_amount" json:"base_amount"`
//	EndAmount float64 `gorm:"end_amount" json:"end_amount"`
//	Rate float64 `gorm:"rate" json:"rate"`
//	Method int `gorm:"method" json:"method"`
//	Remarks string `gorm:"remarks" json:"remarks"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Finance_profit_drawings) TableName() string {
//	return "finance_profit_drawings"
//}
//
//type Schedules struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"` // 待办事项名称
//	Category string `gorm:"category" json:"category"` // 类型
//	Note string `gorm:"note" json:"note"` // 备注
//	Status string `gorm:"status" json:"status"` // 状态
//	UserId int `gorm:"user_id" json:"user_id"` // 提醒人id
//	UserName string `gorm:"user_name" json:"user_name"` // 提醒人姓名
//	ObjectId int `gorm:"object_id" json:"object_id"` // 数据源id
//	ObjectName string `gorm:"object_name" json:"object_name"` // 数据源名称
//	ObjectUrl string `gorm:"object_url" json:"object_url"` // 数据源url
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//}
//
//func (*Schedules) TableName() string {
//	return "schedules"
//}
//
//type Bargain_extend_infos struct {
//	Id int `gorm:"id" json:"id"`
//	Remarks1 string `gorm:"remarks_1" json:"remarks_1"` // 存放优途价格（经加价），客户议价，供应商价（原价）
//	Remarks2 string `gorm:"remarks_2" json:"remarks_2"`
//	Remarks3 string `gorm:"remarks_3" json:"remarks_3"`
//	Remarks4 string `gorm:"remarks_4" json:"remarks_4"`
//	Remarks5 string `gorm:"remarks_5" json:"remarks_5"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	BargainMainId int `gorm:"bargain_main_id" json:"bargain_main_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Bargain_extend_infos) TableName() string {
//	return "bargain_extend_infos"
//}
//
//type Former_customs_orders struct {
//	Id int `gorm:"id" json:"id"`
//	CustomNo string `gorm:"custom_no" json:"custom_no"` // 报关单号
//	LicenceNo string `gorm:"licence_no" json:"licence_no"` // 许可证号
//	InspectNo string `gorm:"inspect_no" json:"inspect_no"` // 商检编号
//	HsCode string `gorm:"hs_code" json:"hs_code"` // 商品编码
//	CustomBrokerType string `gorm:"custom_broker_type" json:"custom_broker_type"`
//	CustomBrokerId int `gorm:"custom_broker_id" json:"custom_broker_id"`
//	ApplicationDate string `gorm:"application_date" json:"application_date"` // 报关日期
//	BaseDataCurrencyId int `gorm:"base_data_currency_id" json:"base_data_currency_id"` // 申报货币
//	BaseDataCiqtypeId int `gorm:"base_data_ciqtype_id" json:"base_data_ciqtype_id"` // 报关类型
//	BaseDataSeaPortId int `gorm:"base_data_sea_port_id" json:"base_data_sea_port_id"` // 港口
//	ConfirmDate string `gorm:"confirm_date" json:"confirm_date"` // 确认日期
//	SubmitDate string `gorm:"submit_date" json:"submit_date"`
//	BaseDataDockId int `gorm:"base_data_dock_id" json:"base_data_dock_id"` // 码头
//	ContractNo string `gorm:"contract_no" json:"contract_no"` // 合同编号
//	PreRecordNo string `gorm:"pre_record_no" json:"pre_record_no"` // 预录编号
//	CancelNo string `gorm:"cancel_no" json:"cancel_no"` // 核销单号
//	HasDrawback int `gorm:"has_drawback" json:"has_drawback"` // 是否退税
//	DrawbackAddress string `gorm:"drawback_address" json:"drawback_address"` // 退税地址
//	FileDeliverAddress string `gorm:"file_deliver_address" json:"file_deliver_address"` // 报关资料送至地址
//	Remark string `gorm:"remark" json:"remark"` // 备注
//	OriginCountryId int `gorm:"origin_country_id" json:"origin_country_id"` // 原产地（国家）
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	SerialNo string `gorm:"serial_no" json:"serial_no"` // 序列号
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	Status string `gorm:"status" json:"status"`
//	InvoiceStatus string `gorm:"invoice_status" json:"invoice_status"`
//	InvoiceNo string `gorm:"invoice_no" json:"invoice_no"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"`
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"`
//	CompanyInstructionContent string `gorm:"company_instruction_content" json:"company_instruction_content"`
//	CustomBrokerContent string `gorm:"custom_broker_content" json:"custom_broker_content"`
//	DepartureCountryId int `gorm:"departure_country_id" json:"departure_country_id"`
//	DestinationAirportId int `gorm:"destination_airport_id" json:"destination_airport_id"`
//	WaysOfTransportationId int `gorm:"ways_of_transportation_id" json:"ways_of_transportation_id"`
//	DestinationCountryId int `gorm:"destination_country_id" json:"destination_country_id"`
//	IsGoodsAttachment int `gorm:"is_goods_attachment" json:"is_goods_attachment"` // 导出货描列表
//}
//
//func (*Former_customs_orders) TableName() string {
//	return "former_customs_orders"
//}
//
//type Schema_migrations struct {
//	Version string `gorm:"version" json:"version"`
//}
//
//func (*Schema_migrations) TableName() string {
//	return "schema_migrations"
//}
//
//type User_share_data struct {
//	Id int `gorm:"id" json:"id"`
//	UserId int `gorm:"user_id" json:"user_id"` // 所属员工
//	ShareId int `gorm:"share_id" json:"share_id"` // 被共享的id
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*User_share_data) TableName() string {
//	return "user_share_data"
//}
//
//type Website_other_sets struct {
//	Id int `gorm:"id" json:"id"`
//	CompanyName string `gorm:"company_name" json:"company_name"` // 公司名称
//	CompanyNameEn string `gorm:"company_name_en" json:"company_name_en"` // 公司英文名称
//	CompanyAddress string `gorm:"company_address" json:"company_address"` // 联系方式
//	CompanyAddressEn string `gorm:"company_address_en" json:"company_address_en"` // 英文联系方式
//	RecordNumber string `gorm:"record_number" json:"record_number"` // 备案号
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Website_other_sets) TableName() string {
//	return "website_other_sets"
//}
//
//type Bargain_platform_inquiries struct {
//	Id int `gorm:"id" json:"id"`
//	InquiryNo string `gorm:"inquiry_no" json:"inquiry_no"` // 询价编号
//	Status string `gorm:"status" json:"status"` // 状态
//	LastedQuotationTime string `gorm:"lasted_quotation_time" json:"lasted_quotation_time"` // 最近报价时间
//	SupplyId int `gorm:"supply_id" json:"supply_id"` // 供应商
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	Remarks string `gorm:"remarks" json:"remarks"` // 询价备注
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	BargainMainId int `gorm:"bargain_main_id" json:"bargain_main_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Bargain_platform_inquiries) TableName() string {
//	return "bargain_platform_inquiries"
//}
//
//type Base_data_code_of_twos struct {
//	Id int `gorm:"id" json:"id"`
//	Code string `gorm:"code" json:"code"` // 编码
//	Airport string `gorm:"airport" json:"airport"` // 机场
//	Name string `gorm:"name" json:"name"` // 机场名
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	AirportPya string `gorm:"airport_pya" json:"airport_pya"` // 机场拼音全写
//	AirportPyf string `gorm:"airport_pyf" json:"airport_pyf"` // 机场拼音简写
//	NamePya string `gorm:"name_pya" json:"name_pya"` // 机场名全写
//	NamePyf string `gorm:"name_pyf" json:"name_pyf"` // 机场名简写
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//	Website string `gorm:"website" json:"website"` // 网站
//}
//
//func (*Base_data_code_of_twos) TableName() string {
//	return "base_data_code_of_twos"
//}
//
//type Base_data_countries struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Name string `gorm:"name" json:"name"` // 国家中文名
//	NameEn string `gorm:"name_en" json:"name_en"` // 国家英文名
//	NamePya string `gorm:"name_pya" json:"name_pya"` // 国家中文名全拼
//	NamePyf string `gorm:"name_pyf" json:"name_pyf"`
//	Code string `gorm:"code" json:"code"` // 国家代码
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//	NameEnShort string `gorm:"name_en_short" json:"name_en_short"`
//	NameCnShort string `gorm:"name_cn_short" json:"name_cn_short"` // 中文简称
//}
//
//func (*Base_data_countries) TableName() string {
//	return "base_data_countries"
//}
//
//type Former_courier_orders struct {
//	Id int `gorm:"id" json:"id"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	SerialNo string `gorm:"serial_no" json:"serial_no"` // 序列号
//	BaseDataInstructionTypeId int `gorm:"base_data_instruction_type_id" json:"base_data_instruction_type_id"` // 委托公司类型
//	WaysOfDeclarationId int `gorm:"ways_of_declaration_id" json:"ways_of_declaration_id"` // 报关方式
//	BaseDataTradeTermsId int `gorm:"base_data_trade_terms_id" json:"base_data_trade_terms_id"` // 贸易条款
//	ServiceContractNo string `gorm:"service_contract_no" json:"service_contract_no"` // 合同编号
//	InvoiceNo string `gorm:"invoice_no" json:"invoice_no"` // 发票
//	CourierNumber string `gorm:"courier_number" json:"courier_number"` // 快递单号
//	ReceiverAddress string `gorm:"receiver_address" json:"receiver_address"` // 收货地址
//	Address string `gorm:"address" json:"address"` // 提货地址
//	Note string `gorm:"note" json:"note"` // 备注
//	Insurance int `gorm:"insurance" json:"insurance"` // 是否需要保险
//	InsuranceAmount float64 `gorm:"insurance_amount" json:"insurance_amount"` // 保险
//	Status string `gorm:"status" json:"status"` // 转态
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	CourierCodeId int `gorm:"courier_code_id" json:"courier_code_id"`
//	ConsigneeId int `gorm:"consignee_id" json:"consignee_id"`
//	ConsigneeType string `gorm:"consignee_type" json:"consignee_type"`
//	ConsigneeContent string `gorm:"consignee_content" json:"consignee_content"`
//	ShipperId int `gorm:"shipper_id" json:"shipper_id"`
//	ShipperType string `gorm:"shipper_type" json:"shipper_type"`
//	ShipperContent string `gorm:"shipper_content" json:"shipper_content"`
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"`
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"`
//	CompanyInstructionContent string `gorm:"company_instruction_content" json:"company_instruction_content"`
//	Marks string `gorm:"marks" json:"marks"`
//	Number string `gorm:"number" json:"number"`
//	DescriptionOfGood string `gorm:"description_of_good" json:"description_of_good"`
//	GrossWeight string `gorm:"gross_weight" json:"gross_weight"`
//	Size string `gorm:"size" json:"size"`
//	EstimatedTimeOfDeparture string `gorm:"estimated_time_of_departure" json:"estimated_time_of_departure"`
//	EstimatedTimeOfArrival string `gorm:"estimated_time_of_arrival" json:"estimated_time_of_arrival"`
//	CodeOfTwoId int `gorm:"code_of_two_id" json:"code_of_two_id"`
//	OceanChangesPaytypeId int `gorm:"ocean_changes_paytype_id" json:"ocean_changes_paytype_id"`
//	OtherChangesPaytypeId int `gorm:"other_changes_paytype_id" json:"other_changes_paytype_id"`
//	DestinationAirportId int `gorm:"destination_airport_id" json:"destination_airport_id"`
//	TransshipmentAirportId int `gorm:"transshipment_airport_id" json:"transshipment_airport_id"`
//	DepartureAirportId int `gorm:"departure_airport_id" json:"departure_airport_id"`
//	NotifyPartyId int `gorm:"notify_party_id" json:"notify_party_id"`
//	NotifyPartyType string `gorm:"notify_party_type" json:"notify_party_type"`
//	NotifyPartyContent string `gorm:"notify_party_content" json:"notify_party_content"`
//	TransshipmentCode string `gorm:"transshipment_code" json:"transshipment_code"`
//	TransshipmentDate string `gorm:"transshipment_date" json:"transshipment_date"`
//	BaseDataPackageType int `gorm:"base_data_package_type" json:"base_data_package_type"`
//	CodeList string `gorm:"code_list" json:"code_list"` // 快递可能存在多个小包，多个快递的集合
//	Ratio float64 `gorm:"ratio" json:"ratio"` // 材积换算系数
//	RatioWeight float64 `gorm:"ratio_weight" json:"ratio_weight"` // 材积重
//	Bubble float64 `gorm:"bubble" json:"bubble"` // 分泡%
//	Dimension string `gorm:"dimension" json:"dimension"` // 体积
//	ChargedWeight float64 `gorm:"charged_weight" json:"charged_weight"` // 计费重
//	IsGoodsAttachment int `gorm:"is_goods_attachment" json:"is_goods_attachment"` // 导出货描列表
//}
//
//func (*Former_courier_orders) TableName() string {
//	return "former_courier_orders"
//}
//
//type Freight_search_logs struct {
//	Id int `gorm:"id" json:"id"`
//	DestinationAddress string `gorm:"destination_address" json:"destination_address"` // 目的地
//	DepartureCity string `gorm:"departure_city" json:"departure_city"` // 城市
//	DepartureDistrict string `gorm:"departure_district" json:"departure_district"` // 地区
//	DepartureAddress string `gorm:"departure_address" json:"departure_address"` // 出发地
//	CityName string `gorm:"city_name" json:"city_name"` // 城市
//	DepartureSeaPortName string `gorm:"departure_sea_port_name" json:"departure_sea_port_name"` // 起运港
//	CityId string `gorm:"city_id" json:"city_id"` // 城市
//	DepartureSeaPortId string `gorm:"departure_sea_port_id" json:"departure_sea_port_id"` // 起运港
//	SeaLineId string `gorm:"sea_line_id" json:"sea_line_id"` // 航线
//	DestinationPortId string `gorm:"destination_port_id" json:"destination_port_id"` // 目的港
//	BoatCompanyId string `gorm:"boat_company_id" json:"boat_company_id"` // 船公司
//	QueryType string `gorm:"query_type" json:"query_type"` // 查询类型【拖车，拼箱，整柜】
//	IpAddress string `gorm:"ip_address" json:"ip_address"` // 客户登录ip
//	UserId int `gorm:"user_id" json:"user_id"` // 系统用户id,可为空
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	RawInfoJson string `gorm:"raw_info_json" json:"raw_info_json"`
//}
//
//func (*Freight_search_logs) TableName() string {
//	return "freight_search_logs"
//}
//
//type Plan_customers struct {
//	Id int `gorm:"id" json:"id"`
//	PlanMainId int `gorm:"plan_main_id" json:"plan_main_id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"` // 委托单位
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"` // 委托单位类型
//	OceanChangesPaytypeId int `gorm:"ocean_changes_paytype_id" json:"ocean_changes_paytype_id"` // 运费付款方式
//	OtherChangesPaytypeId int `gorm:"other_changes_paytype_id" json:"other_changes_paytype_id"` // 其他运费付款方式
//	TradeTermId int `gorm:"trade_term_id" json:"trade_term_id"` // 贸易条款
//	TermId int `gorm:"term_id" json:"term_id"` // 装运条款
//	BusinessTypeId int `gorm:"business_type_id" json:"business_type_id"` // 业务类型
//	BillProduceId int `gorm:"bill_produce_id" json:"bill_produce_id"` // 出单方式
//	HblNumber string `gorm:"hbl_number" json:"hbl_number"` // hbl编号
//	BoxSizeCount string `gorm:"box_size_count" json:"box_size_count"` // 柜型柜量
//	IsFile int `gorm:"is_file" json:"is_file"` // 是否随机文件
//	Enabled int `gorm:"enabled" json:"enabled"` // 有效状态
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Plan_customers) TableName() string {
//	return "plan_customers"
//}
//
//type Active_admin_comments struct {
//	Id int `gorm:"id" json:"id"`
//	Namespace string `gorm:"namespace" json:"namespace"`
//	Body string `gorm:"body" json:"body"`
//	ResourceType string `gorm:"resource_type" json:"resource_type"`
//	ResourceId int `gorm:"resource_id" json:"resource_id"`
//	AuthorType string `gorm:"author_type" json:"author_type"`
//	AuthorId int `gorm:"author_id" json:"author_id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//}
//
//func (*Active_admin_comments) TableName() string {
//	return "active_admin_comments"
//}
//
//type Active_storage_attachments struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"`
//	RecordType string `gorm:"record_type" json:"record_type"`
//	RecordId int `gorm:"record_id" json:"record_id"`
//	BlobId int `gorm:"blob_id" json:"blob_id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	Uploader int `gorm:"uploader" json:"uploader"` // 上传者
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//}
//
//func (*Active_storage_attachments) TableName() string {
//	return "active_storage_attachments"
//}
//
//type Base_data_cities struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"` // 城市名
//	ZoneId int `gorm:"zone_id" json:"zone_id"` // 区域ID
//	NameEn string `gorm:"name_en" json:"name_en"` // 城市英文名
//	NamePya string `gorm:"name_pya" json:"name_pya"` // 城市拼音全写
//	NamePyf string `gorm:"name_pyf" json:"name_pyf"` // 城市拼音简拼
//	HasFreight int `gorm:"has_freight" json:"has_freight"`
//	Position int `gorm:"position" json:"position"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//}
//
//func (*Base_data_cities) TableName() string {
//	return "base_data_cities"
//}
//
//type Finance_month_statements struct {
//	Id int `gorm:"id" json:"id"`
//	ClosingUnitId int `gorm:"closing_unit_id" json:"closing_unit_id"` // 结算单位,方便检索
//	Status string `gorm:"status" json:"status"` // 前台对账状态
//	ConfirmBy string `gorm:"confirm_by" json:"confirm_by"` // 前台确认还是后台确认
//	Remarks string `gorm:"remarks" json:"remarks"` // 备注
//	TemplateId int `gorm:"template_id" json:"template_id"` // 月结账单模板的id
//	FeeIdJson string `gorm:"fee_id_json" json:"fee_id_json"` // 費用的id
//	CustomStatementHtml string `gorm:"custom_statement_html" json:"custom_statement_html"` // 如果操作在模板的基础上修改生成的账单，保存该html
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//}
//
//func (*Finance_month_statements) TableName() string {
//	return "finance_month_statements"
//}
//
//type Former_seas_importor_securities struct {
//	Id int `gorm:"id" json:"id"`
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	SerialNo string `gorm:"serial_no" json:"serial_no"` // 表单序号
//	SellerId int `gorm:"seller_id" json:"seller_id"` // 卖方信息，从cooperator中选择,这个是需要关联的
//	BuyerId int `gorm:"buyer_id" json:"buyer_id"` // 买方信息，从cooperator中选择,这个需要关联
//	PartyId int `gorm:"party_id" json:"party_id"` // party信息，从cooperator中选择，与以上保持一致格式
//	ManufacturerId int `gorm:"manufacturer_id" json:"manufacturer_id"` // 制造商信息，从cooperator中选择
//	OriginCountryId int `gorm:"origin_country_id" json:"origin_country_id"` // 原产地国家
//	ImportorIrsNo string `gorm:"importor_irs_no" json:"importor_irs_no"` // 进口税号
//	ConsigneeIrsNo string `gorm:"consignee_irs_no" json:"consignee_irs_no"` // 收货人编号(联邦税务局数字)
//	ResponsiblePersonName string `gorm:"responsible_person_name" json:"responsible_person_name"` // 责任人名字
//	ResponsiblePersonPhone string `gorm:"responsible_person_phone" json:"responsible_person_phone"` // 责任人联系方式
//	HbLNo string `gorm:"hb_l_no" json:"hb_l_no"` // 货代提单
//	MbLNo string `gorm:"mb_l_no" json:"mb_l_no"` // 船东提单
//	CtrnSeal string `gorm:"ctrn_seal" json:"ctrn_seal"` // 柜封号
//	ScacCode string `gorm:"scac_code" json:"scac_code"` // 自申报号
//	AmsBlNo string `gorm:"ams_bl_no" json:"ams_bl_no"` // ams船东提单号
//	ContainerStuffingLocation string `gorm:"container_stuffing_location" json:"container_stuffing_location"` // 货物装箱位置
//	Consolidator int `gorm:"consolidator" json:"consolidator"` // 拼箱公司
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Status string `gorm:"status" json:"status"` // 状态
//	InvoiceNo string `gorm:"invoice_no" json:"invoice_no"` // 发票号
//	SellerContent string `gorm:"seller_content" json:"seller_content"`
//	BuyerContent string `gorm:"buyer_content" json:"buyer_content"`
//	PartyContent string `gorm:"party_content" json:"party_content"`
//	ManufacturerContent string `gorm:"manufacturer_content" json:"manufacturer_content"`
//	CompanyInstructionId int `gorm:"company_instruction_id" json:"company_instruction_id"`
//	CompanyInstructionType string `gorm:"company_instruction_type" json:"company_instruction_type"`
//	CompanyInstructionContent string `gorm:"company_instruction_content" json:"company_instruction_content"`
//	IsGoodsAttachment int `gorm:"is_goods_attachment" json:"is_goods_attachment"` // 导出货描列表
//	ConsolidatorContent string `gorm:"consolidator_content" json:"consolidator_content"` // 拼箱公司名称和地址
//	TariffSchedule string `gorm:"tariff_schedule" json:"tariff_schedule"` // HS商品编号（前六位）
//	SascCode string `gorm:"sasc_code" json:"sasc_code"` // SASC Code
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Former_seas_importor_securities) TableName() string {
//	return "former_seas_importor_securities"
//}
//
//type Website_articles struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Title string `gorm:"title" json:"title"` // 文章标题
//	Type string `gorm:"type" json:"type"` // 分类
//	Keyword string `gorm:"keyword" json:"keyword"` // 关键字
//	Details string `gorm:"details" json:"details"` // 详情内容
//	Enabled int `gorm:"enabled" json:"enabled"` // 是否可见
//	Locale string `gorm:"locale" json:"locale"` // 语言类型
//	ReadCount int `gorm:"read_count" json:"read_count"` // 已读次数
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Website_articles) TableName() string {
//	return "website_articles"
//}
//
//type Approval_applications struct {
//	Id int `gorm:"id" json:"id"`
//	OperatorId int `gorm:"operator_id" json:"operator_id"` // 申请人id
//	OperatorName string `gorm:"operator_name" json:"operator_name"` // 申请人姓名
//	Note string `gorm:"note" json:"note"` // 备注
//	Status string `gorm:"status" json:"status"` // 状态
//	ApprovalTimes int `gorm:"approval_times" json:"approval_times"` // 当前审核次数
//	Number string `gorm:"number" json:"number"` // 申请单号
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	PayOrReceive string `gorm:"pay_or_receive" json:"pay_or_receive"`
//	AuditorId int `gorm:"auditor_id" json:"auditor_id"`
//	AuditorName string `gorm:"auditor_name" json:"auditor_name"`
//	BatchNumber string `gorm:"batch_number" json:"batch_number"` // 批量申请单号
//	OrderMasterId int `gorm:"order_master_id" json:"order_master_id"`
//	ClosingUnitId int `gorm:"closing_unit_id" json:"closing_unit_id"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Approval_applications) TableName() string {
//	return "approval_applications"
//}
//
//type Base_data_departure_sea_ports struct {
//	Id int `gorm:"id" json:"id"`
//	City string `gorm:"city" json:"city"` // 起运港城市
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Port string `gorm:"port" json:"port"` // 起运港
//	Position int `gorm:"position" json:"position"` // 位置
//	CityId string `gorm:"city_id" json:"city_id"` // 城市ID
//	PortEn string `gorm:"port_en" json:"port_en"` // 港口英文名
//	PortPya string `gorm:"port_pya" json:"port_pya"` // 港口拼音全写
//	PortPyf string `gorm:"port_pyf" json:"port_pyf"` // 港口拼音简写
//	CityEn string `gorm:"city_en" json:"city_en"` // 城市英文名
//	CityPya string `gorm:"city_pya" json:"city_pya"` // 城市拼音全写
//	CityPyf string `gorm:"city_pyf" json:"city_pyf"` // 城市拼音简写
//	HasFreight int `gorm:"has_freight" json:"has_freight"` // 附加费
//	Longitude string `gorm:"longitude" json:"longitude"` // 经度
//	Latitude string `gorm:"latitude" json:"latitude"` // 纬度
//	LocationAddress string `gorm:"location_address" json:"location_address"` // 定位地址
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//}
//
//func (*Base_data_departure_sea_ports) TableName() string {
//	return "base_data_departure_sea_ports"
//}
//
//type Base_data_supplier_business_types struct {
//	Id int `gorm:"id" json:"id"`
//	Name string `gorm:"name" json:"name"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Base_data_supplier_business_types) TableName() string {
//	return "base_data_supplier_business_types"
//}
//
//type Base_data_voyages struct {
//	Id int `gorm:"id" json:"id"`
//	BaseDataBoatCompanyId int `gorm:"base_data_boat_company_id" json:"base_data_boat_company_id"`
//	Name string `gorm:"name" json:"name"`
//	Enabled int `gorm:"enabled" json:"enabled"` // 是否有效
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Base_data_voyages) TableName() string {
//	return "base_data_voyages"
//}
//
//type Fee_verifications struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	FinanceFeesId int `gorm:"finance_fees_id" json:"finance_fees_id"`
//	VerificationsId int `gorm:"verifications_id" json:"verifications_id"`
//	AuditAmount float64 `gorm:"audit_amount" json:"audit_amount"` // 核销金额
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Fee_verifications) TableName() string {
//	return "fee_verifications"
//}
//
//type Finance_currency_histories struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	FinanceCurrencyId int `gorm:"finance_currency_id" json:"finance_currency_id"`
//	RateRealtime float64 `gorm:"rate_realtime" json:"rate_realtime"` // 实时汇率
//	RateFix float64 `gorm:"rate_fix" json:"rate_fix"` // 固定汇率
//	RateResult float64 `gorm:"rate_result" json:"rate_result"` // 实际汇率
//	ValidTime string `gorm:"valid_time" json:"valid_time"` // 汇率生效时间
//	InvalidTime string `gorm:"invalid_time" json:"invalid_time"` // 汇率失效时间
//	UserId int `gorm:"user_id" json:"user_id"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//}
//
//func (*Finance_currency_histories) TableName() string {
//	return "finance_currency_histories"
//}
//
//type Message_contants struct {
//	Id int `gorm:"id" json:"id"`
//	Category string `gorm:"category" json:"category"` // 类型(group群组、personal个人、copy抄送)
//	MailboxId int `gorm:"mailbox_id" json:"mailbox_id"` // 对应信息id
//	UserId int `gorm:"user_id" json:"user_id"` // 接收者id(group_id、person_id)
//	UserName string `gorm:"user_name" json:"user_name"` // 接收者姓名
//	CompanyId int `gorm:"company_id" json:"company_id"` // 接收者公司id
//	CompanyName string `gorm:"company_name" json:"company_name"` // 接收者公司名称
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Email string `gorm:"email" json:"email"`
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	Mark string `gorm:"mark" json:"mark"`
//	LockVersion int `gorm:"lock_version" json:"lock_version"`
//}
//
//func (*Message_contants) TableName() string {
//	return "message_contants"
//}
//
//type Setting_numbers struct {
//	Id int `gorm:"id" json:"id"`
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	Prefix string `gorm:"prefix" json:"prefix"` // 前缀
//	Special string `gorm:"special" json:"special"` // 特殊字符
//	NumberLength int `gorm:"number_length" json:"number_length"` // 流水号长度
//	YearRule string `gorm:"year_rule" json:"year_rule"` // 年设置
//	MonthRule string `gorm:"month_rule" json:"month_rule"` // 月设置
//	DayRule string `gorm:"day_rule" json:"day_rule"` // 日设置
//	UserNumber string `gorm:"user_number" json:"user_number"` // 用户工号
//	ApplicationNo string `gorm:"application_no" json:"application_no"` // 应用
//	DefaultNumber int `gorm:"default_number" json:"default_number"` // 默认的流水号长度
//	ClearMethod string `gorm:"clear_method" json:"clear_method"` // 清空方式
//	CurrentNumber string `gorm:"current_number" json:"current_number"` // 当前流水号
//	Rule string `gorm:"rule" json:"rule"` // 规则
//	OldRule string `gorm:"old_rule" json:"old_rule"` // 旧规则
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsChange int `gorm:"is_change" json:"is_change"` // 是否修改
//	NumberRule string `gorm:"number_rule" json:"number_rule"`
//	UserCompanyId int `gorm:"user_company_id" json:"user_company_id"`
//}
//
//func (*Setting_numbers) TableName() string {
//	return "setting_numbers"
//}
//
//type Base_data_destinations struct {
//	Id int `gorm:"id" json:"id"`
//	City string `gorm:"city" json:"city"` // 城市
//	Address string `gorm:"address" json:"address"` // 地址
//	CreatedAt string `gorm:"created_at" json:"created_at"`
//	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
//	CityPya string `gorm:"city_pya" json:"city_pya"` // 城市拼音全写
//	CityPyf string `gorm:"city_pyf" json:"city_pyf"` // 城市拼音简写
//	AddressPya string `gorm:"address_pya" json:"address_pya"` // 地址拼音全写
//	AddressPyf string `gorm:"address_pyf" json:"address_pyf"` // 地址拼音简写
//	DeletedAt string `gorm:"deleted_at" json:"deleted_at"`
//	IsLocalChanged int `gorm:"is_local_changed" json:"is_local_changed"` // 本地修改
//}
//
//func (*Base_data_destinations) TableName() string {
//	return "base_data_destinations"
//}
//
