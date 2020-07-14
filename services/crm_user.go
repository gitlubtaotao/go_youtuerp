package services

import (
	"errors"
	"youtuerp/models"
	"youtuerp/repositories"
)

type ICrmUser interface {
	Delete(id uint) error
	Update(id uint, user models.CrmContact,language string) error
	Create(user models.CrmContact, language string) (models.CrmContact, error)
	Find(per, page int, filter map[string]interface{}, selectKeys []string,
		order []string) ([]models.CrmContact, int64, error)
}
type CrmUser struct {
	BaseService
	repo repositories.ICrmUser
}

func (c CrmUser) Delete(id uint) error {
	return c.repo.Delete(id)
}

func (c CrmUser) Update(id uint, user models.CrmContact, language string) error {
	valid := NewValidatorService(user)
	if message := valid.ResultError(language); message != "" {
		return errors.New(message)
	}
	return c.repo.Update(id, user)
}

func (c CrmUser) Find(per, page int, filter map[string]interface{}, selectKeys []string,
	order []string) ([]models.CrmContact, int64, error) {
	return c.repo.Find(per, page, filter, selectKeys, order, true)
}

func (c CrmUser) Create(user models.CrmContact, language string) (models.CrmContact, error) {
	var err error
	valid := NewValidatorService(user)
	//user.CompanyType = 1
	if message := valid.ResultError(language); message != "" {
		return models.CrmContact{}, errors.New(message)
	}
	if user, err = c.repo.Create(user); err != nil {
		return models.CrmContact{}, err
	}
	//	进行redis 缓存
	return user, err
}

func NewCrmUser() ICrmUser {
	return CrmUser{repo: repositories.NewCrmUser()}
}
