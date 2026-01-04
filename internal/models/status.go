package models

type Status struct {
	Status StatusState `json:"status"`
}

type StatusState int

const (
	Away StatusState = iota
	Busy
	Free
	HeadDown
	Meeting
)

func (s StatusState) String() string {
	return [...]string{
		"Away",
		"Busy",
		"Free",
		"Headdown",
		"Meeting",
	}[s]
}

func FromString(s string) StatusState {
	switch s {
	case "Away":
		return Away
	case "Busy":
		return Busy
	case "Free":
		return Free
	case "Headdown":
		return HeadDown
	case "Meeting":
		return Meeting
	default:
		return Away
	}
}
