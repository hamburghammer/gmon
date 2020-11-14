package analyse

import (
	"fmt"

	"github.com/hamburghammer/gmon/stats"
)

// DiskRule holds one rule to analyse the disk usage.
type DiskRule struct {
	Rule
	Warning int
	Alert   int
}

// Analyse analyses the given data based on the rules.
func (dr DiskRule) Analyse(data stats.Data) (Result, error) {
	notification := Result{Title: dr.Name, Description: dr.Description}
	var cf compareIntFunc

	switch dr.Compare {
	case ">":
		cf = func(want int) bool {
			return data.Disk.Used > want
		}
	case "<":
		cf = func(want int) bool {
			return data.Disk.Used < want
		}
	case "=":
		cf = func(want int) bool {
			return data.Disk.Used == want
		}
	case "!=":
		cf = func(want int) bool {
			return data.Disk.Used != want
		}
	default:
		return Result{}, fmt.Errorf("Disk rule '%s': %w", dr.Name, ErrCompareMatching)
	}

	notification.Status = dr.compare(cf, dr.Compare)
	return notification, nil
}

func (dr DiskRule) compare(cf compareIntFunc, compareChar string) Status {
	if dr.Alert != 0 && cf(dr.Alert) {
		return Status{AlertStatus: StatusAlert, StatusMessage: fmt.Sprintf("Disk usage %s as %d", compareChar, dr.Alert)}
	} else if dr.Warning != 0 && cf(dr.Warning) {
		return Status{AlertStatus: StatusWarning, StatusMessage: fmt.Sprintf("Disk usage %s as %d", compareChar, dr.Warning)}
	}

	return Status{AlertStatus: StatusOK, StatusMessage: "Disk usage is OK"}
}
