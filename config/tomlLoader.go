package config

import (
	"io"

	"github.com/pelletier/go-toml"
)

// TomlLoader handels a toml configuration file
type TomlLoader struct {
	reader io.Reader
}

// Load implementation of the Loader interface to load the config from the file
func (tl TomlLoader) Load() (Data, error) {
	data := Data{}
	err := toml.NewDecoder(tl.reader).Decode(&data)
	if err != nil {
		return Data{}, err
	}

	return data, nil
}

// NewTomlLoader constructor for the TomlFileLoader struct
func NewTomlLoader(reader io.Reader) TomlLoader {
	return TomlLoader{reader: reader}
}
