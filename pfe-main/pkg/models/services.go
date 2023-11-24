package models

type Service struct {
	Name   string `json:"name"`
	Host   string `json:"host"`
	Status Status
}

// Status is an enum for the status of a service

type Status int64

const (
	ServiceUp Status = iota
	ServiceDown
)

func (s Status) String() string {
	switch s {
	case ServiceUp:
		return "Service Up"
	case ServiceDown:
		return "Service Down"
	default:
		return "Unknown"
	}
}
