package validation

import "fmt"

// Error struct
type Error struct {
	FieldName string
	RuleID    Rule
	RuleIDStr string
	Params    []string
	Value     interface{}
}

// Description function
func (e *Error) Description() string {
	// const op errors.Operation = "pkg.libs.validation.Error.Description"

	switch e.RuleID {
	case List:
	case Struct:
	case Required:
		return fmt.Sprintf("Campo requerido '%s'", e.FieldName)
	case Same:
	case Alpha:
	case Numeric:
	case AlphaNumeric:
	case Int:
	case IntNotRequired:
	case MinValue:
	case MaxValue:
	case BetweenValue:
	case Length:
	case MinLength:
		return fmt.Sprintf("El campo '%s' debe ser al menos '%v' caracteres", e.FieldName, e.Params[0])
	case MaxLength:
		return fmt.Sprintf("El campo '%s' es muy largo", e.FieldName)
	case Date:
		return fmt.Sprintf("El campo '%s' es una fecha inválida", e.FieldName)
	case Time:
		return fmt.Sprintf("El campo '%s' es una hora inválida", e.FieldName)
	case DateTime:
		return fmt.Sprintf("El campo '%s' es una fecha inválida", e.FieldName)
	case EMail:
		return fmt.Sprintf("El campo '%s' es correo electrónico inválida", e.FieldName)
	case In:
	case DBExists:
	case DBNotExists:
	case DBUnique:
	case OptionParam:
	case OptionParamNotRequired:
	case OptionParamList:
	case OptionParamListNotRequired:
	}

	return fmt.Sprintf("El campo '%s' es inválido", e.FieldName)
}
