package drivingliscencegenerator

type (
	ValidApplicant struct {
		initials string
		dob      string
	}
	SpyLogger struct {
		Callcount   int
		Lastmessage string
	}
)

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
