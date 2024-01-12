package drivingliscencegenerator

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type (
	UnderAgeApplicant       struct{}
	LiscenceHolderApplicant struct{}
	SpyLogger               struct {
		Callcount   int
		Lastmessage string
	}
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

func (s *SpyLogger) LogStuff(v string) {
	s.Callcount++
	s.Lastmessage = v
}

type DrivingLiscenceSuite struct {
	suite.Suite
}

func TestDrivingLiscenceSuite(t *testing.T) {
	suite.Run(t, new(DrivingLiscenceSuite))
}

func (s *DrivingLiscenceSuite) TestUnderAgeApplicant() {
	a := UnderAgeApplicant{}
	l := &SpyLogger{}
	lg := NewDrivingLiscenceNumberGenerator(l)
	_, err := lg.Generate(a)
	s.Error(err)
	s.Contains(err.Error(), "Underaged")

	s.Equal(1, l.Callcount)
	s.Contains(l.Lastmessage, "Underaged")
}

func (s *DrivingLiscenceSuite) TestNoSecondLiscence() {
	a := LiscenceHolderApplicant{}
	l := &SpyLogger{}
	lg := NewDrivingLiscenceNumberGenerator(l)
	_, err := lg.Generate(a)
	s.Error(err)
	s.Contains(err.Error(), "Duplicate")

	s.Equal(1, l.Callcount)
	s.Contains(l.Lastmessage, "Duplicate")
}
