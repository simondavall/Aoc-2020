package aoc

import "errors"

type Queue[T any] []T

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Count() int {
	return len(*q)
}

func (q *Queue[T]) Enqueue(data T) {
	*q = append((*q), data)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var dequeued T
	if q.IsEmpty() {
		return dequeued, errors.New("Queue is empty")
	} else {
		dequeued = (*q)[0]
		*q = (*q)[1:]
	}
	return dequeued, nil
}
