package valueobjects

type PaymentType int

const (
	DebitPayment  = 1
	CreditPayment = 2
)

func (p PaymentType) String() string {
	switch p {
	case DebitPayment:
		return "Debit"
	case CreditPayment:
		return "Credit"
	default:
		return ""
	}
}
