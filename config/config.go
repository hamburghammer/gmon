package config

import "errors"

// Loader interface to load a config
type Loader interface {
	Load() (Config, error)
}

// Config the configuration for all clients and the rules for the evaluations
type Config struct {
	Stats  Stats
	Gotify Gotify

	// Interval configures the rate in with it should check for updates.
	// It's given as minutes -> ones every x minutes.
	Interval int
}

// Stats the configuration to configure a stats client
type Stats struct {
	Endpoint string
	Hostname string
	Token    string
}

// Validate checks each field of Stats
func (s *Stats) Validate() []error {
	errs := make([]error, 0)
	if s.Endpoint == "" {
		errs = append(errs, errors.New("Missing Endpoint inside the Stats config"))
	}
	if s.Hostname == "" {
		errs = append(errs, errors.New("Missing Hostname inside the Stats config"))
	}
	if s.Token == "" {
		errs = append(errs, errors.New("Missing Token inside the Stats config"))
	}

	return errs
}

// Gotify configuration for the gotify client
type Gotify struct {
	Endpoint string
	Token    string
}

// Validate checks each field of Gotify
func (g *Gotify) Validate() []error {
	errs := make([]error, 0)
	if g.Endpoint == "" {
		errs = append(errs, errors.New("Missing Endpoint inside the Gotify config"))
	}
	if g.Token == "" {
		errs = append(errs, errors.New("Missing Token inside the Gotify config"))
	}

	return errs
}
