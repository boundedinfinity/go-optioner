package optioner

import (
	"fmt"
)

// String implments the Stringer interface.
func (t Optioner[T]) String() string {
	return fmt.Sprintf("%v", t.Get())
}
