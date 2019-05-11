package convnet

import (
	"errors"
)

// auxiliaries funcs for math

// minAndMaxOfSlice returns the min and the max of a slice
func minAndMaxOfSlice(slice []float64) (float64, float64, error) {
	if len(slice) == 0 {
		return 0.0, 0.0, errors.New("slice is empty")
	}

	max := slice[0]
	min := slice[0]
	for _, val := range slice {
		if max < val {
			max = val
		}

		if min > val {
			min = val
		}
	}

	return min, max, nil
}
