package optioner

import (
	"encoding/json"
	"fmt"
)

// Uint16Option contains a initialized or empty copy of the uint16 type.
type Uint16Option struct {
	v *uint16
}

// NewUint16Value creates a new Uint16Option type with an initialized value.
func NewUint16Value(v uint16) Uint16Option {
	return Uint16Option{
		v: &v,
	}
}

// NewUint16Value creates a new Uint16Option type with the given value.
func NewUint16Ptr(v *uint16) Uint16Option {
	return Uint16Option{
		v: v,
	}
}

// Uint162Ptr create a pointer from the given type.
func Uint162Ptr(v uint16) *uint16 {
	return &v
}

// NewUint16Value creates a new Uint16Option type with an empty value.
func NewUint16Empty() Uint16Option {
	return Uint16Option{}
}

// NewUint16EmptyIfZeroValue creates a new initialized Uint16Option type if the input uint16 doesn't equal the uint16's zero value, or an empty Uint16Option otherwise.
func NewUint16EmptyIfZeroValue(v uint16) Uint16Option {
	var e uint16

	if v == e {
		return NewUint16Empty()
	}

	return NewUint16Value(v)
}

//IsEmpty returns true if the contained uint16 value is empty, false otherwise.
func (t Uint16Option) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained uint16 value is initialized, false otherwise.
func (t Uint16Option) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained uint16 value.
//NOTE: If the value is empty, this will return the uint16 zero value.
func (t Uint16Option) Get() uint16 {
	var v uint16

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained uint16 value if the contained value is initialized or the input value is the value is not initialized.
func (t Uint16Option) GetOrElse(v uint16) uint16 {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the Uint16Option type into the JSON representation.
func (t Uint16Option) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the Uint16Option type.
func (t *Uint16Option) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v uint16

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

//MarshalYAML marshals the Uint16Option type into the YAML representation.
func (t Uint16Option) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the Uint16Option type.
func (t *Uint16Option) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v uint16

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

func (t *Uint16Option) String() string {
	return fmt.Sprintf("%v", t.Get())
}
