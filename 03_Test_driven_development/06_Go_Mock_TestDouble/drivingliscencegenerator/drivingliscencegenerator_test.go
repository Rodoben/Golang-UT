package drivingliscencegenerator

import (
	"tdd/06_Go_Mock_TestDouble/drivingliscencegenerator/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

var (
	ctrl *gomock.Controller
	a    *mocks.MockDrivingLiscenceApplicants
	l    *mocks.MockLogger
	r    *mocks.MockRandomNumberGenerator
	lg   *DrivingLiscenceNumberGenerator
)

func (s *DrivingLiscenceSuite) SetupTest() {
	ctrl = gomock.NewController(s.T())
	a = mocks.NewMockDrivingLiscenceApplicants(ctrl)
	l = mocks.NewMockLogger(ctrl)
	r = mocks.NewMockRandomNumberGenerator(ctrl)
	lg = NewDrivingLiscenceNumberGenerator(l, r)
}

func (s *DrivingLiscenceSuite) TearDownTest() {
	ctrl.Finish()
}

type DrivingLiscenceSuite struct {
	suite.Suite
}

func TestDrivingLiscenceSuite(t *testing.T) {
	suite.Run(t, new(DrivingLiscenceSuite))
}

func (s *DrivingLiscenceSuite) TestUnderAgeApplicant() {
	a.EXPECT().IsAbove18().Return(false)
	a.EXPECT().HoldsLiscence().Return(false)
	l.EXPECT().LogStuff("Underaged Applicant, you must be 18 to hold liscence").Times(1)
	_, err := lg.Generate(a)
	s.Error(err)
	s.Contains(err.Error(), "Underaged")

}

func (s *DrivingLiscenceSuite) TestNoSecondLiscence() {
	a.EXPECT().HoldsLiscence().Return(true)
	l.EXPECT().LogStuff("Duplicate Applicant, you can only hold one liscence").Times(1)
	_, err := lg.Generate(a)
	s.Error(err)
	s.Contains(err.Error(), "Duplicate")
}

func (s *DrivingLiscenceSuite) TestLiscenceGenerator() {
	a.EXPECT().HoldsLiscence().Return(false)
	a.EXPECT().IsAbove18().Return(true)

	a.EXPECT().GetInitials().Return("JH")
	a.EXPECT().GetDOB().Return("07051997")
	r.EXPECT().GetRandomNumbers(gomock.Any()).Return("00000")

	lg := NewDrivingLiscenceNumberGenerator(l, r)
	ln, err := lg.Generate(a)

	s.NoError(err)
	s.Equal("JH0705199700000", ln)

}
func (s *DrivingLiscenceSuite) TestLiscenceGeneratorShorterInitials() {
	a := ValidApplicant{initials: "JH", dob: "07051997"}
	l := &SpyLogger{}
	//r := FakeRand{}
	lg := NewDrivingLiscenceNumberGenerator(l, r)
	ln, err := lg.Generate(a)

	s.NoError(err)
	s.Equal("JH07051997000000", ln)

}
