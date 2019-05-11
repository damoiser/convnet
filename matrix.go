package convnet

// various functions for matrix operations
// trivial maths, they could be improved!

type Matrix struct {
	content [][]int32
}

// Multiply returns the result matrix calculation from matrix1 * matrix2
// or error if the shapes are not correct
func (matrix1 *Matrix) Multiply(matrix2 *Matrix) (error, *Matrix) {

}

func (matrix *Matrix) Flattern() (error, *Matrix) {
	// TODO
	return nil, nil
}

func (matrix *Matrix) Transpose() (error, *Matrix) {
	// TODO
	return nil, nil
}
