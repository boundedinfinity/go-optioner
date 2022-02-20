package optioner

// MarshalYAML implements the gopkg.in/yaml.v2#Marshal interface
func (t Option[T]) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

// UnmarshalYAML implements the gopkg.in/yaml.v2#Unmarshal interface
func (t *Option[T]) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v T

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}