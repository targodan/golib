package math

import "github.com/targodan/golib/constraints"

func Max[T constraints.Ordered](a, b T) T {
	if a < b {
		return b
	}
	return a
}

func Min[T constraints.Ordered](a, b T) T {
	if a > b {
		return b
	}
	return a
}
