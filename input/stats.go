package input

import "time"

// StatsClient is the interface to communicat with the client implementation for extracting the data
type StatsClient interface {
	GetData() (StatsData, error)
}

// StatsData is a datapoint representing a snapshot of the resources of a system to a given time.
type StatsData struct {
	Date      time.Time
	CPU       float64
	Processes []Process
	Disk      Space
	Mem       Space
}

// Process is the representation of a UNIX process with some of it's information
type Process struct {
	Name string
	Pid  int
	CPU  float64
}

// Space represents the usage of disk or RAM space.
// It shows the used and the total space available.
type Space struct {
	Used  int
	Total int
}
