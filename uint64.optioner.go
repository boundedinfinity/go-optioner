package optioner

import (
	"encoding/json"
	"fmt"
)

// Uint64Option contains a initialized or empty copy of the uint64 type.
type Uint64Option struct {
	v *uint64
}

// NewUint64Value creates a new Uint64Option type with an initialized value.
func NewUint64Value(v uint64) Uint64Option {
	return Uint64Option{
		v: &v,
	}
}

// NewUint64Value creates a new Uint64Option type with the given value.
func NewUint64Ptr(v *uint64) Uint64Option {
	return Uint64Option{
		v: v,
	}
}

// Uint642Ptr create a pointer from the given type.
func Uint642Ptr(v uint64) *uint64 {
	return &v
}

// NewUint64Value creates a new Uint64Option type with an empty value.
func NewUint64Empty() Uint64Option {
	return Uint64Option{}
}

// NewUint64EmptyIfZeroValue creates a new initialized Uint64Option type if the input uint64 doesn't equal the uint64's zero value, or an empty Uint64Option otherwise.
func NewUint64EmptyIfZeroValue(v uint64) Uint64Option {
	var e uint64

	if v == e {
		return NewUint64Empty()
	}

	return NewUint64Value(v)
}

//IsEmpty returns true if the contained uint64 value is empty, false otherwise.
func (t Uint64Option) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained uint64 value is initialized, false otherwise.
func (t Uint64Option) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained uint64 value.
//NOTE: If the value is empty, this will return the uint64 zero value.
func (t Uint64Option) Get() uint64 {
	var v uint64

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained uint64 value if the contained value is initialized or the input value is the value is not initialized.
func (t Uint64Option) GetOrElse(v uint64) uint64 {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the Uint64Option type into the JSON representation.
func (t Uint64Option) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the Uint64Option type.
func (t *Uint64Option) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v uint64

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

//MarshalYAML marshals the Uint64Option type into the YAML representation.
func (t Uint64Option) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the Uint64Option type.
func (t *Uint64Option) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v uint64

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

func (t *Uint64Option) String() string {
	return fmt.Sprintf("%v", t.Get())
}
