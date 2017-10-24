package valueobjects

type ValidationType int

const (
	RequiredField  = 1
	InvalidField   = 2
	CannotLessThan = 3
	CannotMoreThan = 4
	MustMatchField = 5
)

func (vt ValidationType) GetMessage(fieldValue string) string {
	switch vt {
	case RequiredField:
		return "is Required"
	case InvalidField:
		return "is Invalid"
	case CannotLessThan:
		return "cannot less than " + fieldValue
	case CannotMoreThan:
		return "cannot more than " + fieldValue
	case MustMatchField:
		return "must be match with " + fieldValue
	default:
		return "Undefined validation"
	}
}
