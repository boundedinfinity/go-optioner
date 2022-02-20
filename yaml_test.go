package optioner_test

import (
	"github.com/boundedinfinity/optioner"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

var _ = Describe("YAML Serialization / Deserialization", func() {
	Describe("Serialize/Deserialize string", func() {
		It("should serialize a string", func() {
			input := optioner.Some("s")
			expected := []byte("s\n")
			actual, err := yaml.Marshal(input)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})

		It("should serialize an empty string", func() {
			input := optioner.None[string]()
			expected := []byte("null\n")
			actual, err := yaml.Marshal(input)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})

		It("should deserialize a string", func() {
			input := []byte("s\n")
			expected := optioner.Some("s")
			var actual optioner.Option[string]

			err := yaml.Unmarshal(input, &actual)

			Expect(err).To(BeNil())
			Expect(actual.Get()).To(Equal(expected.Get()))
			Expect(actual.Empty()).To(Equal(expected.Empty()))
			Expect(actual.Defined()).To(Equal(expected.Defined()))
		})

		It("should deserialize a null string", func() {
			input := []byte("null\n")
			expected := optioner.None[string]()
			var actual optioner.Option[string]

			err := yaml.Unmarshal(input, &actual)

			Expect(err).To(BeNil())
			Expect(actual.Get()).To(Equal(expected.Get()))
			Expect(actual.Empty()).To(Equal(expected.Empty()))
			Expect(actual.Defined()).To(Equal(expected.Defined()))
		})
	})

	Describe("Serialize/Deserialize int", func() {
		It("should serialize an int", func() {
			input := optioner.Some(1)
			expected := []byte("1\n")
			actual, err := yaml.Marshal(input)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})

		It("should serialize an empty int", func() {
			input := optioner.None[int]()
			expected := []byte("null\n")
			actual, err := yaml.Marshal(input)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})

		It("should deserialize an int", func() {
			input := []byte("1\n")
			expected := optioner.Some(1)
			var actual optioner.Option[int]

			err := yaml.Unmarshal(input, &actual)

			Expect(err).To(BeNil())
			Expect(actual.Get()).To(Equal(expected.Get()))
			Expect(actual.Empty()).To(Equal(expected.Empty()))
			Expect(actual.Defined()).To(Equal(expected.Defined()))
		})

		It("should deserialize a null int", func() {
			input := []byte("null\n")
			expected := optioner.None[int]()
			var actual optioner.Option[int]

			err := yaml.Unmarshal(input, &actual)

			Expect(err).To(BeNil())
			Expect(actual.Get()).To(Equal(expected.Get()))
			Expect(actual.Empty()).To(Equal(expected.Empty()))
			Expect(actual.Defined()).To(Equal(expected.Defined()))
		})
	})
})
