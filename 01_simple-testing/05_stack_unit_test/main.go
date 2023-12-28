package main

import (
	"fmt"
	"simple-testing/05_stack/stack"
)

func main() {

	stack1 := stack.NewStack()
	stack1.Push(stack.Employee{Name: "Ronald", Salary: 123.4})
	stack1.Push(stack.Employee{Name: "Donald", Salary: 123.4})
	stack1.Push(stack.Employee{Name: "Derek", Salary: 123.4})
	fmt.Println(stack1.Size())
	stack1.DisplayStack()
	stack1.Pop()
	fmt.Println(stack1.Size())
	stack1.DisplayStack()
}
