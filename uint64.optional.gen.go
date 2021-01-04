package optional

import (
	"encoding/json"
)

// Uint64Optional contains a initialized or empty copy of the uint64 type.
type Uint64Optional struct {
	v *uint64
}

// NewUint64Value creates a new Uint64Optional type with an initialized value.
func NewUint64Value(v uint64) Uint64Optional {
	return Uint64Optional{
		v: &v,
	}
}

// NewUint64Value creates a new Uint64Optional type with the given value.
func NewUint64Ptr(v *uint64) Uint64Optional {
	return Uint64Optional{
		v: v,
	}
}

// Uint642Ptr create a pointer from the given type.
func Uint642Ptr(v uint64) *uint64 {
	return &v
}

// NewUint64Value creates a new Uint64Optional type with an empty value.
func NewUint64Empty() Uint64Optional {
	return Uint64Optional{}
}

// NewUint64EmptyIfZeroValue creates a new initialized Uint64Optional type if the input uint64 doesn't equal the uint64's zero value, or an empty Uint64Optional otherwise.
func NewUint64EmptyIfZeroValue(v uint64) Uint64Optional {
	var e uint64

	if v == e {
		return NewUint64Empty()
	}

	return NewUint64Value(v)
}

//IsEmpty returns true if the contained uint64 value is empty, false otherwise.
func (t Uint64Optional) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained uint64 value is initialized, false otherwise.
func (t Uint64Optional) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained uint64 value.
//NOTE: If the value is empty, this will return the uint64 zero value.
func (t Uint64Optional) Get() uint64 {
	var v uint64

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained uint64 value if the contained value is initialized or the input value is the value is not initialized.
func (t Uint64Optional) GetOrElse(v uint64) uint64 {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the contained uint64 value into JSON representation.
func (t Uint64Optional) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON marshals the contained uint64 value into JSON representation.
func (t *Uint64Optional) UnmarshalJSON(data []byte) error {
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
