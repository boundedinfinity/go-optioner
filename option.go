package optioner

// Optioner[T] is a container for zero or one element of a given type.
type Optioner[T any] struct {
	v *T
}

// Empty returns true if the contained value is empty, false otherwise
func (t Optioner[T]) Empty() bool {
	return t.v == nil
}

// Empty returns true if the contained value is empty, false otherwise
// Alias for Empty()
func (t Optioner[T]) IsEmpty() bool {
	return t.Empty()
}

// Defined returns true if the contained value is not empty, false otherwise
func (t Optioner[T]) Defined() bool {
	return !t.Empty()
}

// IsDefined returns true if the contained value is not empty, false otherwise
// Alias for Defined()
func (t Optioner[T]) IsDefined() bool {
	return t.Defined()
}

// Get returns the contained value.  If called and Empty() is true, the
// value will be the type's zero value.
func (t Optioner[T]) Get() T {
	if t.IsDefined() {
		return *t.v
	}

	var zero T
	return zero
}

// OrElse returns the contained value if Defined() is true or returns
// the given value otherwise.
//
// Alias for GetOrElse()
func (t Optioner[T]) OrElse(v T) T {
	return t.GetOrElse(v)
}

// GetOrElse returns the contained value if Defined() is true or returns
// the given value otherwise.
func (t Optioner[T]) GetOrElse(v T) T {
	if t.Empty() {
		return v
	}

	return *t.v
}
