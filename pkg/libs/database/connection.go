package database

import (
	"reflect"

	"gorm.io/gorm"
)

// DBConnection struct
type DBConnection struct {
	connection        *gorm.DB
	tablesInfoByTypes map[reflect.Type]tableInfo
	tablesInfoByName  map[string]tableInfo
}
