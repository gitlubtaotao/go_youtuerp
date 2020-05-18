package services

import (
	"errors"
	"youtuerp/models"
	"youtuerp/repositories"
)

type ICrmUser interface {
	Delete(id uint) error
	Update(id uint, user models.CrmUser,language string) error
	Create(user models.CrmUser, language string) (models.CrmUser, error)
	Find(per, page uint, filter map[string]interface{}, selectKeys []string,
		order []string) ([]models.CrmUser, uint, error)
}
type CrmUser struct {
	BaseService
	repo repositories.ICrmUser
}

func (c CrmUser) Delete(id uint) error {
	return c.repo.Delete(id)
}

func (c CrmUser) Update(id uint, user models.CrmUser, language string) error {
	valid := NewValidatorService(user)
	if message := valid.ResultError(language); message != "" {
		return errors.New(message)
	}
	return c.repo.Update(id, user)
}

func (c CrmUser) Find(per, page uint, filter map[string]interface{}, selectKeys []string,
	order []string) ([]models.CrmUser, uint, error) {
	return c.repo.Find(per, page, filter, selectKeys, order, true)
}

func (c CrmUser) Create(user models.CrmUser, language string) (models.CrmUser, error) {
	var err error
	valid := NewValidatorService(user)
	user.CompanyType = 1
	if message := valid.ResultError(language); message != "" {
		return models.CrmUser{}, errors.New(message)
	}
	if user, err = c.repo.Create(user); err != nil {
		return models.CrmUser{}, err
	}
	//	进行redis 缓存
	return user, err
}

func NewCrmUser() ICrmUser {
	return CrmUser{repo: repositories.NewCrmUser()}
}
