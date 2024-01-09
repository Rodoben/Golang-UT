package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FizzBuzzSuite struct {
	suite.Suite
}

func TestFizzBuzzSuite(t *testing.T) {
	suite.Run(t, new(FizzBuzzSuite))
}

func (s *FizzBuzzSuite) TestOneUnchanged() {
	r := Run([]int{1})
	s.Equal([]string{"1"}, r)
}
func (s *FizzBuzzSuite) TestTwoUnchanged() {
	r := Run([]int{2})
	s.Equal([]string{"2"}, r)
}
func (s *FizzBuzzSuite) TestThreeUnchanged() {
	r := Run([]int{3})
	s.Equal([]string{"fizz"}, r)
}
func (s *FizzBuzzSuite) TestFiveUnchanged() {
	r := Run([]int{5})
	s.Equal([]string{"buzz"}, r)
}
func (s *FizzBuzzSuite) TestFifteenUnchanged() {
	r := Run([]int{15})
	s.Equal([]string{"fizzbuzz"}, r)
}
func (s *FizzBuzzSuite) TestManyNumbersUnchanged() {
	r := Run([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15, 1})
	s.Equal([]string{"1", "2", "fizz", "4", "buzz", "6", "7", "8", "9", "fizzbuzz", "1"}, r)
}
