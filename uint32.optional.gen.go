package optional

import (
	"encoding/json"
)

// Uint32Optional contains a initialized or empty copy of the uint32 type.
type Uint32Optional struct {
	v *uint32
}

// NewUint32Value creates a new Uint32Optional type with an initialized value.
func NewUint32Value(v uint32) Uint32Optional {
	return Uint32Optional{
		v: &v,
	}
}

// NewUint32Value creates a new Uint32Optional type with the given value.
func NewUint32Ptr(v *uint32) Uint32Optional {
	return Uint32Optional{
		v: v,
	}
}

// Uint322Ptr create a pointer from the given type.
func Uint322Ptr(v uint32) *uint32 {
	return &v
}

// NewUint32Value creates a new Uint32Optional type with an empty value.
func NewUint32Empty() Uint32Optional {
	return Uint32Optional{}
}

// NewUint32EmptyIfZeroValue creates a new initialized Uint32Optional type if the input uint32 doesn't equal the uint32's zero value, or an empty Uint32Optional otherwise.
func NewUint32EmptyIfZeroValue(v uint32) Uint32Optional {
	var e uint32

	if v == e {
		return NewUint32Empty()
	}

	return NewUint32Value(v)
}

//IsEmpty returns true if the contained uint32 value is empty, false otherwise.
func (t Uint32Optional) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained uint32 value is initialized, false otherwise.
func (t Uint32Optional) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained uint32 value.
//NOTE: If the value is empty, this will return the uint32 zero value.
func (t Uint32Optional) Get() uint32 {
	var v uint32

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained uint32 value if the contained value is initialized or the input value is the value is not initialized.
func (t Uint32Optional) GetOrElse(v uint32) uint32 {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the contained uint32 value into JSON representation.
func (t Uint32Optional) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON marshals the contained uint32 value into JSON representation.
func (t *Uint32Optional) UnmarshalJSON(data []byte) error {
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
