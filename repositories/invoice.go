package repositories

import (
	"github.com/jinzhu/gorm"
	"youtuerp/database"
	"youtuerp/models"
)

type IInvoiceRepository interface {
	Delete(id uint) error
	UpdateById(id uint, updateContent models.Invoice) (models.Invoice, error)
	FindByOa(per, page uint, filter map[string]interface{}, selectKeys []string, order []string) (accounts []models.Invoice,
		total uint, err error)
	FindByCrm(per, page uint, filter map[string]interface{}, selectKeys []string, orders []string) (accounts []models.Invoice,
		total uint, err error)
	Create(account models.Invoice) (models.Invoice, error)
	First(id uint) (models.Invoice, error)
}

type InvoiceRepository struct {
	BaseRepository
}

func (i InvoiceRepository) Delete(id uint) error {
	return i.crud.Delete(&models.Invoice{}, id)
}

func (i InvoiceRepository) UpdateById(id uint, updateContent models.Invoice) (models.Invoice, error) {
	invoice, err := i.First(id)
	if err != nil {
		return invoice, err
	}
	err = database.GetDBCon().Model(&invoice).Update(updateContent).Error
	return invoice, err
}

func (i InvoiceRepository) First(id uint) (models.Invoice, error) {
	var data models.Invoice
	err := database.GetDBCon().First(&data, id).Error
	return data, err
}

func (i InvoiceRepository) FindByOa(per, page uint, filter map[string]interface{}, selectKeys []string,
	order []string) (invoices []models.Invoice, total uint, err error) {
	sqlCon := database.GetDBCon().Model(&models.Invoice{})
	sqlCon = sqlCon.Scopes(i.defaultOaScoped)
	return i.Find(sqlCon, per, page, filter, selectKeys, order, true)
}
func (i InvoiceRepository) FindByCrm(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string) (accounts []models.Invoice, total uint, err error) {
	sqlCon := database.GetDBCon().Model(&models.Invoice{})
	sqlCon = sqlCon.Scopes(i.defaultCrmScoped)
	return i.Find(sqlCon, per, page, filter, selectKeys, orders, true)
}

//创建银行账户信息
func (i InvoiceRepository) Create(invoice models.Invoice) (models.Invoice, error) {
	err := database.GetDBCon().Set("gorm:association_autocreate", false).Create(&invoice).Error
	if err != nil {
		return models.Invoice{}, err
	}
	return invoice, err
}

func (i InvoiceRepository) Find(sqlCon *gorm.DB, per, page uint, filter map[string]interface{}, selectKeys []string,
	order []string, isCount bool) (invoices []models.Invoice,
	total uint, err error) {
	if len(filter) > 0 {
		sqlCon = sqlCon.Scopes(i.Ransack(filter))
	}
	if isCount {
		err = sqlCon.Count(&total).Error
		if err != nil {
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = []string{"invoices.*"}
	}
	rows, err := sqlCon.Scopes(i.Paginate(per, page), i.OrderBy(order)).Select(selectKeys).Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var data models.Invoice
		_ = sqlCon.ScanRows(rows, &data)
		invoices = append(invoices, data)
	}
	return
}

func (i InvoiceRepository) defaultOaScoped(db *gorm.DB) *gorm.DB {
	db = db.Joins("inner join user_companies on user_companies.id = invoices.user_company_id and user_companies.company_type = 4")
	return db
}

func (a InvoiceRepository) defaultCrmScoped(db *gorm.DB) *gorm.DB {
	db = db.Joins("inner join user_companies on user_companies.id = invoices.user_company_id and user_companies.company_type in (?) ", []int{1, 2, 3})
	return db
}

func NewInvoice() IInvoiceRepository {
	return &InvoiceRepository{}
}
