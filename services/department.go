package services

import "youtuerp/repositories"

type IDepartmentService interface {
}

type DepartmentService struct {
	repo repositories.IDepartmentRepository
	BaseService
}
