package linkedlist

import (
	"errors"
	"fmt"
)

type Node[T comparable] struct {
	data T
	next *Node[T]
	prev *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (list *LinkedList[T]) Add(element T) bool {
	newNode := &Node[T]{data: element}

	if list.size == 0 {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}

	list.size++
	return true
}

func (list *LinkedList[T]) AddFirst(element T) {
	newNode := &Node[T]{data: element}

	if list.size == 0 {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}

	list.size++
}

func (list *LinkedList[T]) AddLast(element T) {
	list.Add(element)
}

func (list *LinkedList[T]) AddAt(index int, element T) error {
	if index < 0 || index > list.size {
		return errors.New("index out of range")
	}

	if index == 0 {
		list.AddFirst(element)
		return nil
	}

	if index == list.size {
		list.AddLast(element)
		return nil
	}

	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	newNode := &Node[T]{
		data: element,
		next: current,
		prev: current.prev,
	}

	current.prev.next = newNode
	current.prev = newNode
	list.size++

	return nil
}

func (list *LinkedList[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *LinkedList[T]) Contains(element T) bool {
	current := list.head
	for current != nil {
		if current.data == element {
			return true
		}
		current = current.next
	}

	return false
}

func (list *LinkedList[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= list.size {
		return zero, errors.New("index out of range")
	}

	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.data, nil
}

func (list *LinkedList[T]) GetFirst() (T, error) {
	var zero T
	if list.size == 0 {
		return zero, errors.New("list is empty")
	}
	return list.head.data, nil
}

func (list *LinkedList[T]) GetLast() (T, error) {
	var zero T
	if list.size == 0 {
		return zero, errors.New("list is empty")
	}
	return list.tail.data, nil
}

func (list *LinkedList[T]) IndexOf(element T) int {
	current := list.head
	index := 0

	for current != nil {
		if current.data == element {
			return index
		}
		current = current.next
		index++
	}

	return -1
}

func (list *LinkedList[T]) Remove(element T) bool {
	current := list.head

	for current != nil {
		if current.data == element {
			if current.prev == nil {
				return list.RemoveFirst()
			}
			if current.next == nil {
				return list.RemoveLast()
			}

			current.prev.next = current.next
			current.next.prev = current.prev
			list.size--
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList[T]) RemoveAt(index int) (T, error) {
	var zero T
	if index < 0 || index >= list.size {
		return zero, errors.New("index out of range")
	}

	var element T
	if index == 0 {
		element = list.head.data
		list.RemoveFirst()
		return element, nil
	}

	if index == list.size-1 {
		element = list.tail.data
		list.RemoveLast()
		return element, nil
	}

	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	element = current.data
	current.prev.next = current.next
	current.next.prev = current.prev
	list.size--

	return element, nil
}

func (list *LinkedList[T]) RemoveFirst() bool {
	if list.size == 0 {
		return false
	}

	list.head = list.head.next
	if list.head != nil {
		list.head.prev = nil
	} else {
		list.tail = nil
	}

	list.size--
	return true
}

func (list *LinkedList[T]) RemoveLast() bool {
	if list.size == 0 {
		return false
	}

	list.tail = list.tail.prev
	if list.tail != nil {
		list.tail.next = nil
	} else {
		list.head = nil
	}

	list.size--
	return true
}

func (list *LinkedList[T]) Size() int {
	return list.size
}

func (list *LinkedList[T]) IsEmpty() bool {
	return list.size == 0
}

func (list *LinkedList[T]) ToArray() []T {
	array := make([]T, list.size)
	current := list.head
	for i := 0; current != nil; i++ {
		array[i] = current.data
		current = current.next
	}
	return array
}

func (list *LinkedList[T]) ToString() string {
	if list.size == 0 {
		return "[]"
	}

	result := "["
	current := list.head
	for current != nil {
		result += fmt.Sprintf("%v", current.data)
		if current.next != nil {
			result += ", "
		}
		current = current.next
	}
	result += "]"
	return result
}
