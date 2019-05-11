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

// StepPooling add the pooling
func StepPooling() *Step {
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
