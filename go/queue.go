package typed

type QueueOptions struct {
	LimitOptions // Shared shrink options
}

type QueueOption func(*QueueOptions)

func defaultQueueOptions() QueueOptions {
	return QueueOptions{
		LimitOptions: DefaultLimitOptions(),
	}
}

func WithQueueLimitOptions(limitOpts ...LimitOption) QueueOption {
	return func(qo *QueueOptions) {
		for _, opt := range limitOpts {
			opt(&qo.LimitOptions)
		}
	}
}

type Queue[T any] struct {
	queue []T
	start int
	end   int
	size  int
	opts  QueueOptions
}

func NewQueue[T any](opts ...QueueOption) *Queue[T] {
	o := defaultQueueOptions()
	for _, opt := range opts {
		opt(&o)
	}
	return &Queue[T]{
		queue: make([]T, 2),
		opts:  o,
	}
}

func (q *Queue[T]) Push(val T) {
	if q.size+1 == len(q.queue) || len(q.queue) == 0 {
		q.resize()
	}
	q.queue[q.end] = val
	q.size++
	q.end = (q.end + 1) % len(q.queue)
}

func (q *Queue[T]) Pop() (T, bool) {
	var zero T
	if q.IsEmpty() {
		return zero, false
	}
	val := q.queue[q.start]
	q.queue[q.start] = zero
	q.start = (q.start + 1) % len(q.queue)
	q.size--
	if q.shouldShrink() {
		q.shrink()
	}
	return val, true
}

func (q *Queue[T]) Peek() (T, bool) {
	var zero T
	if q.IsEmpty() {
		return zero, false
	}
	return q.queue[q.start], true
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) Cap() int {
	return cap(q.queue)
}

func (q *Queue[T]) Reset() {
	var zero T
	for i := 0; i < q.size; i++ {
		q.queue[(q.start+i)%len(q.queue)] = zero
	}
	q.start = 0
	q.end = 0
	q.size = 0
	if q.opts.EnableAutoShrink {
		q.shrink()
	}
}

func (q *Queue[T]) shouldShrink() bool {
	return q.opts.EnableAutoShrink &&
		cap(q.queue) > q.opts.ShrinkThresholdCap &&
		float64(q.size) < float64(cap(q.queue))*q.opts.ShrinkUsageRatio
}

func (q *Queue[T]) shrink() {
	newQueue := make([]T, q.size)
	for i := 0; i < q.size; i++ {
		newQueue[i] = q.queue[(q.start+i)%len(q.queue)]
	}
	q.queue = newQueue
	q.start = 0
	q.end = q.size
}

func (q *Queue[T]) resize() {
	oldCap := len(q.queue)
	newCap := (oldCap + 1) * 2
	newQueue := make([]T, newCap)
	for i := 0; i < q.size; i++ {
		newQueue[i] = q.queue[(q.start+i)%oldCap]
	}
	q.queue = newQueue
	q.start = 0
	q.end = q.size
}
