package typed

type Stack[T any] struct {
	items   []T
	pointer int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	if s.pointer == len(s.items) {
		s.items = append(s.items, item)
	} else {
		s.items[s.pointer] = item
	}
	s.pointer++
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if s.pointer == 0 {
		return zero, false
	}
	s.pointer--
	val := s.items[s.pointer]
	s.items[s.pointer] = zero
	return val, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if s.pointer == 0 {
		return zero, false
	}
	return s.items[s.pointer-1], true
}

func (s *Stack[T]) Reset() {
	var zero T
	for i := 0; i < s.pointer; i++ {
		s.items[i] = zero
	}
	s.pointer = 0
}

func (s *Stack[T]) IsEmpty() bool {
	return s.pointer == 0
}
