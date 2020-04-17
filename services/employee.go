package services

import (
	"youtuerp/models"
	"youtuerp/repositories"
)

type IEmployeeService interface {
	FirstByPhoneAndEmail(phone string, email string) (*models.Employee, error)
	FirstByPhoneOrEmail(account string) (*models.Employee, error)
	UpdatePassword(user *models.Employee, password string) error
	UpdateColumn(user *models.Employee, updateColumn map[string]interface{}) error
	UpdateRecord(user *models.Employee, employee models.Employee) error
}

type EmployeeService struct {
	repo repositories.IEmployeeRepository
	BaseService
}

func (e *EmployeeService) UpdateRecord(user *models.Employee, employee models.Employee) error {
	//validate := NewValidatorService(user)
	//errs, _ := validate.HandlerError("zh-CN")
	//if len(errs) > 0 {
	//	return errors.New("保存失败")
	//}
	return e.repo.UpdateRecordByModel(user, employee)
}

func (e *EmployeeService) UpdateColumn(user *models.Employee, updateColumn map[string]interface{}) error {
	return e.repo.UpdateColumn(user, updateColumn)
}

func (e *EmployeeService) FirstByPhoneAndEmail(phone string, email string) (*models.Employee, error) {
	return e.repo.FirstByPhoneAndEmail(phone, email)
}

func (e *EmployeeService) UpdatePassword(user *models.Employee, password string) error {
	updateColumn := map[string]interface{}{
		"encrypted_password": password,
	}
	return e.repo.UpdateColumn(user, updateColumn)
}

func (e *EmployeeService) FirstByPhoneOrEmail(account string) (*models.Employee, error) {
	return e.repo.FirstByPhoneOrEmail(account)
}

func NewEmployeeService() IEmployeeService {
	return &EmployeeService{repo: repositories.NewEmployeeRepository()}
}
