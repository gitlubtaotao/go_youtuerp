package repositories

import (
	"gorm.io/gorm"
	"youtuerp/database"
	"youtuerp/models"
)

type IDepartmentRepository interface {
	Find(per, page int, attr map[string]interface{}, selectKeys []string,
		order []string, isCount bool) ([]interface{}, int64, error)
	First(id uint) (*models.Department, error)
	Update(id uint, updateData models.Department) error
	Create(department models.Department) (models.Department, error)
	Delete(id uint) error
}
type DepartmentRepository struct {
	BaseRepository
}

func (d DepartmentRepository) Delete(id uint) error {
	var readData models.Department
	return database.GetDBCon().Find(&readData).Delete(&readData).Error
}

func (d DepartmentRepository) Create(department models.Department) (models.Department, error) {
	err := database.GetDBCon().Set("gorm:association_autocreate", false).Create(&department).Error
	if err != nil {
		return models.Department{}, err
	}
	return department, err
}

func (d DepartmentRepository) Find(per, page int, filter map[string]interface{}, selectKeys []string,
	order []string, isCount bool) (departments []interface{}, total int64, err error) {
	sqlCon := database.GetDBCon().Model(&models.Department{}).Scopes(d.defaultScope)
	selectKeys = []string{
		"departments.id",
		"departments.created_at",
		"departments.updated_at",
		"departments.name_en",
		"departments.name_cn",
		"user_companies.name_nick as user_companies_name_nick",
		"departments.user_company_id",
	}
	if isCount {
		countCon := database.GetDBCon().Model(&models.Department{}).Scopes(d.defaultScope).Scopes(d.Ransack(filter))
		if err = countCon.Count(&total).Error; err != nil {
			return
		}
	}
	rows, err  := sqlCon.Scopes(d.CustomerWhere(filter,selectKeys,d.Paginate(per, page), d.OrderBy(order))).Where("departments.deleted_at is NULL").Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var department models.ResponseDepartment
		_ = sqlCon.ScanRows(rows, &department)
		departments = append(departments, &department)
	}
	return departments, total, err
}

func (d DepartmentRepository) First(id uint) (department *models.Department, err error) {
	var data models.Department
	err = database.GetDBCon().First(&data, "id = ?", id).Error
	return &data, err
}

func (d DepartmentRepository) Update(id uint, updateData models.Department) error {
	return database.GetDBCon().Model(&models.Department{ID: id}).Updates(updateData).Error
}

func (d DepartmentRepository) defaultScope(db *gorm.DB) *gorm.DB {
	return db.Joins("INNER JOIN user_companies on departments.user_company_id = user_companies.id")
}

func NewDepartmentRepository() IDepartmentRepository {
	return &DepartmentRepository{}
}
