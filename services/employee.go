package services

import (
	"errors"
	"strconv"
	"sync"
	"youtuerp/models"
	"youtuerp/redis"
	"youtuerp/repositories"
)

type IEmployeeService interface {
	FindRedis() []map[string]string
	First(id uint) (*models.Employee, error)
	FirstByPhoneAndEmail(phone string, email string) (*models.Employee, error)
	FirstByPhoneOrEmail(account string) (*models.Employee, error)
	UpdatePassword(user *models.Employee, password string) error
	UpdateColumn(user *models.Employee, updateColumn map[string]interface{}) error
	UpdateRecord(userId uint, employee models.Employee) error
	Find(per, page int, filter map[string]interface{}, selectKeys []string, order []string, isCount bool) ([]models.ResultEmployee, int64, error)
	Create(employee models.Employee) (models.Employee, error)
	Delete(id uint) error
}

type EmployeeService struct {
	repo repositories.IEmployeeRepository
	BaseService
	mu sync.Mutex
}

func (e EmployeeService) FindRedis() []map[string]string {
	red := redis.NewRedis()
	tableName := models.Employee{}.TableName()
	data := make([]map[string]string, 0)
	data = red.HCollectOptions(tableName)
	if len(data) > 0 {
		return data
	}
	employees, _, err := e.Find(0, 0, map[string]interface{}{}, []string{}, []string{}, false)
	if err != nil {
		return []map[string]string{}
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	for _, k := range employees {
		temp := map[string]string{
			"id":              strconv.Itoa(int(k.ID)),
			"user_company_id": strconv.Itoa(k.UserCompanyId),
			"name":            k.Name,
		}
		go e.SaveRedisData(toolOther.StructToMap(k))
		data = append(data, temp)
	}
	return data
}

func (e EmployeeService) Delete(id uint) error {
	err := e.repo.Delete(id)
	if err != nil {
		return err
	}
	go redis.NewRedis().SRemove(models.User{}.TableName(), id)
	return nil
}

func (e EmployeeService) First(id uint) (*models.Employee, error) {
	return e.repo.First(id)
}

func (e EmployeeService) Create(employee models.Employee) (models.Employee, error) {
	valid := NewValidatorService(employee)
	if message := valid.ResultError("zh-CN"); message != "" {
		return employee, errors.New(message)
	}
	data, err := e.repo.Create(employee)
	if err != nil {
		return data, err
	}
	go e.SaveRedisData(toolOther.StructToMap(data))
	return data, nil
}

func (e EmployeeService) Find(per, page int, filter map[string]interface{}, selectKeys []string, order []string, isCount bool) ([]models.ResultEmployee, int64, error) {
	return e.repo.Find(per, page, filter, selectKeys, order, isCount)
}

func (e *EmployeeService) UpdateRecord(userId uint, employee models.Employee) error {
	validate := NewValidatorService(employee)
	if message := validate.ResultError("zh-CN"); message != "" {
		return errors.New(message)
	}
	err := e.repo.UpdateRecordByModel(userId, employee)
	if err != nil {
		return err
	}
	go e.SaveRedisData(toolOther.StructToMap(employee))
	return nil
}

func (e EmployeeService) UpdateColumn(user *models.Employee, updateColumn map[string]interface{}) error {
	err := e.repo.UpdateColumn(user, updateColumn)
	if err != nil {
		return err
	}
	go e.SaveRedisData(toolOther.StructToMap(*user))
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

//
func (e EmployeeService) SaveRedisData(result map[string]interface{}) {
	redis.HSetValue(models.User{}.TableName(), result["id"], map[string]interface{}{
		"id":              result["id"],
		"name":            result["name"],
		"user_company_id": result["user_company_id"],
	})
}

func NewEmployeeService() IEmployeeService {
	return &EmployeeService{repo: repositories.NewEmployeeRepository()}
}
