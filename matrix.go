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

func (matrix *Matrix) Flattern() (*Matrix, error) {
	// TODO
	return nil, nil
}

// Transpose the input matrix (A^T)
func (matrix *Matrix) Transpose() (*Matrix, error) {
	result := &Matrix{
		content: make([][]float64, len(matrix.content[0])),
	}

	for i := 0; i < len(matrix.content); i++ {
		for j := 0; j < len(matrix.content[0]); j++ {
			result.content[j] = append(result.content[j], matrix.content[i][j])
		}
	}

	return result, nil
}
