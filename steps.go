package convnet

func baseStep(input [][]int32) *Step {
	return &Step{
		Result: input,
	}
}

// StepConvolution create a convolutional step
func StepConvolution() *Step {
	// TODO
	return nil
}

// StepReLU apply the activation ReLU
func StepReLU() *Step {
	// TODO
	return nil
}

// StepPooling add the pooling => max pooling
// this apply a filter_size (square) to the original matrix and a
// "move" of the filter on the original matrix of stride-distance
func StepPooling(filter_size int, stride int) *Step {
	// TODO
	return nil
}

// StepFlatten flat the volume to a matrix 1xN
func StepFlatten() *Step {
	// TODO
	return nil
}

// StepSoftmax add classification
func StepSoftmax() *Step {
	// TODO
	return nil
}

func (step *Step) process(prev_step *Step) {

}
