package typed

type StackOptions struct {
	LimitOptions
}

type StackOption func(*StackOptions)

func defaultStackOptions() StackOptions {
	return StackOptions{
		LimitOptions: DefaultLimitOptions(),
	}
}

func WithStackLimitOptions(limitOpts ...LimitOption) StackOption {
	return func(so *StackOptions) {
		for _, opt := range limitOpts {
			opt(&so.LimitOptions)
		}
	}
}

type Stack[T any] struct {
	items   []T
	pointer int
	opts    StackOptions
}

func NewStack[T any](options ...StackOption) *Stack[T] {
	opts := defaultStackOptions()
	for _, opt := range options {
		opt(&opts)
	}
	return &Stack[T]{opts: opts}
}

func (s *Stack[T]) Len() int      { return s.pointer }
func (s *Stack[T]) Cap() int      { return cap(s.items) }
func (s *Stack[T]) IsEmpty() bool { return s.pointer == 0 }

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
	if s.shouldShrink() {
		s.Shrink()
	}
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
	if s.opts.EnableAutoShrink {
		s.Shrink()
	}
}

func (s *Stack[T]) Shrink() {
	if s.pointer < cap(s.items) {
		newItems := make([]T, s.pointer)
		copy(newItems, s.items[:s.pointer])
		s.items = newItems
	}
}

func (s *Stack[T]) shouldShrink() bool {
	return s.opts.EnableAutoShrink &&
		cap(s.items) > s.opts.ShrinkThresholdCap &&
		float64(s.pointer) < float64(cap(s.items))*s.opts.ShrinkUsageRatio
}

func (s *Stack[T]) ItemsCopy() []T {
	cp := make([]T, s.pointer)
	copy(cp, s.items[:s.pointer])
	return cp
}
