package stats

import "time"

// Client is the interface to be implement by a client to access the external data.
type Client interface {
	GetData() (Data, error)
}

// Data is a datapoint representing a snapshot of the resources of a system to a given time.
type Data struct {
	Hostname  string
	Date      time.Time
	CPU       float64
	Processes []Process
	Disk      Memory
	Mem       Memory
}

// Process is the representation of a UNIX process with some of its information.
type Process struct {
	Name string
	Pid  int
	CPU  float64
}

// Memory represents the usage of disk or RAM space.
// It shows the used and the total available space.
type Memory struct {
	Used  int
	Total int
}
