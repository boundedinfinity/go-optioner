package optioner

// Some[T] creates an Option[T] with the given value.
func Some[T any](v T) Option[T] {
	var o Option[T]
	o.v = &v
	return o
}

// None[T] creates an Option[T] with no value.
func None[T any]() Option[T] {
	var o Option[T]
	return o
}

// Of[T] creates a Option[T] that may or may not have a value.
// If *T is nil the return Option[T] is empty (equivalent to None[T]),
// If *T is non-nil return Option[T] will contain value (equivalent to Some[T]).
//
// This is an alias for Option[T]
func Of[T any](v *T) Option[T] {
	return Option[T]{
		v: v,
	}
}
