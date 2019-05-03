package convnet


func baseStep(input [][]int32) *Step {
  return &Step{
    Result: input,
  }
}

// StepConvolution create a convolutional step
func StepConvolution() *Step {

}

// StepReLU apply the activation ReLU
func StepReLU() *Step {

}

// StepPooling add the pooling
func StepPooling() *Step {

}

// StepFlatten flat the volume to a matrix 1xN
func StepFlatten() *Step {

}

// StepSoftmax add classification
func StepSoftmax() *Step {

}

func (step *Step) process(prev_step *Step) {

}