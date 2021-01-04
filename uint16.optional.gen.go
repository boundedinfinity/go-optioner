package optional

import (
	"encoding/json"
)

// Uint16Optional contains a initialized or empty copy of the uint16 type.
type Uint16Optional struct {
	v *uint16
}

// NewUint16Value creates a new Uint16Optional type with an initialized value.
func NewUint16Value(v uint16) Uint16Optional {
	return Uint16Optional{
		v: &v,
	}
}

// NewUint16Value creates a new Uint16Optional type with the given value.
func NewUint16Ptr(v *uint16) Uint16Optional {
	return Uint16Optional{
		v: v,
	}
}

// Uint162Ptr create a pointer from the given type.
func Uint162Ptr(v uint16) *uint16 {
	return &v
}

// NewUint16Value creates a new Uint16Optional type with an empty value.
func NewUint16Empty() Uint16Optional {
	return Uint16Optional{}
}

// NewUint16EmptyIfZeroValue creates a new initialized Uint16Optional type if the input uint16 doesn't equal the uint16's zero value, or an empty Uint16Optional otherwise.
func NewUint16EmptyIfZeroValue(v uint16) Uint16Optional {
	var e uint16

	if v == e {
		return NewUint16Empty()
	}

	return NewUint16Value(v)
}

//IsEmpty returns true if the contained uint16 value is empty, false otherwise.
func (t Uint16Optional) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained uint16 value is initialized, false otherwise.
func (t Uint16Optional) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained uint16 value.
//NOTE: If the value is empty, this will return the uint16 zero value.
func (t Uint16Optional) Get() uint16 {
	var v uint16

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained uint16 value if the contained value is initialized or the input value is the value is not initialized.
func (t Uint16Optional) GetOrElse(v uint16) uint16 {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the contained uint16 value into JSON representation.
func (t Uint16Optional) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON marshals the contained uint16 value into JSON representation.
func (t *Uint16Optional) UnmarshalJSON(data []byte) error {
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
