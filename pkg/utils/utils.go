package utils

import (
	"math"
	"math/rand/v2"
)

func GenerateRandomFloat(scale, decimalPlaces int) float64 {
	randNumber := rand.Float64() * float64(scale)
	return math.Round(randNumber*100*float64(decimalPlaces)) / 100 * float64(decimalPlaces)
}

func RandomSliceIndex(sliceLenght int) int {
	return int(math.Round(rand.Float64() * float64(sliceLenght-1)))
}
