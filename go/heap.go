package typed

import (
	"container/heap"

	"github.com/tauki/typed/go/internal"
)

// Comparator is a function that determines priority.
// Should return true if a has higher priority than b.
type Comparator[T any] internal.Comparator[T]

type Heap[T any] struct {
	inner *internal.Heap[T]
}

// NewHeap creates a new heap using the provided comparator.
func NewHeap[T any](cmp Comparator[T]) *Heap[T] {
	return &Heap[T]{inner: internal.NewHeap(cmp)}
}

func (h *Heap[T]) Push(x T) {
	heap.Push(h.inner, x)
}

func (h *Heap[T]) Pop() (T, bool) {
	var zero T
	if h.inner.Len() == 0 {
		return zero, false
	}
	return heap.Pop(h.inner).(T), true
}

func (h *Heap[T]) Peek() (T, bool) {
	return h.inner.Peek()
}

func (h *Heap[T]) Size() int {
	return h.inner.Len()
}

func (h *Heap[T]) ItemsCopy() []T {
	return h.inner.ItemsCopy()
}
