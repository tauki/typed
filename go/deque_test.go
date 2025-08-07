package typed

import (
	"testing"
)

func TestDeque(t *testing.T) {
	type step struct {
		op       string
		value    any
		expected any
	}

	tests := []struct {
		name     string
		dataType string
		steps    []step
	}{
		{
			name:     "basic operations",
			dataType: "int",
			steps: []step{
				{"pushBack", 1, nil},
				{"pushBack", 2, nil},
				{"pushFront", 0, nil},
				{"peekFront", nil, 0},
				{"peekBack", nil, 2},
				{"popFront", nil, 0},
				{"popBack", nil, 2},
				{"popFront", nil, 1},
				{"isEmpty", nil, true},
			},
		},
		{
			name:     "empty operations",
			dataType: "int",
			steps: []step{
				{"isEmpty", nil, true},
				{"size", nil, 0},
				{"popFront", nil, false},
				{"popBack", nil, false},
				{"peekFront", nil, false},
				{"peekBack", nil, false},
			},
		},
		{
			name:     "size and capacity",
			dataType: "int",
			steps: []step{
				{"checkInitialCap", nil, true},
				{"pushBack", 1, nil},
				{"pushFront", 2, nil},
				{"pushBack", 3, nil},
				{"size", nil, 3},
				{"growCapacity", 10, true},
			},
		},
		{
			name:     "reset operations",
			dataType: "string",
			steps: []step{
				{"pushBack", "first", nil},
				{"pushBack", "second", nil},
				{"pushFront", "zero", nil},
				{"reset", nil, nil},
				{"isEmpty", nil, true},
				{"size", nil, 0},
				{"pushBack", "new", nil},
				{"peekBack", nil, "new"},
			},
		},
		{
			name:     "items copy",
			dataType: "int",
			steps: []step{
				{"itemsCopy", nil, []int{}},
				{"pushBack", 10, nil},
				{"pushBack", 20, nil},
				{"pushFront", 0, nil},
				{"itemsCopy", nil, []int{0, 10, 20}},
				{"modifyCopy", nil, false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var intDeque *Deque[int]
			var stringDeque *Deque[string]

			if tt.dataType == "int" {
				intDeque = NewDeque[int]()
			} else {
				stringDeque = NewDeque[string]()
			}

			for i, step := range tt.steps {
				switch step.op {
				case "pushBack":
					if tt.dataType == "int" {
						intDeque.PushBack(step.value.(int))
					} else {
						stringDeque.PushBack(step.value.(string))
					}
				case "pushFront":
					if tt.dataType == "int" {
						intDeque.PushFront(step.value.(int))
					} else {
						stringDeque.PushFront(step.value.(string))
					}
				case "popFront":
					if tt.dataType == "int" {
						val, ok := intDeque.PopFront()
						if step.expected != false {
							if !ok || val != step.expected.(int) {
								t.Errorf("step %d: popFront expected %v, got %v (ok=%v)", i, step.expected, val, ok)
							}
						} else if ok {
							t.Errorf("step %d: popFront expected to fail but succeeded with %v", i, val)
						}
					} else {
						val, ok := stringDeque.PopFront()
						if step.expected != false {
							if !ok || val != step.expected.(string) {
								t.Errorf("step %d: popFront expected %v, got %v (ok=%v)", i, step.expected, val, ok)
							}
						} else if ok {
							t.Errorf("step %d: popFront expected to fail but succeeded with %v", i, val)
						}
					}
				case "popBack":
					if tt.dataType == "int" {
						val, ok := intDeque.PopBack()
						if step.expected != false {
							if !ok || val != step.expected.(int) {
								t.Errorf("step %d: popBack expected %v, got %v (ok=%v)", i, step.expected, val, ok)
							}
						} else if ok {
							t.Errorf("step %d: popBack expected to fail but succeeded with %v", i, val)
						}
					} else {
						val, ok := stringDeque.PopBack()
						if step.expected != false {
							if !ok || val != step.expected.(string) {
								t.Errorf("step %d: popBack expected %v, got %v (ok=%v)", i, step.expected, val, ok)
							}
						} else if ok {
							t.Errorf("step %d: popBack expected to fail but succeeded with %v", i, val)
						}
					}
				case "peekFront":
					if tt.dataType == "int" {
						val, ok := intDeque.PeekFront()
						if step.expected != false {
							if !ok || val != step.expected.(int) {
								t.Errorf("step %d: peekFront expected %v, got %v (ok=%v)", i, step.expected, val, ok)
							}
						} else if ok {
							t.Errorf("step %d: peekFront expected to fail but succeeded with %v", i, val)
						}
					} else {
						val, ok := stringDeque.PeekFront()
						if step.expected != false {
							if !ok || val != step.expected.(string) {
								t.Errorf("step %d: peekFront expected %v, got %v (ok=%v)", i, step.expected, val, ok)
							}
						} else if ok {
							t.Errorf("step %d: peekFront expected to fail but succeeded with %v", i, val)
						}
					}
				case "peekBack":
					if tt.dataType == "int" {
						val, ok := intDeque.PeekBack()
						if step.expected != false {
							if !ok || val != step.expected.(int) {
								t.Errorf("step %d: peekBack expected %v, got %v (ok=%v)", i, step.expected, val, ok)
							}
						} else if ok {
							t.Errorf("step %d: peekBack expected to fail but succeeded with %v", i, val)
						}
					} else {
						val, ok := stringDeque.PeekBack()
						if step.expected != false {
							if !ok || val != step.expected.(string) {
								t.Errorf("step %d: peekBack expected %v, got %v (ok=%v)", i, step.expected, val, ok)
							}
						} else if ok {
							t.Errorf("step %d: peekBack expected to fail but succeeded with %v", i, val)
						}
					}
				case "isEmpty":
					var got bool
					if tt.dataType == "int" {
						got = intDeque.IsEmpty()
					} else {
						got = stringDeque.IsEmpty()
					}
					if got != step.expected.(bool) {
						t.Errorf("step %d: isEmpty expected %v, got %v", i, step.expected, got)
					}
				case "size":
					var got int
					if tt.dataType == "int" {
						got = intDeque.Size()
					} else {
						got = stringDeque.Size()
					}
					if got != step.expected.(int) {
						t.Errorf("step %d: size expected %v, got %v", i, step.expected, got)
					}
				case "reset":
					if tt.dataType == "int" {
						intDeque.Reset()
					} else {
						stringDeque.Reset()
					}
				case "itemsCopy":
					if tt.dataType == "int" {
						got := intDeque.ItemsCopy()
						expected := step.expected.([]int)

						if len(got) != len(expected) {
							t.Errorf("step %d: itemsCopy expected length %d, got %d", i, len(expected), len(got))
							continue
						}

						if len(expected) == 0 {
							continue
						}

						for j, v := range expected {
							if got[j] != v {
								t.Errorf("step %d: itemsCopy[%d] expected %v, got %v", i, j, v, got[j])
							}
						}
					}
				case "checkInitialCap":
					if tt.dataType == "int" {
						initialCap := intDeque.Cap()
						if initialCap == 0 {
							t.Error("Expected initial capacity to be non-zero")
						}
					} else {
						initialCap := stringDeque.Cap()
						if initialCap == 0 {
							t.Error("Expected initial capacity to be non-zero")
						}
					}
				case "growCapacity":
					if tt.dataType == "int" {
						initialCap := intDeque.Cap()
						count := step.value.(int)
						for i := 0; i < count; i++ {
							intDeque.PushBack(i)
						}
						if intDeque.Cap() <= initialCap {
							t.Errorf("Expected capacity to increase, got %d (was %d)", intDeque.Cap(), initialCap)
						}
					}
				case "modifyCopy":
					if tt.dataType == "int" {
						items := intDeque.ItemsCopy()
						if len(items) > 0 {
							items[0] = 999
							val, _ := intDeque.PeekFront()
							if val == 999 {
								t.Error("Modifying ItemsCopy should not affect the original deque")
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

func TestDeque_Options(t *testing.T) {
	type step struct {
		op       string
		value    any
		expected any
	}

	tests := []struct {
		name  string
		steps []step
	}{
		{
			name: "auto-shrink enabled",
			steps: []step{
				{"createWithOptions", nil, nil},
				{"pushMany", 20, nil},
				{"checkCapacity", nil, nil},
				{"popMany", 15, nil},
				{"verifyShrink", nil, true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d *Deque[int]
			var expandedCap int

			for i, step := range tt.steps {
				switch step.op {
				case "createWithOptions":
					d = NewDeque[int](WithDequeLimitOptions(
						WithAutoShrink(true),
						WithShrinkThresholdCap(10),
						WithShrinkUsageRatio(0.25),
					))
				case "pushMany":
					count := step.value.(int)
					for j := 0; j < count; j++ {
						d.PushBack(j)
					}
				case "checkCapacity":
					expandedCap = d.Cap()
				case "popMany":
					count := step.value.(int)
					for j := 0; j < count; j++ {
						d.PopFront()
					}
				case "verifyShrink":
					if d.Cap() >= expandedCap {
						t.Errorf("step %d: expected capacity to shrink below %d, got %d",
							i, expandedCap, d.Cap())
					}
				default:
					t.Fatalf("step %d: unknown op %s", i, step.op)
				}
			}
		})
	}
}

// Example of using Deque
func ExampleDeque() {
	// Create a new deque of integers
	d := NewDeque[int]()

	// Add elements to the back
	d.PushBack(20)
	d.PushBack(30)

	// Add elements to the front
	d.PushFront(10)
	d.PushFront(0)

	// Check if the deque is empty
	isEmpty := d.IsEmpty() // false

	// Get the number of elements
	size := d.Size() // 4

	// Get the capacity of the underlying slice
	capacity := d.Cap()

	// Peek at the front and back elements without removing them
	front, _ := d.PeekFront() // front = 0
	back, _ := d.PeekBack()   // back = 30

	// Remove elements from the front and back
	frontVal, _ := d.PopFront() // frontVal = 0
	backVal, _ := d.PopBack()   // backVal = 30

	// Get a copy of all items in the deque
	items := d.ItemsCopy() // [10, 20]

	// Reset the deque (clear all elements)
	d.Reset()

	// Prevent unused variable warnings in example
	_, _, _, _, _, _, _, _ = isEmpty, size, capacity, front, back, frontVal, backVal, items
}
