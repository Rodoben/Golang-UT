package drivingliscencegenerator

import "errors"

type (
	DrivingLiscenceApplicants interface {
		IsAbove18() bool
		HoldsLiscence() bool
	}
	DrivingLiscenceNumberGenerator struct{}
)

func NewDrivingLiscenceNumberGenerator() *DrivingLiscenceNumberGenerator {

	return &DrivingLiscenceNumberGenerator{}
}

func (g *DrivingLiscenceNumberGenerator) Generate(dlh DrivingLiscenceApplicants) (string, error) {

	if dlh.HoldsLiscence() {
		return "", errors.New("Duplicate Applicant, you can only hold one liscence")
	}

	return "", errors.New("Underaged Applicant, you must be 18 to hold liscence")
}
