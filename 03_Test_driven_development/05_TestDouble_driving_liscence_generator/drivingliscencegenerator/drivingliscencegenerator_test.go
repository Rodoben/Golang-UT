package drivingliscencegenerator

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type (
	UnderAgeApplicant       struct{}
	LiscenceHolderApplicant struct{}
	FakeRand                struct{}
	ValidApplicant          struct {
		initials string
		dob      string
	}
	SpyLogger struct {
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
func (u UnderAgeApplicant) GetInitials() string {
	return ""
}
func (u UnderAgeApplicant) GetDOB() string {
	return ""
}

func (l LiscenceHolderApplicant) IsAbove18() bool {
	return true
}

func (l LiscenceHolderApplicant) HoldsLiscence() bool {
	return true
}
func (l LiscenceHolderApplicant) GetInitials() string {
	return ""
}
func (l LiscenceHolderApplicant) GetDOB() string {
	return ""
}
func (v ValidApplicant) IsAbove18() bool {
	return true
}

func (v ValidApplicant) HoldsLiscence() bool {
	return false
}
func (v ValidApplicant) GetInitials() string {
	return v.initials
}
func (v ValidApplicant) GetDOB() string {
	return v.dob
}

func (s *SpyLogger) LogStuff(v string) {
	s.Callcount++
	s.Lastmessage = v
}
func (f FakeRand) GetRandomNumbers(len int) string {
	fmt.Println(strings.Repeat("0", len))
	return strings.Repeat("0", len)
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
