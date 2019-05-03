package convnet

func New(input [][]int32) (*ConvNet, error) {
  // do init

  conv := &ConvNet{
    steps: []*Steps{baseStep(input)},
    input: input,
  }

  return conv
}

func (conv *ConvNet) Train() error {

}

func (conv *ConvNet) Run() error {

}