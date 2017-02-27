package round

import "math"

func Round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 || frac >= -0.5 && frac < 0{
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}
	t := rounder / pow
	return t
}