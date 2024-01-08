package set_test

import (
	"tdd/03_set_gingko_gomega/set"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var s *set.Set

var _ = Describe("Set", func() {
	BeforeEach(func() {
		s = set.NewSet()
	})

	Describe("Emptiness", func() {
		Context("when the set is empty", func() {
			It("should be empty", func() {
				Expect(s.IsEmpty()).To(BeTrue())
			})
		})
		Context("when the set is empty", func() {
			It("should be empty", func() {
				s.Add("red")
				Expect(s.IsEmpty()).To(BeFalse())
			})
		})
	})

	Describe("Size", func() {
		Context("As the items are added", func() {
			It("should return an increasing size", func() {
				By("Empty set size being 0", func() {
					Expect(s.Size()).To(BeZero())
				})
				By("Adding first item", func() {
					s.Add("red")
					Expect(s.Size()).To(Equal(1))
				})
				By("Adding Second item", func() {
					s.Add("blue")
					Expect(s.Size()).To(Equal(2))
				})
			})
		})
	})

	Describe("Contains", func() {
		Context("When red is not added", func() {
			It("should not contain red", func() {
				Expect(s.Contains("red")).To(BeFalse())
			})
		})
		Context("When red is added", func() {
			It("should contain red", func() {
				s.Add("red")
				Expect(s.Contains("red")).To(BeTrue())
			})
		})
	})

})
