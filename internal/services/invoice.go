package services

import (
	"errors"
	"youtuerp/internal/dao"
	"youtuerp/internal/models"
)

type IInvoiceService interface {
	Delete(id uint) error
	UpdateById(id uint, updateContent models.Invoice, language string) (models.Invoice, error)
	First(id uint) (models.Invoice, error)
	Create(account models.Invoice, language string) (models.Invoice, error)
	FindByOa(per, page int, filter map[string]interface{}, selectKeys []string,
		order []string) (accounts []models.Invoice, total int64, err error)
	FindByCrm(per, page int, filter map[string]interface{}, selectKeys []string,
		orders []string) (accounts []models.Invoice, total int64, err error)
}

type InvoiceService struct {
	BaseService
	repo dao.IInvoiceRepository
}

func (a InvoiceService) Delete(id uint) error {
	return a.repo.Delete(id)
}

func (a InvoiceService) UpdateById(id uint, updateContent models.Invoice, language string) (models.Invoice, error) {
	validate := NewValidatorService(updateContent)
	if message := validate.ResultError(language); message != "" {
		return models.Invoice{}, errors.New(message)
	}
	return a.repo.UpdateById(id, updateContent)
}

func (a InvoiceService) First(id uint) (models.Invoice, error) {
	return a.repo.First(id)
}

func (a InvoiceService) FindByOa(per, page int, filter map[string]interface{}, selectKeys []string,
	order []string) (accounts []models.Invoice, total int64, err error) {
	return a.repo.FindByOa(per, page, filter, selectKeys, order)
}
func (a InvoiceService) FindByCrm(per, page int, filter map[string]interface{}, selectKeys []string,
	orders []string) (accounts []models.Invoice, total int64, err error) {
	return a.repo.FindByCrm(per, page, filter, selectKeys, orders)
}

func (a InvoiceService) Create(account models.Invoice, language string) (models.Invoice, error) {
	validate := NewValidatorService(account)
	if message := validate.ResultError(language); message != "" {
		return models.Invoice{}, errors.New(message)
	}
	return a.repo.Create(account)
}

func NewInvoiceService() IInvoiceService {
	return InvoiceService{repo: dao.NewInvoice()}
}
