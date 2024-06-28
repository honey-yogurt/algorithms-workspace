package main

type ArrayQueue struct {
	items    []string
	capacity int
	head     int // 下一个出队的下标
	tail     int // 下一个入队的下标
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		items:    make([]string, capacity, capacity),
		capacity: capacity,
		head:     0,
		tail:     0,
	}
}

// 入队
func (q *ArrayQueue) Enqueue(value string) bool {
	if q.tail == q.capacity {
		// 判断是否可以搬迁，集中搬迁
		if q.head == 0 { // 没有空闲空间，不能入队
			return false
		}
		// 所有元素前移 q.head 位
		// 改变 head 以及 tail
		for i := q.head; i < q.tail; i++ {
			q.items[i-q.head] = q.items[i]
		}
		q.tail = q.tail - q.head
		q.head = 0
	}
	q.items[q.tail] = value
	q.tail++
	return true
}

// 出队，出队不做任何数据搬迁
func (q *ArrayQueue) Dequeue() (string, bool) {
	if q.head == q.tail {
		return "", false
	}
	item := q.items[q.head]
	q.head++
	return item, true
}
