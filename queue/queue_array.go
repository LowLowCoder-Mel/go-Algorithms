package queue

type ArrayQueue struct {
	q []interface{}
	capacity int
	head int
	tail int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}

func (q *ArrayQueue) EnQueue(v interface{}) bool {
	if q.tail == q.capacity {
		return false
	}

	q.q[q.tail] = v
	q.tail++
	return true
}

func (q *ArrayQueue) DeQueue() interface{} {
	if q.head == q.tail {
		return nil
	}

	v := q.q[q.head]
	q.head++
	return v
}