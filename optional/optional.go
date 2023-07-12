package optional

import "encoding/json"

type T[V any] struct {
	value *V
}

func Some[V any](value V) T[V] {
	return T[V]{
		value: &value,
	}
}

func None[V any]() T[V] {
	return T[V]{}
}

// Value returns the stored value if it exists or panics.
func (o T[V]) Value() V {
	return *o.value
}

func (o T[V]) ValueOr(d V) V {
	if o.value == nil {
		return d
	}
	return *o.value
}

func (o T[V]) Exists() bool {
	return o.value != nil
}

func (o *T[V]) Set(v V) {
	o.value = &v
}

func (o *T[V]) UnmarshalJSON(b []byte) error {
	var value *V
	err := json.Unmarshal(b, &value)
	if err != nil {
		return err
	}
	o.value = value
	return nil
}

type Bool struct {
	T[bool]
}

func (o Bool) Bool() bool {
	return o.ValueOr(false)
}
