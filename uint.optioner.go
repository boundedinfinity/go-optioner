package optioner

import (
	"encoding/json"
	"fmt"
)

// UintOption contains a initialized or empty copy of the uint type.
type UintOption struct {
	v *uint
}

// NewUintValue creates a new UintOption type with an initialized value.
func NewUintValue(v uint) UintOption {
	return UintOption{
		v: &v,
	}
}

// NewUintValue creates a new UintOption type with the given value.
func NewUintPtr(v *uint) UintOption {
	return UintOption{
		v: v,
	}
}

// Uint2Ptr create a pointer from the given type.
func Uint2Ptr(v uint) *uint {
	return &v
}

// NewUintValue creates a new UintOption type with an empty value.
func NewUintEmpty() UintOption {
	return UintOption{}
}

// NewUintEmptyIfZeroValue creates a new initialized UintOption type if the input uint doesn't equal the uint's zero value, or an empty UintOption otherwise.
func NewUintEmptyIfZeroValue(v uint) UintOption {
	var e uint

	if v == e {
		return NewUintEmpty()
	}

	return NewUintValue(v)
}

//IsEmpty returns true if the contained uint value is empty, false otherwise.
func (t UintOption) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained uint value is initialized, false otherwise.
func (t UintOption) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained uint value.
//NOTE: If the value is empty, this will return the uint zero value.
func (t UintOption) Get() uint {
	var v uint

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained uint value if the contained value is initialized or the input value is the value is not initialized.
func (t UintOption) GetOrElse(v uint) uint {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the UintOption type into the JSON representation.
func (t UintOption) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the UintOption type.
func (t *UintOption) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v uint

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

//MarshalYAML marshals the UintOption type into the YAML representation.
func (t UintOption) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the UintOption type.
func (t *UintOption) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v uint

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

func (t *UintOption) String() string {
	return fmt.Sprintf("%v", t.Get())
}
