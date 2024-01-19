package drivingliscencegenerator

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DrivingLiscenceSuite struct {
	suite.Suite
}

func TestDrivingLiscenceSuite(t *testing.T) {
	suite.Run(t, new(DrivingLiscenceSuite))
}

func (s *DrivingLiscenceSuite) TestUnderAgeApplicant() {
	a := UnderAgeApplicant{}
	l := &SpyLogger{}
	r := FakeRand{}
	lg := NewDrivingLiscenceNumberGenerator(l, r)
	_, err := lg.Generate(a)
	s.Error(err)
	s.Contains(err.Error(), "Underaged")

	s.Equal(1, l.Callcount)
	s.Contains(l.Lastmessage, "Underaged")
}

func (s *DrivingLiscenceSuite) TestNoSecondLiscence() {
	a := LiscenceHolderApplicant{}
	l := &SpyLogger{}
	r := FakeRand{}
	lg := NewDrivingLiscenceNumberGenerator(l, r)
	_, err := lg.Generate(a)
	s.Error(err)
	s.Contains(err.Error(), "Duplicate")

	s.Equal(1, l.Callcount)
	s.Contains(l.Lastmessage, "Duplicate")
}

func (s *DrivingLiscenceSuite) TestLiscenceGenerator() {
	a := ValidApplicant{initials: "JH", dob: "07051997"}
	l := &SpyLogger{}
	r := FakeRand{}
	lg := NewDrivingLiscenceNumberGenerator(l, r)
	ln, err := lg.Generate(a)

	s.NoError(err)
	s.Equal("JH07051997000000", ln)

}
func (s *DrivingLiscenceSuite) TestLiscenceGeneratorShorterInitials() {
	a := ValidApplicant{initials: "JH", dob: "07051997"}
	l := &SpyLogger{}
	r := FakeRand{}
	lg := NewDrivingLiscenceNumberGenerator(l, r)
	ln, err := lg.Generate(a)

	s.NoError(err)
	s.Equal("JH07051997000000", ln)

}
