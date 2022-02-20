package optioner_test

import (
	"github.com/boundedinfinity/optioner"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Constructors", func() {
	Describe("Some()", func() {
		It("should work with a string", func() {
			actual := optioner.Some("s")

			Expect(actual.Empty()).To(BeFalse())
			Expect(actual.Defined()).To(BeTrue())
			Expect(actual.Get()).To(Equal("s"))
			Expect(actual.OrElse("x")).To(Equal("s"))
		})

		It("should work with an int", func() {
			actual := optioner.Some(1)

			Expect(actual.Empty()).To(BeFalse())
			Expect(actual.Defined()).To(BeTrue())
			Expect(actual.Get()).To(Equal(1))
			Expect(actual.OrElse(2)).To(Equal(1))
		})

		It("should work with a boolean", func() {
			actual := optioner.Some(false)

			Expect(actual.Empty()).To(BeFalse())
			Expect(actual.Defined()).To(BeTrue())
			Expect(actual.Get()).To(Equal(false))
			Expect(actual.OrElse(true)).To(Equal(false))
		})
	})

	Describe("None()", func() {
		It("should work with a string", func() {
			actual := optioner.None[string]()

			Expect(actual.IsEmpty()).To(BeTrue())
			Expect(actual.Defined()).To(BeFalse())
			Expect(actual.Get()).To(Equal(""))
			Expect(actual.OrElse("x")).To(Equal("x"))
		})

		It("should work with an int", func() {
			actual := optioner.None[int]()

			Expect(actual.IsEmpty()).To(BeTrue())
			Expect(actual.Defined()).To(BeFalse())
			Expect(actual.Get()).To(Equal(0))
			Expect(actual.OrElse(1)).To(Equal(1))
		})

		It("should work with a boolean", func() {
			actual := optioner.None[bool]()

			Expect(actual.IsEmpty()).To(BeTrue())
			Expect(actual.Defined()).To(BeFalse())
			Expect(actual.Get()).To(Equal(false))
			Expect(actual.OrElse(true)).To(Equal(true))
		})
	})

	Describe("Optional()", func() {
		It("should work with a nil string pointer", func() {
			actual := optioner.Option[string](nil)

			Expect(actual.IsEmpty()).To(BeTrue())
			Expect(actual.Defined()).To(BeFalse())
			Expect(actual.Get()).To(Equal(""))
			Expect(actual.OrElse("x")).To(Equal("x"))
		})

		It("should work with a non-nil string pointer", func() {
			v := "s"
			actual := optioner.Option[string](&v)

			Expect(actual.IsEmpty()).To(BeFalse())
			Expect(actual.Defined()).To(BeTrue())
			Expect(actual.Get()).To(Equal("s"))
			Expect(actual.OrElse("x")).To(Equal("s"))
		})

		It("should work with a nil int pointer", func() {
			actual := optioner.Option[int](nil)

			Expect(actual.IsEmpty()).To(BeTrue())
			Expect(actual.Defined()).To(BeFalse())
			Expect(actual.Get()).To(Equal(0))
			Expect(actual.OrElse(2)).To(Equal(2))
		})

		It("should work with a non-nil int pointer", func() {
			v := 1
			actual := optioner.Option[int](&v)

			Expect(actual.IsEmpty()).To(BeFalse())
			Expect(actual.Defined()).To(BeTrue())
			Expect(actual.Get()).To(Equal(1))
			Expect(actual.OrElse(2)).To(Equal(1))
		})
	})
})
