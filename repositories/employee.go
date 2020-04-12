package repositories

type IEmployeeRepository interface {
}
type EmployeeRepository struct {
	BaseRepository
}

func NewEmployeeRepository() IEmployeeRepository {
	return &EmployeeRepository{}
}
