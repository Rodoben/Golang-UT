package main

import (
	"fmt"
	"simple-testing/06_queue/queue"
)

func main() {

	queue1 := queue.NewQueue()
	fmt.Println(queue1)
	queue1.Enqueue(queue.Mobile{BrandName: "Apple", Model: "iphone 14 pro", Storage: 512, Ram: 12, Price: 123405.5})
	queue1.Enqueue(queue.Mobile{BrandName: "Apple", Model: "iphone 13 pro", Storage: 512, Ram: 12, Price: 123405.5})
	fmt.Println(queue1.Size())
	queue1.Displaystack()
	queue1.Dequeue()
	fmt.Println(queue1.Size())
	queue1.Displaystack()
}
