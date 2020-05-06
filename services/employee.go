package services

import (
	"youtuerp/models"
	"youtuerp/repositories"
)

type IEmployeeService interface {
	First(id uint) (*models.Employee, error)
	FirstByPhoneAndEmail(phone string, email string) (*models.Employee, error)
	FirstByPhoneOrEmail(account string) (*models.Employee, error)
	UpdatePassword(user *models.Employee, password string) error
	UpdateColumn(user *models.Employee, updateColumn map[string]interface{}) error
	UpdateRecord(user *models.Employee, employee models.Employee) error
	Find(per, page uint, filter map[string]interface{}, selectKeys []string, order []string, isCount bool) ([]models.ResultEmployee, uint, error)
	Create(employee models.Employee) (models.Employee, error)
	Delete(id uint) error
}

type EmployeeService struct {
	repo repositories.IEmployeeRepository
	BaseService
}

func (e *EmployeeService) Delete(id uint) error {
	return e.repo.Delete(id)
}

func (e *EmployeeService) First(id uint) (*models.Employee, error) {
	return e.repo.First(id)
}

func (e *EmployeeService) Create(employee models.Employee) (models.Employee, error) {
	return e.repo.Create(employee)
}

func (e *EmployeeService) Find(per, page uint, filter map[string]interface{}, selectKeys []string, order []string, isCount bool) ([]models.ResultEmployee, uint, error) {
	return e.repo.Find(per, page, filter, selectKeys, order, isCount)
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
