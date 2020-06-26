# slice
--
    import "github.com/iamtabrezkhan/godash/slice"


## Usage

#### func  Chunk

```go
func Chunk(slice interface{}, size int) interface{}
```
Chunk returns a slice of elements split into groups the length of size. If slice
can't be split evenly, the final chunk will be the remaining elements.

Examples:

    Chunk([]int{1, 2, 3, 4}, 2) // ==> [[1, 2], [3, 4]]
