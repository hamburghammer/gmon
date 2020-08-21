package config

// Loader interface to load a config
type Loader interface {
	Load() (Data, error)
}

// Data the configuration for all clients and the rules for the evaluations
type Data struct {
	Stats  Stats
	Gotify Gotify
}

// Stats the configuration to configure a stats client
type Stats struct {
	Endpoint string
	Hostname string
	Token    string
}

// Gotify configuration for the gotify client
type Gotify struct {
	Endpoint string
	Token    string
}
