## [Mergesort](https://en.wikipedia.org/wiki/Merge_sort)

[![GoDoc](https://godoc.org/github.com/chasestarr/mergesort?status.svg)](https://godoc.org/github.com/chasestarr/mergesort)

```go
import (
  "fmt"

  "github.com/chasestarr/mergesort"
)

func main() {
  sorted := mergesort.Sort([]int{4, 7, 3, 9, 1, 2})
  fmt.Println(sorted) // [1, 2, 3, 4, 7, 9]
}