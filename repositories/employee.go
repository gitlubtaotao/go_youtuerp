package repositories

import "sync"

type IEmployeeRepository interface {
}
type EmployeeRepository struct {
	mu sync.Mutex
}

func NewEmployeeRepository() IEmployeeRepository {
	return &EmployeeRepository{mu: sync.Mutex{}}
}
