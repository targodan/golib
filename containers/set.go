package containers

import (
	"github.com/targodan/golib/constraints"
	"golang.org/x/exp/slices"
)

type CompareFunc[T any] func(T, T) int

type Set[T any] struct {
	Values      []T
	CompareFunc CompareFunc[T]
}

func NewSet[T any](compareFunc CompareFunc[T]) *Set[T] {
	return &Set[T]{
		Values:      make([]T, 0),
		CompareFunc: compareFunc,
	}
}

func NewOrderedSet[T constraints.Ordered]() *Set[T] {
	return &Set[T]{
		Values: make([]T, 0),
		CompareFunc: func(v1 T, v2 T) int {
			if v1 < v2 {
				return -1
			}
			if v1 == v2 {
				return 0
			}
			return 1
		},
	}
}

func (s *Set[T]) Insert(v T) bool {
	pos, exists := slices.BinarySearchFunc(s.Values, v, s.CompareFunc)
	if exists {
		return false
	}
	if pos == len(s.Values) {
		s.Values = append(s.Values)
	} else {
		s.Values = slices.Insert(s.Values, pos, v)
	}
	return true
}

func (s *Set[T]) Contains(v T) bool {
	_, exists := slices.BinarySearchFunc(s.Values, v, s.CompareFunc)
	return exists
}
