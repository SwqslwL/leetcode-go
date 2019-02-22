package queue


type Queue struct {
	front int //下标
	rear  int
	data  []interface{}
}

func (q *Queue)InitQueue() {
	q.data = make([]interface{},100)
	q.rear = 0
	q.front =0
}

func (q *Queue)IsEmpty() bool{
	if q.rear == q.front {
		return true
	} else {
		return false
	}
}

func(q *Queue) EnQueue(Elem interface{}) {
	q.rear += 1
	q.data[q.rear-1] = Elem
}

func (q *Queue) DeQueue()interface{} {
	q.front += 1
	return q.data[q.front-1]
}