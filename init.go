package convnet

func New(input [][]int32) (*ConvNet, error) {
	// do init

	conv := &ConvNet{
		steps: []*Step{baseStep(input)},
		input: input,
	}

	return conv, nil
}

func (conv *ConvNet) Train() error {
	// TODO
	return nil
}

func (conv *ConvNet) Run() error {
	// TODO
	return nil
}
