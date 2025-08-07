package typed

import (
	"testing"
)

func TestStack(t *testing.T) {
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
			name: "basic push/pop",
			steps: []step{
				{"push", 1, nil},
				{"push", 2, nil},
				{"push", 3, nil},
				{"pop", nil, 3},
				{"pop", nil, 2},
				{"pop", nil, 1},
			},
		},
		{
			name: "peek and underflow",
			steps: []step{
				{"push", 42, nil},
				{"peek", nil, 42},
				{"pop", nil, 42},
				{"pop", nil, 0}, // underflow returns zero
			},
		},
		{
			name: "interleaved operations",
			steps: []step{
				{"push", 5, nil},
				{"pop", nil, 5},
				{"push", 6, nil},
				{"push", 7, nil},
				{"peek", nil, 7},
				{"pop", nil, 7},
				{"pop", nil, 6},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stack[int]{}
			for i, step := range tt.steps {
				switch step.op {
				case "push":
					s.Push(step.value.(int))
				case "pop":
					got, _ := s.Pop()
					if got != step.expected {
						t.Errorf("step %d: pop expected %v, got %v", i, step.expected, got)
					}
				case "peek":
					got, _ := s.Peek()
					if got != step.expected {
						t.Errorf("step %d: peek expected %v, got %v", i, step.expected, got)
					}
				default:
					t.Fatalf("step %d: unknown op %s", i, step.op)
				}
			}
		})
	}
}

func TestStack_AdditionalMethods(t *testing.T) {
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
			name: "empty stack operations",
			steps: []step{
				{"isEmpty", nil, true},
				{"len", nil, 0},
				{"cap", nil, 0},
			},
		},
		{
			name: "stack with items",
			steps: []step{
				{"push", 10, nil},
				{"push", 20, nil},
				{"push", 30, nil},
				{"isEmpty", nil, false},
				{"len", nil, 3},
				{"itemsCopy", nil, []int{10, 20, 30}},
				{"modifyCopy", nil, false},
				{"reset", nil, nil},
				{"isEmpty", nil, true},
				{"len", nil, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack[int]()

			for i, step := range tt.steps {
				switch step.op {
				case "push":
					s.Push(step.value.(int))
				case "isEmpty":
					got := s.IsEmpty()
					if got != step.expected.(bool) {
						t.Errorf("step %d: isEmpty expected %v, got %v", i, step.expected, got)
					}
				case "len":
					got := s.Len()
					if got != step.expected.(int) {
						t.Errorf("step %d: len expected %v, got %v", i, step.expected, got)
					}
				case "cap":
					got := s.Cap()
					if got != 0 && step.expected.(int) == 0 {
						t.Errorf("step %d: cap expected %v, got %v", i, step.expected, got)
					}
				case "itemsCopy":
					got := s.ItemsCopy()
					expected := step.expected.([]int)

					if len(got) != len(expected) {
						t.Errorf("step %d: itemsCopy expected length %d, got %d", i, len(expected), len(got))
						continue
					}

					for j, v := range expected {
						if got[j] != v {
							t.Errorf("step %d: itemsCopy[%d] expected %v, got %v", i, j, v, got[j])
						}
					}
				case "modifyCopy":
					items := s.ItemsCopy()
					if len(items) > 0 {
						items[0] = 999
						peek, _ := s.Peek()
						if peek == 999 {
							t.Error("Modifying ItemsCopy should not affect the original stack")
						}
					}
				case "reset":
					s.Reset()
				default:
					t.Fatalf("step %d: unknown op %s", i, step.op)
				}
			}
		})
	}
}

func TestStack_Options(t *testing.T) {
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
			name:     "auto-shrink enabled",
			dataType: "int",
			steps: []step{
				{"createWithOptions", nil, nil},
				{"pushMany", 20, nil},
				{"popMany", 15, nil},
				{"checkCapacity", 10, true},
			},
		},
		{
			name:     "manual shrink",
			dataType: "string",
			steps: []step{
				{"create", nil, nil},
				{"pushMany", 10, "item"},
				{"saveCapacity", nil, nil},
				{"popMany", 10, nil},
				{"shrink", nil, nil},
				{"verifyCapacityReduced", nil, true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var intStack *Stack[int]
			var stringStack *Stack[string]
			var initialCap int

			for i, step := range tt.steps {
				switch step.op {
				case "createWithOptions":
					if tt.dataType == "int" {
						intStack = NewStack[int](WithStackLimitOptions(
							WithAutoShrink(true),
							WithShrinkThresholdCap(10),
							WithShrinkUsageRatio(0.25),
						))
					}
				case "create":
					if tt.dataType == "string" {
						stringStack = NewStack[string]()
					}
				case "pushMany":
					count := step.value.(int)
					if tt.dataType == "int" {
						for j := 0; j < count; j++ {
							intStack.Push(j)
						}
					} else {
						item := step.expected.(string)
						for j := 0; j < count; j++ {
							stringStack.Push(item)
						}
					}
				case "popMany":
					count := step.value.(int)
					if tt.dataType == "int" {
						for j := 0; j < count; j++ {
							intStack.Pop()
						}
					} else {
						for j := 0; j < count; j++ {
							stringStack.Pop()
						}
					}
				case "checkCapacity":
					threshold := step.value.(int)
					if intStack.Cap() > threshold {
						t.Errorf("step %d: expected capacity to shrink below %d, got %d",
							i, threshold, intStack.Cap())
					}
				case "saveCapacity":
					if tt.dataType == "string" {
						initialCap = stringStack.Cap()
					}
				case "shrink":
					if tt.dataType == "string" {
						stringStack.Shrink()
					}
				case "verifyCapacityReduced":
					if tt.dataType == "string" {
						if stringStack.Cap() >= initialCap {
							t.Errorf("step %d: expected capacity to decrease from %d, got %d",
								i, initialCap, stringStack.Cap())
						}
					}
				default:
					t.Fatalf("step %d: unknown op %s", i, step.op)
				}
			}
		})
	}
}

// Example of using Stack
func ExampleStack() {
	// Create a new stack of integers
	s := NewStack[int]()

	// Push elements onto the stack
	s.Push(10)
	s.Push(20)
	s.Push(30)

	// Check if the stack is empty
	isEmpty := s.IsEmpty() // false

	// Get the number of elements
	length := s.Len() // 3

	// Peek at the top element without removing it
	top, ok := s.Peek() // top = 30, ok = true

	// Pop elements from the stack (LIFO order)
	val1, _ := s.Pop() // val1 = 30
	val2, _ := s.Pop() // val2 = 20

	// Get a copy of all items in the stack
	items := s.ItemsCopy()

	// Reset the stack (clear all elements)
	s.Reset()

	// Prevent unused variable warnings in example
	_, _, _, _, _, _, _ = isEmpty, length, top, ok, val1, val2, items
}
