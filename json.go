package optioner

import "encoding/json"

// MarshalJSON implements the encoding/json#Marshaler interface
func (t Optioner[T]) MarshalJSON() ([]byte, error) {
	if t.Defined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

// UnmarshalJSON implements the encoding/json#Unmarshaler interface
func (t *Optioner[T]) UnmarshalJSON(data []byte) error {
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
