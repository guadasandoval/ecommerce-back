package errors

import "go.uber.org/zap"

// Fields function
func Fields(err error) []zap.Field {
	e, ok := err.(*Error)
	if !ok {
		return []zap.Field{}
	}

	return append(e.fields, Fields(e.err)...)
}
