package optioner

import (
	"encoding/json"
	"fmt"
)

// Float32Option contains a initialized or empty copy of the float32 type.
type Float32Option struct {
	v *float32
}

// NewFloat32Value creates a new Float32Option type with an initialized value.
func NewFloat32Value(v float32) Float32Option {
	return Float32Option{
		v: &v,
	}
}

// NewFloat32Value creates a new Float32Option type with the given value.
func NewFloat32Ptr(v *float32) Float32Option {
	return Float32Option{
		v: v,
	}
}

// Float322Ptr create a pointer from the given type.
func Float322Ptr(v float32) *float32 {
	return &v
}

// NewFloat32Value creates a new Float32Option type with an empty value.
func NewFloat32Empty() Float32Option {
	return Float32Option{}
}

// NewFloat32EmptyIfZeroValue creates a new initialized Float32Option type if the input float32 doesn't equal the float32's zero value, or an empty Float32Option otherwise.
func NewFloat32EmptyIfZeroValue(v float32) Float32Option {
	var e float32

	if v == e {
		return NewFloat32Empty()
	}

	return NewFloat32Value(v)
}

//IsEmpty returns true if the contained float32 value is empty, false otherwise.
func (t Float32Option) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained float32 value is initialized, false otherwise.
func (t Float32Option) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained float32 value.
//NOTE: If the value is empty, this will return the float32 zero value.
func (t Float32Option) Get() float32 {
	var v float32

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained float32 value if the contained value is initialized or the input value is the value is not initialized.
func (t Float32Option) GetOrElse(v float32) float32 {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the Float32Option type into the JSON representation.
func (t Float32Option) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the Float32Option type.
func (t *Float32Option) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v float32

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

//MarshalYAML marshals the Float32Option type into the YAML representation.
func (t Float32Option) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the Float32Option type.
func (t *Float32Option) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v float32

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

func (t *Float32Option) String() string {
	return fmt.Sprintf("%v", t.Get())
}
