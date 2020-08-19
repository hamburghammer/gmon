package stats

import "time"

// Client is the interface to communicat with the client implementation for extracting the data
type Client interface {
	GetData() (Data, error)
}

// Data is a datapoint representing a snapshot of the resources of a system to a given time.
type Data struct {
	Date      time.Time
	CPU       float64
	Processes []Process
	Disk      Memory
	Mem       Memory
}

// Process is the representation of a UNIX process with some of it's information
type Process struct {
	Name string
	Pid  int
	CPU  float64
}

// Memory represents the usage of disk or RAM space.
// It shows the used and the total space available.
type Memory struct {
	Used  int
	Total int
}
