package optioner_test

import (
	"testing"

	o "github.com/boundedinfinity/go-optioner"
	"github.com/stretchr/testify/assert"
)

func Test_OrElse(t *testing.T) {
	assert.Equal(t, "a", o.Some("a").OrElse("b"))
	assert.Equal(t, "b", o.None[string]().OrElse("b"))
}
