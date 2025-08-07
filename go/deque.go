package typed

type DequeOptions struct {
	LimitOptions
}

type DequeOption func(*DequeOptions)

func defaultDequeOptions() DequeOptions {
	return DequeOptions{
		LimitOptions: DefaultLimitOptions(),
	}
}

func WithDequeLimitOptions(limitOpts ...LimitOption) DequeOption {
	return func(do *DequeOptions) {
		for _, opt := range limitOpts {
			opt(&do.LimitOptions)
		}
	}
}

type Deque[T any] struct {
	data        []T
	front, back int
	size        int
	opts        DequeOptions
}

func NewDeque[T any](opts ...DequeOption) *Deque[T] {
	o := defaultDequeOptions()
	for _, opt := range opts {
		opt(&o)
	}
	return &Deque[T]{
		data: make([]T, 4),
		opts: o,
	}
}

func (d *Deque[T]) PushFront(val T) {
	if d.size == len(d.data) {
		d.grow()
	}
	d.front = (d.front - 1 + len(d.data)) % len(d.data)
	d.data[d.front] = val
	d.size++
}

func (d *Deque[T]) PushBack(val T) {
	if d.size == len(d.data) {
		d.grow()
	}
	d.data[d.back] = val
	d.back = (d.back + 1) % len(d.data)
	d.size++
}

func (d *Deque[T]) PopFront() (T, bool) {
	var zero T
	if d.size == 0 {
		return zero, false
	}
	val := d.data[d.front]
	d.front = (d.front + 1) % len(d.data)
	d.size--
	d.maybeShrink()
	return val, true
}

func (d *Deque[T]) PopBack() (T, bool) {
	var zero T
	if d.size == 0 {
		return zero, false
	}
	d.back = (d.back - 1 + len(d.data)) % len(d.data)
	val := d.data[d.back]
	d.size--
	d.maybeShrink()
	return val, true
}

func (d *Deque[T]) PeekFront() (T, bool) {
	var zero T
	if d.size == 0 {
		return zero, false
	}
	return d.data[d.front], true
}

func (d *Deque[T]) PeekBack() (T, bool) {
	var zero T
	if d.size == 0 {
		return zero, false
	}
	return d.data[(d.back-1+len(d.data))%len(d.data)], true
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) Cap() int {
	return cap(d.data)
}

func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}

func (d *Deque[T]) Reset() {
	var zero T
	for i := 0; i < d.size; i++ {
		d.data[(d.front+i)%len(d.data)] = zero
	}
	d.front = 0
	d.back = 0
	d.size = 0
	if d.opts.EnableAutoShrink {
		d.shrink()
	}
}

func (d *Deque[T]) grow() {
	newCap := len(d.data) * 2
	if newCap == 0 {
		newCap = 4
	}
	newData := make([]T, newCap)
	for i := 0; i < d.size; i++ {
		newData[i] = d.data[(d.front+i)%len(d.data)]
	}
	d.data = newData
	d.front = 0
	d.back = d.size
}

func (d *Deque[T]) maybeShrink() {
	if !d.opts.EnableAutoShrink {
		return
	}
	if cap(d.data) > d.opts.ShrinkThresholdCap &&
		float64(d.size) < float64(cap(d.data))*d.opts.ShrinkUsageRatio {
		d.shrink()
	}
}

func (d *Deque[T]) shrink() {
	newData := make([]T, d.size)
	for i := 0; i < d.size; i++ {
		newData[i] = d.data[(d.front+i)%len(d.data)]
	}
	d.data = newData
	d.front = 0
	d.back = d.size
}

func (d *Deque[T]) ItemsCopy() []T {
	cp := make([]T, d.size)
	for i := 0; i < d.size; i++ {
		cp[i] = d.data[(d.front+i)%len(d.data)]
	}
	return cp
}
