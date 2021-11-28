package optioner

import (
	"encoding/json"
	"fmt"
)

// IntOption contains a initialized or empty copy of the int type.
type IntOption struct {
	v *int
}

// NewIntValue creates a new IntOption type with an initialized value.
func NewIntValue(v int) IntOption {
	return IntOption{
		v: &v,
	}
}

// NewIntValue creates a new IntOption type with the given value.
func NewIntPtr(v *int) IntOption {
	return IntOption{
		v: v,
	}
}

// Int2Ptr create a pointer from the given type.
func Int2Ptr(v int) *int {
	return &v
}

// NewIntValue creates a new IntOption type with an empty value.
func NewIntEmpty() IntOption {
	return IntOption{}
}

// NewIntEmptyIfZeroValue creates a new initialized IntOption type if the input int doesn't equal the int's zero value, or an empty IntOption otherwise.
func NewIntEmptyIfZeroValue(v int) IntOption {
	var e int

	if v == e {
		return NewIntEmpty()
	}

	return NewIntValue(v)
}

//IsEmpty returns true if the contained int value is empty, false otherwise.
func (t IntOption) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained int value is initialized, false otherwise.
func (t IntOption) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained int value.
//NOTE: If the value is empty, this will return the int zero value.
func (t IntOption) Get() int {
	var v int

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained int value if the contained value is initialized or the input value is the value is not initialized.
func (t IntOption) GetOrElse(v int) int {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the IntOption type into the JSON representation.
func (t IntOption) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the IntOption type.
func (t *IntOption) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v int

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

//MarshalYAML marshals the IntOption type into the YAML representation.
func (t IntOption) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the IntOption type.
func (t *IntOption) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v int

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

func (t *IntOption) String() string {
	return fmt.Sprintf("%v", t.Get())
}
