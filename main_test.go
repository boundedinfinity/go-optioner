package optional_test

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/boundedinfinity/optional"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOptionals(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Optional Suite")
}

type StrStruct struct {
	V optional.StringOptional
}

var _ = Describe("StringOptional", func() {
	Describe("constructed with NewStringEmpty", func() {
		actual := optional.NewStringEmpty()

		It("be empty", func() {
			Expect(actual.IsEmpty()).To(BeTrue())
			Expect(actual.IsDefined()).To(BeFalse())
			Expect(actual.Get()).To(BeEmpty())
			Expect(actual.GetOrElse("turd")).To(Equal("turd"))
		})

		It("marshal to JSON value of null", func() {
			bs, err := json.Marshal(actual)
			Expect(err).To(BeNil())
			Expect(bs).NotTo(BeNil())
			Expect(string(bs)).To(Equal("null"))
		})

		It("unmarshal from JSON value of null", func() {
			var actual optional.StringOptional
			bs := []byte("null")
			err := json.Unmarshal(bs, &actual)
			Expect(err).To(BeNil())
			Expect(actual.IsEmpty()).To(BeTrue())
			Expect(actual.IsDefined()).To(BeFalse())
		})

		It("unmarshal from JSON embedded value of null", func() {
			var actual StrStruct
			bs1 := []byte(`{"V": null}`)
			err := json.Unmarshal(bs1, &actual)
			Expect(err).To(BeNil())
			Expect(actual.V.IsEmpty()).To(BeTrue())
			Expect(actual.V.IsDefined()).To(BeFalse())
		})

		It("unmarshal from JSON missing value", func() {
			var actual StrStruct
			bs1 := []byte(`{}`)
			err := json.Unmarshal(bs1, &actual)
			Expect(err).To(BeNil())
			Expect(actual.V.IsEmpty()).To(BeTrue())
			Expect(actual.V.IsDefined()).To(BeFalse())
		})
	})

	Describe("constructed with NewStringValue", func() {
		input := "a string"
		actual := optional.NewStringValue(input)

		It("be defined", func() {
			Expect(actual.IsDefined()).To(BeTrue())
			Expect(actual.IsEmpty()).To(BeFalse())
			Expect(actual.Get()).To(Equal(input))
			Expect(actual.GetOrElse("turd")).To(Equal(input))
		})

		It(fmt.Sprintf("marshal to JSON value of '%v'", input), func() {
			bs, err := json.Marshal(actual)
			Expect(err).To(BeNil())
			Expect(bs).NotTo(BeNil())
			Expect(string(bs)).To(Equal(`"` + input + `"`))
		})

		It(fmt.Sprintf("unmarshal from JSON value of '%v'", input), func() {
			var actual optional.StringOptional
			bs := []byte(`"` + input + `"`)
			err := json.Unmarshal(bs, &actual)
			Expect(err).To(BeNil())
			Expect(actual.IsDefined()).To(BeTrue())
		})

		It(fmt.Sprintf("unmarshal from JSON embedded value of '%v'", input), func() {
			var actual StrStruct
			bs1 := []byte(fmt.Sprintf(`{"V": "%v"}`, input))
			err := json.Unmarshal(bs1, &actual)
			Expect(err).To(BeNil())
			Expect(actual.V.IsEmpty()).To(BeFalse())
			Expect(actual.V.IsDefined()).To(BeTrue())
			Expect(actual.V.Get()).To(Equal(input))
		})
	})
})

var _ = Describe("IntOptional", func() {
	Describe("constructed with NewIntEmpty", func() {
		actual := optional.NewIntEmpty()

		It("be empty", func() {
			Expect(actual.IsEmpty()).To(BeTrue())
			Expect(actual.IsDefined()).To(BeFalse())
			Expect(actual.Get()).To(Equal(0))
			Expect(actual.GetOrElse(55)).To(Equal(55))
		})

		It("marshal to JSON value of null", func() {
			bs, err := json.Marshal(actual)
			Expect(err).To(BeNil())
			Expect(bs).NotTo(BeNil())
			Expect(string(bs)).To(Equal("null"))
		})

		It("unmarshal from JSON value of null", func() {
			var actual optional.IntOptional
			bs := []byte("null")
			err := json.Unmarshal(bs, &actual)
			Expect(err).To(BeNil())
			Expect(actual.IsEmpty()).To(BeTrue())
			Expect(actual.IsDefined()).To(BeFalse())
			Expect(actual.Get()).To(Equal(0))
			Expect(actual.GetOrElse(55)).To(Equal(55))
		})
	})

	Describe("constructed with NewIntValue", func() {
		input := 42
		actual := optional.NewIntValue(input)

		It("be defined", func() {
			Expect(actual.IsDefined()).To(BeTrue())
			Expect(actual.IsEmpty()).To(BeFalse())
			Expect(actual.Get()).To(Equal(input))
			Expect(actual.GetOrElse(55)).To(Equal(input))
		})

		It(fmt.Sprintf("marshal to JSON value of '%v'", input), func() {
			bs, err := json.Marshal(actual)
			Expect(err).To(BeNil())
			Expect(bs).NotTo(BeNil())
			Expect(string(bs)).To(Equal(strconv.Itoa(input)))
		})

		It(fmt.Sprintf("unmarshal from JSON value of '%v'", input), func() {
			var actual optional.IntOptional
			bs := []byte(strconv.Itoa(input))
			err := json.Unmarshal(bs, &actual)
			Expect(err).To(BeNil())
			Expect(actual.IsDefined()).To(BeTrue())
			Expect(actual.Get()).To(Equal(input))
			Expect(actual.GetOrElse(55)).To(Equal(input))
		})
	})
})

var _ = Describe("BoolOptional", func() {
	Describe("constructed with NewBoolEmpty", func() {
		actual := optional.NewBoolEmpty()

		It("be empty", func() {
			Expect(actual.IsEmpty()).To(BeTrue())
			Expect(actual.IsDefined()).To(BeFalse())
			Expect(actual.Get()).To(BeFalse())
			Expect(actual.GetOrElse(true)).To(BeTrue())
		})

		It("marshal to JSON value of null", func() {
			bs, err := json.Marshal(actual)
			Expect(err).To(BeNil())
			Expect(bs).NotTo(BeNil())
			Expect(string(bs)).To(Equal("null"))
		})

		It("unmarshal from JSON value of null", func() {
			var actual optional.BoolOptional
			bs := []byte("null")
			err := json.Unmarshal(bs, &actual)
			Expect(err).To(BeNil())
			Expect(actual.IsEmpty()).To(BeTrue())
			Expect(actual.IsDefined()).To(BeFalse())
			Expect(actual.Get()).To(BeFalse())
			Expect(actual.GetOrElse(true)).To(BeTrue())
		})
	})

	Describe("constructed with NewIntValue", func() {
		input := true
		actual := optional.NewBoolValue(input)

		It("be defined", func() {
			Expect(actual.IsDefined()).To(BeTrue())
			Expect(actual.IsEmpty()).To(BeFalse())
			Expect(actual.Get()).To(Equal(input))
			Expect(actual.GetOrElse(false)).To(Equal(input))
		})

		It(fmt.Sprintf("marshal to JSON value of '%v'", input), func() {
			bs, err := json.Marshal(actual)
			Expect(err).To(BeNil())
			Expect(bs).NotTo(BeNil())
			Expect(string(bs)).To(Equal(strconv.FormatBool(input)))
		})

		It(fmt.Sprintf("unmarshal from JSON value of '%v'", input), func() {
			var actual optional.BoolOptional
			bs := []byte(strconv.FormatBool(input))
			err := json.Unmarshal(bs, &actual)
			Expect(err).To(BeNil())
			Expect(actual.IsDefined()).To(BeTrue())
			Expect(actual.Get()).To(Equal(input))
			Expect(actual.GetOrElse(false)).To(Equal(input))
		})
	})
})
