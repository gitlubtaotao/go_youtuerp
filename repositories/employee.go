package repositories

import (
	"github.com/jinzhu/gorm"
	"youtuerp/database"
	"youtuerp/models"
)

type IEmployeeRepository interface {
	//通过员工的昵称或者用户信息进行查询
	FirstByNameOrEmail(account string) (employee *models.Employee, err error)
}
type EmployeeRepository struct {
	BaseRepository
}

func (e *EmployeeRepository) FirstByNameOrEmail(account string) (employee *models.Employee, err error) {
	var user models.Employee
	err = database.GetDBCon().Scopes(e.defaultScoped).Where("name = ?", account).Or("email = ?", account).First(&user).Error
	employee = &user
	return
}

func (e *EmployeeRepository) defaultScoped(db *gorm.DB) *gorm.DB {
	return db.Joins("inner join user_companies on user_companies.id = users.user_company_id and user_companies.company_type = 4")
}

func NewEmployeeRepository() IEmployeeRepository {
	return &EmployeeRepository{}
}
