package convnet

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Matrix", func() {
	Context("when testing matrix multiplication", func() {
		It("calculates correctly (1st)", func() {
			matrix1 := Matrix{
				content: [][]float64{
					[]float64{1.0, 2.0, -1.0},
					[]float64{2.0, 0.0, 1.0},
				},
			}

			matrix2 := Matrix{
				content: [][]float64{
					[]float64{3.0, 1.0},
					[]float64{0.0, -1.0},
					[]float64{-2.0, 3.0},
				},
			}

			res, err := matrix1.Multiply(&matrix2)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.content[0]).To(Equal([]float64{5.0, -4.0}))
			Expect(res.content[1]).To(Equal([]float64{4.0, 5.0}))
		})

		It("returns an error if the matrix don't have the correct shapes", func() {
			matrix1 := Matrix{
				content: [][]float64{
					[]float64{1.0, 2.0, -1.0},
					[]float64{2.0, 0.0, 1.0},
				},
			}

			matrix2 := Matrix{
				content: [][]float64{
					[]float64{3.0, 1.0},
					[]float64{0.0, -1.0},
				},
			}

			res, err := matrix1.Multiply(&matrix2)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("matrix cols are not equals to matrix2 rows"))
			Expect(res).To(BeNil())
		})
	})

	Context("when testing matrix transposion", func() {
		It("transpose correctly", func() {
			matrix := Matrix{
				content: [][]float64{
					[]float64{-1.0, 2.0},
					[]float64{3.0, 4.0},
					[]float64{5.0, -6.0},
					[]float64{7.0, 8.0},
					[]float64{9.0, 10.0},
				},
			}

			res := matrix.Transpose()
			Expect(res).NotTo(BeNil())
			Expect(len(res.content)).To(Equal(2))
			Expect(res.content[0]).To(Equal([]float64{-1.0, 3.0, 5.0, 7.0, 9.0}))
			Expect(res.content[1]).To(Equal([]float64{2.0, 4.0, -6.0, 8.0, 10.0}))
		})
	})

	Context("when testing matrix flattening", func() {
		It("flattens correctly", func() {
			matrix := Matrix{
				content: [][]float64{
					[]float64{-1.0, 2.0, 3.0},
					[]float64{4.0, 5.0, -6.0},
					[]float64{7.0, 8.0, 9.0},
				},
			}

			res := matrix.Flattening()
			Expect(res).NotTo(BeNil())
			Expect(len(res.content)).To(Equal(1))
			Expect(res.content[0]).To(Equal([]float64{-1.0, 2.0, 3.0, 4.0, 5.0, -6.0, 7.0, 8.0, 9.0}))
		})
	})
})
