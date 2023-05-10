package errors

import "fmt"

// Msg function
func Msg(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
