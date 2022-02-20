package optioner

import (
	"fmt"
)

// Option[T] is a container for zero or one element of a given type. 
type Option[T comparable] struct {
	v *T
}

// String implments the Stringer interface.
func (t Option[T]) String() string {
	return fmt.Sprintf("%v", t.Get())
}