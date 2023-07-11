package pair

import (
	"github.com/targodan/golib/math"

	opt "github.com/repeale/fp-go/option"
)

type Pair[L any, R any] struct {
	Left  L
	Right R
}

func (p Pair[L, R]) Values() (L, R) {
	return p.Left, p.Right
}

func Reduce[L any, R any, T any](f func(L, R) T) func(Pair[L, R]) T {
	return func(p Pair[L, R]) T {
		return f(p.Left, p.Right)
	}
}

func Zip[L any, R any](left []L, right []R) []Pair[L, R] {
	size := math.Min(len(left), len(right))

	result := make([]Pair[L, R], size)
	for i := 0; i < size; i++ {
		result[i].Left = left[i]
		result[i].Right = right[i]
	}
	return result
}

func ZipLongest[L any, R any](left []L, right []R) []Pair[opt.Option[L], opt.Option[R]] {
	size := math.Max(len(left), len(right))

	result := make([]Pair[opt.Option[L], opt.Option[R]], size)
	for i := 0; i < size; i++ {
		if i < len(left) {
			result[i].Left = opt.Some(left[i])
		} else {
			result[i].Left = opt.None[L]()
		}
		if i < len(right) {
			result[i].Right = opt.Some(right[i])
		} else {
			result[i].Right = opt.None[R]()
		}
	}
	return result
}

func Map[L any, R any, TL any, TR any](callbackL func(L) TL, callbackR func(R) TR) func([]Pair[L, R]) []Pair[TL, TR] {
	return func(xs []Pair[L, R]) []Pair[TL, TR] {
		result := make([]Pair[TL, TR], len(xs))

		for i, x := range xs {
			result[i].Left = callbackL(x.Left)
			result[i].Right = callbackR(x.Right)
		}

		return result
	}
}

func MapRight[L any, R any, T any](callback func(R) T) func([]Pair[L, R]) []Pair[L, T] {
	return func(xs []Pair[L, R]) []Pair[L, T] {
		result := make([]Pair[L, T], len(xs))

		for i, x := range xs {
			result[i].Left = x.Left
			result[i].Right = callback(x.Right)
		}

		return result
	}
}

func MapLeft[L any, R any, T any](callback func(L) T) func([]Pair[L, R]) []Pair[T, R] {
	return func(xs []Pair[L, R]) []Pair[T, R] {
		result := make([]Pair[T, R], len(xs))

		for i, x := range xs {
			result[i].Left = callback(x.Left)
			result[i].Right = x.Right
		}

		return result
	}
}
