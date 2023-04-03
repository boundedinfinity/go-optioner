package optioner

import "encoding/json"

// MarshalJSON implements the encoding/json#Marshaler interface
//
// If Some[T], T is marshelled normally
//
// If a None[T] a JSON 'null' is marshalled
func (t Option[T]) MarshalJSON() ([]byte, error) {
	if t.Defined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

// UnmarshalJSON implements the encoding/json#Unmarshaler interface
//
// If the value is contained, its value is unmarshalled normally
//
// If the value is missing, or its value is 'null', is is unmarshalled
// as a None[T]
func (t *Option[T]) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v T

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}
