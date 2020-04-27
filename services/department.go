package services

import (
	"youtuerp/models"
	"youtuerp/repositories"
)

type IDepartmentService interface {
	Find(per, page uint, attr map[string]interface{}, selectKeys []string,
		order []string, isCount bool) ([]*models.Department, uint, error)
	Create(department models.Department) (models.Department, error)
}

type DepartmentService struct {
	repo repositories.IDepartmentRepository
	BaseService
}

func (d *DepartmentService) Create(department models.Department) (models.Department, error) {
	return d.repo.Create(department)
}

func NewDepartmentService() IDepartmentService {
	return &DepartmentService{repo: repositories.NewDepartmentRepository()}
}

func (d *DepartmentService) Find(per, page uint, attr map[string]interface{}, selectKeys []string,
	order []string, isCount bool) ([]*models.Department, uint, error) {
	return d.repo.Find(per, page, attr, selectKeys, order, isCount)
}
