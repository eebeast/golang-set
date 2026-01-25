# Set

A generic, type-safe set data structure implementation for Go.

## Features

- ✅ **Generic**: Works with any comparable type (string, int, float64, custom types, etc.)
- ✅ **Type-safe**: Compile-time type safety with Go generics
- ✅ **Efficient**: O(1) average-case operations using `map[T]struct{}`
- ✅ **Zero dependencies**: No external dependencies required
- ✅ **Simple API**: Clean and intuitive interface

## Installation

```bash
go get github.com/eebeast/set
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/eebeast/set"
)

func main() {
    // Create a new empty set
    s := set.New[string]()
    s.Add("hello")
    s.Add("world")
    
    fmt.Println(s.Has("hello")) // true
    fmt.Println(s.Len())        // 2
    
    // Create a set with initial values
    numbers := set.NewWith(1, 2, 3, 4, 5)
    fmt.Println(numbers.Len()) // 5
    
    // Remove an element
    numbers.Remove(3)
    fmt.Println(numbers.Has(3)) // false
}
```

## API Documentation

### Types

#### `Set[T comparable]`

A set of unique elements of type `T`. The zero value is a ready-to-use empty set.

**Note**: Sets are not safe for concurrent use by multiple goroutines without additional synchronization.

### Functions

#### `New[T comparable]() Set[T]`

Creates and returns a new empty set.

```go
s := set.New[string]()
```

#### `NewWith[T comparable](values ...T) Set[T]`

Creates a new set with the given initial values. Duplicate values are automatically deduplicated.

```go
s := set.NewWith("a", "b", "c")
s2 := set.NewWith(1, 2, 3, 2, 1) // Results in {1, 2, 3}
```

### Methods

#### `Add(k T)`

Adds an element to the set. If the element already exists, the set remains unchanged.

**Time complexity**: O(1) average case

```go
s := set.New[string]()
s.Add("hello")
s.Add("world")
```

#### `Remove(k T)`

Removes an element from the set. If the element doesn't exist, the operation has no effect.

**Time complexity**: O(1) average case

```go
s := set.NewWith("a", "b", "c")
s.Remove("b")
```

#### `Has(k T) bool`

Returns `true` if the element exists in the set, `false` otherwise.

**Time complexity**: O(1) average case

```go
s := set.NewWith("a", "b", "c")
if s.Has("a") {
    fmt.Println("Found!")
}
```

#### `Len() int`

Returns the number of elements in the set.

**Time complexity**: O(1)

```go
s := set.NewWith("a", "b", "c")
fmt.Println(s.Len()) // 3
```

#### `Foreach(fun func(T))`

Calls the provided function for each element in the set. The iteration order is not guaranteed.

```go
s := set.NewWith(1, 2, 3, 4, 5)
sum := 0
s.Foreach(func(value int) {
    sum += value
})
fmt.Println(sum) // 15
```

#### `ToSlice() []T`

Returns a slice containing all elements in the set. The order of elements is not guaranteed.

```go
s := set.NewWith("a", "b", "c")
slice := s.ToSlice()
fmt.Println(slice) // ["a", "b", "c"] (order may vary)
```

## Examples

### Working with Different Types

```go
// String set
strSet := set.NewWith("apple", "banana", "cherry")

// Integer set
intSet := set.New[int]()
intSet.Add(1)
intSet.Add(2)
intSet.Add(3)

// Float set
floatSet := set.NewWith(1.1, 2.2, 3.3)

// Custom type set
type Person struct {
    Name string
    Age  int
}

personSet := set.New[Person]()
personSet.Add(Person{Name: "Alice", Age: 30})
personSet.Add(Person{Name: "Bob", Age: 25})
```

### Iterating Over Elements

```go
s := set.NewWith("red", "green", "blue")

// Using Foreach
s.Foreach(func(color string) {
    fmt.Println(color)
})

// Using ToSlice
colors := s.ToSlice()
for _, color := range colors {
    fmt.Println(color)
}

// Direct range (since Set is a map)
for color := range s {
    fmt.Println(color)
}
```

### Working with Empty Sets

```go
s := set.New[string]()
fmt.Println(s.Len())        // 0
fmt.Println(s.Has("test"))  // false

// Safe to call methods on empty set
s.Remove("nonexistent")     // No error
slice := s.ToSlice()        // Returns empty slice
```

## Thread Safety

Sets are **not** thread-safe. For concurrent access, use external synchronization:

```go
import "sync"

var mu sync.RWMutex
s := set.New[string]()

// Safe concurrent access
mu.Lock()
s.Add("value")
mu.Unlock()

mu.RLock()
exists := s.Has("value")
mu.RUnlock()
```

## Requirements

- Go 1.21 or later (for generics support)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

[eebeast](https://github.com/eebeast)
