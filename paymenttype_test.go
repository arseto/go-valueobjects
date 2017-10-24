package valueobjects

import "testing"

func TestPaymentTypeString(t *testing.T) {
	pt := PaymentType(DebitPayment)
	expected := "Debit"

	if pt.String() != expected {
		t.Errorf("Expected %s, found %s", expected, pt.String())
	}

	pt = PaymentType(CreditPayment)
	expected = "Credit"

	if pt.String() != expected {
		t.Errorf("Expected %s, found %s", expected, pt.String())
	}

}
