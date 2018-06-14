package structures

import "errors"

// queue is a basic data structure for integers.
// It follows a FIFO pattern.
// Not suitable for concurrent use
type queue struct {
	q []int
}

func InitQueue() *queue {
	return &queue{make([]int, 0)}
}

func (q *queue) Push(element int) {
	q.q = append(q.q, element)
}

func (q *queue) Pop() (int, error) {
	if len(q.q) == 0 {
		return 0, errors.New("can not pop empty queue")
	}

	e := q.q[0]
	q.q = q.q[1:]
	return e, nil
}

func (q *queue) Size() int {
	return len(q.q)
}
