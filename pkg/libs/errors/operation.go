package errors

// Operation type
type Operation string

// Operations function
func Operations(err error) []Operation {
	ret := []Operation{}
	for {
		e, ok := err.(*Error)
		if !ok {
			return ret
		}

		ret = append(ret, e.op)
		err = e.err
	}
}
