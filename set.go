package set

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(Set[T])
}

func NewWith[T comparable](value ...T) Set[T] {
	set := make(Set[T], len(value))
	for _, s := range value {
		set.Add(s)
	}
	return set
}

func (s Set[T]) Add(k T) {
	s[k] = struct{}{}
}

func (s Set[T]) Remove(k T) {
	delete(s, k)
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Has(k T) bool {
	_, exists := s[k]
	return exists
}

func (s Set[T]) Foreach(fun func(T)) {
	for value := range s {
		fun(value)
	}
}

func (s Set[T]) ToSlice() []T {
	list := make([]T, 0, s.Len())
	for value := range s {
		list = append(list, value)
	}
	return list
}
