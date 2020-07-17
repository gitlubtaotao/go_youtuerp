package repositories

import (
	"database/sql"
	"fmt"
	"github.com/kataras/golog"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
	"youtuerp/database"
	"youtuerp/models"
)

type INumberSettingRepository interface {
	Search(attr models.NumberSetting, orders string) (models.NumberSetting, error)
	//查询历史流水号规则
	//订单可以进行补录,不同补录时间，对应的当前流水号存在不一致
	Create(numberSetting models.NumberSetting) (models.NumberSetting, error)
	Find(per, page int, filter map[string]interface{}, selectKeys []string, order []string, isCount bool) (numberSettings []models.ResponseNumberSetting,
		total int64, err error)
	Delete(id uint) error
	//生成订单号
	GenerateOrderNo(create time.Time) (string, error)
}

type NumberSettingRepository struct {
	BaseRepository
}

//查询历史流水号规则
func (n NumberSettingRepository) FindOrCreateHistory(attr models.NumberSettingHistory) (record models.NumberSettingHistory, err error) {
	err = database.GetDBCon().FirstOrCreate(&record, attr).Error
	return record, err
}

func (n NumberSettingRepository) UpdateNumberHistory(history models.NumberSettingHistory, attr map[string]interface{}) error {
	return database.GetDBCon().Model(&history).Updates(attr).Error
}

func (n NumberSettingRepository) Search(attr models.NumberSetting, order string) (result models.NumberSetting, err error) {
	if order == "" {
		order = "id desc"
	}
	err = database.GetDBCon().Where(&attr).Order(order).First(&result).Error
	return
}

func (n NumberSettingRepository) Delete(id uint) error {
	return n.crud.Delete(&models.NumberSetting{}, id)
}

func (n NumberSettingRepository) Find(per, page int, filter map[string]interface{}, selectKeys []string, order []string, isCount bool) (numberSettings []models.ResponseNumberSetting,
	total int64, err error) {
	var rows *sql.Rows
	sqlCon := database.GetDBCon().Model(&models.NumberSetting{}).Scopes(n.defaultJoin)
	if isCount {
		totalCon := database.GetDBCon().Model(&models.NumberSetting{}).Scopes(n.defaultJoin)
		if total, err = n.Count(totalCon, filter); err != nil {
			return
		}
	}
	if len(selectKeys) == 0 {
		selectKeys = []string{"number_settings.*", "user_companies.name_nick as user_companies_name_nick",}
	}
	rows, err = sqlCon.Scopes(n.CustomerWhere(filter, selectKeys, n.Paginate(per, page), n.OrderBy(order))).Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var data models.ResponseNumberSetting
		_ = sqlCon.ScanRows(rows, &data)
		numberSettings = append(numberSettings, data)
	}
	return
}

func (n NumberSettingRepository) defaultJoin(db *gorm.DB) *gorm.DB {
	return db.Joins("inner join user_companies on user_companies.id = number_settings.user_company_id and user_companies.company_type = 4")
}

func (n NumberSettingRepository) Create(numberSetting models.NumberSetting) (models.NumberSetting, error) {
	err := n.crud.Create(&numberSetting)
	return numberSetting, err
}

//对当前current_number进行更新
func (n NumberSettingRepository) UpdateNumber(setting models.NumberSetting, attr map[string]interface{}) error {
	return database.GetDBCon().Model(&setting).Updates(attr).Error
}

type searchNumber func(numberSetting models.NumberSetting, create time.Time) (int, error)

func (n NumberSettingRepository) GenerateOrderNo(create time.Time) (string, error) {
	var (
		err           error
		numberSetting models.NumberSetting
		currentNumber string // 当前流水号长度
	)
	numberSetting, err = n.Search(models.NumberSetting{ApplicationNo: models.NumberSettingOrderNumber}, "")
	if err != nil {
		return currentNumber, err
	}
	currentNumber = numberSetting.Prefix // 获取当前流水号前缀
	//查询当前流水号
	if numberSetting.ClearRule != models.NumberSettingNonZero {
		return n.generateNumber(numberSetting, create, n.searchNumberFromSetting)
	} else {
		return n.generateNumber(numberSetting, create, n.searchNumberFromHistory)
	}
}

//组装流水号
func (n NumberSettingRepository) generateNumber(setting models.NumberSetting, create time.Time,
	searchData searchNumber) (string, error) {
	var (
		ruleArray     []string
		buffer        strings.Builder
		err           error
		currentNumber int
	)
	ruleArray = strings.Split(setting.DefaultRule, ",")
	buffer.Grow(len(ruleArray) * 3)
	for _, item := range ruleArray {
		switch item {
		case "YY":
			buffer.WriteString(create.Format("06"))
		case "YYYY":
			buffer.WriteString(create.Format("2006"))
		case "MM":
			month := strconv.Itoa(int(create.Month()))
			if len(month) == 1 {
				month = "0" + month
			}
			buffer.WriteString(month)
		case "DD":
			day := strconv.Itoa(int(create.Day()))
			if len(day) == 1 {
				day = "0" + day
			}
			buffer.WriteString(day)
		case "special":
			buffer.WriteString(setting.Special)
		default:
			currentNumber, err = searchData(setting, create)
			if err != nil {
				return "", err
			}
			buffer.WriteString(fmt.Sprintf("%0*d", len(item), currentNumber))
		}
	}
	golog.Infof("current number is %v\n", buffer.String())
	return buffer.String(), nil
}

//从历史订单中查询订单号
func (n NumberSettingRepository) searchNumberFromHistory(setting models.NumberSetting, create time.Time) (int, error) {
	var (
		history = models.NumberSettingHistory{
			NumberSettingId: setting.ID,
		}
		err           error
		currentNumber int
	)
	clearMethod := setting.ClearRule
	//根据不同的清零方式，查询不同的历史流水号，获取当前最新的流水号
	switch clearMethod {
	case models.NumberSettingYearClear:
		history.Year = create.Year()
	case models.NumberSettingMonthClear:
		history.Year = create.Year()
		history.Month = int(create.Month())
	case models.NumberSettingDayClear:
		history.Year = create.Year()
		history.Month = int(create.Month())
		history.Day = create.Day()
	}
	history, err = n.FindOrCreateHistory(history)
	if err != nil {
		return 0, err
	}
	//说明是第一个
	if history.CurrentNumber == 0 {
		currentNumber = setting.DefaultNumber + 1
	} else {
		currentNumber = history.CurrentNumber + 1
	}
	err = n.UpdateNumberHistory(history, map[string]interface{}{
		"current_number": currentNumber,
	})
	if err != nil {
		return 0, err
	} else {
		return currentNumber, nil
	}
}

func (n NumberSettingRepository) searchNumberFromSetting(setting models.NumberSetting, create time.Time) (int, error) {
	currentNumber := setting.CurrentNumber
	if currentNumber == 0 {
		currentNumber = setting.DefaultNumber + 1
	} else {
		currentNumber = setting.CurrentNumber + 1
	}
	err := n.UpdateNumber(setting, map[string]interface{}{
		"current_number": currentNumber,
	})
	if err != nil {
		return 0, err
	} else {
		return currentNumber, err
	}
}

func NewNumberSetting() INumberSettingRepository {
	return &NumberSettingRepository{}
}
