package database

import "reflect"

type tableInfoFields struct {
	Name          string
	FieldType     reflect.Type
	AutoIncrement bool
	Ignore        bool
	AutoValue     bool
	Tag           reflect.StructTag
}
type tableInfo struct {
	Name               string
	Prefix             string
	ModelType          reflect.Type
	Fields             []tableInfoFields
	AutoIncrementIndex int
}
