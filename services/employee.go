package services

import (
	"errors"
	"youtuerp/models"
	"youtuerp/redis"
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

func (e EmployeeService) Delete(id uint) error {
	err := e.repo.Delete(id)
	if err != nil {
		return err
	}
	return redis.NewRedis().SRemove(models.User{}.TableName(), id)
}

func (e EmployeeService) First(id uint) (*models.Employee, error) {
	return e.repo.First(id)
}

func (e EmployeeService) Create(employee models.Employee) (models.Employee, error) {
	data, err := e.repo.Create(employee)
	if err != nil {
		return data, err
	}
	e.SaveRedisData(data)
	return data, nil
}

func (e EmployeeService) Find(per, page uint, filter map[string]interface{}, selectKeys []string, order []string, isCount bool) ([]models.ResultEmployee, uint, error) {
	return e.repo.Find(per, page, filter, selectKeys, order, isCount)
}

func (e *EmployeeService) UpdateRecord(user *models.Employee, employee models.Employee) error {
	validate := NewValidatorService(employee)
	if message := validate.ResultError("zh-CN"); message != "" {
		return errors.New(message)
	}
	err := e.repo.UpdateRecordByModel(user, employee)
	if err != nil {
		return err
	}
	e.SaveRedisData(*user)
	return nil
}

func (e EmployeeService) UpdateColumn(user *models.Employee, updateColumn map[string]interface{}) error {
	err := e.repo.UpdateColumn(user, updateColumn)
	if err != nil {
		return err
	}
	e.SaveRedisData(*user)
	return nil
}

func (e EmployeeService) FirstByPhoneAndEmail(phone string, email string) (*models.Employee, error) {
	return e.repo.FirstByPhoneAndEmail(phone, email)
}

func (e EmployeeService) UpdatePassword(user *models.Employee, password string) error {
	updateColumn := map[string]interface{}{
		"encrypted_password": password,
	}
	return e.repo.UpdateColumn(user, updateColumn)
}

func (e EmployeeService) FirstByPhoneOrEmail(account string) (*models.Employee, error) {
	return e.repo.FirstByPhoneOrEmail(account)
}

func (e EmployeeService) SaveRedisData(result models.Employee) {
	redis.HSetValue(models.User{}.TableName(), result.ID, map[string]interface{}{
		"id":              result.ID,
		"name":            result.Name,
		"user_company_id": result.UserCompanyId,
	})
}

func NewEmployeeService() IEmployeeService {
	return &EmployeeService{repo: repositories.NewEmployeeRepository()}
}
