package errors

import (
	"fmt"

	"go.uber.org/zap"
)

// Error struct
type Error struct {
	level  LevelType
	op     Operation
	kind   KindType
	err    error
	fields []zap.Field
}

// Error function
func (e *Error) Error() string {
	str := fmt.Sprintf("%s [%s]: %s", e.level, e.kind, e.op)
	if e.err != nil {
		str += fmt.Sprintf(" - %v", e.err.Error())
	}
	return str
}

// BaseError function
func (e *Error) BaseError() string {
	subError, ok := e.err.(*Error)
	if !ok {
		return e.err.Error()
	}

	return subError.BaseError()
}

// New function
func (e *Error) New(op Operation, fields ...zap.Field) *Error {
	return new(e.level, op, e.kind, e, fields...)
}

// NewDebug function
func (e *Error) NewDebug(op Operation, fields ...zap.Field) *Error {
	return new(LevelDebug, op, e.kind, e, fields...)
}

// NewInfo function
func (e *Error) NewInfo(op Operation, fields ...zap.Field) *Error {
	return new(LevelInfo, op, e.kind, e, fields...)
}

// NewWarn function
func (e *Error) NewWarn(op Operation, fields ...zap.Field) *Error {
	return new(LevelWarn, op, e.kind, e, fields...)
}

// NewError function
func (e *Error) NewError(op Operation, fields ...zap.Field) *Error {
	return new(LevelError, op, e.kind, e, fields...)
}

func new(level LevelType, op Operation, kind KindType, err error, fields ...zap.Field) *Error {
	if err == nil {
		panic("errors.New: err cannot be nil")
	}

	for i := range fields {
		fields[i].Key = string(op) + ":" + fields[i].Key
	}

	return &Error{
		level:  level,
		op:     op,
		kind:   kind,
		err:    err,
		fields: fields,
	}
}

// NewDebug function
func NewDebug(op Operation, kind KindType, err error, fields ...zap.Field) *Error {
	return new(LevelDebug, op, kind, err, fields...)
}

// NewInfo function
func NewInfo(op Operation, kind KindType, err error, fields ...zap.Field) *Error {
	return new(LevelInfo, op, kind, err, fields...)
}

// NewWarn function
func NewWarn(op Operation, kind KindType, err error, fields ...zap.Field) *Error {
	return new(LevelWarn, op, kind, err, fields...)
}

// NewError function
func NewError(op Operation, kind KindType, err error, fields ...zap.Field) *Error {
	return new(LevelError, op, kind, err, fields...)
}

func QueryArgs(key string, values []interface{}) zap.Field {
	newValue := ""
	for _, value := range values {
		if newValue != "" {
			newValue += ", "
		}

		strValue := fmt.Sprintf("%v", value)
		if strValue != "NULL" {
			strValue = "'" + strValue + "'"
		}
		newValue += strValue
	}
	return zap.String(key, newValue)
}
