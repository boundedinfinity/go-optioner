package optioner


func (t Option[T]) Empty() bool {
	return t.v == nil
}

func (t Option[T]) IsEmpty() bool {
	return t.Empty()
}

func (t Option[T]) Defined() bool {
	return !t.Empty()
}

func (t Option[T]) IsDefined() bool {
	return t.Defined()
}

func (t Option[T]) Get() T {
	if t.IsDefined() {
		return *t.v
	}
	
	var zero T
	return zero
}

func (t Option[T]) OrElse(v T) T {
	if t.Empty() {
		return v
	}

	return *t.v
}

func (t Option[T]) GetOrElse(v T) T {
	return t.OrElse(v)
}

