package optioner

func (t Option[T]) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

func (t *Option[T]) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v T

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}