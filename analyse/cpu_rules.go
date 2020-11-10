package analyse

import (
	"fmt"

	"github.com/hamburghammer/gmon/stats"
)

// CPURule holds one rule to check for the cpu value.
type CPURule struct {
	Rule
	Warning float64
	Alert   float64
}

// Analyse executes the rule on a datapoint.
func (c *CPURule) Analyse(data stats.Data) (Result, error) {
	notification := Result{Title: c.Name, Description: c.Description}
	var cf compareFloatFunc

	switch c.Compare {
	case ">":
		cf = func(want float64) bool {
			return data.CPU > want
		}
	case "<":
		cf = func(want float64) bool {
			return data.CPU < want
		}
	case "=":
		cf = func(want float64) bool {
			return data.CPU == want
		}
	case "!=":
		cf = func(want float64) bool {
			return data.CPU != want
		}
	default:
		return Result{}, fmt.Errorf("CPU rule '%s': %w", c.Name, CompareMatchingError)
	}

	notification.Status = c.compare(cf, c.Compare)
	return notification, nil
}

func (c *CPURule) compare(cf compareFloatFunc, compareChar string) Status {
	if c.Alert != 0 && cf(c.Alert) {
		return Status{AlertStatus: StatusAlert, StatusMessage: fmt.Sprintf("CPU usage %s as %f", compareChar, c.Alert)}
	} else if c.Warning != 0 && cf(c.Warning) {
		return Status{AlertStatus: StatusWarning, StatusMessage: fmt.Sprintf("CPU usage %s as %f", compareChar, c.Warning)}
	}
	return Status{AlertStatus: StatusOK, StatusMessage: "CPU usage is OK"}
}
