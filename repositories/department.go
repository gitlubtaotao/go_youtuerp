package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"youtuerp/database"
	"youtuerp/models"
	"youtuerp/tools"
)

type IDepartmentRepository interface {
	Find(per, page int, attr map[string]interface{}, selectKeys []string,
		order []string, isCount bool) ([]interface{}, int64, error)
	First(id uint) (*models.Department, error)
	Update(department *models.Department, updateData models.Department) error
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
	sqlCon := database.GetDBCon().Model(&models.Department{})
	sqlCon = sqlCon.Scopes(d.defaultScope)
	if len(filter) > 0 {
		sqlCon = sqlCon.Scopes(d.Ransack(filter))
	}
	if isCount {
		err = sqlCon.Count(&total).Error
		fmt.Print(err)
		if err != nil {
			return departments, total, err
		}
	}
	sqlCon = sqlCon.Scopes(d.Paginate(per, page), d.OrderBy(order))
	if len(selectKeys) == 0 {
		selectKeys = []string{"departments.id", "departments.created_at",
			"departments.updated_at",
			"departments.name_en",
			"departments.name_cn",
			"user_companies.name_nick as user_companies_name_nick",
			"departments.user_company_id",
		}
	}
	rows, err := sqlCon.Select(selectKeys).Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var department models.ResultDepartment
		_ = sqlCon.ScanRows(rows, &department)
		departments = append(departments, &department)
	}
	return departments, total, err
}

func (d *DepartmentRepository) First(id uint) (department *models.Department, err error) {
	var data models.Department
	err = database.GetDBCon().First(&data, "id = ?", id).Error
	return &data, err
}

func (d *DepartmentRepository) Update(department *models.Department, updateData models.Department) error {
	return database.GetDBCon().Model(&department).Updates(tools.StructToChange(updateData)).Error
}

func (d *DepartmentRepository) defaultScope(db *gorm.DB) *gorm.DB {
	return db.Joins("INNER JOIN user_companies " +
		"on departments.user_company_id = user_companies.id")
	
}

func NewDepartmentRepository() IDepartmentRepository {
	return &DepartmentRepository{}
}
