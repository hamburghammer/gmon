package config

import "github.com/hamburghammer/gmon/analyse"

// RulesLoader is the interface that should be implemented if you want to load rules.
type RulesLoader interface {
	Load() (Rules, error)
}

// Rules holds all configured rules.
type Rules struct {
	CPU  []analyse.CPURule
	Disk []analyse.DiskRule
	RAM  []analyse.RAMRule
}

// GetCPU returns the cpu rules as analyser interface
func (r Rules) GetCPU() []analyse.Analyser {
	rules := make([]analyse.Analyser, len(r.CPU))
	for i, r := range r.CPU {
		rules[i] = r
	}

	return rules
}

// GetDisk returns the disk rules as analyser interface
func (r Rules) GetDisk() []analyse.Analyser {
	rules := make([]analyse.Analyser, len(r.Disk))
	for i, r := range r.Disk {
		rules[i] = r
	}

	return rules
}

// GetRAM returns the ram rules as analyser interface
func (r Rules) GetRAM() []analyse.Analyser {
	rules := make([]analyse.Analyser, len(r.RAM))
	for i, r := range r.RAM {
		rules[i] = r
	}

	return rules
}
