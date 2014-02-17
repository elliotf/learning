package learning_test

import (
  "github.com/elliotf/learning"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Learning", func() {
  Describe("Sum", func() {
    Context("When given one number", func() {
      It("returns that number", func() {
        Expect(learning.Sum(2)).To(Equal(2))
      })
    })
    Context("When given two numbers", func() {
      It("returns the sum of those numbers", func() {
        Expect(learning.Sum(1, 2)).To(Equal(3))
      })
    })
  })
})
