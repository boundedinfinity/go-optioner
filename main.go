package optioner

import (
	"fmt"
)

type Option[T comparable] struct {
	v *T
}

func (t Option[T]) String() string {
	return fmt.Sprintf("%v", t.Get())
}