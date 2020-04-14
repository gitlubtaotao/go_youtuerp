package services

import (
	"youtuerp/models"
	"youtuerp/repositories"
)

type IEmployeeService interface {
	FirstByNameOrEmail(account string) (*models.Employee, error)
}

type EmployeeService struct {
	repo repositories.IEmployeeRepository
	BaseService
}

func (e *EmployeeService) FirstByNameOrEmail(account string) (*models.Employee, error) {
	return e.repo.FirstByNameOrEmail(account)
}

func NewEmployeeService() IEmployeeService {
	return &EmployeeService{repo: repositories.NewEmployeeRepository()}
}
