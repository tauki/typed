package typed

import "testing"

func TestSet_Basic(t *testing.T) {
	s := NewSet[int]()

	if s.Contains(1) {
		t.Error("Expected set to not contain 1 initially")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if !s.Contains(2) {
		t.Error("Expected set to contain 2 after adding")
	}

	s.Remove(2)
	if s.Contains(2) {
		t.Error("Expected set to not contain 2 after removal")
	}

	if s.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", s.Size())
	}

	s.Clear()
	if s.Size() != 0 {
		t.Error("Expected set to be empty after clear")
	}
}
