package analyse

import (
	"github.com/hamburghammer/gmon/stats"
)

// AlertStatus represents the result of an analyses.
type AlertStatus int

const (
	// StatusOK if the result is ok.
	StatusOK = AlertStatus(iota)
	// StatusWarning if the result should produce a warning.
	StatusWarning
	// StatusAlert if the result should produce an alert.
	StatusAlert
)

// Status representation of the result of an analysis with a status message and a status.
type Status struct {
	StatusMessage string
	AlertStatus
}

// Result holds all information of an analysis.
type Result struct {
	Title       string
	Description string
	Status
}

// Analyser is an interface that should be implemented by actors that analyse data based on some rules.
type Analyser interface {
	Analyse(stats.Data) (Result, error)
}
