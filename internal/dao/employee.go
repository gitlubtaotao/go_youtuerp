package dao

import (
	"gorm.io/gorm"
	"youtuerp/global"
	"youtuerp/internal/models"
)

type IEmployeeRepository interface {
	FirstByPhoneAndEmail(phone string, email string) (employee *models.Employee, err error)
	//通过员工的昵称或者用户信息进行查询
	First(id uint) (*models.Employee, error)
	FirstByPhoneOrEmail(account string) (employee *models.Employee, err error)
	UpdateColumnByID(employeeID uint, updateColumn map[string]interface{}) error
	UpdateColumn(employee *models.Employee, updateColumn map[string]interface{}) error
	UpdateRecordByModel(userId uint, updateModel models.Employee) error
	Find(per, page int, filter map[string]interface{}, selectKeys []string, order []string, isCount bool) (
		employees []models.ResponseEmployee,
		total int64, err error)
	Create(employee models.Employee) (models.Employee, error)
	Delete(id uint) error
}
type EmployeeRepository struct {
	BaseRepository
}

func (e EmployeeRepository) Delete(id uint) error {
	var readData models.Employee
	return global.DataEngine.Find(&readData).Delete(&readData).Error
}

func (e EmployeeRepository) First(id uint) (*models.Employee, error) {
	var data models.Employee
	err := global.DataEngine.First(&data, "id = ?", id).Error
	return &data, err
}

func (e EmployeeRepository) Create(employee models.Employee) (models.Employee, error) {
	err := global.DataEngine.Set("gorm:association_autocreate", false).Create(&employee).Error
	if err != nil {
		return models.Employee{}, err
	}
	return employee, err
}

func (e EmployeeRepository) Find(per, page int, filter map[string]interface{},
	selectKeys []string, order []string, isCount bool) (
	employees []models.ResponseEmployee, total int64, err error) {
	sqlCon := global.DataEngine.Model(&models.Employee{}).Scopes(e.defaultScoped)
	if isCount {
		countCon := global.DataEngine.Model(&models.Employee{}).Scopes(e.defaultScoped)
		total, err = e.Count(countCon, filter)
		if err != nil {
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = []string{"users.*",
			"user_companies.name_nick as user_companies_name_nick",
			"departments.name_cn as departments_name_cn",
		}
	}
	rows, err := sqlCon.Scopes(e.CustomerWhere(filter, selectKeys, e.Paginate(per, page), e.OrderBy(order))).Where("users.deleted_at is NULL").Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var data models.ResponseEmployee
		_ = sqlCon.ScanRows(rows, &data)
		employees = append(employees, data)
	}
	return
}

func (e EmployeeRepository) UpdateRecordByModel(userId uint, updateModel models.Employee) error {
	sqlCon := global.DataEngine.Model(&models.Employee{ID: userId})
	return sqlCon.Omit("last_sign_in_ip", "last_sign_in_at", "sign_in_count", "created_at", "current_sign_in_at").Updates(updateModel).Error
}

//通过手机号码和邮箱查询当前用户
func (e EmployeeRepository) FirstByPhoneAndEmail(phone string, email string) (employee *models.Employee, err error) {
	var user models.Employee
	err = global.DataEngine.Where(&models.Employee{Phone: phone, Email: email}).First(&user).Error
	employee = &user
	return
}

func (e EmployeeRepository) UpdateColumnByID(employeeID uint, updateColumn map[string]interface{}) error {
	user := models.Employee{ID: employeeID}
	return global.DataEngine.Model(&user).Updates(updateColumn).Error
}

func (e EmployeeRepository) UpdateColumn(employee *models.Employee, updateColumn map[string]interface{}) error {
	return global.DataEngine.Model(&employee).Updates(updateColumn).Error
}

func (e EmployeeRepository) FirstByPhoneOrEmail(account string) (employee *models.Employee, err error) {
	var user models.Employee
	err = global.DataEngine.Scopes(e.defaultScoped).Where("users.phone = ?", account).Or("users.email = ?", account).First(&user).Error
	employee = &user
	return
}

func (e EmployeeRepository) defaultScoped(db *gorm.DB) *gorm.DB {
	db = db.Joins("inner join user_companies on user_companies.id = users.user_company_id and user_companies.company_type = 4")
	db = db.Joins("left join departments on departments.id = users.department_id")
	return db
}

func NewEmployeeRepository() IEmployeeRepository {
	return &EmployeeRepository{}
}
