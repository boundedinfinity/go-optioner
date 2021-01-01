
package optional

type StringOptional struct {
	v *string
}

func NewStringValue(v string) StringOptional {
	return StringOptional{
		v: &v,
	}
}

func NewStringEmpty() StringOptional {
	return StringOptional{}
}

func (t StringOptional) IsEmpty() bool {
	return t.v == nil
}

func (t StringOptional) IsDefined() bool {
	return !t.IsEmpty()
}

func (t StringOptional) Get() string {
	return *t.v
}

func (t StringOptional) GetOrElse(v string) string {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

func (t StringOptional) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

func (t *StringOptional) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	var v string

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}
