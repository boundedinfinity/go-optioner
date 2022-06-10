package optioner_test

import (
	"encoding/json"
	"testing"

	"github.com/boundedinfinity/optioner"
	"github.com/stretchr/testify/assert"
)

func Test_JSON_serialize_string(t *testing.T) {
	input := optioner.Some("s")
	expected := []byte(`"s"`)
	actual, err := json.Marshal(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_JSON_serialize_empty_string(t *testing.T) {
	input := optioner.None[string]()
	expected := []byte(`null`)
	actual, err := json.Marshal(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_JSON_deserialize_string(t *testing.T) {
	input := []byte(`"s"`)
	expected := optioner.Some("s")
	var actual optioner.Optioner[string]
	err := json.Unmarshal(input, &actual)

	assert.Nil(t, err)
	assert.Equal(t, expected.Get(), actual.Get())
	assert.Equal(t, expected.Empty(), actual.Empty())
	assert.Equal(t, expected.Defined(), actual.Defined())
}

func Test_JSON_deserialize_nil_string(t *testing.T) {
	input := []byte(`null`)
	expected := optioner.None[string]()
	var actual optioner.Optioner[string]
	err := json.Unmarshal(input, &actual)

	assert.Nil(t, err)
	assert.Equal(t, expected.Get(), actual.Get())
	assert.Equal(t, expected.Empty(), actual.Empty())
	assert.Equal(t, expected.Defined(), actual.Defined())
}

func Test_JSON_serialize_int(t *testing.T) {
	input := optioner.Some("s")
	expected := []byte(`"s"`)
	actual, err := json.Marshal(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_JSON_serialize_int_empty(t *testing.T) {
	input := optioner.None[int]()
	expected := []byte(`null`)
	actual, err := json.Marshal(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_JSON_deserialize_int(t *testing.T) {
	input := []byte(`null`)
	expected := optioner.None[string]()
	var actual optioner.Optioner[string]
	err := json.Unmarshal(input, &actual)

	assert.Nil(t, err)
	assert.Equal(t, expected.Get(), actual.Get())
	assert.Equal(t, expected.Empty(), actual.Empty())
	assert.Equal(t, expected.Defined(), actual.Defined())
}

func Test_JSON_deserialize_int_empty(t *testing.T) {
	input := []byte(`null`)
	expected := optioner.None[string]()
	var actual optioner.Optioner[string]
	err := json.Unmarshal(input, &actual)

	assert.Nil(t, err)
	assert.Equal(t, expected.Get(), actual.Get())
	assert.Equal(t, expected.Empty(), actual.Empty())
	assert.Equal(t, expected.Defined(), actual.Defined())
}

func Test_JSON_deserialize_nil_int(t *testing.T) {
	input := []byte(`null`)
	expected := optioner.None[int]()
	var actual optioner.Optioner[int]
	err := json.Unmarshal(input, &actual)

	assert.Nil(t, err)
	assert.Equal(t, expected.Get(), actual.Get())
	assert.Equal(t, expected.Empty(), actual.Empty())
	assert.Equal(t, expected.Defined(), actual.Defined())
}
