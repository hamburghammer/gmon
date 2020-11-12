package analyse

import (
	"fmt"

	"github.com/hamburghammer/gmon/stats"
)

// RAMRule represents a rule to analyse the ram usage.
type RAMRule struct {
	Rule
	Warning int
	Alert   int
}

// Analyse analyses the given data based on the rules.
func (rr RAMRule) Analyse(data stats.Data) (Result, error) {
	notification := Result{Title: rr.Name, Description: rr.Description}
	var cf compareIntFunc

	switch rr.Compare {
	case ">":
		cf = func(want int) bool {
			return data.Mem.Used > want
		}
	case "<":
		cf = func(want int) bool {
			return data.Mem.Used < want
		}
	case "=":
		cf = func(want int) bool {
			return data.Mem.Used == want
		}
	case "!=":
		cf = func(want int) bool {
			return data.Mem.Used != want
		}
	default:
		return Result{}, fmt.Errorf("Disk rule '%s': %w", rr.Name, ErrCompareMatching)
	}

	notification.Status = rr.compare(cf, rr.Compare)
	return notification, nil
}

func (rr RAMRule) compare(cf compareIntFunc, compareChar string) Status {
	if rr.Alert != 0 && cf(rr.Alert) {
		return Status{AlertStatus: StatusAlert, StatusMessage: fmt.Sprintf("Disk usage %s as %d", compareChar, rr.Alert)}
	} else if rr.Warning != 0 && cf(rr.Warning) {
		return Status{AlertStatus: StatusWarning, StatusMessage: fmt.Sprintf("Disk usage %s as %d", compareChar, rr.Warning)}
	}

	return Status{AlertStatus: StatusOK, StatusMessage: "Disk usage is OK"}
}
