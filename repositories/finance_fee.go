package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
	"youtuerp/database"
	"youtuerp/models"
	"youtuerp/tools"
)

type IFinanceFee interface {
	//根据不同的查询条件查询费用
	FindFinanceFees(per, page uint, filter map[string]interface{},
		selectKeys []string, orders []string) ([]models.ResultFinanceFee, uint, error)
	//根据不同的结算查询历史费用
	GetHistoryFee(filter map[string]interface{}, limit int, selectKeys []string) ([]models.FinanceFee, error)
	//主要通过费用ID进行查询
	FindFeesById(ids []uint, otherKeys ...string) ([]models.FinanceFee, error)
	//查询费用信息
	FindFees(per, page uint, filter map[string]interface{}, selectKeys []string,
		orders []string)
	//更改费用状态
	ChangeStatusFees(ids []uint, status string) error
	//删除订单费用信息
	DeleteFees([]uint) error
	//批量插入订单费用信息或者更新费用信息
	BulkInsertOrUpdate(financeFees []models.FinanceFee) ([]models.FinanceFee, error)
	OrderFees(attr map[string]interface{}, payOrReceive ...string) (map[string][]models.FinanceFee, error)
}

type FinanceFee struct {
	BaseRepository
}

func (f FinanceFee) FindFinanceFees(per, page uint, filter map[string]interface{},
	selectKeys []string, orders []string) (financeFees []models.ResultFinanceFee, total uint, err error) {
	var keys []string
	sqlCon := database.GetDBCon().Model(&models.FinanceFee{})
	sqlCon = sqlCon.Joins("INNER JOIN order_masters ON order_masters.id = finance_fees.order_master_id")
	sqlCon = sqlCon.Joins("INNER JOIN order_extend_infos ON order_extend_infos.order_master_id = order_masters.id ")
	if len(filter) > 0 {
		sqlCon = sqlCon.Scopes(f.Ransack(filter))
	}
	if err = sqlCon.Count(&total).Error; err != nil {
		return
	}
	if len(selectKeys) == 0 {
		keys, err = tools.GetStructFieldByJson(models.FinanceFee{})
		if err != nil {
			return
		}
		for i := 0; i < len(keys); i++ {
			keys[i] = "finance_fees." + keys[i]
		}
		keys = append(keys,"order_masters.serial_number")
		golog.Infof("select keys is %v", keys)
		selectKeys = keys
	}
	rows, err := sqlCon.Scopes(f.Paginate(per, page), f.OrderBy(orders)).Select(selectKeys).Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var data models.ResultFinanceFee
		_ = sqlCon.ScanRows(rows, &data)
		financeFees = append(financeFees, data)
	}
	return
}

//获取历史费用
func (f FinanceFee) GetHistoryFee(filter map[string]interface{}, limit int, selectKeys []string) ([]models.FinanceFee, error) {
	var financeFees []models.FinanceFee
	if limit == 0 {
		limit = 50
	}
	sqlCon := database.GetDBCon().Model(&models.FinanceFee{})
	if len(filter) > 0 {
		sqlCon = sqlCon.Scopes(f.Ransack(filter))
	}
	if len(selectKeys) > 0 {
		sqlCon = sqlCon.Select(selectKeys)
	}
	rows, err := sqlCon.Order("id desc").Group("name").Limit(limit).Rows()
	if err != nil {
		return financeFees, err
	}
	for rows.Next() {
		var data models.FinanceFee
		_ = sqlCon.ScanRows(rows, &data)
		financeFees = append(financeFees, data)
	}
	return financeFees, nil
}

func (f FinanceFee) FindFeesById(ids []uint, otherKeys ...string) ([]models.FinanceFee, error) {
	var financeFees []models.FinanceFee
	sqlConn := database.GetDBCon().Where("id IN (?)", ids)
	if len(otherKeys) >= 1 {
		sqlConn = sqlConn.Select(otherKeys)
	}
	err := sqlConn.Find(&financeFees).Error
	return financeFees, err
}

func (f FinanceFee) FindFees(per, page uint, filter map[string]interface{}, selectKeys []string,
	orders []string) {
	
}

func (f FinanceFee) ChangeStatusFees(ids []uint, status string) error {
	return database.GetDBCon().Where("id IN (?)", ids).Model(&models.FinanceFee{}).Updates(map[string]interface{}{"status": status}).Error
}

func (f FinanceFee) DeleteFees(ids []uint) error {
	return database.GetDBCon().Where("id IN (?) and status IN (?)", ids, []string{
		models.FinanceFeeStatusInit,
		models.FinanceFeeStatusDismiss,
		models.FinanceFeeStatusVerify,
	}).Delete(models.FinanceFee{}).Error
}

func (f FinanceFee) BulkInsertOrUpdate(financeFees []models.FinanceFee) ([]models.FinanceFee, error) {
	sqlConn := database.GetDBCon()
	var data []models.FinanceFee
	err := sqlConn.Transaction(func(tx *gorm.DB) error {
		for _, item := range financeFees {
			if item.ID != 0 {
				if err := tx.Model(&models.FinanceFee{ID: item.ID}).Update(tools.StructToChange(item)).Error; err != nil {
					return err
				}
			} else {
				if err := tx.Create(&item).Error; err != nil {
					return err
				}
				data = append(data, item)
			}
		}
		// 在事务中做一些数据库操作 (这里应该使用 'tx' ，而不是 'db')
		return nil
	})
	return data, err
}

func (f FinanceFee) OrderFees(attr map[string]interface{}, payOrReceive ...string) (map[string][]models.FinanceFee, error) {
	var (
		financeFees = make(map[string][]models.FinanceFee)
		temp        []models.FinanceFee
	)
	if len(payOrReceive) == 0 {
		return financeFees, nil
	}
	sqlConn := database.GetDBCon().Where(attr)
	for _, item := range payOrReceive {
		if err := sqlConn.Where("pay_or_receive = ?", item).Find(&temp).Error; err != nil {
			return financeFees, err
		}
		financeFees[item] = temp
	}
	return financeFees, nil
}

func NewFinanceFee() IFinanceFee {
	return FinanceFee{}
}
