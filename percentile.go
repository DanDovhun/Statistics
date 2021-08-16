package statistics

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Percentile(arr []float64, per float64) float64 {
	sort.Float64s(arr)

	var (
		ri, nx, r float64
		split     []string
	)

	per /= 100

	r = float64(len(arr)-1)*per + 1
	split = strings.Split(fmt.Sprintf("%v", r), ".")
	pos, _ := strconv.ParseInt(split[0], 10, 64)

	ri = arr[pos-1]
	nx = arr[pos]

	rf := r - float64(pos)

	return ri + rf*(nx-ri)
}

func Quartile(arr []float64, quart int64) (float64, error) {
	switch quart {
	case 1:
		return Percentile(arr, 25), nil

	case 2:
		return Percentile(arr, 50), nil

	case 3:
		return Percentile(arr, 75), nil

	default:
		return 0, errors.New("invalid quartile")
	}
}

func FindPercentile(arr []float64, num float64) float64 {
	sort.Float64s(arr)
	var (
		pos       int64
		ri, nx, q float64
	)

	for i, j := range arr {
		if j > num {
			if i != 0 {
				pos = int64(i) - 1
			} else {
				pos = int64(i)
			}
			ri = arr[int(pos)]

			break
		}
	}

	nx = arr[pos+1]

	num -= ri
	num /= (nx - ri)

	q = 100 * (float64(pos) + num) / (float64(len(arr)) - 1)

	return q
}
