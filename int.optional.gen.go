package optional

import (
	"encoding/json"
)

// IntOptional contains a initialized or empty copy of the int type.
type IntOptional struct {
	v *int
}

// NewIntValue creates a new IntOptional type with an initialized value.
func NewIntValue(v int) IntOptional {
	return IntOptional{
		v: &v,
	}
}

// NewIntValue creates a new IntOptional type with the given value.
func NewIntPtr(v *int) IntOptional {
	return IntOptional{
		v: v,
	}
}

// Int2Ptr create a pointer from the given type.
func Int2Ptr(v int) *int {
	return &v
}

// NewIntValue creates a new IntOptional type with an empty value.
func NewIntEmpty() IntOptional {
	return IntOptional{}
}

// NewIntEmptyIfZeroValue creates a new initialized IntOptional type if the input int doesn't equal the int's zero value, or an empty IntOptional otherwise.
func NewIntEmptyIfZeroValue(v int) IntOptional {
	var e int

	if v == e {
		return NewIntEmpty()
	}

	return NewIntValue(v)
}

//IsEmpty returns true if the contained int value is empty, false otherwise.
func (t IntOptional) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained int value is initialized, false otherwise.
func (t IntOptional) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained int value.
//NOTE: If the value is empty, this will return the int zero value.
func (t IntOptional) Get() int {
	var v int

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained int value if the contained value is initialized or the input value is the value is not initialized.
func (t IntOptional) GetOrElse(v int) int {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the contained int value into JSON representation.
func (t IntOptional) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON marshals the contained int value into JSON representation.
func (t *IntOptional) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v int

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}
