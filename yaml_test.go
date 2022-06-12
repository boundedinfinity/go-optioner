package optioner_test

import (
	"testing"

	"github.com/boundedinfinity/optioner"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func Test_YAML_serialize_string(t *testing.T) {
	input := optioner.Some("s")
	expected := []byte("s\n")
	actual, err := yaml.Marshal(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_YAML_serialize_empty_string(t *testing.T) {
	input := optioner.None[string]()
	expected := []byte("null\n")
	actual, err := yaml.Marshal(input)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_YAML_deserialize_string(t *testing.T) {
	input := []byte("s\n")
	expected := optioner.Some("s")
	var actual optioner.Option[string]

	err := yaml.Unmarshal(input, &actual)

	assert.Nil(t, err)
	assert.Equal(t, expected.Get(), actual.Get())
	assert.Equal(t, expected.Empty(), actual.Empty())
	assert.Equal(t, expected.Defined(), actual.Defined())
}

func Test_YAML_deserialize_nil_string(t *testing.T) {
	input := []byte("null\n")
	expected := optioner.None[string]()
	var actual optioner.Option[string]

	err := yaml.Unmarshal(input, &actual)

	assert.Nil(t, err)
	assert.Equal(t, expected.Get(), actual.Get())
	assert.Equal(t, expected.Empty(), actual.Empty())
	assert.Equal(t, expected.Defined(), actual.Defined())
}

func Test_YAML_serialize_int(t *testing.T) {
	input := optioner.Some(1)
	expected := []byte("1\n")
	actual, err := yaml.Marshal(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_YAML_serialize_int_empty(t *testing.T) {
	input := optioner.None[int]()
	expected := []byte("null\n")
	actual, err := yaml.Marshal(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_YAML_deserialize_int(t *testing.T) {
	input := []byte("1\n")
	expected := optioner.Some(1)
	var actual optioner.Option[int]

	err := yaml.Unmarshal(input, &actual)

	assert.Nil(t, err)
	assert.Equal(t, expected.Get(), actual.Get())
	assert.Equal(t, expected.Empty(), actual.Empty())
	assert.Equal(t, expected.Defined(), actual.Defined())
}

func Test_YAML_deserialize_int_empty(t *testing.T) {
	input := []byte("null\n")
	expected := optioner.None[int]()
	var actual optioner.Option[int]

	err := yaml.Unmarshal(input, &actual)

	assert.Nil(t, err)
	assert.Equal(t, expected.Get(), actual.Get())
	assert.Equal(t, expected.Empty(), actual.Empty())
	assert.Equal(t, expected.Defined(), actual.Defined())
}
