package stack

import "testing"

func Test_Empty(t *testing.T) {
	s := NewStack()
	if s.isEmpty() != true {
		t.Error("stack was not empty")
	}
}

func Test_NotEmpty(t *testing.T) {
	s := NewStack()
	s.Push(Employee{Name: "ronald", Salary: 1232.4})
	if s.isEmpty() != false {
		t.Error("stack was empty")
	}
}

func TestSizeZero(t *testing.T) {
	s := NewStack()
	if s.Size() != 0 {
		t.Error("stack was not empty")
	}
}

func TestSizeOne(t *testing.T) {
	s := NewStack()
	s.Push(Employee{Name: "Ronald", Salary: 23.4})
	if s.Size() != 1 {
		t.Errorf("expected: %d, Actual: %d", 1, s.Size())
	}
}

func TestSizeThree(t *testing.T) {
	s := NewStack()
	s.Push(Employee{Name: "Ronald", Salary: 23.4})
	s.Push(Employee{Name: "Ronald", Salary: 23.4})
	s.Push(Employee{Name: "Ronald", Salary: 23.4})
	if s.Size() != 3 {
		t.Errorf("expected: %d, Actual: %d", 1, s.Size())
	}
}

func TestPopOne(t *testing.T) {
	s := NewStack()
	s.Push(Employee{Name: "Ronald", Salary: 23.4})
	v := s.Pop()
	if v != "Ronald" {
		t.Errorf("Expected Ronald, Actual %s", s.Pop())
	}
	if s.Size() != 0 {
		t.Errorf("Expected 0, Actual %d", s.Size())
	}
}

func TestPopTwo(t *testing.T) {
	s := NewStack()
	s.Push(Employee{Name: "Ronald", Salary: 23.4})
	s.Push(Employee{Name: "Derek", Salary: 23.4})
	v := s.Pop()
	if v != "Derek" {
		t.Errorf("Expected Derek , Actual %s", v)
	}

	v2 := s.Pop()
	if v2 != "Ronald" {
		t.Errorf("Expected Ronald , Actual %s", v2)
	}

}
