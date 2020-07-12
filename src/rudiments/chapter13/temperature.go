package temperature

func CtoF(c float64) float64 {
	return (c * (9 / 5)) + 32
}

func FtoC(c float64) float64 {
	return (c - 32) * (9 / 5)
}
