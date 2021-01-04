package optional

import (
	"encoding/json"
)

// Float32Optional contains a initialized or empty copy of the float32 type.
type Float32Optional struct {
	v *float32
}

// NewFloat32Value creates a new Float32Optional type with an initialized value.
func NewFloat32Value(v float32) Float32Optional {
	return Float32Optional{
		v: &v,
	}
}

// NewFloat32Value creates a new Float32Optional type with the given value.
func NewFloat32Ptr(v *float32) Float32Optional {
	return Float32Optional{
		v: v,
	}
}

// Float322Ptr create a pointer from the given type.
func Float322Ptr(v float32) *float32 {
	return &v
}

// NewFloat32Value creates a new Float32Optional type with an empty value.
func NewFloat32Empty() Float32Optional {
	return Float32Optional{}
}

// NewFloat32EmptyIfZeroValue creates a new initialized Float32Optional type if the input float32 doesn't equal the float32's zero value, or an empty Float32Optional otherwise.
func NewFloat32EmptyIfZeroValue(v float32) Float32Optional {
	var e float32

	if v == e {
		return NewFloat32Empty()
	}

	return NewFloat32Value(v)
}

//IsEmpty returns true if the contained float32 value is empty, false otherwise.
func (t Float32Optional) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained float32 value is initialized, false otherwise.
func (t Float32Optional) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained float32 value.
//NOTE: If the value is empty, this will return the float32 zero value.
func (t Float32Optional) Get() float32 {
	var v float32

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained float32 value if the contained value is initialized or the input value is the value is not initialized.
func (t Float32Optional) GetOrElse(v float32) float32 {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the contained float32 value into JSON representation.
func (t Float32Optional) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON marshals the contained float32 value into JSON representation.
func (t *Float32Optional) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v float32

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}
