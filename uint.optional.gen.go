package optional

import (
	"encoding/json"
)

// UintOptional contains a initialized or empty copy of the uint type.
type UintOptional struct {
	v *uint
}

// NewUintValue creates a new UintOptional type with an initialized value.
func NewUintValue(v uint) UintOptional {
	return UintOptional{
		v: &v,
	}
}

// NewUintValue creates a new UintOptional type with the given value.
func NewUintPtr(v *uint) UintOptional {
	return UintOptional{
		v: v,
	}
}

// Uint2Ptr create a pointer from the given type.
func Uint2Ptr(v uint) *uint {
	return &v
}

// NewUintValue creates a new UintOptional type with an empty value.
func NewUintEmpty() UintOptional {
	return UintOptional{}
}

// NewUintEmptyIfZeroValue creates a new initialized UintOptional type if the input uint doesn't equal the uint's zero value, or an empty UintOptional otherwise.
func NewUintEmptyIfZeroValue(v uint) UintOptional {
	var e uint

	if v == e {
		return NewUintEmpty()
	}

	return NewUintValue(v)
}

//IsEmpty returns true if the contained uint value is empty, false otherwise.
func (t UintOptional) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained uint value is initialized, false otherwise.
func (t UintOptional) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained uint value.
//NOTE: If the value is empty, this will return the uint zero value.
func (t UintOptional) Get() uint {
	var v uint

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained uint value if the contained value is initialized or the input value is the value is not initialized.
func (t UintOptional) GetOrElse(v uint) uint {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the contained uint value into JSON representation.
func (t UintOptional) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON marshals the contained uint value into JSON representation.
func (t *UintOptional) UnmarshalJSON(data []byte) error {
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
