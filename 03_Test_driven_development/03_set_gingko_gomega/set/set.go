package set

type Set struct {
	value  map[string]bool
	length int
}

func NewSet() *Set {
	return &Set{make(map[string]bool), 0}
}

func (s *Set) IsEmpty() bool {
	return len(s.value) == 0
}
func (s *Set) Add(value string) {
	s.value[value] = true
	s.length++
}

func (s *Set) Size() int {
	return s.length
}

func (s *Set) Contains(value string) bool {
	if _, ok := s.value[value]; !ok {
		return false
	}
	return true
}
