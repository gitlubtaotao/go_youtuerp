package repositories

import (
	"database/sql"
	"fmt"
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
	FindCompany(per, page uint, attr map[string]interface{}, selectKeys []string, order []string, isCount bool) (companies []*models.UserCompany, total uint, err error)
}

type CompanyRepository struct {
	BaseRepository
}

//
////默认查询company
func (c CompanyRepository) DefaultScope(temp *gorm.DB) *gorm.DB {
	return temp.Where("company_type = ?", 4)
}

//
func (c *CompanyRepository) FindCompany(per, page uint, filters map[string]interface{},
	selectKeys []string, orders []string,
	isCount bool) (companies []*models.UserCompany, total uint, err error) {
	var rows *sql.Rows
	sqlCon := database.GetDBCon()
	temp := sqlCon.Scopes(c.DefaultScope)
	fmt.Println(len(filters),"sssss")
	if len(filters) > 0 {
		temp = c.Ransack(temp, filters)
	}
	//limit
	if isCount {
		temp.Model(&models.UserCompany{}).Count(&total)
	}
	if page == 0 && per > 0 {
		temp = temp.Limit(per)
	} else if page > 0 && per > 0 {
		temp = temp.Limit(per).Offset((page - 1) * per)
	}
	if len(orders) == 0 {
		temp = temp.Order("id desc")
	} else {
		for _, order := range orders {
			temp = temp.Order(order)
		}
	}
	if len(selectKeys) > 0 {
		temp = temp.Select(selectKeys)
	}
	rows, err = temp.Model(&models.UserCompany{}).Rows()
	if err != nil {
		return
	}
	
	for rows.Next() {
		var userCompany models.UserCompany
		_ = sqlCon.ScanRows(rows, &userCompany)
		companies = append(companies, &userCompany)
	}
	return
}

//

func NewCompanyRepository() ICompanyRepository {
	return &CompanyRepository{}
}
