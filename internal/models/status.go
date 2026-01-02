package models

import "time"

type Status struct {
	Status  StatusState `json:"status"`
	TimeSet time.Time   `json:"time"`
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
		"HeadDown",
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
	case "HeadDown":
		return HeadDown
	case "Meeting":
		return Meeting
	default:
		return Away
	}
}
