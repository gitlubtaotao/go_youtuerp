package repositories

import (
	"github.com/jinzhu/gorm"
	"youtuerp/database"
	"youtuerp/models"
	"youtuerp/tools"
)

type IFinanceFee interface {
	//删除订单费用信息
	DeleteFees([]uint) error
	//批量插入订单费用信息或者更新费用信息
	BulkInsertOrUpdate(financeFees []models.FinanceFee) ([]models.FinanceFee, error)
	OrderFees(attr map[string]interface{}, payOrReceive ...string) (map[string][]models.FinanceFee, error)
}

type FinanceFee struct {
}

func (f FinanceFee) DeleteFees(ids []uint) error {
	return database.GetDBCon().Where("id IN (?) and status IN (?)", ids,[]string{models.FinanceFeeStatusInit,models.FinanceFeeStatusDismiss}).Delete(models.FinanceFee{}).Error
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
