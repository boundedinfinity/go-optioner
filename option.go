package optioner

// Option[T] is a container for zero or one element of a given type.
type Option[T any] struct {
	v *T
}

// Empty returns true if the contained value is empty, false otherwise
//
// Inverse of Defined()
func (t Option[T]) Empty() bool {
	return t.v == nil
}

// Defined returns true if the contained value is not empty, false otherwise
//
// Inverse of Empty()
func (t Option[T]) Defined() bool {
	return !t.Empty()
}

// Get returns the contained value.
//
// NOTE: If called and Empty() is true, the value will be the type's zero value.
func (t Option[T]) Get() T {
	if t.Defined() {
		return *t.v
	}

	var zero T
	return zero
}

// OrElse returns the contained value if Defined() is true or returns
// the provided argument.
func (t Option[T]) OrElse(v T) T {
	if t.Empty() {
		return v
	}

	return *t.v
}

func (t Option[T]) OrFirst(vs ...Option[T]) Option[T] {
	for _, v := range vs {
		if v.Defined() {
			return v
		}
	}

	return t
}

func (t Option[T]) OrLast(vs ...Option[T]) Option[T] {
	for i := len(vs) - 1; i >= 0; i-- {
		if vs[i].Defined() {
			return vs[i]
		}
	}

	return t
}
