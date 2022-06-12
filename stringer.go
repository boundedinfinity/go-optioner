package optioner

import (
	"fmt"
)

// String implments the Stringer interface.
func (t Option[T]) String() string {
	return fmt.Sprintf("%v", t.Get())
}
