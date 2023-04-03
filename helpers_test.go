package optioner_test

import (
	"testing"

	o "github.com/boundedinfinity/go-optioner"
	"github.com/stretchr/testify/assert"
)

func Test_First(t *testing.T) {
	assert.Equal(t, o.Some("a"), o.First(o.Some("a"), o.Some("b")))
	assert.Equal(t, o.Some("b"), o.First(o.None[string](), o.Some("b")))
	assert.Equal(t, o.None[string](), o.First(o.None[string]()))
}

func Test_OrLast(t *testing.T) {
	assert.Equal(t, o.Some("b"), o.Last(o.Some("a"), o.Some("b")))
	assert.Equal(t, o.Some("a"), o.Last(o.Some("a"), o.None[string]()))
	assert.Equal(t, o.None[string](), o.Last(o.None[string]()))
}
