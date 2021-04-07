package optional

import (
	"encoding/json"
)

// Int16Optional contains a initialized or empty copy of the int16 type.
type Int16Optional struct {
	v *int16
}

// NewInt16Value creates a new Int16Optional type with an initialized value.
func NewInt16Value(v int16) Int16Optional {
	return Int16Optional{
		v: &v,
	}
}

// NewInt16Value creates a new Int16Optional type with the given value.
func NewInt16Ptr(v *int16) Int16Optional {
	return Int16Optional{
		v: v,
	}
}

// Int162Ptr create a pointer from the given type.
func Int162Ptr(v int16) *int16 {
	return &v
}

// NewInt16Value creates a new Int16Optional type with an empty value.
func NewInt16Empty() Int16Optional {
	return Int16Optional{}
}

// NewInt16EmptyIfZeroValue creates a new initialized Int16Optional type if the input int16 doesn't equal the int16's zero value, or an empty Int16Optional otherwise.
func NewInt16EmptyIfZeroValue(v int16) Int16Optional {
	var e int16

	if v == e {
		return NewInt16Empty()
	}

	return NewInt16Value(v)
}

//IsEmpty returns true if the contained int16 value is empty, false otherwise.
func (t Int16Optional) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained int16 value is initialized, false otherwise.
func (t Int16Optional) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained int16 value.
//NOTE: If the value is empty, this will return the int16 zero value.
func (t Int16Optional) Get() int16 {
	var v int16

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained int16 value if the contained value is initialized or the input value is the value is not initialized.
func (t Int16Optional) GetOrElse(v int16) int16 {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the Int16Optional type into the JSON representation.
func (t Int16Optional) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the Int16Optional type.
func (t *Int16Optional) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v int16

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

//MarshalYAML marshals the Int16Optional type into the YAML representation.
func (t Int16Optional) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the Int16Optional type.
func (t *Int16Optional) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v int16

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}
