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
		key := b.keyString(splitArray)
		fmt.Println(key)
		switch item {
		case "gt":
			db = db.Where(key+" > ? ", v)
		case "gtEq":
			db = db.Where(key+" >= ? ", v)
		case "lt":
			db = db.Where(key+" < ? ", v)
		case "ltEq":
			db = db.Where(key+" <= ?", v)
		case "in":
			db = db.Where(key+" IN (?)", v)
		case "eq":
			db = db.Where(key+" = ? ", v)
		case "notEq":
			db = db.Where(key+" <> ? ", v)
		case "cont":
			db = db.Where(key+" LIKE ? ", "%"+v.(string)+"%")
		case "lCount":
			db = db.Where(key+" LIKE ? ", "%"+v.(string))
		case "rCount":
			db = db.Where(key+" LIKE ? ", v.(string)+"%")
		}
	}
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
