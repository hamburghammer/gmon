package config

import "github.com/hamburghammer/gmon/analyse"

// RulesLoader is the interface that should be implemented if you want to load rules.
type RulesLoader interface {
	Load() (Rules, error)
}

// Rules holds all configured rules.
type Rules struct {
	CPU []analyse.CPURule
}
