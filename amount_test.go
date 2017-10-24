package valueobjects

import "testing"

func TestAmountFractionByPercentage(t *testing.T) {
	amount := Amount(1000)

	actual := amount.FractionByPercentage(10)
	expected := Amount(100)

	if actual != expected {
		t.Errorf("Expected fraction by percentage: %g, found %g",
			expected, actual,
		)
	}

	actual = amount.FractionByPercentage(20)
	expected = Amount(200)

	if actual != expected {
		t.Errorf("Expected fraction by percentage: %g, found %g",
			expected, actual,
		)
	}
}

func TestAmountDiscount(t *testing.T) {
	amount := Amount(1000)
	amount.Discount(10)

	expected := Amount(900)

	if amount != expected {
		t.Errorf("Expected Amount after discount: %g, found %g",
			expected, amount,
		)
	}

	amount.Discount(20)
	expected = Amount(720)

	if amount != expected {
		t.Errorf("Expected Amount after discount: %g, found %g",
			expected, amount,
		)
	}
}
