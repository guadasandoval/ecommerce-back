package errors

// Is function
func Is(err error, target error) bool {
	var errCompare error
	errCompare = err

	for {
		if errCompare == target {
			return true
		}

		eerrCompare, ok := errCompare.(*Error)
		if !ok {
			return errCompare == target
		}

		if eerrCompare.err == nil {
			return false
		}
		errCompare = eerrCompare.err
	}
}

// IsKind function
func IsKind(err error, target KindType) bool {
	return Kind(err) == target
}
