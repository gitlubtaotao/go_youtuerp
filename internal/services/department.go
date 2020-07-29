package services

import (
	"youtuerp/internal/dao"
	"youtuerp/internal/models"
)

type IDepartmentService interface {
	Find(per, page int, attr map[string]interface{}, selectKeys []string,
		order []string, isCount bool) ([]interface{}, int64, error)
	FindAll(filter map[string]interface{}, selectKeys []string, order []string, isCount bool) ([]interface{}, int64, error)
	First(id uint) (*models.Department, error)
	Update(id uint, updateData models.Department) error
	Create(department models.Department) (models.Department, error)
	Delete(id uint) error
}

type DepartmentService struct {
	repo dao.IDepartmentRepository
	BaseService
}

func (d *DepartmentService) Find(per, page int, attr map[string]interface{}, selectKeys []string,
	order []string, isCount bool) ([]interface{}, int64, error) {
	return d.repo.Find(per, page, attr, selectKeys, order, isCount)
}
func (d *DepartmentService) FindAll(filter map[string]interface{}, selectKeys []string, order []string, isCount bool) ([]interface{}, int64, error) {
	return d.repo.Find(0, 0, filter, selectKeys, order, isCount)
}

func (d *DepartmentService) Delete(id uint) error {
	return d.repo.Delete(id)
}

func (d *DepartmentService) Update(id uint, updateData models.Department) error {
	return d.repo.Update(id, updateData)
}

func (d *DepartmentService) First(id uint) (*models.Department, error) {
	return d.repo.First(id)
}

func (d *DepartmentService) Create(department models.Department) (models.Department, error) {
	return d.repo.Create(department)
}

func NewDepartmentService() IDepartmentService {
	return &DepartmentService{repo: dao.NewDepartmentRepository()}
}
