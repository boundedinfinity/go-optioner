package optional

import (
	"encoding/json"
)

// Int64Optional contains a initialized or empty copy of the int64 type.
type Int64Optional struct {
	v *int64
}

// NewInt64Value creates a new Int64Optional type with an initialized value.
func NewInt64Value(v int64) Int64Optional {
	return Int64Optional{
		v: &v,
	}
}

// NewInt64Value creates a new Int64Optional type with the given value.
func NewInt64Ptr(v *int64) Int64Optional {
	return Int64Optional{
		v: v,
	}
}

// Int642Ptr create a pointer from the given type.
func Int642Ptr(v int64) *int64 {
	return &v
}

// NewInt64Value creates a new Int64Optional type with an empty value.
func NewInt64Empty() Int64Optional {
	return Int64Optional{}
}

// NewInt64EmptyIfZeroValue creates a new initialized Int64Optional type if the input int64 doesn't equal the int64's zero value, or an empty Int64Optional otherwise.
func NewInt64EmptyIfZeroValue(v int64) Int64Optional {
	var e int64

	if v == e {
		return NewInt64Empty()
	}

	return NewInt64Value(v)
}

//IsEmpty returns true if the contained int64 value is empty, false otherwise.
func (t Int64Optional) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained int64 value is initialized, false otherwise.
func (t Int64Optional) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained int64 value.
//NOTE: If the value is empty, this will return the int64 zero value.
func (t Int64Optional) Get() int64 {
	var v int64

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained int64 value if the contained value is initialized or the input value is the value is not initialized.
func (t Int64Optional) GetOrElse(v int64) int64 {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the contained int64 value into JSON representation.
func (t Int64Optional) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON marshals the contained int64 value into JSON representation.
func (t *Int64Optional) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v int64

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}
