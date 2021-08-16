package statistics

import (
	"errors"
	"fmt"
	"math"
)

func Round(num float64) float64 {
	return math.Round(10000*num) / 10000
}

func SquareSum(arr []float64) float64 {
	var sum, sumSqr, n float64

	n = float64(len(arr))

	for _, i := range arr {
		sum += i
		sumSqr += i * i
	}

	return sumSqr - (math.Pow(sum, 2) / n)
}

func correlationIndex(x, y []float64) (float64, error) {
	var (
		n, sumXY, sqrXY  float64 = float64(len(x)), 0, 0
		r, sumX, sumXSqr float64
		sumY, sumYSqr    float64
		sqrXX, sqrYY     float64
	)

	if len(x) != len(y) {
		return 0, errors.New("error: lists are not equal length")
	} else {
		for i := 0; i < len(x); i++ {
			sumX += x[i]
			sumXSqr += x[i] * x[i]

			sumY += y[i]
			sumYSqr += y[i] * y[i]

			sumXY += x[i] * y[i]
		}

		sqrXX = sumXSqr - (math.Pow(sumX, 2) / n)
		sqrYY = sumYSqr - (math.Pow(sumY, 2) / n)
		sqrXY = sumXY - (sumX * sumY / n)
		r = sqrXY / math.Sqrt(sqrYY*sqrXX)

		return r, nil
	}
}

func Formula(m, c float64) string {
	if m > 0 {
		if c > 0 {
			return fmt.Sprintf("y = %vx + %v", Round(m), Round(c))
		} else if c < 0 {
			return fmt.Sprintf("y = %vx - %v", Round(m), -1*Round(c))
		}

		return fmt.Sprintf("y = %vx", Round(m))
	}

	if c > 0 {
		return fmt.Sprintf("y = %v - %vx", Round(c), -1*Round(m))
	} else if c < 0 {
		return fmt.Sprintf("y = -(%vx + %v)", Round(m), Round(c))
	}

	return fmt.Sprintf("y = -%vx", -1*Round(m))
}

func Regression(x, y []float64) (string, float64, float64, error) {
	if len(x) != len(y) {
		return "", 0, 0, errors.New("error: lists are not equal length")
	} else {
		var (
			n, m, c, xM, yM float64
			sumXY, sqrXY    float64
			sumX, sumXSqr   float64
			sumY, sqrXX     float64
		)

		n = float64(len(x))

		xM = Mean(x)
		yM = Mean(y)

		for i := 0; i < len(x); i++ {
			sumX += x[i]
			sumY += y[i]
			sumXSqr += x[i] * x[i]

			sumXY += x[i] * y[i]
		}

		sqrXX = sumXSqr - (math.Pow(sumX, 2) / n)
		sqrXY = sumXY - (sumX * sumY / n)

		m = sqrXY / sqrXX
		c = yM - m*xM

		return Formula(m, c), m, c, nil
	}
}

func Mean(x []float64) float64 {
	var sum float64

	for _, i := range x {
		sum += i
	}

	return sum / float64(len(x))
}
