package optioner

// First[T] returns the first non-Empty[T] value from the given
// variadic arguments.
//
// If no arguments are non-Empty[T], an Empty[T] value is returned.
func First[T any](elems ...Option[T]) Option[T] {
	for _, elem := range elems {
		if elem.Defined() {
			return elem
		}
	}

	return None[T]()
}

// Last[T] returns the first non-Empty[T] value from the given
// variadic arguments from right to left (reverse of the input).
//
// If no arguments are non-Empty[T], an Empty[T] value is returned.
func Last[T any](vs ...Option[T]) Option[T] {
	for i := len(vs) - 1; i >= 0; i-- {
		if vs[i].Defined() {
			return vs[i]
		}
	}

	return None[T]()
}
