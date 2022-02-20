package optioner

// Empty returns true if the contained value is empty, false otherwise
func (t Option[T]) Empty() bool {
	return t.v == nil
}

// Empty returns true if the contained value is empty, false otherwise
// Alias for Empty()
func (t Option[T]) IsEmpty() bool {
	return t.Empty()
}

// Defined returns true if the contained value is not empty, false otherwise
func (t Option[T]) Defined() bool {
	return !t.Empty()
}

// IsDefined returns true if the contained value is not empty, false otherwise
// Alias for Defined()
func (t Option[T]) IsDefined() bool {
	return t.Defined()
}

// Get returns the contained value.  If called and Empty() is true, the
// value will be the types zero value.
func (t Option[T]) Get() T {
	if t.IsDefined() {
		return *t.v
	}
	
	var zero T
	return zero
}

// OrElse returns the contained value if Defined() is true or returns
// the given value otherwise.
func (t Option[T]) OrElse(v T) T {
	if t.Empty() {
		return v
	}

	return *t.v
}

// GetOrElse returns the contained value if Defined() is true or returns
// the given value otherwise.
// Alias for OrElse()
func (t Option[T]) GetOrElse(v T) T {
	return t.OrElse(v)
}
