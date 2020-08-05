package dao

import (
	"gorm.io/gorm"
	"youtuerp/global"
	"youtuerp/internal/models"
)

type IAccountRepository interface {
	Delete(id uint) error
	UpdateById(id uint, updateContent models.Account) (models.Account, error)
	FindByOa(per, page int, filter map[string]interface{}, selectKeys []string, order []string) (accounts []models.Account,
		total int64, err error)
	FindByCrm(per, page int, filter map[string]interface{}, selectKeys []string, orders []string) (accounts []models.Account,
		total int64, err error)
	Create(account models.Account) (models.Account, error)
	First(id uint) (models.Account, error)
}

type AccountRepository struct {
	BaseRepository
}

func (a AccountRepository) Delete(id uint) error {
	account, err := a.First(id)
	if err != nil {
		return err
	}
	return global.DataEngine.Delete(&account).Error
}

func (a AccountRepository) UpdateById(id uint, updateContent models.Account) (models.Account, error) {
	account, err := a.First(id)
	if err != nil {
		return account, err
	}
	err = global.DataEngine.Model(&account).Updates(updateContent).Error
	return account, err
}

func (a AccountRepository) First(id uint) (models.Account, error) {
	var data models.Account
	err := global.DataEngine.First(&data, id).Error
	return data, err
}

func (a AccountRepository) FindByOa(per, page int, filter map[string]interface{}, selectKeys []string,
	order []string) (accounts []models.Account, total int64, err error) {
	sqlCon := global.DataEngine.Model(&models.Account{})
	sqlCon = sqlCon.Scopes(a.defaultOaScoped)
	accounts, err = a.FindRecord(sqlCon, per, page, filter, selectKeys, order)
	if err != nil {
		return
	}
	countCon := global.DataEngine.Model(&models.Account{}).Scopes(a.defaultOaScoped)
	total, err = a.Count(countCon, filter)
	return
}

func (a AccountRepository) FindByCrm(per, page int, filter map[string]interface{}, selectKeys []string,
	orders []string) (accounts []models.Account, total int64, err error) {
	sqlCon := global.DataEngine.Model(&models.Account{})
	sqlCon = sqlCon.Scopes(a.defaultCrmScoped)
	accounts, err = a.FindRecord(sqlCon, per, page, filter, selectKeys, orders)
	if err != nil {
		return
	}
	countCon := global.DataEngine.Model(&models.Account{}).Scopes(a.defaultCrmScoped)
	total, err = a.Count(countCon, filter)
	return
}

//创建银行账户信息
func (a AccountRepository) Create(account models.Account) (models.Account, error) {
	err := global.DataEngine.Set("gorm:association_autocreate", false).Create(&account).Error
	if err != nil {
		return models.Account{}, err
	}
	return account, err
}

func (a AccountRepository) FindRecord(sqlCon *gorm.DB, per, page int, filter map[string]interface{}, selectKeys []string,
	order []string) (accounts []models.Account, err error) {
	if len(selectKeys) == 0 {
		selectKeys = a.selectKeys()
	}
	rows, err := sqlCon.Scopes(a.CustomerWhere(filter, selectKeys, a.Paginate(per, page), a.OrderBy(order))).Where("accounts.deleted_at IS NULL").Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var data models.Account
		_ = sqlCon.ScanRows(rows, &data)
		accounts = append(accounts, data)
	}
	return
}

func (a AccountRepository) selectKeys() []string {
	return []string{"accounts.*", "user_companies.name_nick as user_companies_name_nick"}
}

func (a AccountRepository) defaultOaScoped(db *gorm.DB) *gorm.DB {
	db = db.Joins("inner join user_companies on user_companies.id = accounts.user_company_id and user_companies.company_type = 4")
	return db
}

func (a AccountRepository) defaultCrmScoped(db *gorm.DB) *gorm.DB {
	db = db.Joins("inner join user_companies on user_companies.id = accounts.user_company_id and user_companies.company_type in (?) ", []int{1, 2, 3})
	return db
}

func NewAccountRepository() IAccountRepository {
	return &AccountRepository{}
}
