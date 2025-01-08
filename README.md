# stack

A concurrent-safe LIFO (Last In, First Out) stack implementation in Go using generics. This package provides a flexible and type-safe stack that supports basic operations like pushing and popping elements, as well as concurrent access using synchronization mechanisms.
Features

    Generics Support: The stack works with any data type, providing type safety and flexibility.
    Concurrency-Safe: The stack is safe for concurrent use by multiple goroutines, using sync.Mutex (or sync.RWMutex for read-heavy workloads).
    Resizable: The stack automatically resizes when elements are added or removed.
    Simple API: Easy-to-use API with clear and predictable behavior.

### Installation

To install the package, run:

```bash
go get github.com/josuedeavila/stack
```

### Usage

Below is an example of how to use the `stack` package.

#### Basic Example

```go
package main

import (
	"fmt"
	"your_module/stack"
)

func main() {
	// Create a stack for integers
	intStack := stack.Stack[int]{}

	// Push elements onto the stack
	intStack.Push(1, 2, 3)

	// Pop elements from the stack
	fmt.Println(intStack.Pop()) // Outputs: 3 true
	fmt.Println(intStack.Pop()) // Outputs: 2 true
	fmt.Println(intStack.Pop()) // Outputs: 1 true
}
```

### Concurrent Usage Example

This package is designed to be safe for concurrent access. Below is an example demonstrating how to use the stack with multiple goroutines.

```go
package main

import (
	"fmt"
	"sync"
	"your_module/stack"
)

func main() {
	var wg sync.WaitGroup
	intStack := stack.Stack[int]{}

	// Launching multiple goroutines to push elements concurrently
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			intStack.Push(i)
			fmt.Printf("Pushed: %d\n", i)
		}(i)
	}

	// Launching multiple goroutines to pop elements concurrently
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if val, ok := intStack.Pop(); ok {
				fmt.Printf("Popped: %d\n", val)
			}
		}()
	}

	wg.Wait() // Wait for all goroutines to finish
}
```

### API Documentation

```go
Stack[T any]
```
A generic stack data structure that holds elements of type T.

#### Methods

`Push(c ...T)`: Adds one or more elements to the stack.

Example:

```go
stack.Push(1, 2, 3) // or stack.Push([]int{1, 2, 3}...)
```

`Pop()`: Removes and returns the top element of the stack. Returns the element and a boolean indicating whether the operation was successful (i.e., the stack was not empty).

Example:

```go
value, ok := stack.Pop()
if ok {
    fmt.Println("Popped:", value)
}
```

`Len()`: Returns the current size of the stack (the number of elements in the stack).

Example:

```go
fmt.Println("Stack size:", stack.Len())
```

### License

This package is released under the MIT License. See the LICENSE file for more information.
