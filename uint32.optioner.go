package optioner

import (
	"encoding/json"
	"fmt"
)

// Uint32Option contains a initialized or empty copy of the uint32 type.
type Uint32Option struct {
	v *uint32
}

// NewUint32Value creates a new Uint32Option type with an initialized value.
func NewUint32Value(v uint32) Uint32Option {
	return Uint32Option{
		v: &v,
	}
}

// NewUint32Value creates a new Uint32Option type with the given value.
func NewUint32Ptr(v *uint32) Uint32Option {
	return Uint32Option{
		v: v,
	}
}

// Uint322Ptr create a pointer from the given type.
func Uint322Ptr(v uint32) *uint32 {
	return &v
}

// NewUint32Value creates a new Uint32Option type with an empty value.
func NewUint32Empty() Uint32Option {
	return Uint32Option{}
}

// NewUint32EmptyIfZeroValue creates a new initialized Uint32Option type if the input uint32 doesn't equal the uint32's zero value, or an empty Uint32Option otherwise.
func NewUint32EmptyIfZeroValue(v uint32) Uint32Option {
	var e uint32

	if v == e {
		return NewUint32Empty()
	}

	return NewUint32Value(v)
}

//IsEmpty returns true if the contained uint32 value is empty, false otherwise.
func (t Uint32Option) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained uint32 value is initialized, false otherwise.
func (t Uint32Option) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained uint32 value.
//NOTE: If the value is empty, this will return the uint32 zero value.
func (t Uint32Option) Get() uint32 {
	var v uint32

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained uint32 value if the contained value is initialized or the input value is the value is not initialized.
func (t Uint32Option) GetOrElse(v uint32) uint32 {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the Uint32Option type into the JSON representation.
func (t Uint32Option) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the Uint32Option type.
func (t *Uint32Option) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v uint32

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

//MarshalYAML marshals the Uint32Option type into the YAML representation.
func (t Uint32Option) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the Uint32Option type.
func (t *Uint32Option) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v uint32

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

func (t *Uint32Option) String() string {
	return fmt.Sprintf("%v", t.Get())
}
