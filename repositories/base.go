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
	if len(selectKeys) > 0 {
		sqlCon = sqlCon.Select(selectKeys)
	}
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
		if len(selectColumn) == 0 {
			return db
		}
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
	switch v.Kind() {
	case reflect.Array, reflect.Struct, reflect.Slice:
		return value == nil
	case reflect.Bool:
		return value.(bool)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint64:
		return value == 0
	default:
		return value == ""
	}
}

type BulkInsert struct {
	db *gorm.DB
}

func (b *BulkInsert) BatchInsert(objAttr []interface{}) (total int64, err error) {
	if len(objAttr) == 0 {
		return
	}
	mainObj := objAttr[0]
	mainScope := b.db.NewScope(mainObj)
	mainFields := mainScope.Fields()
	quoted := make([]string, 0, len(mainFields))
	for i := range mainFields {
		// If primary key has blank value (0 for int, "" for string, nil for interface ...), skip it.
		// If field is ignore field, skip it.
		if (mainFields[i].IsPrimaryKey && mainFields[i].IsBlank) || (mainFields[i].IsIgnored) {
			continue
		}
		quoted = append(quoted, mainScope.Quote(mainFields[i].DBName))
	}
	placeholdersArr := make([]string, 0, len(objAttr))
	for _, obj := range objAttr {
		scope := b.db.NewScope(obj)
		fields := scope.Fields()
		placeholders := make([]string, 0, len(fields))
		for i := range fields {
			if (fields[i].IsPrimaryKey && fields[i].IsBlank) || (fields[i].IsIgnored) {
				continue
			}
			var vars interface{}
			if (fields[i].Name == "CreatedAt" || fields[i].Name == "UpdatedAt") && fields[i].IsBlank {
				vars = gorm.NowFunc()
			} else {
				vars = fields[i].Field.Interface()
			}
			placeholders = append(placeholders, scope.AddToVars(vars))
		}
		placeholdersStr := "(" + strings.Join(placeholders, ", ") + ")"
		placeholdersArr = append(placeholdersArr, placeholdersStr)
		// add real variables for the replacement of placeholders' '?' letter later.
		mainScope.SQLVars = append(mainScope.SQLVars, scope.SQLVars...)
	}
	mainScope.Raw(fmt.Sprintf("INSERT INTO %s (%s) VALUES %s",
		mainScope.QuotedTableName(),
		strings.Join(quoted, ", "),
		strings.Join(placeholdersArr, ", "),
	))
	if err := mainScope.Exec().DB().Error; err != nil {
		return 0, err
	}
	return mainScope.DB().RowsAffected, nil
}
