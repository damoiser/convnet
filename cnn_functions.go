package convnet

import (
	"errors"
	"math"
)

// MaxPool apply max pooling to the matrix with a max-pool filter of size filter_size (square)
// and a stride (movement of the filter), the filter applied outside the original matrix use as paddings "0s"
func (matrix *Matrix) MaxPool(filterSize int, stride int) (*Matrix, error) {
	if filterSize <= 0 {
		return nil, errors.New("filterSize is invalid")
	}
	if stride <= 0 {
		return nil, errors.New("stride is invalid")
	}

	result := &Matrix{
		content: [][]float64{},
	}

	resIndex := 0

	// iterate through the main matrix-x jumping from stride_size
	for i := 0; i < len(matrix.content); i += stride {

		// create new result row
		result.content = append(result.content, []float64{})

		// iterate through the main matrix-y jumping from stride_size
		for j := 0; j < len(matrix.content[0]); j += stride {

			// avoid out of bound (simulating padding = 0), getting min of current index+filter_size or matrix-len
			endRow := int(math.Min(float64(i+filterSize-1), float64(len(matrix.content)-1)))
			endCol := int(math.Min(float64(j+filterSize-1), float64(len(matrix.content[0])-1)))

			// getting row subset for filter
			subset := matrix.content[i : endRow+1]

			// getting tmp max (first element of the subset)
			max := subset[0][j]

			// iterate through the elements of the subset searching for the max
			for k := 0; k < len(subset); k++ {
				// get max of the subset-array
				_, maxRow, err := minAndMaxOfSlice(subset[k][j : endCol+1])
				panicIfErr(err)

				if max < maxRow {
					max = maxRow
				}
			}

			result.content[resIndex] = append(result.content[resIndex], max)
		}

		// update result index
		resIndex++
	}

	return result, nil
}

// ReLU (rectified linear unit) applies non-satutaring activation functions => f(x) = max(0, x)
// it removes negative value from an activation map setting to 0s
func (matrix *Matrix) ReLU() (*Matrix, error) {
	result := &Matrix{
		content: [][]float64{},
	}

	for i := 0; i < len(matrix.content); i++ {
		for j := 0; j < len(matrix.content[0]); j++ {
			if matrix.content[i][j] < 0.0 {
				result.content[i][j] = 0.0
			} else {
				result.content[i][j] = matrix.content[i][j]
			}
		}
	}

	return result, nil
}
