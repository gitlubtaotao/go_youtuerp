package repositories

import (
	"github.com/jinzhu/gorm"
	"youtuerp/database"
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
//
////默认查询company
func (c CompanyRepository) DefaultScope(temp *gorm.DB) *gorm.DB {
	return temp.Where("company_type = ?", 4)
}
//
func (c *CompanyRepository) FindCompany(per, page uint, attr map[string]interface{}, selectKeys []string, orders []string) (companies []*models.UserCompany, err error) {
	temp := database.GetDBCon()
	temp = temp.Scopes(c.DefaultScope)
	if len(attr) > 0 {
		temp = temp.Where(attr)
	}
	temp = temp.Joins("INNER JOIN companies on companies.source_id = user_companies.id")
	//limit
	if page == 0 && per > 0 {
		temp = temp.Limit(per)
	} else if page > 0 && per > 0 {
		temp = temp.Limit(per).Offset((page - 1) * per)
	}
	if len(orders) == 0 {
		temp = temp.Order("id desc")
	}
	for _, order := range orders {
		temp = temp.Order(order)
	}
	if len(selectKeys) > 0 {
		temp = temp.Select(selectKeys)
	}
	err = temp.Scan(&companies).Error
	return
}
//

func NewCompanyRepository() ICompanyRepository {
	return &CompanyRepository{}
}
