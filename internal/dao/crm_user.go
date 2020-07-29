package dao

import (
	"gorm.io/gorm"
	"youtuerp/database"
	"youtuerp/internal/models"
	"youtuerp/pkg/util"
)

type ICrmUser interface {
	Delete(id uint) error
	Update(id uint, user models.CrmContact) error
	Create(user models.CrmContact) (models.CrmContact, error)
	Find(per, page int, filter map[string]interface{}, selectKeys []string,
		order []string, isTotal bool) ([]models.CrmContact, int64, error)
}
type CrmUser struct {
	BaseRepository
}

func (c CrmUser) Delete(id uint) error {
	return c.crud.Delete(&models.CrmContact{}, id)
}
func (c CrmUser) Update(id uint, user models.CrmContact) error {
	return database.GetDBCon().Model(&models.CrmContact{ID: id}).Updates(util.StructToChange(user)).Error
}

func (c CrmUser) Find(per, page int, filter map[string]interface{}, selectKeys []string,
	order []string, isTotal bool) ([]models.CrmContact, int64, error) {
	var (
		users []models.CrmContact
		total int64
		err   error
	)
	return users, total, err
}

func (c CrmUser) Create(user models.CrmContact) (models.CrmContact, error) {
	err := database.GetDBCon().Create(&user).Error
	if err != nil {
		return models.CrmContact{}, err
	}
	return user, err
}

func (c CrmUser) defaultScope(db *gorm.DB) *gorm.DB {
	return db.Where("company_type in (?)", []int{1, 2, 3})
}

func NewCrmUser() ICrmUser {
	return &CrmUser{}
}
