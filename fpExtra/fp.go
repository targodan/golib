// Package fpExtra contains functional programming helpers based on github.com/repeale/fp-go
package fpExtra

import (
	opt "github.com/repeale/fp-go/option"
)

func GetOrElseValue[T any](onNone T) func(opt.Option[T]) T {
	return func(o opt.Option[T]) T {
		if opt.IsNone(o) {
			return onNone
		}

		return o.Value
	}
}

func StringLen[T ~string](s T) int {
	return len(s)
}

func SliceLen[T any](s []T) int {
	return len(s)
}

func SliceCap[T any](s []T) int {
	return cap(s)
}

func MapLen[K comparable, V any](m map[K]V) int {
	return len(m)
}
