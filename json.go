package optioner

import "encoding/json"

func (t Option[T]) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

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
