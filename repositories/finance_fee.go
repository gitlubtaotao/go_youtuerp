package repositories

import (
	"gorm.io/gorm"
	"youtuerp/database"
	"youtuerp/models"
	"youtuerp/tools"
)

type IFinanceFee interface {
	//根据不同的查询条件查询费用
	FindFinanceFees(per, page int, filter map[string]interface{},
		selectKeys []string, orders []string) ([]models.ResponseFinanceFee, int64, error)
	//根据不同的结算查询历史费用
	GetHistoryFee(filter map[string]interface{}, limit int, selectKeys []string) ([]models.FinanceFee, error)
	//主要通过费用ID进行查询
	FindFeesById(ids []uint, otherFilter map[string]interface{}, otherKeys ...string) ([]models.FinanceFee, error)
	//查询费用信息
	FindFees(per, page int, filter map[string]interface{}, selectKeys []string,
		orders []string)
	//更改费用状态
	ChangeStatusFees(ids []uint, otherFilter map[string]interface{}, status string, ) error
	//删除订单费用信息
	DeleteFees([]uint) error
	//批量插入订单费用信息或者更新费用信息
	BulkInsertOrUpdate(financeFees []models.FinanceFee) ([]models.FinanceFee, error)
	OrderFees(attr map[string]interface{}, payOrReceive ...string) (map[string][]models.FinanceFee, error)
}

type FinanceFee struct {
	BaseRepository
}

func (f FinanceFee) FindFinanceFees(per, page int, filter map[string]interface{},
	selectKeys []string, orders []string) (financeFees []models.ResponseFinanceFee, total int64, err error) {
	sqlCon := database.GetDBCon().Model(&models.FinanceFee{}).Scopes(f.defaultJoinTables).Where("finance_fees.deleted_at IS NULL")
	if total, err = f.Count(database.GetDBCon().Model(&models.FinanceFee{}).Scopes(f.defaultJoinTables), filter); err != nil {
		return
	}
	if len(selectKeys) == 0 {
		selectKeys = append(selectKeys, "finance_fees.*", "order_masters.serial_number as serial_number")
	}
	rows, err := sqlCon.Scopes(f.CustomerWhere(filter, selectKeys, f.Paginate(per, page), f.OrderBy(orders))).Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var data models.ResponseFinanceFee
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

func (f FinanceFee) FindFeesById(ids []uint, otherFilter map[string]interface{}, otherKeys ...string) ([]models.FinanceFee, error) {
	var financeFees []models.FinanceFee
	sqlConn := database.GetDBCon().Where("id IN (?)", ids)
	if len(otherFilter) > 0 {
		sqlConn.Scopes(f.crud.ransack(otherFilter))
	}
	if len(otherKeys) >= 1 {
		sqlConn = sqlConn.Select(otherKeys)
	}
	err := sqlConn.Find(&financeFees).Error
	return financeFees, err
}

func (f FinanceFee) FindFees(per, page int, filter map[string]interface{}, selectKeys []string,
	orders []string) {
	
}

func (f FinanceFee) ChangeStatusFees(ids []uint, otherFilter map[string]interface{}, status string) error {
	sqlCon := database.GetDBCon().Model(&models.FinanceFee{})
	if len(otherFilter) > 0 {
		sqlCon.Scopes(f.crud.ransack(otherFilter))
	}
	return sqlCon.Where("id IN (?)", ids).Updates(map[string]interface{}{"status": status}).Error
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
	var updateFinance []models.FinanceFee
	var createFinance []models.FinanceFee
	for _, item := range financeFees {
		if item.ID != 0 {
			updateFinance = append(updateFinance, item)
		} else {
			createFinance = append(createFinance, item)
		}
	}
	err := sqlConn.Transaction(func(tx *gorm.DB) error {
		for _, item := range updateFinance {
			if err := tx.Model(&models.FinanceFee{ID: item.ID}).Updates(tools.StructToChange(item)).Error; err != nil {
				return err
			}
		}
		if err := tx.Create(&createFinance).Error; err != nil {
			return err
		}
		// 在事务中做一些数据库操作 (这里应该使用 'tx' ，而不是 'db')
		return nil
	})
	return createFinance, err
}

func (f FinanceFee) OrderFees(attr map[string]interface{}, payOrReceive ...string) (map[string][]models.FinanceFee, error) {
	var (
		financeFees = make(map[string][]models.FinanceFee)
		temp        []models.FinanceFee
	)
	if len(payOrReceive) == 0 {
		return financeFees, nil
	}
	for _, item := range payOrReceive {
		if err := database.GetDBCon().Where(attr).Where("pay_or_receive = ?", item).Find(&temp).Error; err != nil {
			return financeFees, err
		}
		financeFees[item] = temp
	}
	return financeFees, nil
}

func (f FinanceFee) defaultJoinTables(db *gorm.DB) *gorm.DB {
	db.Joins("INNER JOIN order_masters ON order_masters.id = finance_fees.order_master_id")
	db = db.Joins("INNER JOIN order_extend_infos ON order_extend_infos.order_master_id = finance_fees.order_master_id")
	return db
}

func NewFinanceFee() IFinanceFee {
	return FinanceFee{}
}
