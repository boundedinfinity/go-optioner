package optioner

import (
	"encoding/json"
	"fmt"
)

// Int32Option contains a initialized or empty copy of the int32 type.
type Int32Option struct {
	v *int32
}

// NewInt32Value creates a new Int32Option type with an initialized value.
func NewInt32Value(v int32) Int32Option {
	return Int32Option{
		v: &v,
	}
}

// NewInt32Value creates a new Int32Option type with the given value.
func NewInt32Ptr(v *int32) Int32Option {
	return Int32Option{
		v: v,
	}
}

// Int322Ptr create a pointer from the given type.
func Int322Ptr(v int32) *int32 {
	return &v
}

// NewInt32Value creates a new Int32Option type with an empty value.
func NewInt32Empty() Int32Option {
	return Int32Option{}
}

// NewInt32EmptyIfZeroValue creates a new initialized Int32Option type if the input int32 doesn't equal the int32's zero value, or an empty Int32Option otherwise.
func NewInt32EmptyIfZeroValue(v int32) Int32Option {
	var e int32

	if v == e {
		return NewInt32Empty()
	}

	return NewInt32Value(v)
}

//IsEmpty returns true if the contained int32 value is empty, false otherwise.
func (t Int32Option) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained int32 value is initialized, false otherwise.
func (t Int32Option) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained int32 value.
//NOTE: If the value is empty, this will return the int32 zero value.
func (t Int32Option) Get() int32 {
	var v int32

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained int32 value if the contained value is initialized or the input value is the value is not initialized.
func (t Int32Option) GetOrElse(v int32) int32 {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the Int32Option type into the JSON representation.
func (t Int32Option) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the Int32Option type.
func (t *Int32Option) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v int32

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

//MarshalYAML marshals the Int32Option type into the YAML representation.
func (t Int32Option) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the Int32Option type.
func (t *Int32Option) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v int32

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

func (t *Int32Option) String() string {
	return fmt.Sprintf("%v", t.Get())
}
