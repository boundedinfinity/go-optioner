package optioner_test

import (
	"testing"

	o "github.com/boundedinfinity/go-optioner"
	"github.com/stretchr/testify/assert"
)

func Test_Some_with_string(t *testing.T) {
	actual := o.Some("s")

	assert.Equal(t, actual.Empty(), false)
	assert.Equal(t, actual.Defined(), true)
	assert.Equal(t, actual.Get(), "s")
	assert.Equal(t, actual.OrElse("x"), "s")
}

func Test_Some_with_int(t *testing.T) {
	actual := o.Some(1)

	assert.Equal(t, actual.Empty(), false)
	assert.Equal(t, actual.Defined(), true)
	assert.Equal(t, actual.Get(), 1)
	assert.Equal(t, actual.OrElse(1), 1)
}

func Test_Some_with_boolean(t *testing.T) {
	actual := o.Some(false)

	assert.Equal(t, actual.Empty(), false)
	assert.Equal(t, actual.Defined(), true)
	assert.Equal(t, actual.Get(), false)
	assert.Equal(t, actual.OrElse(true), false)
}

func Test_None_with_string(t *testing.T) {
	actual := o.None[string]()

	assert.Equal(t, actual.Empty(), true)
	assert.Equal(t, actual.Defined(), false)
	assert.Equal(t, actual.Get(), "")
	assert.Equal(t, actual.OrElse("x"), "x")
}

func Test_None_with_int(t *testing.T) {
	actual := o.None[int]()

	assert.Equal(t, actual.Empty(), true)
	assert.Equal(t, actual.Defined(), false)
	assert.Equal(t, actual.Get(), 0)
	assert.Equal(t, actual.OrElse(1), 1)
}

func Test_None_with_bool(t *testing.T) {
	actual := o.None[bool]()

	assert.Equal(t, actual.Empty(), true)
	assert.Equal(t, actual.Defined(), false)
	assert.Equal(t, actual.Get(), false)
	assert.Equal(t, actual.OrElse(true), true)
}

func Test_OfP_nil_string(t *testing.T) {
	actual := o.OfP[string](nil)

	assert.Equal(t, actual.Empty(), true)
	assert.Equal(t, actual.Defined(), false)
	assert.Equal(t, actual.Get(), "")
	assert.Equal(t, actual.OrElse("x"), "x")
}

func Test_OfP_with_string(t *testing.T) {
	v := "s"
	actual := o.OfP(&v)

	assert.Equal(t, actual.Empty(), false)
	assert.Equal(t, actual.Defined(), true)
	assert.Equal(t, actual.Get(), "s")
	assert.Equal(t, actual.OrElse("x"), "s")
}

func Test_OfP_with_nil_int(t *testing.T) {
	actual := o.OfP[int](nil)

	assert.Equal(t, actual.Empty(), true)
	assert.Equal(t, actual.Defined(), false)
	assert.Equal(t, actual.Get(), 0)
	assert.Equal(t, actual.OrElse(1), 1)
}

func Test_OpP_with_int(t *testing.T) {
	v := 1
	actual := o.OfP(&v)

	assert.Equal(t, actual.Empty(), false)
	assert.Equal(t, actual.Defined(), true)
	assert.Equal(t, actual.Get(), 1)
	assert.Equal(t, actual.OrElse(2), 1)
}
