package errors

// LevelType type
type LevelType int

// Level constants
const (
	LevelDebug LevelType = 1
	LevelInfo  LevelType = 2
	LevelWarn  LevelType = 3
	LevelError LevelType = 4
)

// Level function
func Level(err error) LevelType {
	e, ok := err.(*Error)
	if !ok {
		return LevelError
	}

	if e.level != 0 {
		return e.level
	}

	return Level(e.err)
}

func (l LevelType) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return "----"
	}
}
