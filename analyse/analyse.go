package analyse

import (
	"github.com/hamburghammer/gmon/stats"
)

type AlertStatus int

const (
	StatusOK = AlertStatus(iota)
	StatusWarning
	StatusAlert
)

type Status struct {
	StatusMessage string
	AlertStatus
}

type Result struct {
	Title       string
	Description string
	Status
}

type Analyser interface {
	Analyse(stats.Data) (Result, error)
}
