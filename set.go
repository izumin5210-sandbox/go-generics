package main

func makeSet[T comparable](capacity int) set[T] {
	values := make(map[T]struct{}, capacity)
	return set[T]{values: values}
}

func setFromSlice[T comparable](slice []T) set[T] {
	values := make(map[T]struct{}, len(slice))
	for _, v := range slice {
		values[v] = struct{}{}
	}
	return set[T]{values: values}
}

type set[T comparable] struct {
	values map[T]struct{}
}

func (s set[T]) Add(value T) {
	s.values[value] = struct{}{}
	// NOTE: compiler panic when use variadic args
	// func (s set[T]) Add(values ...T) {
	// for _, v := range values {
	//   s.values[v] = struct{}{}
	// }
}

func (s set[T]) Contains(value T) bool {
	_, ok := s.values[value]
	return ok
}
