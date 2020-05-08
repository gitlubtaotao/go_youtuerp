package repositories

import (
	"github.com/jinzhu/gorm"
	"youtuerp/database"
	"youtuerp/models"
)

type IAccountRepository interface {
	Delete(id uint) error
	UpdateById(id uint, updateContent models.Account) (models.Account, error)
	FindByOa(per, page uint, filter map[string]interface{}, selectKeys []string, order []string) (accounts []models.ResultAccount,
		total uint, err error)
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
	return database.GetDBCon().Delete(&account).Error
}

func (a AccountRepository) UpdateById(id uint, updateContent models.Account) (models.Account, error) {
	account, err := a.First(id)
	if err != nil {
		return account, err
	}
	err = database.GetDBCon().Model(&account).Update(updateContent).Error
	return account, err
}

func (a AccountRepository) First(id uint) (models.Account, error) {
	var data models.Account
	err := database.GetDBCon().First(&data, id).Error
	return data, err
}

func (a AccountRepository) FindByOa(per, page uint, filter map[string]interface{}, selectKeys []string,
	order []string) (accounts []models.ResultAccount, total uint, err error) {
	sqlCon := database.GetDBCon().Model(&models.Account{})
	sqlCon = sqlCon.Scopes(a.defaultOaScoped)
	return a.find(sqlCon, per, page, filter, selectKeys, order, true)
}

//创建银行账户信息
func (a AccountRepository) Create(account models.Account) (models.Account, error) {
	err := database.GetDBCon().Set("gorm:association_autocreate", false).Create(&account).Error
	if err != nil {
		return models.Account{}, err
	}
	return account, err
}

func (a AccountRepository) find(sqlCon *gorm.DB, per, page uint, filter map[string]interface{}, selectKeys []string,
	order []string, isCount bool) (accounts []models.ResultAccount,
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
		selectKeys = a.selectKeys()
	}
	rows, err := sqlCon.Scopes(a.Paginate(per, page), a.OrderBy(order)).Select(selectKeys).Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var data models.ResultAccount
		_ = sqlCon.ScanRows(rows, &data)
		accounts = append(accounts, data)
	}
	return
}

func (a AccountRepository) selectKeys() []string {
	return []string{"accounts.*", "user_companies.name_nick as user_companies_name_nick",}
}

func (a AccountRepository) defaultOaScoped(db *gorm.DB) *gorm.DB {
	db = db.Joins("inner join user_companies on user_companies.id = accounts.user_company_id and user_companies.company_type = 4")
	return db
}

func NewAccountRepository() IAccountRepository {
	return &AccountRepository{}
}
