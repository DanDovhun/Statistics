package statistics

import "math"

func Sum(lst []float64) float64 {
	var sum float64

	for _, i := range lst {
		sum += i
	}

	return sum
}

func SquareSum(lst []float64) float64 {
	var sum, sumSqr, square float64

	sum = Sum(lst)

	for _, i := range lst {
		sumSqr += i * i
	}

	square = sumSqr - (math.Pow(sum, 2) / float64(len(lst)))

	return square
}

func Correlation(X []float64, Y []float64) float64 {
	var (
		x, y, r float64
		prodXY  float64
		xy      []float64
	)

	for i := 0; i < len(X); i++ {
		xy = append(xy, X[i], Y[i])
	}

	x = SquareSum(X)
	y = SquareSum(Y)
	prodXY = SquareSum(xy)

	r = prodXY / math.Sqrt(x*y)

	return r
}

func LinearRegression(X []float64, Y []float64) (float64, float64) {
	var (
		x, y   float64
		a, b   float64
		prodXY float64
		xy     []float64
	)

	for i := 0; i < len(X); i++ {
		xy = append(xy, X[i], Y[i])
	}

	prodXY = SquareSum(xy)
	x = SquareSum(X)
	y = SquareSum(Y)

	a = prodXY / math.Sqrt(x*y)
	b = Mean(Y) - a*Mean(X)

	return a, b
}
