package drivingliscencegenerator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type (
	UnderAgeApplicant       struct{}
	LiscenceHolderApplicant struct{}
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

func (s *DrivingLiscenceSuite) TestLiscenceGenerator() {
	a := ValidApplicant{initials: "JH", dob: "07051997"}
	l := &SpyLogger{}

	lg := NewDrivingLiscenceNumberGenerator(l)
	ln, err := lg.Generate(a)
	fmt.Println(ln)
	s.NoError(err)
	s.Equal("JH07051997", ln)

}
