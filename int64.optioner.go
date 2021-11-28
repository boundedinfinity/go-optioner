package optioner

import (
	"encoding/json"
	"fmt"
)

// Int64Option contains a initialized or empty copy of the int64 type.
type Int64Option struct {
	v *int64
}

// NewInt64Value creates a new Int64Option type with an initialized value.
func NewInt64Value(v int64) Int64Option {
	return Int64Option{
		v: &v,
	}
}

// NewInt64Value creates a new Int64Option type with the given value.
func NewInt64Ptr(v *int64) Int64Option {
	return Int64Option{
		v: v,
	}
}

// Int642Ptr create a pointer from the given type.
func Int642Ptr(v int64) *int64 {
	return &v
}

// NewInt64Value creates a new Int64Option type with an empty value.
func NewInt64Empty() Int64Option {
	return Int64Option{}
}

// NewInt64EmptyIfZeroValue creates a new initialized Int64Option type if the input int64 doesn't equal the int64's zero value, or an empty Int64Option otherwise.
func NewInt64EmptyIfZeroValue(v int64) Int64Option {
	var e int64

	if v == e {
		return NewInt64Empty()
	}

	return NewInt64Value(v)
}

//IsEmpty returns true if the contained int64 value is empty, false otherwise.
func (t Int64Option) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained int64 value is initialized, false otherwise.
func (t Int64Option) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained int64 value.
//NOTE: If the value is empty, this will return the int64 zero value.
func (t Int64Option) Get() int64 {
	var v int64

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained int64 value if the contained value is initialized or the input value is the value is not initialized.
func (t Int64Option) GetOrElse(v int64) int64 {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the Int64Option type into the JSON representation.
func (t Int64Option) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the Int64Option type.
func (t *Int64Option) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v int64

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

//MarshalYAML marshals the Int64Option type into the YAML representation.
func (t Int64Option) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the Int64Option type.
func (t *Int64Option) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v int64

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

func (t *Int64Option) String() string {
	return fmt.Sprintf("%v", t.Get())
}
