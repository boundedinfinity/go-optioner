package optioner_test

import (
	"encoding/json"

	"github.com/boundedinfinity/optioner"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("JSON Serialization / Deserialization", func() {
	Describe("Serialize/Deserialize string", func() {
		It("should serialize a string", func() {
			input := optioner.Some("s")
			expected := []byte(`"s"`)
			actual, err := json.Marshal(input)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})

		It("should serialize an empty string", func() {
			input := optioner.None[string]()
			expected := []byte(`null`)
			actual, err := json.Marshal(input)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})

		It("should deserialize a string", func() {
			input := []byte(`"s"`)
			expected := optioner.Some("s")
			var actual optioner.Option[string]

			err := json.Unmarshal(input, &actual)

			Expect(err).To(BeNil())
			Expect(actual.Get()).To(Equal(expected.Get()))
			Expect(actual.Empty()).To(Equal(expected.Empty()))
			Expect(actual.Defined()).To(Equal(expected.Defined()))
		})

		It("should deserialize a null string", func() {
			input := []byte(`null`)
			expected := optioner.None[string]()
			var actual optioner.Option[string]

			err := json.Unmarshal(input, &actual)

			Expect(err).To(BeNil())
			Expect(actual.Get()).To(Equal(expected.Get()))
			Expect(actual.Empty()).To(Equal(expected.Empty()))
			Expect(actual.Defined()).To(Equal(expected.Defined()))
		})
	})

	Describe("Serialize/Deserialize int", func() {
		It("should serialize an int", func() {
			input := optioner.Some("s")
			expected := []byte(`"s"`)
			actual, err := json.Marshal(input)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})

		It("should serialize an empty int", func() {
			input := optioner.None[int]()
			expected := []byte(`null`)
			actual, err := json.Marshal(input)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})

		It("should deserialize an int", func() {
			input := []byte(`1`)
			expected := optioner.Some(1)
			var actual optioner.Option[int]

			err := json.Unmarshal(input, &actual)

			Expect(err).To(BeNil())
			Expect(actual.Get()).To(Equal(expected.Get()))
			Expect(actual.Empty()).To(Equal(expected.Empty()))
			Expect(actual.Defined()).To(Equal(expected.Defined()))
		})

		It("should deserialize a null int", func() {
			input := []byte(`null`)
			expected := optioner.None[int]()
			var actual optioner.Option[int]

			err := json.Unmarshal(input, &actual)

			Expect(err).To(BeNil())
			Expect(actual.Get()).To(Equal(expected.Get()))
			Expect(actual.Empty()).To(Equal(expected.Empty()))
			Expect(actual.Defined()).To(Equal(expected.Defined()))
		})
	})
})
