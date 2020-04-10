package oa

import (
	"youtuerp/conf"
	"youtuerp/models"
)

type ICompanyRepository interface {
	/*查询公司信息
	per: 每次分页的记录数
	page: 当前页数
	attr: 查询公司信息
	*/
	FindCompany(per, page uint, attr map[string]interface{}, selectKeys []string, order []string) (companies []*models.UserCompany, err error)
}

type CompanyRepository struct {
}

func (c CompanyRepository) FindCompany(per, page uint, attr map[string]interface{}, selectKeys []string, order []string) (companies []*models.UserCompany, err error) {
	if per == 0 {
		per = conf.Configuration.PerPage
	}
	return
}


func NewCompanyRepository() ICompanyRepository {
	return &CompanyRepository{}
}
