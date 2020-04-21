package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type IBaseRepository interface {
	//将map函数解析成对应的where查询条件
	Ransack(db *gorm.DB, selectColumn map[string]interface{}) *gorm.DB
}

var RexGrep = [...]string{"gt", "gtEq", "lt", "ltEq", "in",
	"eq", "notEq", "cont", "lCont", "rCont"}

type BaseRepository struct {
}

func (b BaseRepository) Ransack(db *gorm.DB, selectColumn map[string]interface{}) *gorm.DB {
	var (
		searchString string
		searchValue  []interface{}
		count        = 0
	)
	for k, v := range selectColumn {
		splitArray := strings.Split(k, "-")
		if len(splitArray) > 3 {
			continue
		}
		n := len(splitArray) - 1
		item := splitArray[n]
		splitArray = splitArray[:n]
		fmt.Println(splitArray)
		if !b.isExist(item) {
			continue
		}
		count++
		key := b.keyString(splitArray)
		fmt.Println(key)
		switch item {
		case "gt":
			searchString += key + " > ? "
			searchValue = append(searchValue, v)
		case "gtEq":
			searchString += key + " >= ?"
			searchValue = append(searchValue, v)
		case "lt":
			searchString += key + " < ?"
			searchValue = append(searchValue, v)
		case "ltEq":
			searchString += key + " <=?"
			searchValue = append(searchValue, v)
		case "in":
			searchString += key + " IN ?"
			searchValue = append(searchValue, v)
		case "eq":
			searchString += key + " = ?"
			searchValue = append(searchValue, v)
		case "notEq":
			searchString += key + " <> ?"
			searchValue = append(searchValue, v)
		case "cont":
			searchString += key + " LIKE ?"
			searchValue = append(searchValue, "%"+v.(string)+"%")
		case "lCount":
			searchString += key + " LIKE ?"
			searchValue = append(searchValue, "%"+v.(string))
		case "rCount":
			searchString += key + " LIKE ?"
			searchValue = append(searchValue, v.(string)+"%")
		}
		if count < len(selectColumn) {
			searchString += " and "
		}
		
	}
	fmt.Println(searchValue, searchString)
	db = db.Where(searchString, searchValue)
	return db
}

//元素是否存在
func (b BaseRepository) isExist(dst string) bool {
	for _, item := range RexGrep {
		if item == dst {
			return true
		}
	}
	return false
}

func (b BaseRepository) keyString(splitArray []string) string {
	if len(splitArray) == 1 {
		return splitArray[0]
	} else {
		return splitArray[0] + splitArray[1]
	}
}
