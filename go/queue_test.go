package typed

import (
	"testing"
)

func TestQueue_BasicOperations(t *testing.T) {
	q := NewQueue[int]()

	t.Run("empty queue should return false on Pop and Peek", func(t *testing.T) {
		if _, ok := q.Pop(); ok {
			t.Error("Expected Pop to return false on empty queue")
		}
		if _, ok := q.Peek(); ok {
			t.Error("Expected Peek to return false on empty queue")
		}
	})

	t.Run("push and peek", func(t *testing.T) {
		q.Push(10)
		if val, ok := q.Peek(); !ok || val != 10 {
			t.Errorf("Expected Peek to return 10, got %v (ok=%v)", val, ok)
		}
	})

	t.Run("push multiple and pop in order", func(t *testing.T) {
		q.Push(20)
		q.Push(30)

		expected := []int{10, 20, 30}
		for i, exp := range expected {
			val, ok := q.Pop()
			if !ok {
				t.Fatalf("Expected value at index %d, but queue was empty", i)
			}
			if val != exp {
				t.Errorf("Expected %d, got %d", exp, val)
			}
		}
	})

	t.Run("queue should be empty again", func(t *testing.T) {
		if !q.IsEmpty() {
			t.Error("Expected queue to be empty after popping all elements")
		}
	})
}

func TestQueue_ResizeAndWrapAround(t *testing.T) {
	q := NewQueue[int]()

	t.Run("push enough elements to force resize", func(t *testing.T) {
		for i := 1; i <= 10; i++ {
			q.Push(i)
		}
		if q.Size() != 10 {
			t.Errorf("Expected size to be 10, got %d", q.Size())
		}
	})

	t.Run("pop half and then push again to test wrap-around", func(t *testing.T) {
		for i := 1; i <= 5; i++ {
			val, ok := q.Pop()
			if !ok || val != i {
				t.Errorf("Expected %d, got %d", i, val)
			}
		}
		for i := 11; i <= 15; i++ {
			q.Push(i)
		}
	})

	t.Run("check remaining order", func(t *testing.T) {
		expected := []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
		for i, exp := range expected {
			val, ok := q.Pop()
			if !ok {
				t.Fatalf("Expected value at index %d, but queue was empty", i)
			}
			if val != exp {
				t.Errorf("Expected %d, got %d", exp, val)
			}
		}
	})
}
