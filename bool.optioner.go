package optioner

import (
	"encoding/json"
	"fmt"
)

// BoolOption contains a initialized or empty copy of the bool type.
type BoolOption struct {
	v *bool
}

// NewBoolValue creates a new BoolOption type with an initialized value.
func NewBoolValue(v bool) BoolOption {
	return BoolOption{
		v: &v,
	}
}

// NewBoolValue creates a new BoolOption type with the given value.
func NewBoolPtr(v *bool) BoolOption {
	return BoolOption{
		v: v,
	}
}

// Bool2Ptr create a pointer from the given type.
func Bool2Ptr(v bool) *bool {
	return &v
}

// NewBoolValue creates a new BoolOption type with an empty value.
func NewBoolEmpty() BoolOption {
	return BoolOption{}
}

// NewBoolEmptyIfZeroValue creates a new initialized BoolOption type if the input bool doesn't equal the bool's zero value, or an empty BoolOption otherwise.
func NewBoolEmptyIfZeroValue(v bool) BoolOption {
	var e bool

	if v == e {
		return NewBoolEmpty()
	}

	return NewBoolValue(v)
}

//IsEmpty returns true if the contained bool value is empty, false otherwise.
func (t BoolOption) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained bool value is initialized, false otherwise.
func (t BoolOption) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained bool value.
//NOTE: If the value is empty, this will return the bool zero value.
func (t BoolOption) Get() bool {
	var v bool

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained bool value if the contained value is initialized or the input value is the value is not initialized.
func (t BoolOption) GetOrElse(v bool) bool {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the BoolOption type into the JSON representation.
func (t BoolOption) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the BoolOption type.
func (t *BoolOption) UnmarshalJSON(data []byte) error {
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

//MarshalYAML marshals the BoolOption type into the YAML representation.
func (t BoolOption) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the BoolOption type.
func (t *BoolOption) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v bool

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

func (t *BoolOption) String() string {
	return fmt.Sprintf("%v", t.Get())
}
