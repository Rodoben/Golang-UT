package queue

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// first import the testify package
//go get github.com/stretchr/testify/suite

type QueueSuite struct {
	suite.Suite
	Queue *Queue
}

func (q *QueueSuite) SetupTest() {
	q.Queue = NewQueue()

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(QueueSuite))
}

func (q *QueueSuite) Test_QueueEmpty() {
	q.True(q.Queue.IsEmpty())
}

func (q *QueueSuite) Test_QueueNotEmpty() {
	q.Queue.Enqueue(Mobile{BrandName: "Apple", Model: "iphone 14 pro", Storage: 512, Ram: 12, Price: 123405.5})
	q.False(q.Queue.IsEmpty())
}

func (q *QueueSuite) Test_QueueSizeZero() {

	q.Zero(q.Queue.Size())
}
func (q *QueueSuite) Test_QueueSizeOne() {

	q.Queue.Enqueue(Mobile{BrandName: "Apple", Model: "iphone 14 pro", Storage: 512, Ram: 12, Price: 123405.5})
	q.Equal(1, q.Queue.Size())
}
func (q *QueueSuite) Test_QueueSizeTwo() {
	q.Queue.Enqueue(Mobile{BrandName: "Apple", Model: "iphone 14 pro", Storage: 512, Ram: 12, Price: 123405.5})
	q.Queue.Enqueue(Mobile{BrandName: "Apple", Model: "iphone 14 pro", Storage: 512, Ram: 12, Price: 123405.5})
	q.Equal(2, q.Queue.Size())
}

func (q *QueueSuite) Test_Queue_Dequeue_One() {

	q.Queue.Enqueue(Mobile{BrandName: "Apple", Model: "iphone 14 pro", Storage: 512, Ram: 12, Price: 123405.5})
	q.Queue.Enqueue(Mobile{BrandName: "Apple", Model: "iphone 13 pro", Storage: 512, Ram: 12, Price: 123405.5})
	q.Equal(Mobile{BrandName: "Apple", Model: "iphone 14 pro", Storage: 512, Ram: 12, Price: 123405.5}, q.Queue.Dequeue())
}

func (q *QueueSuite) Test_Queue_Dequeue_Two() {

	q.Queue.Enqueue(Mobile{BrandName: "Apple", Model: "iphone 14 pro", Storage: 512, Ram: 12, Price: 123405.5})
	q.Queue.Enqueue(Mobile{BrandName: "Apple", Model: "iphone 13 pro", Storage: 512, Ram: 12, Price: 123405.5})
	q.Equal(Mobile{BrandName: "Apple", Model: "iphone 14 pro", Storage: 512, Ram: 12, Price: 123405.5}, q.Queue.Dequeue())
	q.Equal(Mobile{BrandName: "Apple", Model: "iphone 13 pro", Storage: 512, Ram: 12, Price: 123405.5}, q.Queue.Dequeue())
}
