package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"strings"
	"youtuerp/database"
)

var RexGrep = [...]string{"gt", "gtEq", "lt", "ltEq", "in",
	"eq", "notEq", "cont", "lCount", "rCount"}

type IBaseRepository interface {
	//将map函数解析成对应的where查询条件
	Ransack(selectColumn map[string]interface{}) func(db *gorm.DB) *gorm.DB
	//分页方法
	Paginate(per, page uint) func(db *gorm.DB) *gorm.DB
	//进行数据的排序
	OrderBy(orders []string) func(db *gorm.DB) *gorm.DB
}

type BaseRepository struct {
	crud
}

//抽象查询方法
func (b BaseRepository) Ransack(selectColumn map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return b.crud.ransack(selectColumn)
}

//分页方法
//page = 0 表示不进行分页
//默认的分页方法
func (b BaseRepository) Paginate(per, page uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 && per > 0 {
			return db.Limit(per)
		} else if page > 0 && per > 0 {
			return db.Limit(per).Offset((page - 1) * per)
		} else {
			return db
		}
	}
}

//order by
//进行排序
func (b BaseRepository) OrderBy(orders []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(orders) == 0 {
			return db.Order("id desc")
		}
		for _, order := range orders {
			db = db.Order(order)
		}
		return db
	}
}

type crud struct {
}

func (c crud) Delete(value interface{}, id uint) error {
	return database.GetDBCon().Delete(value, "id = ?", id).Error
}

func (c crud) First(value interface{}, id uint) error {
	return database.GetDBCon().First(value, "id = ?", id).Error
}

func (c crud) Create(value interface{}) error {
	return database.GetDBCon().Create(value).Error
}

func (c crud) Where(sqlCon *gorm.DB, filter map[string]interface{}, selectKeys []string, funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	if len(filter) > 0 {
		sqlCon = sqlCon.Scopes(c.ransack(filter))
	}
	sqlCon = sqlCon.Scopes(funcs...)
	sqlCon = sqlCon.Select(selectKeys)
	return sqlCon
}

func (c crud) Count(sqlCon *gorm.DB, filter map[string]interface{}, funcs ...func(*gorm.DB) *gorm.DB) (count uint, err error) {
	if len(filter) > 0 {
		sqlCon = sqlCon.Scopes(c.ransack(filter))
	}
	sqlCon = sqlCon.Scopes(funcs...)
	err = sqlCon.Count(&count).Error
	return
}

func (c crud) ransack(selectColumn map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for k, v := range selectColumn {
			if c.notSearchValue(v) {
				continue
			}
			splitArray := strings.Split(k, "-")
			if len(splitArray) > 3 {
				continue
			}
			n := len(splitArray) - 1
			item := splitArray[n]
			splitArray = splitArray[:n]
			if !c.isExist(item) {
				continue
			}
			key := c.keyString(splitArray)
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
				fmt.Println("33333")
				db = db.Where(key+" LIKE ? ", v.(string)+"%")
			}
		}
		return db
	}
}

//元素匹配
func (c crud) isExist(dst string) bool {
	for _, item := range RexGrep {
		if item == dst {
			return true
		}
	}
	return false
}

//返回sql 对应的column
func (c crud) keyString(splitArray []string) string {
	if len(splitArray) == 1 {
		return splitArray[0]
	} else {
		return splitArray[0] + splitArray[1]
	}
}

//空值不进行查询处理
func (c crud) notSearchValue(value interface{}) bool {
	v := reflect.ValueOf(value)
	kind := v.Kind()
	if kind == reflect.Array || kind == reflect.Struct || kind == reflect.Slice {
		return value == nil
	}
	if kind == reflect.Bool {
		return false
	}
	return value == ""
}
