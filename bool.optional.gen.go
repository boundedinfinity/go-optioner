package optional

import (
	"encoding/json"
)

// BoolOptional contains a initialized or empty copy of the bool type.
type BoolOptional struct {
	v *bool
}

// NewBoolValue creates a new BoolOptional type with an initialized value.
func NewBoolValue(v bool) BoolOptional {
	return BoolOptional{
		v: &v,
	}
}

// NewBoolValue creates a new BoolOptional type with the given value.
func NewBoolPtr(v *bool) BoolOptional {
	return BoolOptional{
		v: v,
	}
}

// Bool2Ptr create a pointer from the given type.
func Bool2Ptr(v bool) *bool {
	return &v
}

// NewBoolValue creates a new BoolOptional type with an empty value.
func NewBoolEmpty() BoolOptional {
	return BoolOptional{}
}

// NewBoolEmptyIfZeroValue creates a new initialized BoolOptional type if the input bool doesn't equal the bool's zero value, or an empty BoolOptional otherwise.
func NewBoolEmptyIfZeroValue(v bool) BoolOptional {
	var e bool

	if v == e {
		return NewBoolEmpty()
	}

	return NewBoolValue(v)
}

//IsEmpty returns true if the contained bool value is empty, false otherwise.
func (t BoolOptional) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained bool value is initialized, false otherwise.
func (t BoolOptional) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained bool value.
//NOTE: If the value is empty, this will return the bool zero value.
func (t BoolOptional) Get() bool {
	var v bool

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained bool value if the contained value is initialized or the input value is the value is not initialized.
func (t BoolOptional) GetOrElse(v bool) bool {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the BoolOptional type into the JSON representation.
func (t BoolOptional) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the BoolOptional type.
func (t *BoolOptional) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v bool

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

//MarshalYAML marshals the BoolOptional type into the YAML representation.
func (t BoolOptional) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the BoolOptional type.
func (t *BoolOptional) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v bool

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}
