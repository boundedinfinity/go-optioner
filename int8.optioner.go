package optioner

import (
	"encoding/json"
	"fmt"
)

// Int8Option contains a initialized or empty copy of the int8 type.
type Int8Option struct {
	v *int8
}

// NewInt8Value creates a new Int8Option type with an initialized value.
func NewInt8Value(v int8) Int8Option {
	return Int8Option{
		v: &v,
	}
}

// NewInt8Value creates a new Int8Option type with the given value.
func NewInt8Ptr(v *int8) Int8Option {
	return Int8Option{
		v: v,
	}
}

// Int82Ptr create a pointer from the given type.
func Int82Ptr(v int8) *int8 {
	return &v
}

// NewInt8Value creates a new Int8Option type with an empty value.
func NewInt8Empty() Int8Option {
	return Int8Option{}
}

// NewInt8EmptyIfZeroValue creates a new initialized Int8Option type if the input int8 doesn't equal the int8's zero value, or an empty Int8Option otherwise.
func NewInt8EmptyIfZeroValue(v int8) Int8Option {
	var e int8

	if v == e {
		return NewInt8Empty()
	}

	return NewInt8Value(v)
}

//IsEmpty returns true if the contained int8 value is empty, false otherwise.
func (t Int8Option) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained int8 value is initialized, false otherwise.
func (t Int8Option) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained int8 value.
//NOTE: If the value is empty, this will return the int8 zero value.
func (t Int8Option) Get() int8 {
	var v int8

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained int8 value if the contained value is initialized or the input value is the value is not initialized.
func (t Int8Option) GetOrElse(v int8) int8 {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the Int8Option type into the JSON representation.
func (t Int8Option) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the Int8Option type.
func (t *Int8Option) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v int8

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

//MarshalYAML marshals the Int8Option type into the YAML representation.
func (t Int8Option) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the Int8Option type.
func (t *Int8Option) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v int8

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

func (t *Int8Option) String() string {
	return fmt.Sprintf("%v", t.Get())
}
