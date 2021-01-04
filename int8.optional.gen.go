package optional

import (
	"encoding/json"
)

// Int8Optional contains a initialized or empty copy of the int8 type.
type Int8Optional struct {
	v *int8
}

// NewInt8Value creates a new Int8Optional type with an initialized value.
func NewInt8Value(v int8) Int8Optional {
	return Int8Optional{
		v: &v,
	}
}

// NewInt8Value creates a new Int8Optional type with the given value.
func NewInt8Ptr(v *int8) Int8Optional {
	return Int8Optional{
		v: v,
	}
}

// Int82Ptr create a pointer from the given type.
func Int82Ptr(v int8) *int8 {
	return &v
}

// NewInt8Value creates a new Int8Optional type with an empty value.
func NewInt8Empty() Int8Optional {
	return Int8Optional{}
}

// NewInt8EmptyIfZeroValue creates a new initialized Int8Optional type if the input int8 doesn't equal the int8's zero value, or an empty Int8Optional otherwise.
func NewInt8EmptyIfZeroValue(v int8) Int8Optional {
	var e int8

	if v == e {
		return NewInt8Empty()
	}

	return NewInt8Value(v)
}

//IsEmpty returns true if the contained int8 value is empty, false otherwise.
func (t Int8Optional) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained int8 value is initialized, false otherwise.
func (t Int8Optional) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained int8 value.
//NOTE: If the value is empty, this will return the int8 zero value.
func (t Int8Optional) Get() int8 {
	var v int8

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained int8 value if the contained value is initialized or the input value is the value is not initialized.
func (t Int8Optional) GetOrElse(v int8) int8 {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the contained int8 value into JSON representation.
func (t Int8Optional) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON marshals the contained int8 value into JSON representation.
func (t *Int8Optional) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v int8

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}
