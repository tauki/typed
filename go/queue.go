package typed

type Queue[T any] struct {
	queue []T
	start int
	end   int
	size  int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{queue: make([]T, 2)}
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

func (q *Queue[T]) resize() {
	l := len(q.queue)
	newCap := (l + 1) * 2
	newQueue := make([]T, newCap)
	for i := 0; i < q.size; i++ {
		newQueue[i] = q.queue[(q.start+i)%l]
	}
	q.queue = newQueue
	q.start = 0
	q.end = q.size
}
