package optioner

import (
	"fmt"
)

// Optioner[T] is a container for zero or one element of a given type.
type Optioner[T any] struct {
	v *T
}

// String implments the Stringer interface.
func (t Optioner[T]) String() string {
	return fmt.Sprintf("%v", t.Get())
}
