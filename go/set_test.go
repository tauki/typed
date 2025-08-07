package typed

import (
	"sort"
	"testing"
)

func TestSet(t *testing.T) {
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
			name:     "basic operations with integers",
			dataType: "int",
			steps: []step{
				{"contains", 1, false},
				{"add", 1, nil},
				{"add", 2, nil},
				{"add", 3, nil},
				{"contains", 2, true},
				{"remove", 2, nil},
				{"contains", 2, false},
				{"size", nil, 2},
				{"clear", nil, nil},
				{"size", nil, 0},
			},
		},
		{
			name:     "values operations",
			dataType: "int",
			steps: []step{
				{"values", nil, []int{}},
				{"add", 3, nil},
				{"add", 1, nil},
				{"add", 2, nil},
				{"values", nil, []int{1, 2, 3}},
			},
		},
		{
			name:     "edge cases with strings",
			dataType: "string",
			steps: []step{
				{"add", "", nil},
				{"contains", "", true},
				{"add", "hello", nil},
				{"add", "hello", nil}, // Adding duplicate
				{"size", nil, 2},
				{"remove", "world", nil}, // Removing non-existent element
				{"size", nil, 2},
				{"remove", "hello", nil},
				{"contains", "hello", false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var intSet *Set[int]
			var stringSet *Set[string]

			if tt.dataType == "int" {
				intSet = NewSet[int]()
			} else {
				stringSet = NewSet[string]()
			}

			for i, step := range tt.steps {
				switch step.op {
				case "add":
					if tt.dataType == "int" {
						intSet.Add(step.value.(int))
					} else {
						stringSet.Add(step.value.(string))
					}
				case "remove":
					if tt.dataType == "int" {
						intSet.Remove(step.value.(int))
					} else {
						stringSet.Remove(step.value.(string))
					}
				case "contains":
					var got bool
					if tt.dataType == "int" {
						got = intSet.Contains(step.value.(int))
					} else {
						got = stringSet.Contains(step.value.(string))
					}
					if got != step.expected.(bool) {
						t.Errorf("step %d: contains expected %v, got %v", i, step.expected, got)
					}
				case "size":
					var got int
					if tt.dataType == "int" {
						got = intSet.Size()
					} else {
						got = stringSet.Size()
					}
					if got != step.expected.(int) {
						t.Errorf("step %d: size expected %v, got %v", i, step.expected, got)
					}
				case "clear":
					if tt.dataType == "int" {
						intSet.Clear()
					} else {
						stringSet.Clear()
					}
				case "values":
					if tt.dataType == "int" {
						got := intSet.Values()
						expected := step.expected.([]int)

						sort.Ints(got)

						if len(got) != len(expected) {
							t.Errorf("step %d: values expected length %d, got %d", i, len(expected), len(got))
							continue
						}
						if len(expected) == 0 {
							continue
						}
						for j, v := range expected {
							if j < len(got) && got[j] != v {
								t.Errorf("step %d: values[%d] expected %v, got %v", i, j, v, got[j])
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

// Example of using Set
func ExampleSet() {
	// Create a new set of strings
	s := NewSet[string]()

	// Add elements
	s.Add("apple")
	s.Add("banana")
	s.Add("cherry")

	// Check if element exists
	if s.Contains("banana") {
		// Element exists
	}

	// Get the size of the set
	size := s.Size() // 3

	// Get all values (order not guaranteed)
	values := s.Values()

	// Remove an element
	s.Remove("banana")

	// Clear the set
	s.Clear()

	// Prevent unused variable warnings in example
	_, _, _ = size, values, s
}
