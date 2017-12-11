package utils

// Queue is FIFO
type Queue struct {
	head, tail *element
	size       int
}

// IsEmpty check if the q is empty or not
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Enroll put a value in the queue
func (q *Queue) Enroll(value interface{}) {
	if q.head == nil {
		q.head = &element{value: value}
		q.tail = q.head
	} else {
		temp := &element{value: value}
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
}

// Pull from the head
func (q *Queue) Pull() interface{} {
	if q.size == 0 {
		panic("must enroll before pull")
	}
	v := q.head
	if q.head == q.tail {
		// just one element
		q.tail = q.tail.next

	}
	q.head = q.head.next
	q.size--
	return v.value

}

// Size returns the size of q
func (q *Queue) Size() int {
	return q.size
}

// NewQueue returns a *Queue
func NewQueue() *Queue {
	return &Queue{}
}
