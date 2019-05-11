package convnet

// various functions for matrix operations
// trivial maths performance, they could be improved!

import (
	"errors"
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
func (matrix *Matrix) MaxPool(filter_size int, stride int) *Matrix {
	result := &Matrix{
		content: [][]float64{},
	}

	for i := 0; i < len(matrix.content)+stride; i += stride {
		for j := 0; j < len(matrix.content[0])+stride; j += stride {

		}
	}

}
