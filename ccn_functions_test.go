package convnet

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CcnFunctions", func() {
	Context("when testing maxpool", func() {
		It("complains if filterSize is invalid", func() {
			matrix := &Matrix{}
			res, err := matrix.MaxPool(0, 2)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("filterSize is invalid"))
			Expect(res).To(BeNil())
		})

		It("complains if stride is invalid", func() {
			matrix := &Matrix{}
			res, err := matrix.MaxPool(2, 0)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("stride is invalid"))
			Expect(res).To(BeNil())
		})

		It("calculates correctly with filter 2x2 and stride 2", func() {
			matrix := &Matrix{
				content: [][]float64{
					[]float64{2.0, 3.0, 4.0, 0.0},
					[]float64{1.0, 5.0, 3.0, 2.0},
					[]float64{0.0, 4.0, 2.0, 3.0},
					[]float64{1.0, 0.0, 6.0, 1.0},
				},
			}

			res, err := matrix.MaxPool(2, 2)
			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
			Expect(res.content[0]).To(Equal([]float64{5.0, 4.0}))
			Expect(res.content[1]).To(Equal([]float64{4.0, 6.0}))
		})

		It("calculates correctly with filter 3x3 and stride 1", func() {
			matrix := &Matrix{
				content: [][]float64{
					[]float64{2.0, 3.0, 4.0, 0.0},
					[]float64{1.0, 5.0, 3.0, 2.0},
					[]float64{0.0, 4.0, 2.0, 3.0},
					[]float64{1.0, 0.0, 6.0, 1.0},
				},
			}

			res, err := matrix.MaxPool(3, 1)
			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
			Expect(res.content[0]).To(Equal([]float64{5.0, 5.0, 4.0, 3.0}))
			Expect(res.content[1]).To(Equal([]float64{6.0, 6.0, 6.0, 3.0}))
			Expect(res.content[2]).To(Equal([]float64{6.0, 6.0, 6.0, 3.0}))
			Expect(res.content[3]).To(Equal([]float64{6.0, 6.0, 6.0, 1.0}))
		})

		It("calculates correctly with filter 3x3 and stride 2", func() {
			matrix := &Matrix{
				content: [][]float64{
					[]float64{2.0, 3.0, 4.0, 0.0},
					[]float64{1.0, 5.0, 3.0, 2.0},
					[]float64{0.0, 4.0, 2.0, 3.0},
					[]float64{1.0, 0.0, 6.0, 1.0},
				},
			}

			res, err := matrix.MaxPool(3, 2)
			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
			Expect(res.content[0]).To(Equal([]float64{5.0, 4.0}))
			Expect(res.content[1]).To(Equal([]float64{6.0, 6.0}))
		})
	})

	Context("when testing ReLU", func() {
		It("applies correctly the max(0,x)", func() {
			matrix := &Matrix{
				content: [][]float64{
					[]float64{2.0, 3.0, 4.0, 0.0},
					[]float64{1.0, -5.0, 3.0, 2.0},
					[]float64{-0.5, 4.2, -2.0, 3.0},
					[]float64{-1.0, 0.0, -6.0, 1.0},
				},
			}

			res, err := matrix.ReLU()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
			Expect(res.content[0]).To(Equal([]float64{2.0, 3.0, 4.0, 0.0}))
			Expect(res.content[1]).To(Equal([]float64{1.0, 0.0, 3.0, 2.0}))
			Expect(res.content[2]).To(Equal([]float64{0.0, 4.2, 0.0, 3.0}))
			Expect(res.content[3]).To(Equal([]float64{0.0, 0.0, 0.0, 1.0}))
		})
	})
})
