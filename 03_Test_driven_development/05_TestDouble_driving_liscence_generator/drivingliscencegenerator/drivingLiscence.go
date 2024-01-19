package drivingliscencegenerator

import (
	"fmt"
	"strings"
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
