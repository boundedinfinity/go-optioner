package optioner

import (
	"encoding/json"
	"fmt"
)

// Float64Option contains a initialized or empty copy of the float64 type.
type Float64Option struct {
	v *float64
}

// NewFloat64Value creates a new Float64Option type with an initialized value.
func NewFloat64Value(v float64) Float64Option {
	return Float64Option{
		v: &v,
	}
}

// NewFloat64Value creates a new Float64Option type with the given value.
func NewFloat64Ptr(v *float64) Float64Option {
	return Float64Option{
		v: v,
	}
}

// Float642Ptr create a pointer from the given type.
func Float642Ptr(v float64) *float64 {
	return &v
}

// NewFloat64Value creates a new Float64Option type with an empty value.
func NewFloat64Empty() Float64Option {
	return Float64Option{}
}

// NewFloat64EmptyIfZeroValue creates a new initialized Float64Option type if the input float64 doesn't equal the float64's zero value, or an empty Float64Option otherwise.
func NewFloat64EmptyIfZeroValue(v float64) Float64Option {
	var e float64

	if v == e {
		return NewFloat64Empty()
	}

	return NewFloat64Value(v)
}

//IsEmpty returns true if the contained float64 value is empty, false otherwise.
func (t Float64Option) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained float64 value is initialized, false otherwise.
func (t Float64Option) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained float64 value.
//NOTE: If the value is empty, this will return the float64 zero value.
func (t Float64Option) Get() float64 {
	var v float64

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained float64 value if the contained value is initialized or the input value is the value is not initialized.
func (t Float64Option) GetOrElse(v float64) float64 {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the Float64Option type into the JSON representation.
func (t Float64Option) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the Float64Option type.
func (t *Float64Option) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v float64

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

//MarshalYAML marshals the Float64Option type into the YAML representation.
func (t Float64Option) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the Float64Option type.
func (t *Float64Option) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v float64

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

func (t *Float64Option) String() string {
	return fmt.Sprintf("%v", t.Get())
}
