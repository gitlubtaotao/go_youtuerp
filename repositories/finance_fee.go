package repositories

import (
	"youtuerp/database"
	"youtuerp/models"
)

type IFinanceFee interface {
	OrderFees(attr map[string]interface{}, payOrReceive ...string) (map[string][]models.FinanceFee, error)
}

type FinanceFee struct {
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
