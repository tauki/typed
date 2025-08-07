# typed

Generic, reusable data structures in Go.

## Available Data Structures

This package provides the following generic data structures:

- **Set**: A collection of unique elements with O(1) lookup, insertion, and deletion.
- **Stack**: A Last-In-First-Out (LIFO) collection.
- **Queue**: A First-In-First-Out (FIFO) collection.
- **Deque**: A double-ended queue supporting operations at both ends.
- **Heap**: A priority queue implementation.

## Examples

### Set

```
// Import the package
import "github.com/tauki/typed/go"

// Create a new set of strings
s := typed.NewSet[string]()

// Add elements
s.Add("apple")
s.Add("banana")

// Check if element exists
if s.Contains("apple") {
    // Element exists
}

// Get all values
values := s.Values()

// Remove an element
s.Remove("apple")

// Get the size
size := s.Size()

// Clear the set
s.Clear()
```

### Stack

```
// Import the package
import "github.com/tauki/typed/go"

// Create a new stack of integers
s := typed.NewStack[int]()

// Push elements
s.Push(10)
s.Push(20)

// Peek at the top element without removing it
top, ok := s.Peek() // top = 20, ok = true

// Pop elements (LIFO order)
val, ok := s.Pop() // val = 20, ok = true

// Check if empty and get size
isEmpty := s.IsEmpty()
size := s.Len()
```

### Queue

```
// Import the package
import "github.com/tauki/typed/go"

// Create a new queue of integers
q := typed.NewQueue[int]()

// Add elements
q.Push(10)
q.Push(20)

// Peek at the front element without removing it
front, ok := q.Peek() // front = 10, ok = true

// Remove elements (FIFO order)
val, ok := q.Pop() // val = 10, ok = true

// Check if empty and get size
isEmpty := q.IsEmpty()
size := q.Size()
```

### Deque

```
// Import the package
import "github.com/tauki/typed/go"

// Create a new deque of integers
d := typed.NewDeque[int]()

// Add elements to both ends
d.PushBack(20)
d.PushFront(10)

// Peek at both ends
front, _ := d.PeekFront() // front = 10
back, _ := d.PeekBack()   // back = 20

// Remove from both ends
frontVal, _ := d.PopFront() // frontVal = 10
backVal, _ := d.PopBack()   // backVal = 20
```

### Heap

```
// Import the package
import "github.com/tauki/typed/go"

// Create a min-heap for integers
h := typed.NewHeap[int](func(a, b int) bool {
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
```

## More Examples

For more comprehensive examples, see the example functions in the test files:

- `ExampleSet` in [set_test.go](set_test.go)
- `ExampleStack` in [stack_test.go](stack_test.go)
- `ExampleQueue` in [queue_test.go](queue_test.go)
- `ExampleDeque` in [deque_test.go](deque_test.go)
- `ExampleHeap` in [heap_test.go](heap_test.go)
