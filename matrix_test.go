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
})
