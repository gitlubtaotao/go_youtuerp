package repositories

import (
	"github.com/jinzhu/gorm"
	"youtuerp/database"
	"youtuerp/models"
)

type IAddressRepository interface {
	Delete(id uint) error
	UpdateById(id uint, updateContent models.Address) (models.Address, error)
	FindByOa(per, page uint, filter map[string]interface{}, selectKeys []string, order []string) (accounts []models.Address,
		total uint, err error)
	FindByCrm(per, page uint, filter map[string]interface{}, selectKeys []string, orders []string) (accounts []models.Address,
		total uint, err error)
	Create(account models.Address) (models.Address, error)
	First(id uint) (models.Address, error)
}

type AddressRepository struct {
	BaseRepository
}

func (a AddressRepository) Delete(id uint) error {
	return a.crud.Delete(&models.Address{}, id)
}

func (a AddressRepository) UpdateById(id uint, updateContent models.Address) (models.Address, error) {
	address, err := a.First(id)
	if err != nil {
		return address, err
	}
	err = database.GetDBCon().Model(&address).Update(updateContent).Error
	return address, err
}

func (a AddressRepository) First(id uint) (models.Address, error) {
	var data models.Address
	err := database.GetDBCon().First(&data, id).Error
	return data, err
}

func (a AddressRepository) FindByOa(per, page uint, filter map[string]interface{}, selectKeys []string,
	order []string) (invoices []models.Address, total uint, err error) {
	sqlCon := database.GetDBCon().Model(&models.Address{})
	sqlCon = sqlCon.Scopes(a.defaultOaScoped)
	return a.Find(sqlCon, per, page, filter, selectKeys, order, true)
}
func (a AddressRepository) FindByCrm(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string) (accounts []models.Address, total uint, err error) {
	sqlCon := database.GetDBCon().Model(&models.Address{})
	sqlCon = sqlCon.Scopes(a.defaultCrmScoped)
	return a.Find(sqlCon, per, page, filter, selectKeys, orders, true)
}

//创建银行账户信息
func (a AddressRepository) Create(address models.Address) (models.Address, error) {
	err := database.GetDBCon().Set("gorm:association_autocreate", false).Create(&address).Error
	if err != nil {
		return models.Address{}, err
	}
	return address, err
}

func (a AddressRepository) Find(sqlCon *gorm.DB, per, page uint, filter map[string]interface{}, selectKeys []string,
	order []string, isCount bool) (address []models.Address,
	total uint, err error) {
	if len(filter) > 0 {
		sqlCon = sqlCon.Scopes(a.Ransack(filter))
	}
	if isCount {
		err = sqlCon.Count(&total).Error
		if err != nil {
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = []string{"address.*"}
	}
	rows, err := sqlCon.Scopes(a.Paginate(per, page), a.OrderBy(order)).Select(selectKeys).Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var data models.Address
		_ = sqlCon.ScanRows(rows, &data)
		address = append(address, data)
	}
	return
}

func (a AddressRepository) defaultOaScoped(db *gorm.DB) *gorm.DB {
	db = db.Joins("inner join user_companies on user_companies.id = address.user_company_id and user_companies.company_type = 4")
	return db
}

func (a AddressRepository) defaultCrmScoped(db *gorm.DB) *gorm.DB {
	db = db.Joins("inner join user_companies on user_companies.id = address.user_company_id and user_companies.company_type in (?) ", []int{1, 2, 3})
	return db
}

func NewAddress() IAddressRepository {
	return &AddressRepository{}
}
