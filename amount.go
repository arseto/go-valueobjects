package valueobjects

type Amount float64

func (sub *Amount) Discount(discountPercent float64) {
	disc := sub.FractionByPercentage(discountPercent)
	*sub -= Amount(disc)
}

func (a *Amount) FractionByPercentage(percent float64) Amount {
	frac := float64(*a) * (percent / 100)
	return Amount(frac)
}
