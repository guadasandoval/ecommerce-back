package errors

// KindType type
type KindType int

// Kind constants
const (
	KindUnexpected        KindType = 1
	KindUnauthorized      KindType = 2
	KindNotFound          KindType = 3
	KindDataConversion    KindType = 4
	KindDataDecode        KindType = 5
	KindDataEncode        KindType = 6
	KindDBQuery           KindType = 7
	KindDBQueryForeignKey KindType = 8
	KindInfoOutdated      KindType = 9
	KindMigracionError    KindType = 10
	KindErrorToFront      KindType = 11
)

// Kind function
func Kind(err error) KindType {
	e, ok := err.(*Error)
	if !ok {
		return KindUnexpected
	}

	if e.kind != 0 {
		return e.kind
	}

	return Kind(e.err)
}

func (k KindType) String() string {
	switch k {
	case KindUnexpected:
		return "Unexpected"
	case KindUnauthorized:
		return "Unauthorized"
	case KindNotFound:
		return "NotFound"
	case KindDataConversion:
		return "DataConversion"
	case KindDataDecode:
		return "DataDecode"
	case KindDataEncode:
		return "DataEncode"
	case KindDBQuery:
		return "DBQuery"
	case KindDBQueryForeignKey:
		return "DBQueryForeignKey"
	case KindMigracionError:
		return "MigracionError"
	default:
		return "-----"
	}
}
