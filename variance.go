package statistics

import (
	"math"
)

func Variance(arr []float64) float64 {
	var sumX, sumXSqr, mean, variance float64

	for _, i := range arr {
		sumX += i
		sumXSqr += i * i
	}

	mean = sumX / float64(len(arr))
	variance = sumXSqr/float64(len(arr)) - math.Pow(mean, 2)

	return variance
}

func SDev(arr []float64) float64 {
	return math.Sqrt(Variance(arr))
}
