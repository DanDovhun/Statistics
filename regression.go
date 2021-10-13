package statistics

import (
	"errors"
	"math"
)

func Sum(lst []float64) float64 {
	var sum float64

	for _, i := range lst {
		sum += i
	}

	return sum
}

func SquareSum(arr []float64) float64 {
	var (
		sum       float64 = Sum(arr)
		sumSqr    float64
		squareSum float64
	)

	for _, i := range arr {
		sumSqr += i * i
	}

	squareSum = sumSqr - math.Pow(sum, 2)/float64(len(arr))

	return squareSum
}

func Correlation(X []float64, Y []float64) (float64, error) {
	if len(X) != len(Y) {
		return 0, errors.New("Error: X and Y aren't equal")
	}

	var (
		Sxx, Syy, Sxy float64 = SquareSum(X), SquareSum(Y), 0
		sumX, sumY    float64 = Sum(X), Sum(Y)
		prodXY        []float64
	)

	for i := 0; i < len(X); i++ {
		prodXY = append(prodXY, X[i]*Y[i])
	}

	Sxy = Sum(prodXY) - (sumX*sumY)/float64(len(X))

	return Sxy / math.Sqrt(Sxx*Syy), nil
}

func LinearRegression(X []float64, Y []float64) (float64, float64, error) {
	if len(X) != len(Y) {
		return 0, 0, errors.New("X and Y aren't equal")
	}

	var (
		Sxx, Sxy     float64 = SquareSum(X), 0
		meanX, meanY float64 = Mean(X), Mean(Y)
		sumX, sumY   float64 = Sum(X), Sum(Y)
		grad, inter  float64
		prodXY       []float64
	)

	for i := 0; i < len(X); i++ {
		prodXY = append(prodXY, X[i]*Y[i])
	}

	Sxy = Sum(prodXY) - (sumX*sumY)/float64(len(X))

	grad = Sxy / Sxx
	inter = meanY - meanX*grad

	return grad, inter, nil
}
