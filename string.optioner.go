package optioner

import (
	"encoding/json"
	"fmt"
)

// StringOption contains a initialized or empty copy of the string type.
type StringOption struct {
	v *string
}

// NewStringValue creates a new StringOption type with an initialized value.
func NewStringValue(v string) StringOption {
	return StringOption{
		v: &v,
	}
}

// NewStringValue creates a new StringOption type with the given value.
func NewStringPtr(v *string) StringOption {
	return StringOption{
		v: v,
	}
}

// String2Ptr create a pointer from the given type.
func String2Ptr(v string) *string {
	return &v
}

// NewStringValue creates a new StringOption type with an empty value.
func NewStringEmpty() StringOption {
	return StringOption{}
}

// NewStringEmptyIfZeroValue creates a new initialized StringOption type if the input string doesn't equal the string's zero value, or an empty StringOption otherwise.
func NewStringEmptyIfZeroValue(v string) StringOption {
	var e string

	if v == e {
		return NewStringEmpty()
	}

	return NewStringValue(v)
}

//IsEmpty returns true if the contained string value is empty, false otherwise.
func (t StringOption) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained string value is initialized, false otherwise.
func (t StringOption) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained string value.
//NOTE: If the value is empty, this will return the string zero value.
func (t StringOption) Get() string {
	var v string

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained string value if the contained value is initialized or the input value is the value is not initialized.
func (t StringOption) GetOrElse(v string) string {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the StringOption type into the JSON representation.
func (t StringOption) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the StringOption type.
func (t *StringOption) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v string

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

//MarshalYAML marshals the StringOption type into the YAML representation.
func (t StringOption) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the StringOption type.
func (t *StringOption) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v string

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

func (t *StringOption) String() string {
	return fmt.Sprintf("%v", t.Get())
}
