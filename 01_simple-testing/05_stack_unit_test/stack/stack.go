package stack

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

type Stack struct {
	value  []Employee
	length int
}

func NewStack() *Stack {
	return &Stack{value: []Employee{}, length: 0}
}

func (s *Stack) isEmpty() bool {
	return s.length == 0
}

func (s *Stack) Push(value Employee) {
	s.value = append(s.value, value)
	s.length++

}
func (s *Stack) Size() int {
	return s.length
}

func (s *Stack) Pop() string {
	s.length--
	lastindex := len(s.value) - 1
	element := s.value[lastindex].Name
	s.value = s.value[:lastindex]
	return element
}

func (s *Stack) DisplayStack() {
	for _, v := range s.value {
		fmt.Println(v)
	}

}
