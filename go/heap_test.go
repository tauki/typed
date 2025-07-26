package typed

import (
	"testing"
)

func intMinHeap() *Heap[int] {
	return NewHeap[int](func(a, b int) bool {
		return a < b // Min-heap
	})
}

func intMaxHeap() *Heap[int] {
	return NewHeap[int](func(a, b int) bool {
		return a > b // Max-heap
	})
}

func TestGenericHeap_MinHeap(t *testing.T) {
	h := intMinHeap()

	input := []int{5, 3, 8, 1, 6}
	for _, v := range input {
		h.Push(v)
	}

	expectedOrder := []int{1, 3, 5, 6, 8}
	for i, exp := range expectedOrder {
		val, ok := h.Pop()
		if !ok {
			t.Fatalf("Expected value at index %d, but heap was empty", i)
		}
		if val != exp {
			t.Errorf("Expected %d, got %d", exp, val)
		}
	}
}

func TestGenericHeap_MaxHeap(t *testing.T) {
	h := intMaxHeap()

	input := []int{2, 7, 4, 9, 1}
	for _, v := range input {
		h.Push(v)
	}

	expectedOrder := []int{9, 7, 4, 2, 1}
	for i, exp := range expectedOrder {
		val, ok := h.Pop()
		if !ok {
			t.Fatalf("Expected value at index %d, but heap was empty", i)
		}
		if val != exp {
			t.Errorf("Expected %d, got %d", exp, val)
		}
	}
}

func TestGenericHeap_EmptyPopPeek(t *testing.T) {
	h := intMinHeap()

	_, ok := h.Pop()
	if ok {
		t.Error("Expected PopItem to return false on empty heap")
	}

	_, ok = h.Peek()
	if ok {
		t.Error("Expected Peek to return false on empty heap")
	}
}

func TestGenericHeap_Peek(t *testing.T) {
	h := intMinHeap()
	h.Push(10)
	h.Push(5)
	h.Push(8)

	val, ok := h.Peek()
	if !ok || val != 5 {
		t.Errorf("Expected Peek to return 5, got %d (ok=%v)", val, ok)
	}

	h.Pop()
	val, _ = h.Peek()
	if val != 8 {
		t.Errorf("Expected Peek to return 8 after one PopItem, got %d", val)
	}
}

func TestGenericHeap_ItemsCopy(t *testing.T) {
	h := intMinHeap()
	h.Push(4)
	h.Push(2)
	h.Push(6)

	cp := h.ItemsCopy()
	if len(cp) != h.Size() {
		t.Errorf("Expected copied slice of length %d, got %d", h.Size(), len(cp))
	}

	cp[0] = 999 // Modify copy, not original
	// Ensure modifying the copy does not affect the original heap
	for _, v := range h.ItemsCopy() {
		if v == 999 {
			t.Error("Expected original heap to remain unchanged after modifying copy")
		}
	}
}
