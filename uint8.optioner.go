package optioner

import (
	"encoding/json"
	"fmt"
)

// Uint8Option contains a initialized or empty copy of the uint8 type.
type Uint8Option struct {
	v *uint8
}

// NewUint8Value creates a new Uint8Option type with an initialized value.
func NewUint8Value(v uint8) Uint8Option {
	return Uint8Option{
		v: &v,
	}
}

// NewUint8Value creates a new Uint8Option type with the given value.
func NewUint8Ptr(v *uint8) Uint8Option {
	return Uint8Option{
		v: v,
	}
}

// Uint82Ptr create a pointer from the given type.
func Uint82Ptr(v uint8) *uint8 {
	return &v
}

// NewUint8Value creates a new Uint8Option type with an empty value.
func NewUint8Empty() Uint8Option {
	return Uint8Option{}
}

// NewUint8EmptyIfZeroValue creates a new initialized Uint8Option type if the input uint8 doesn't equal the uint8's zero value, or an empty Uint8Option otherwise.
func NewUint8EmptyIfZeroValue(v uint8) Uint8Option {
	var e uint8

	if v == e {
		return NewUint8Empty()
	}

	return NewUint8Value(v)
}

//IsEmpty returns true if the contained uint8 value is empty, false otherwise.
func (t Uint8Option) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained uint8 value is initialized, false otherwise.
func (t Uint8Option) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained uint8 value.
//NOTE: If the value is empty, this will return the uint8 zero value.
func (t Uint8Option) Get() uint8 {
	var v uint8

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained uint8 value if the contained value is initialized or the input value is the value is not initialized.
func (t Uint8Option) GetOrElse(v uint8) uint8 {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the Uint8Option type into the JSON representation.
func (t Uint8Option) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the Uint8Option type.
func (t *Uint8Option) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v uint8

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

//MarshalYAML marshals the Uint8Option type into the YAML representation.
func (t Uint8Option) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the Uint8Option type.
func (t *Uint8Option) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v uint8

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

func (t *Uint8Option) String() string {
	return fmt.Sprintf("%v", t.Get())
}
