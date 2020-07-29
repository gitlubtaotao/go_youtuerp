package global

import "gorm.io/gorm"

var dataEngine *gorm.DB

//global database connection
func GetDBCon() *gorm.DB {
	return dataEngine
}

