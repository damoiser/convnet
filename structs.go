package convnet

type ConvNet struct {
  steps []*Step
  input [][]int32
}

type Step struct {
  Kind string // e.g. ReLU, Pooling, Softmax,...

  Filter *Filter
  Stride int32 // the matrix movement of the filter

  Result [][]int32 // the result for the next iteration
}

type Filter struct {
  Kind string // Kind of the filter like [[1,0,1], [0,1,0], [1,0,1]]
  Content [][]int32

  SizeX int32 // helper for x
  SizeY int32 // helper for y
}
