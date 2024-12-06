package utils

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](slice []T) Set[T] {
	set := make(Set[T])
	for _, el := range slice {
		set[el] = struct{}{}
	}
	return set
}

func (s Set[T]) Intersect(other Set[T]) Set[T] {
	set := make(Set[T])
	for k := range s {
		set[k] = struct{}{}
	}
	for k := range other {
		set[k] = struct{}{}
	}
	return set
}

func (s Set[T]) Contains(el T) bool {
	_, ok := s[el]
	return ok
}

func (s Set[T]) Add(el T) {
	s[el] = struct{}{}
}
