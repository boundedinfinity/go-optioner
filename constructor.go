package optioner

func Some[T comparable](v T) Option[T] {
	var o Option[T]
	o.v = &v
	return o
}

func None[T comparable]() Option[T] {
	var o Option[T]
	return o
}

func Optional[T comparable](v *T) Option[T] {
	return Option[T] {
		v: v,
	}
}