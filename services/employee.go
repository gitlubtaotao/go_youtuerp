package services

import (
	"youtuerp/models"
	"youtuerp/repositories"
)

type IEmployeeService interface {
	FirstByNameOrEmail(account string) (*models.Employee, error)
	UpdatePassword(user *models.Employee, password string) error
}

type EmployeeService struct {
	repo repositories.IEmployeeRepository
	BaseService
}

func (e *EmployeeService) UpdatePassword(user *models.Employee, password string) error {
	updateColumn := map[string]interface{}{
		"encrypted_password": password,
	}
	return e.repo.UpdateColumn(user, updateColumn)
}

func (e *EmployeeService) FirstByNameOrEmail(account string) (*models.Employee, error) {
	return e.repo.FirstByPhoneOrEmail(account)
}

func NewEmployeeService() IEmployeeService {
	return &EmployeeService{repo: repositories.NewEmployeeRepository()}
}
