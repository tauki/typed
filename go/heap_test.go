package typed

import (
	"math/rand"
	"testing"
)

func TestHeap(t *testing.T) {
	type step struct {
		op       string
		value    any
		expected any
	}

	tests := []struct {
		name     string
		heapType string
		steps    []step
	}{
		{
			name:     "min heap operations",
			heapType: "min",
			steps: []step{
				{"pushMany", []int{5, 3, 8, 1, 6}, nil},
				{"popSequence", nil, []int{1, 3, 5, 6, 8}},
			},
		},
		{
			name:     "max heap operations",
			heapType: "max",
			steps: []step{
				{"pushMany", []int{2, 7, 4, 9, 1}, nil},
				{"popSequence", nil, []int{9, 7, 4, 2, 1}},
			},
		},
		{
			name:     "empty heap operations",
			heapType: "min",
			steps: []step{
				{"pop", nil, false},
				{"peek", nil, false},
			},
		},
		{
			name:     "peek operations",
			heapType: "min",
			steps: []step{
				{"push", 10, nil},
				{"push", 5, nil},
				{"push", 8, nil},
				{"peek", nil, 5},
				{"pop", nil, 5},
				{"peek", nil, 8},
			},
		},
		{
			name:     "items copy",
			heapType: "min",
			steps: []step{
				{"push", 4, nil},
				{"push", 2, nil},
				{"push", 6, nil},
				{"itemsCopy", nil, 3},
				{"modifyCopy", nil, false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var h *Heap[int]

			// Initialize the appropriate heap type
			if tt.heapType == "min" {
				h = NewHeap[int](func(a, b int) bool {
					return a < b // Min-heap
				})
			} else {
				h = NewHeap[int](func(a, b int) bool {
					return a > b // Max-heap
				})
			}

			for i, step := range tt.steps {
				switch step.op {
				case "push":
					h.Push(step.value.(int))
				case "pushMany":
					values := step.value.([]int)
					for _, v := range values {
						h.Push(v)
					}
				case "pop":
					val, ok := h.Pop()
					if step.expected != false {
						if !ok || val != step.expected.(int) {
							t.Errorf("step %d: pop expected %v, got %v (ok=%v)", i, step.expected, val, ok)
						}
					} else if ok {
						t.Errorf("step %d: pop expected to fail but succeeded with %v", i, val)
					}
				case "popSequence":
					expected := step.expected.([]int)
					for j, exp := range expected {
						val, ok := h.Pop()
						if !ok {
							t.Fatalf("step %d: popSequence[%d] heap was empty", i, j)
						}
						if val != exp {
							t.Errorf("step %d: popSequence[%d] expected %d, got %d", i, j, exp, val)
						}
					}
				case "peek":
					val, ok := h.Peek()
					if step.expected != false {
						if !ok || val != step.expected.(int) {
							t.Errorf("step %d: peek expected %v, got %v (ok=%v)", i, step.expected, val, ok)
						}
					} else if ok {
						t.Errorf("step %d: peek expected to fail but succeeded with %v", i, val)
					}
				case "itemsCopy":
					cp := h.ItemsCopy()
					expectedSize := step.expected.(int)
					if len(cp) != expectedSize {
						t.Errorf("step %d: itemsCopy expected length %d, got %d", i, expectedSize, len(cp))
					}
				case "modifyCopy":
					cp := h.ItemsCopy()
					if len(cp) > 0 {
						cp[0] = 999 // Modify copy, not original
						// Ensure modifying the copy does not affect the original heap
						for _, v := range h.ItemsCopy() {
							if v == 999 {
								t.Error("Modifying ItemsCopy should not affect the original heap")
							}
						}
					}
				default:
					t.Fatalf("step %d: unknown op %s", i, step.op)
				}
			}
		})
	}
}

// TestHeap_CustomType tests the heap with a custom struct type for both min and max heap
func TestHeap_CustomType(t *testing.T) {
	type Task struct {
		ID       int
		Priority int
		Name     string
	}

	// Add tasks with different priorities
	tasks := []Task{
		{ID: 1, Priority: 5, Name: "Write documentation"},
		{ID: 2, Priority: 3, Name: "Fix bug"},
		{ID: 3, Priority: 7, Name: "Add feature"},
		{ID: 4, Priority: 1, Name: "Critical security fix"},
		{ID: 5, Priority: 4, Name: "Code review"},
	}

	tests := []struct {
		name          string
		comparator    func(a, b Task) bool
		expectedOrder []int
	}{
		{
			name:          "min heap by priority",
			comparator:    func(a, b Task) bool { return a.Priority < b.Priority },
			expectedOrder: []int{4, 2, 5, 1, 3},
		},
		{
			name:          "max heap by priority",
			comparator:    func(a, b Task) bool { return a.Priority > b.Priority },
			expectedOrder: []int{3, 1, 5, 2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHeap[Task](tt.comparator)
			for _, task := range tasks {
				h.Push(task)
			}
			for i, expectedID := range tt.expectedOrder {
				task, ok := h.Pop()
				if !ok {
					t.Fatalf("Expected task at index %d, but heap was empty", i)
				}
				if task.ID != expectedID {
					t.Errorf("Expected task with ID %d at position %d, got task with ID %d",
						expectedID, i, task.ID)
				}
			}

			// Verify heap is now empty
			if h.Size() != 0 {
				t.Errorf("Expected heap to be empty, but size is %d", h.Size())
			}
		})
	}
}

// TestHeap_PropertyMaintained verifies that the heap property is maintained after operations
func TestHeap_PropertyMaintained(t *testing.T) {
	h := NewHeap[int](func(a, b int) bool {
		return a < b // Min-heap
	})

	// Add random elements
	for i := 0; i < 100; i++ {
		h.Push(rand.Intn(1000))
	}

	// Verify heap property by popping elements
	prev, _ := h.Pop()
	for h.Size() > 0 {
		curr, _ := h.Pop()
		if curr < prev { // For min-heap, elements should come out in ascending order
			t.Errorf("Heap property violated: %d came after %d", curr, prev)
		}
		prev = curr
	}
}

// Example of using Heap
func ExampleHeap() {
	// Create a min-heap for integers
	h := NewHeap[int](func(a, b int) bool {
		return a < b // Min-heap (smallest value has highest priority)
	})

	// Add elements
	h.Push(5)
	h.Push(3)
	h.Push(8)

	// Peek at highest priority element without removing it
	top, _ := h.Peek() // top = 3

	// Remove highest priority element
	val, _ := h.Pop() // val = 3

	// Get the size of the heap
	size := h.Size() // 2

	// Get a copy of all items in the heap
	items := h.ItemsCopy()

	// Prevent unused variable warnings in example
	_, _, _, _ = top, val, size, items
}
