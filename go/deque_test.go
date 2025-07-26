package typed

import "testing"

func TestDeque_Basic(t *testing.T) {
	d := NewDeque[int]()

	d.PushBack(1)
	d.PushBack(2)
	d.PushFront(0)

	if val, _ := d.PeekFront(); val != 0 {
		t.Errorf("Expected PeekFront to be 0, got %d", val)
	}

	if val, _ := d.PeekBack(); val != 2 {
		t.Errorf("Expected PeekBack to be 2, got %d", val)
	}

	if val, _ := d.PopFront(); val != 0 {
		t.Errorf("Expected PopFront to return 0, got %d", val)
	}

	if val, _ := d.PopBack(); val != 2 {
		t.Errorf("Expected PopBack to return 2, got %d", val)
	}

	if val, _ := d.PopFront(); val != 1 {
		t.Errorf("Expected PopFront to return 1, got %d", val)
	}

	if !d.IsEmpty() {
		t.Error("Expected deque to be empty")
	}
}
