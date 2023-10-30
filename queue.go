package goqueue

import "sync"

type Queue[T any] struct {
	lock   sync.Mutex
	items  []T
	head   int
	length int
}

// q.resize, doubles the size of the queue while maintaining
// the current position of all the elements
func (q *Queue[T]) resize() {
	items := make([]T, len(q.items)*2)
	for i := 0; i < len(q.items); i++ {
		items[i] = q.items[q.offset(i)]
	}
	q.items = items
	q.head = 0
}

// q.offset gets the array index for the given offset
// from the current head position
func (q *Queue[T]) offset(i int) int {
	return (q.head + i) % len(q.items)
}

func (q *Queue[T]) isFull() bool {
	return q.length == len(q.items)
}

// q.Add adds a new item to the queue
func (q *Queue[T]) Add(item T) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.isFull() {
		q.resize()
	}
	q.items[q.offset(q.length)] = item
	q.length += 1
}

func (q *Queue[T]) Get() (item T, ok bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.length == 0 {
		ok = false
		return
	}
	item = q.items[q.head]
	ok = true
	q.head += 1
	q.length -= 1
	return
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		lock:   sync.Mutex{},
		items:  make([]T, 8),
		head:   0,
		length: 0,
	}
}
