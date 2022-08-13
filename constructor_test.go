package optioner_test

import (
	"testing"

	"github.com/boundedinfinity/go-optioner"
	"github.com/stretchr/testify/assert"
)

func Test_Some_with_string(t *testing.T) {
	actual := optioner.Some("s")

	assert.Equal(t, actual.Empty(), false)
	assert.Equal(t, actual.Defined(), true)
	assert.Equal(t, actual.Get(), "s")
	assert.Equal(t, actual.OrElse("x"), "s")
}

func Test_Some_with_int(t *testing.T) {
	actual := optioner.Some(1)

	assert.Equal(t, actual.Empty(), false)
	assert.Equal(t, actual.Defined(), true)
	assert.Equal(t, actual.Get(), 1)
	assert.Equal(t, actual.OrElse(1), 1)
}

func Test_Some_with_boolean(t *testing.T) {
	actual := optioner.Some(false)

	assert.Equal(t, actual.Empty(), false)
	assert.Equal(t, actual.Defined(), true)
	assert.Equal(t, actual.Get(), false)
	assert.Equal(t, actual.OrElse(true), false)
}

func Test_None_with_string(t *testing.T) {
	actual := optioner.None[string]()

	assert.Equal(t, actual.Empty(), true)
	assert.Equal(t, actual.Defined(), false)
	assert.Equal(t, actual.Get(), "")
	assert.Equal(t, actual.OrElse("x"), "x")
}

func Test_None_with_int(t *testing.T) {
	actual := optioner.None[int]()

	assert.Equal(t, actual.Empty(), true)
	assert.Equal(t, actual.Defined(), false)
	assert.Equal(t, actual.Get(), 0)
	assert.Equal(t, actual.OrElse(1), 1)
}

func Test_None_with_bool(t *testing.T) {
	actual := optioner.None[bool]()

	assert.Equal(t, actual.Empty(), true)
	assert.Equal(t, actual.Defined(), false)
	assert.Equal(t, actual.Get(), false)
	assert.Equal(t, actual.OrElse(true), true)
}

func Test_Option_with_nil_string(t *testing.T) {
	actual := optioner.Of[string](nil)

	assert.Equal(t, actual.Empty(), true)
	assert.Equal(t, actual.Defined(), false)
	assert.Equal(t, actual.Get(), "")
	assert.Equal(t, actual.OrElse("x"), "x")
}

func Test_Option_with_string(t *testing.T) {
	v := "s"
	actual := optioner.Of(&v)

	assert.Equal(t, actual.Empty(), false)
	assert.Equal(t, actual.Defined(), true)
	assert.Equal(t, actual.Get(), "s")
	assert.Equal(t, actual.OrElse("x"), "s")
}

func Test_Option_with_nil_int(t *testing.T) {
	actual := optioner.Of[int](nil)

	assert.Equal(t, actual.Empty(), true)
	assert.Equal(t, actual.Defined(), false)
	assert.Equal(t, actual.Get(), 0)
	assert.Equal(t, actual.OrElse(1), 1)
}

func Test_Option_with_int(t *testing.T) {
	v := 1
	actual := optioner.Of(&v)

	assert.Equal(t, actual.Empty(), false)
	assert.Equal(t, actual.Defined(), true)
	assert.Equal(t, actual.Get(), 1)
	assert.Equal(t, actual.OrElse(2), 1)
}
