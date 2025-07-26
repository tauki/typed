package typed

import (
	"testing"
)

func intMinHeap() *GenericHeap[int] {
	return NewHeap[int](func(a, b int) bool {
		return a < b // Min-heap
	})
}

func intMaxHeap() *GenericHeap[int] {
	return NewHeap[int](func(a, b int) bool {
		return a > b // Max-heap
	})
}

func TestGenericHeap_MinHeap(t *testing.T) {
	h := intMinHeap()

	// Push elements
	input := []int{5, 3, 8, 1, 6}
	for _, v := range input {
		h.PushItem(v)
	}

	expectedOrder := []int{1, 3, 5, 6, 8}
	for i, exp := range expectedOrder {
		val, ok := h.PopItem()
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
		h.PushItem(v)
	}

	expectedOrder := []int{9, 7, 4, 2, 1}
	for i, exp := range expectedOrder {
		val, ok := h.PopItem()
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

	_, ok := h.PopItem()
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
	h.PushItem(10)
	h.PushItem(5)
	h.PushItem(8)

	val, ok := h.Peek()
	if !ok || val != 5 {
		t.Errorf("Expected Peek to return 5, got %d (ok=%v)", val, ok)
	}

	// Pop and check peek again
	h.PopItem()
	val, _ = h.Peek()
	if val != 8 {
		t.Errorf("Expected Peek to return 8 after one PopItem, got %d", val)
	}
}

func TestGenericHeap_ItemsCopy(t *testing.T) {
	h := intMinHeap()
	h.PushItem(4)
	h.PushItem(2)
	h.PushItem(6)

	cp := h.ItemsCopy()
	if len(cp) != h.Size() {
		t.Errorf("Expected copied slice of length %d, got %d", h.Size(), len(cp))
	}

	cp[0] = 999 // Modify copy, not original
	if h.items[0] == 999 {
		t.Error("ItemsCopy should return a copy, not a reference")
	}
}
