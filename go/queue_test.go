package typed

import (
	"testing"
)

func TestQueue(t *testing.T) {
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
				{"isEmpty", nil, true},
				{"pop", nil, false},
				{"peek", nil, false},
				{"push", 10, nil},
				{"peek", nil, 10},
				{"push", 20, nil},
				{"push", 30, nil},
				{"pop", nil, 10},
				{"pop", nil, 20},
				{"pop", nil, 30},
				{"isEmpty", nil, true},
			},
		},
		{
			name:     "resize and wrap-around",
			dataType: "int",
			steps: []step{
				{"pushMany", 10, 1}, // Push 10 items starting from 1
				{"size", nil, 10},
				{"popMany", 5, 1},                 // Pop 5 items, should start from 1
				{"pushRange", []int{11, 15}, nil}, // Push items 11-15
				{"popSequence", []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, nil},
			},
		},
		{
			name:     "capacity and reset",
			dataType: "int",
			steps: []step{
				{"checkInitialCap", nil, true},
				{"pushMany", 10, 0}, // Push 10 items starting from 0
				{"checkCapIncreased", nil, true},
				{"reset", nil, nil},
				{"isEmpty", nil, true},
				{"size", nil, 0},
				{"push", 42, nil},
				{"peek", nil, 42},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue[int]()

			for i, step := range tt.steps {
				switch step.op {
				case "push":
					q.Push(step.value.(int))
				case "pop":
					val, ok := q.Pop()
					if step.expected != false {
						if !ok || val != step.expected.(int) {
							t.Errorf("step %d: pop expected %v, got %v (ok=%v)", i, step.expected, val, ok)
						}
					} else if ok {
						t.Errorf("step %d: pop expected to fail but succeeded with %v", i, val)
					}
				case "peek":
					val, ok := q.Peek()
					if step.expected != false {
						if !ok || val != step.expected.(int) {
							t.Errorf("step %d: peek expected %v, got %v (ok=%v)", i, step.expected, val, ok)
						}
					} else if ok {
						t.Errorf("step %d: peek expected to fail but succeeded with %v", i, val)
					}
				case "isEmpty":
					got := q.IsEmpty()
					if got != step.expected.(bool) {
						t.Errorf("step %d: isEmpty expected %v, got %v", i, step.expected, got)
					}
				case "size":
					got := q.Size()
					if got != step.expected.(int) {
						t.Errorf("step %d: size expected %v, got %v", i, step.expected, got)
					}
				case "reset":
					q.Reset()
				case "pushMany":
					count := step.value.(int)
					startValue := step.expected.(int)
					for j := 0; j < count; j++ {
						q.Push(startValue + j)
					}
				case "popMany":
					count := step.value.(int)
					startValue := step.expected.(int)
					for j := 0; j < count; j++ {
						val, ok := q.Pop()
						if !ok || val != startValue+j {
							t.Errorf("step %d: popMany[%d] expected %v, got %v (ok=%v)",
								i, j, startValue+j, val, ok)
						}
					}
				case "pushRange":
					values := step.value.([]int)
					start, end := values[0], values[1]
					for j := start; j <= end; j++ {
						q.Push(j)
					}
				case "popSequence":
					expected := step.value.([]int)
					for j, exp := range expected {
						val, ok := q.Pop()
						if !ok {
							t.Fatalf("step %d: popSequence[%d] queue was empty", i, j)
						}
						if val != exp {
							t.Errorf("step %d: popSequence[%d] expected %d, got %d", i, j, exp, val)
						}
					}
				case "checkInitialCap":
					initialCap := q.Cap()
					if initialCap == 0 {
						t.Error("Expected initial capacity to be non-zero")
					}
				case "checkCapIncreased":
					// This assumes pushMany was called earlier
					initialCap := 8 // Default initial capacity
					if q.Cap() <= initialCap {
						t.Errorf("Expected capacity to increase beyond %d, got %d", initialCap, q.Cap())
					}
				default:
					t.Fatalf("step %d: unknown op %s", i, step.op)
				}
			}
		})
	}
}

func TestQueue_Options(t *testing.T) {
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
			var q *Queue[int]
			var expandedCap int

			for i, step := range tt.steps {
				switch step.op {
				case "createWithOptions":
					q = NewQueue[int](WithQueueLimitOptions(
						WithAutoShrink(true),
						WithShrinkThresholdCap(10),
						WithShrinkUsageRatio(0.25),
					))
				case "pushMany":
					count := step.value.(int)
					for j := 0; j < count; j++ {
						q.Push(j)
					}
				case "checkCapacity":
					expandedCap = q.Cap()
				case "popMany":
					count := step.value.(int)
					for j := 0; j < count; j++ {
						q.Pop()
					}
				case "verifyShrink":
					if q.Cap() >= expandedCap {
						t.Errorf("step %d: expected capacity to shrink below %d, got %d",
							i, expandedCap, q.Cap())
					}
				default:
					t.Fatalf("step %d: unknown op %s", i, step.op)
				}
			}
		})
	}
}

// Example of using Queue
func ExampleQueue() {
	// Create a new queue of integers
	q := NewQueue[int]()

	// Add elements to the queue
	q.Push(10)
	q.Push(20)
	q.Push(30)

	// Check if the queue is empty
	isEmpty := q.IsEmpty() // false

	// Get the number of elements
	size := q.Size() // 3

	// Get the capacity of the underlying slice
	capacity := q.Cap()

	// Peek at the front element without removing it
	front, ok := q.Peek() // front = 10, ok = true

	// Remove and get elements from the queue (FIFO order)
	val1, _ := q.Pop() // val1 = 10
	val2, _ := q.Pop() // val2 = 20

	// Reset the queue (clear all elements)
	q.Reset()

	// Prevent unused variable warnings in example
	_, _, _, _, _, _, _ = isEmpty, size, capacity, front, ok, val1, val2
}
