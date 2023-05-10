package validation

// import (
// 	"fmt"
// 	"reflect"

// 	"traductor.mingeneros.gob.ar/pkg/libs/database"
//)

// // Validation struct
// type Validation struct {
// 	fields []validationField
// 	data   interface{}
// 	errors []Error
// 	db     *database.DBConnection
// }

// // CreateValidation - Create a validation
// func CreateValidation(data interface{}) Validation {
// 	// const op errors.Operation = "pkg.libs.validation.CreateValidation"

// 	dataType := reflect.TypeOf(data)
// 	numFields := dataType.NumField()

// 	var ret Validation
// 	ret.fields = make([]validationField, numFields)
// 	for i := 0; i < numFields; i++ {
// 		var field validationField
// 		field.name = dataType.Field(i).Name
// 		field.fieldType = dataType.Field(i).Type
// 		validationString := dataType.Field(i).Tag.Get("valid")
// 		field.validationRules = parseValidationString(validationString)

// 		field.jsonName = dataType.Field(i).Tag.Get("json")
// 		if field.jsonName == "" {
// 			field.jsonName = field.name
// 		}

// 		ret.fields[i] = field
// 	}
// 	ret.data = data
// 	ret.errors = []Error{}

// 	return ret
// }

// // SetDBConnection - Set database connection for validations
// func (v *Validation) SetDBConnection(db *database.DBConnection) {
// 	// const op errors.Operation = "pkg.libs.validation.Validation.SetDBConnection"

// 	v.db = db
// }

// // Valid - Check if it is valid
// func (v *Validation) Valid() bool {
// 	// const op errors.Operation = "pkg.libs.validation.Validation.Valid"

// 	value := reflect.ValueOf(v.data)

// 	numFields := len(v.fields)
// 	v.errors = make([]Error, 0)
// 	for i := 0; i < numFields; i++ {
// 		valueField := value.Field(i).Interface()

// 		v.errors = append(v.errors, validateValue(valueField, v.fields[i].validationRules, v.fields[i].name,
// 			v.fields[i].jsonName, v.db, v.data)...)
// 	}

// 	return (len(v.errors) == 0)
// }

// // ErrorsDebug - Print the error in a readable way
// func (v *Validation) ErrorsDebug() string {
// 	// const op errors.Operation = "pkg.libs.validation.Validation.ErrorsDebug"

// 	var ret string
// 	ret = "Validation Errors:\n"
// 	for _, err := range v.errors {
// 		ret = ret + fmt.Sprintf("Field: %s Rule: %s Params: %v Value: %v \n", err.FieldName, mapNames[err.RuleID],
// 			err.Params, err.Value)
// 	}
// 	return ret
// }

// // ErrorsInfo - List the errors information for the client
// func (v *Validation) ErrorsInfo() []ErrorInfo {
// 	// const op errors.Operation = "pkg.libs.validation.Validation.ErrorsInfo"

// 	var ret []ErrorInfo
// 	for _, err := range v.errors {
// 		ret = append(ret, err.Info())
// 	}
// 	return ret
// }

// // Errors - Get all the errors
// func (v *Validation) Errors() []Error {
// 	// const op errors.Operation = "pkg.libs.validation.Validation.Errors"

// 	return v.errors
// }

// // FirstError - Get the first error
// func (v *Validation) FirstError() Error {
// 	// const op errors.Operation = "pkg.libs.validation.Validation.FirstError"

// 	if len(v.errors) == 0 {
// 		return Error{}
// 	}
// 	return v.errors[0]
// }

// // IsError - Return if there is an error
// func (v *Validation) IsError(fieldName string, ruleID Rule) bool {
// 	// const op errors.Operation = "pkg.libs.validation.Validation.IsError"

// 	for _, err := range v.errors {
// 		if err.FieldName == fieldName && err.RuleID == ruleID {
// 			return true
// 		}
// 	}
// 	return false
// }