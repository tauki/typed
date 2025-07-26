package typed

import (
	"container/heap"
)

type HeapComparator[T any] func(a, b T) bool
type GenericHeap[T any] struct {
	items      []T
	comparator HeapComparator[T] // returns true if a has higher priority than b
}

func NewHeap[T any](comparator HeapComparator[T]) *GenericHeap[T] {
	h := &GenericHeap[T]{comparator: comparator}
	heap.Init(h)
	return h
}

func (h *GenericHeap[T]) Len() int           { return len(h.items) }
func (h *GenericHeap[T]) Less(i, j int) bool { return h.comparator(h.items[i], h.items[j]) }
func (h *GenericHeap[T]) Swap(i, j int)      { h.items[i], h.items[j] = h.items[j], h.items[i] }

func (h *GenericHeap[T]) Push(x any) {
	h.items = append(h.items, x.(T))
}

func (h *GenericHeap[T]) Pop() any {
	n := len(h.items)
	x := h.items[n-1]
	h.items = h.items[:n-1]
	return x
}

func (h *GenericHeap[T]) PushItem(x T) {
	heap.Push(h, x)
}

func (h *GenericHeap[T]) PopItem() (T, bool) {
	var zero T
	if h.Len() == 0 {
		return zero, false
	}
	return heap.Pop(h).(T), true
}

func (h *GenericHeap[T]) Peek() (T, bool) {
	var zero T
	if len(h.items) == 0 {
		return zero, false
	}
	return h.items[0], true
}

func (h *GenericHeap[T]) Size() int {
	return len(h.items)
}

func (h *GenericHeap[T]) ItemsCopy() []T {
	cp := make([]T, len(h.items))
	copy(cp, h.items)
	return cp
}
