# Go Data Structures

A lightweight, generic implementation of common data structures for Go, designed to simplify algorithmic problem solving on platforms like LeetCode, HackerRank, and CodeSignal.

## Features

- **Type-safe**: Built with Go's generics for maximum type safety
- **Simple API**: Clean, intuitive, and consistent interfaces
- **Performant**: Optimized implementations for algorithmic challenges
- **Fully tested**: Comprehensive test coverage
- **Zero dependencies**: No external dependencies required

## Installation

```bash
go get github.com/yourusername/go-data-structures
```

## Quick Start

```go
package main

import (
	"fmt"
	
	"github.com/yourusername/go-data-structures/linkedlist"
	"github.com/yourusername/go-data-structures/queue"
	"github.com/yourusername/go-data-structures/stack"
)

func main() {
	// LinkedList example
	list := linkedlist.NewLinkedList[int]()
	list.Add(1).Add(2).Add(3)
	fmt.Println("LinkedList:", list.ToString()) // [1, 2, 3]
	
	// Queue example
	q := queue.NewQueue[string]()
	q.Enqueue("first")
	q.Enqueue("second")
	item, _ := q.Dequeue()
	fmt.Println("Dequeued item:", item) // first
	
	// Stack example
	s := stack.NewStack[float64]()
	s.Push(1.1)
	s.Push(2.2)
	s.Push(3.3)
	top, _ := s.Pop()
	fmt.Println("Popped item:", top) // 3.3
}
```

## Data Structures

### LinkedList

A doubly-linked list implementation that allows for efficient insertions and deletions.

```go
list := linkedlist.NewLinkedList[int]()
list.Add(10)           // Add to the end
list.AddFirst(5)       // Add to the beginning
list.AddAt(1, 7)       // Insert at specific position
item, _ := list.Get(1) // Get item at index 1
list.Remove(7)         // Remove specific value
list.RemoveAt(0)       // Remove at specific position
array := list.ToArray() // Convert to slice
```

### Queue

A FIFO (First-In-First-Out) data structure.

```go
queue := queue.NewQueue[string]()
queue.Enqueue("task1")       // Add to the end
queue.Enqueue("task2")
item, ok := queue.Dequeue()  // Remove from the front
front, _ := queue.Front()    // Peek at the front item
size := queue.Size()         // Get number of items
isEmpty := queue.IsEmpty()   // Check if empty
```

### Stack

A LIFO (Last-In-First-Out) data structure.

```go
stack := stack.NewStack[int]()
stack.Push(10)           // Add to the top
stack.Push(20)
item, _ := stack.Pop()   // Remove from the top
top, _ := stack.Peek()   // Peek at the top item
size := stack.Size()     // Get number of items
isEmpty := stack.IsEmpty() // Check if empty
```

## Use Cases

These data structures are particularly useful for solving algorithm challenges that require:

- Tracking elements in a specific order
- Implementing BFS (Queue) or DFS (Stack) algorithms
- Managing a collection with frequent insertions/deletions
- Implementing undo/redo functionality (Stack)
- Processing items in arrival order (Queue)

## Why Use This Library?

1. **Focus on the algorithm**: Stop reimplementing basic data structures for every coding challenge
2. **Readability**: Make your solutions more readable and maintainable
3. **Correctness**: Avoid bugs in your data structure implementations
4. **Consistency**: Use a consistent API across different data structures

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Proposal for Go Standard Library

This package is proposed as a potential addition to the Go standard library. The goal is to provide standardized implementations of common data structures that follow Go's design philosophy and make algorithmic problem-solving more accessible to Go developers.

If you find this library useful, please consider expressing your support for its inclusion in the standard library by commenting on the proposal issue [link to proposal issue].