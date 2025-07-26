package typed

import "testing"

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
