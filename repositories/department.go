package repositories

import (
	"fmt"
	"youtuerp/database"
	"youtuerp/models"
)

type IDepartmentRepository interface {
	Find(per, page uint, attr map[string]interface{}, selectKeys []string,
		order []string, isCount bool) ([]*models.Department, uint, error)
}
type DepartmentRepository struct {
	BaseRepository
}

func (d DepartmentRepository) Find(per, page uint, filter map[string]interface{}, selectKeys []string,
	order []string, isCount bool) ([]*models.Department, uint, error) {
	var (
		err         error
		departments []*models.Department
		total       = 0
	)
	sqlCon := database.GetDBCon().Model(&models.Department{})
	if len(filter) > 0 {
		sqlCon = sqlCon.Scopes(d.Ransack(filter))
	}
	if isCount {
		err = sqlCon.Count(&total).Error
		fmt.Print(err)
		if err != nil {
			return departments, uint(total), err
		}
	}
	sqlCon = sqlCon.Scopes(d.Paginate(per, page), d.OrderBy(order))
	if len(selectKeys) > 0 {
		sqlCon = sqlCon.Select(selectKeys)
	}
	err = sqlCon.Preloads(&models.UserCompany{}).Find(&departments).Error
	if err != nil {
		return nil, uint(total), err
	}
	return departments, uint(total), err
}

func NewDepartmentRepository() IDepartmentRepository {
	return &DepartmentRepository{}
}
