package drivingliscencegenerator

import (
	"errors"
	"fmt"
)

type (
	DrivingLiscenceApplicants interface {
		IsAbove18() bool
		HoldsLiscence() bool
		GetInitials() string
		GetDOB() string
	}
	DrivingLiscenceNumberGenerator struct {
		l Logger
	}
	Logger interface {
		LogStuff(v string)
	}
)

func NewDrivingLiscenceNumberGenerator(l Logger) *DrivingLiscenceNumberGenerator {
	return &DrivingLiscenceNumberGenerator{l: l}
}

func (g *DrivingLiscenceNumberGenerator) Generate(dlh DrivingLiscenceApplicants) (string, error) {

	if dlh.HoldsLiscence() {
		g.l.LogStuff("Duplicate Applicant, you can only hold one liscence")
		return "", errors.New("Duplicate Applicant, you can only hold one liscence")
	}
	if !dlh.IsAbove18() {
		g.l.LogStuff("Underaged Applicant, you must be 18 to hold liscence")
		return "", errors.New("Underaged Applicant, you must be 18 to hold liscence")
	}

	return fmt.Sprintf("%s%s", dlh.GetInitials(), dlh.GetDOB()), nil
}