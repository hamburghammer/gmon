package config

import (
	"io"

	"github.com/pelletier/go-toml"
)

// TomlConfigLoader handels a toml configuration file
type TomlConfigLoader struct {
	reader io.Reader
}

// Load implementation of the Loader interface to load the config from the file
func (tl TomlConfigLoader) Load() (Config, error) {
	data := Config{}
	err := toml.NewDecoder(tl.reader).Decode(&data)
	if err != nil {
		return Config{}, err
	}

	return data, nil
}

// NewTomlConfigLoader constructor for the TomlConfigLoader struct
func NewTomlConfigLoader(reader io.Reader) TomlConfigLoader {
	return TomlConfigLoader{reader: reader}
}
