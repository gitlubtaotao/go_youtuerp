package repositories

import (
	"github.com/jinzhu/gorm"
	"youtuerp/database"
	"youtuerp/models"
)

type IEmployeeRepository interface {
	FirstByPhoneAndEmail(phone string, email string) (employee *models.Employee, err error)
	//通过员工的昵称或者用户信息进行查询
	FirstByPhoneOrEmail(account string) (employee *models.Employee, err error)
	UpdateColumnByID(employeeID uint, updateColumn map[string]interface{}) error
	UpdateColumn(employee *models.Employee, updateColumn map[string]interface{}) error
	UpdateRecordByModel(employee *models.Employee, updateModel models.Employee) error
}
type EmployeeRepository struct {
	BaseRepository
}

func (e *EmployeeRepository) UpdateRecordByModel(employee *models.Employee, updateModel models.Employee) error {
	return database.GetDBCon().Model(&employee).Update(updateModel).Error
}

//通过手机号码和邮箱查询当前用户
func (e *EmployeeRepository) FirstByPhoneAndEmail(phone string, email string) (employee *models.Employee, err error) {
	var user models.Employee
	err = database.GetDBCon().Where(&models.Employee{Phone: phone, Email: email}).First(&user).Error
	employee = &user
	return
}

func (e *EmployeeRepository) UpdateColumnByID(employeeID uint, updateColumn map[string]interface{}) error {
	user := models.Employee{ID: employeeID}
	return database.GetDBCon().Model(&user).Updates(updateColumn).Error
}

func (e *EmployeeRepository) UpdateColumn(employee *models.Employee, updateColumn map[string]interface{}) error {
	return database.GetDBCon().Model(&employee).Updates(updateColumn).Error
}

func (e *EmployeeRepository) FirstByPhoneOrEmail(account string) (employee *models.Employee, err error) {
	var user models.Employee
	err = database.GetDBCon().Scopes(e.defaultScoped).Where("phone = ?", account).Or("email = ?", account).First(&user).Error
	employee = &user
	return
}

func (e *EmployeeRepository) defaultScoped(db *gorm.DB) *gorm.DB {
	return db.Joins("inner join user_companies on user_companies.id = users.user_company_id and user_companies.company_type = 4")
}

func NewEmployeeRepository() IEmployeeRepository {
	return &EmployeeRepository{}
}
