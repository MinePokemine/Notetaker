package helpers

type (
	Blank             struct{}
	Set[T comparable] map[T]Blank
)

func (s Set[T]) Add(value T) {
	s[value] = Blank{}
}

func (s Set[T]) Remove(value T) {
	delete(s, value)
}

func (s Set[T]) Contains(value T) bool {
	_, exists := s[value]
	return exists
}

func (s Set[T]) AddAll(other Set[T]) {
	for v, _ := range other {
		s.Add(v)
	}
}

func (s Set[T]) AddAllSlice(slice []T) {
	for _, v := range slice {
		s.Add(v)
	}
}
