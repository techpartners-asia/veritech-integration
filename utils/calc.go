package utils

func CalculateVat(price float64) float64 {
	return NumberPrecision((price / 1.1) * 0.1)
}

func CalculateWithNoVat(price float64) float64 {
	return NumberPrecision(price / 1.1)
}
