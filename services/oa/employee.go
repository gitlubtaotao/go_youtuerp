package oa

import "youtuerp/repositories"

type IEmployeeService interface {
}

type EmployeeService struct {
	repo repositories.IEmployeeRepository
}

func NewEmployeeService() IEmployeeService {
	return &EmployeeService{repo: NewEmployeeService()}
}
