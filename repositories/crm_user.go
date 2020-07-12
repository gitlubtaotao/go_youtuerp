package repositories

import (
	"gorm.io/gorm"
	"youtuerp/database"
	"youtuerp/models"
	"youtuerp/tools"
)

type ICrmUser interface {
	Delete(id uint) error
	Update(id uint, user models.CrmUser) error
	Create(user models.CrmUser) (models.CrmUser, error)
	Find(per, page int, filter map[string]interface{}, selectKeys []string,
		order []string, isTotal bool) ([]models.CrmUser, int64, error)
}
type CrmUser struct {
	BaseRepository
}

func (c CrmUser) Delete(id uint) error {
	return c.crud.Delete(&models.CrmUser{}, id)
}
func (c CrmUser) Update(id uint, user models.CrmUser) error {
	return database.GetDBCon().Model(&models.CrmUser{ID: id}).Updates(tools.StructToChange(user)).Error
}

func (c CrmUser) Find(per, page int, filter map[string]interface{}, selectKeys []string,
	order []string, isTotal bool) (users []models.CrmUser, total int64, err error) {
	sqlConn := database.GetDBCon().Model(&models.CrmUser{}).Scopes(c.defaultScope)
	if isTotal {
		if total, err = c.Count(sqlConn, filter); err != nil {
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = []string{"users.*"}
	}
	sqlConn = c.crud.Where(sqlConn, filter, selectKeys, c.Paginate(per, page),c.OrderBy(order))
	err = sqlConn.Find(&users).Error
	return users, total, err
}

func (c CrmUser) Create(user models.CrmUser) (models.CrmUser, error) {
	err := database.GetDBCon().Create(&user).Error
	if err != nil {
		return models.CrmUser{}, err
	}
	return user, err
}

func (c CrmUser) defaultScope(db *gorm.DB) *gorm.DB {
	return db.Where("company_type in (?)", []int{1, 2, 3})
}

func NewCrmUser() ICrmUser {
	return &CrmUser{}
}
