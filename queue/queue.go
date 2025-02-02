package queue

type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if q.IsEmpty() {
		return zero, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue[T]) Front() (T, bool) {
	var zero T
	if q.IsEmpty() {
		return zero, false
	}

	return q.items[0], true
}

func (q *Queue[T]) Back() (T, bool) {
	var zero T
	if q.IsEmpty() {
		return zero, false
	}

	return q.items[len(q.items)-1], true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}

func (q *Queue[T]) Clear() {
	q.items = make([]T, 0)
}
