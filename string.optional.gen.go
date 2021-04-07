package optional

import (
	"encoding/json"
	"fmt"
)

// StringOptional contains a initialized or empty copy of the string type.
type StringOptional struct {
	v *string
}

// NewStringValue creates a new StringOptional type with an initialized value.
func NewStringValue(v string) StringOptional {
	return StringOptional{
		v: &v,
	}
}

// NewStringValue creates a new StringOptional type with the given value.
func NewStringPtr(v *string) StringOptional {
	return StringOptional{
		v: v,
	}
}

// String2Ptr create a pointer from the given type.
func String2Ptr(v string) *string {
	return &v
}

// NewStringValue creates a new StringOptional type with an empty value.
func NewStringEmpty() StringOptional {
	return StringOptional{}
}

// NewStringEmptyIfZeroValue creates a new initialized StringOptional type if the input string doesn't equal the string's zero value, or an empty StringOptional otherwise.
func NewStringEmptyIfZeroValue(v string) StringOptional {
	var e string

	if v == e {
		return NewStringEmpty()
	}

	return NewStringValue(v)
}

//IsEmpty returns true if the contained string value is empty, false otherwise.
func (t StringOptional) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained string value is initialized, false otherwise.
func (t StringOptional) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained string value.
//NOTE: If the value is empty, this will return the string zero value.
func (t StringOptional) Get() string {
	var v string

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained string value if the contained value is initialized or the input value is the value is not initialized.
func (t StringOptional) GetOrElse(v string) string {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the StringOptional type into the JSON representation.
func (t StringOptional) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the StringOptional type.
func (t *StringOptional) UnmarshalJSON(data []byte) error {
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

//MarshalYAML marshals the StringOptional type into the YAML representation.
func (t StringOptional) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the StringOptional type.
func (t *StringOptional) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v string

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

func (t *StringOptional) String() string {
	return fmt.Sprintf("%v", t.Get())
}
