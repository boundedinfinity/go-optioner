package optioner

// Optioner[T] is a container for zero or one element of a given type.
type Optioner[T any] struct {
	v *T
}

// Empty returns true if the contained value is empty, false otherwise
//
// Inverse of Defined()
func (t Optioner[T]) Empty() bool {
	return t.v == nil
}

// Defined returns true if the contained value is not empty, false otherwise
//
// Inverse of Empty()
func (t Optioner[T]) Defined() bool {
	return !t.Empty()
}

// Get returns the contained value.
//
// NOTE: If called and Empty() is true, the value will be the type's zero value.
func (t Optioner[T]) Get() T {
	if t.Defined() {
		return *t.v
	}

	var zero T
	return zero
}

// OrElse returns the contained value if Defined() is true or returns
// the provided argument.
func (t Optioner[T]) OrElse(v T) T {
	if t.Empty() {
		return v
	}

	return *t.v
}
