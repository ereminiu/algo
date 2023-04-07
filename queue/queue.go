package queue

type Queue[T any] struct {
	q []T
}

type Tqueue[T any] Queue[T]

func NewQueue[T any]() *Tqueue[T] {
	return &Tqueue[T]{q: make([]T, 0)}
}

func (qu *Tqueue[T]) Push(s T) {
	qu.q = append(qu.q, s)
}

func (qu *Tqueue[T]) Pop() T {
	last := qu.q[0]
	qu.q = qu.q[1:len(qu.q)]
	return last
}

func (qu *Tqueue[T]) Empty() bool {
	return len(qu.q) == 0
}
