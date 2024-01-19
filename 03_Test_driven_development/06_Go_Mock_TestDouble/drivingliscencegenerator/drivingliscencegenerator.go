package drivingliscencegenerator

import (
	"errors"
	"fmt"
)

//go:generate mockgen -destination=./mocks/DrivingLiscenceApplicants.go -package=mocks tdd/06_Go_Mock_TestDouble/drivingliscencegenerator DrivingLiscenceApplicants
//go:generate mockgen -destination=./mocks/RandomNumberGenerator.go -package=mocks tdd/06_Go_Mock_TestDouble/drivingliscencegenerator RandomNumberGenerator
//go:generate mockgen -destination=./mocks/Logger.go -package=mocks tdd/06_Go_Mock_TestDouble/drivingliscencegenerator Logger
type (
	DrivingLiscenceApplicants interface {
		IsAbove18() bool
		HoldsLiscence() bool
		GetInitials() string
		GetDOB() string
	}
	RandomNumberGenerator interface {
		GetRandomNumbers(len int) string
	}
	Logger interface {
		LogStuff(v string)
	}
	DrivingLiscenceNumberGenerator struct {
		l Logger
		r RandomNumberGenerator
	}
)

func NewDrivingLiscenceNumberGenerator(l Logger, r RandomNumberGenerator) *DrivingLiscenceNumberGenerator {
	return &DrivingLiscenceNumberGenerator{l: l, r: r}
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
	n := fmt.Sprintf("%s%s", dlh.GetInitials(), dlh.GetDOB())
	num := 16 - len(n)
	return fmt.Sprintf("%s%s", n, g.r.GetRandomNumbers(num)), nil
}
