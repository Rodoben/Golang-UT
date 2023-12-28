package queue

import "fmt"

type Mobile struct {
	BrandName string
	Model     string
	Storage   int
	Ram       int
	Price     float64
}

type Queue struct {
	length int
	value  []Mobile
}

func NewQueue() *Queue {
	return &Queue{value: []Mobile{}, length: 0}
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}
func (q *Queue) Size() int {
	return q.length
}

func (q *Queue) Enqueue(value Mobile) {
	q.value = append(q.value, value)
	q.length++
}

func (q *Queue) Dequeue() Mobile {
	dequeueValue := q.value[0]
	q.value = q.value[1:]
	return dequeueValue

}

func (q *Queue) Displaystack() {
	for _, v := range q.value {
		fmt.Println(v)
	}
}
