package dao

import (
	"gorm.io/gorm"
	"youtuerp/global"
	"youtuerp/internal/models"
)

type IAddressRepository interface {
	Delete(id uint) error
	UpdateById(id uint, updateContent models.Address) (models.Address, error)
	FindByOa(per, page int, filter map[string]interface{}, selectKeys []string, order []string) (accounts []models.Address,
		total int64, err error)
	FindByCrm(per, page int, filter map[string]interface{}, selectKeys []string, orders []string) (accounts []models.Address,
		total int64, err error)
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
	err = global.DataEngine.Model(&address).Updates(updateContent).Error
	return address, err
}

func (a AddressRepository) First(id uint) (models.Address, error) {
	var data models.Address
	err := global.DataEngine.First(&data, id).Error
	return data, err
}

func (a AddressRepository) FindByOa(per, page int, filter map[string]interface{}, selectKeys []string,
	order []string) (address []models.Address, total int64, err error) {
	sqlCon := global.DataEngine.Model(&models.Address{}).Scopes(a.defaultOaScoped)
	address, err = a.Find(sqlCon, per, page, filter, selectKeys, order)
	total, err = a.Count(global.DataEngine.Model(&models.Address{}).Scopes(a.defaultOaScoped), filter)
	return
}

func (a AddressRepository) FindByCrm(per, page int, filter map[string]interface{}, selectKeys []string,
	orders []string) (address []models.Address, total int64, err error) {
	sqlCon := global.DataEngine.Model(&models.Address{}).Scopes(a.defaultCrmScoped)
	address, err = a.Find(sqlCon, per, page, filter, selectKeys, orders)
	total, err = a.Count(global.DataEngine.Model(&models.Address{}).Scopes(a.defaultCrmScoped), filter)
	return
}

//创建银行账户信息
func (a AddressRepository) Create(address models.Address) (models.Address, error) {
	err := global.DataEngine.Set("gorm:association_autocreate", false).Create(&address).Error
	if err != nil {
		return models.Address{}, err
	}
	return address, err
}

func (a AddressRepository) Find(sqlCon *gorm.DB, per, page int, filter map[string]interface{}, selectKeys []string,
	order []string) (address []models.Address, err error) {
	if len(selectKeys) == 0 {
		selectKeys = []string{"address.*"}
	}
	rows, err := sqlCon.Scopes(a.CustomerWhere(filter, selectKeys, a.Paginate(per, page), a.OrderBy(order))).Rows()
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
