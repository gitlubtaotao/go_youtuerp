package services

import "youtuerp/repositories"

type IEmployeeService interface {
}

type EmployeeService struct {
	repo repositories.IEmployeeRepository
	BaseService
}

func NewEmployeeService() IEmployeeService {
	return &EmployeeService{repo: repositories.NewEmployeeRepository()}
}
