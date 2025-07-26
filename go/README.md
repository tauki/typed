# typed

Generic, reusable data structures in Go.

## Example

```go
import "github.com/tauki/typed/go"

s := typed.NewStack[int]()
s.Push(42)
val, _ := s.Pop()
```
