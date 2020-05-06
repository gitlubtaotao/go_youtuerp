package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"youtuerp/database"
	"youtuerp/models"
)

type IEmployeeRepository interface {
	FirstByPhoneAndEmail(phone string, email string) (employee *models.Employee, err error)
	//通过员工的昵称或者用户信息进行查询
	First(id uint) (*models.Employee, error)
	FirstByPhoneOrEmail(account string) (employee *models.Employee, err error)
	UpdateColumnByID(employeeID uint, updateColumn map[string]interface{}) error
	UpdateColumn(employee *models.Employee, updateColumn map[string]interface{}) error
	UpdateRecordByModel(employee *models.Employee, updateModel models.Employee) error
	Find(per, page uint, filter map[string]interface{}, selectKeys []string, order []string, isCount bool) (employees []models.ResultEmployee,
		total uint, err error)
	Create(employee models.Employee) (models.Employee, error)
	Delete(id uint) error
}
type EmployeeRepository struct {
	BaseRepository
}

func (e *EmployeeRepository) Delete(id uint) error {
	var readData models.Employee
	return database.GetDBCon().Find(&readData).Delete(&readData).Error
}

func (e *EmployeeRepository) First(id uint) (*models.Employee, error) {
	var data models.Employee
	err := database.GetDBCon().First(&data, "id = ?", id).Error
	return &data, err
}

func (e *EmployeeRepository) Create(employee models.Employee) (models.Employee, error) {
	err := database.GetDBCon().Set("gorm:association_autocreate", false).Create(&employee).Error
	if err != nil {
		return models.Employee{}, err
	}
	return employee, err
}

func (e *EmployeeRepository) Find(per, page uint, filter map[string]interface{},
	selectKeys []string, order []string, isCount bool) (
	employees []models.ResultEmployee, total uint, err error) {
	sqlCon := database.GetDBCon().Model(&models.Employee{})
	sqlCon = sqlCon.Scopes(e.defaultScoped)
	if len(filter) > 0 {
		sqlCon = sqlCon.Scopes(e.Ransack(filter))
	}
	fmt.Println(filter)
	if isCount {
		err = sqlCon.Count(&total).Error
		if err != nil {
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = []string{"users.*", "user_companies.name_nick as user_companies_name_nick",
			"departments.name_cn as departments_name_cn",
		}
	}
	rows, err := sqlCon.Scopes(e.Paginate(per, page), e.OrderBy(order)).Select(selectKeys).Rows()
	if err != nil {
		fmt.Printf("%v", err)
	}
	
	for rows.Next() {
		var data models.ResultEmployee
		_ = sqlCon.ScanRows(rows, &data)
		employees = append(employees, data)
	}
	return
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
	err = database.GetDBCon().Scopes(e.defaultScoped).Where("users.phone = ?", account).Or("users.email = ?", account).First(&user).Error
	employee = &user
	return
}

func (e *EmployeeRepository) defaultScoped(db *gorm.DB) *gorm.DB {
	db = db.Joins("inner join user_companies on user_companies.id = users.user_company_id and user_companies.company_type = 4")
	db = db.Joins("left join departments on departments.id = users.department_id")
	return db
}

func NewEmployeeRepository() IEmployeeRepository {
	return &EmployeeRepository{}
}
