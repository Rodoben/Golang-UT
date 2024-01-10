package drivingliscencegenerator

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type (
	UnderAgeApplicant       struct{}
	LiscenceHolderApplicant struct{}
)

func (u UnderAgeApplicant) IsAbove18() bool {
	return false
}

func (u UnderAgeApplicant) HoldsLiscence() bool {
	return false
}

func (l LiscenceHolderApplicant) IsAbove18() bool {
	return true
}

func (l LiscenceHolderApplicant) HoldsLiscence() bool {
	return true
}

type DrivingLiscenceSuite struct {
	suite.Suite
}

func TestDrivingLiscenceSuite(t *testing.T) {
	suite.Run(t, new(DrivingLiscenceSuite))
}

func (s *DrivingLiscenceSuite) TestUnderAgeApplicant() {
	a := UnderAgeApplicant{}
	lg := NewDrivingLiscenceNumberGenerator()
	_, err := lg.Generate(a)
	s.Error(err)
	s.Contains(err.Error(), "Underaged")
}

func (s *DrivingLiscenceSuite) TestNoSecondLiscence() {
	a := LiscenceHolderApplicant{}
	lg := NewDrivingLiscenceNumberGenerator()
	_, err := lg.Generate(a)
	s.Error(err)
	s.Contains(err.Error(), "Duplicate")
}
