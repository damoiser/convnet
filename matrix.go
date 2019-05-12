package convnet

// various functions for matrix operations
// trivial maths performance, they could be improved!

import (
	"errors"
	"math"
)

// Matrix basic matrix struct (float64)
type Matrix struct {
	content [][]float64
}

// Multiply returns the result matrix calculation from matrix * matrix2
// or error if the shapes are not correct
func (matrix *Matrix) Multiply(matrix2 *Matrix) (*Matrix, error) {
	if len(matrix.content[1]) != len(matrix2.content) {
		return nil, errors.New("matrix cols are not equals to matrix2 rows")
	}

	result := &Matrix{
		content: make([][]float64, len(matrix.content[1])),
	}

	for i := 0; i < len(matrix.content); i++ {
		result.content[i] = make([]float64, len(matrix2.content[0]))
		for j := 0; j < len(matrix2.content[0]); j++ {
			for k := 0; k < len(matrix2.content); k++ {
				result.content[i][j] += matrix.content[i][k] * matrix2.content[k][j]
			}
		}
	}

	return result, nil
}

// Flattening flat a matrix NxM to a matrix (1, N+M)
func (matrix *Matrix) Flattening() *Matrix {
	result := &Matrix{
		content: make([][]float64, 1),
	}

	for i := 0; i < len(matrix.content); i++ {
		for j := 0; j < len(matrix.content[i]); j++ {
			result.content[0] = append(result.content[0], matrix.content[i][j])
		}
	}

	return result
}

// Transpose the input matrix (A^T)
func (matrix *Matrix) Transpose() *Matrix {
	result := &Matrix{
		content: make([][]float64, len(matrix.content[0])),
	}

	for i := 0; i < len(matrix.content); i++ {
		for j := 0; j < len(matrix.content[0]); j++ {
			result.content[j] = append(result.content[j], matrix.content[i][j])
		}
	}

	return result
}

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
