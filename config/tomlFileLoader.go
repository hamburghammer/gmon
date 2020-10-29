package config

import (
	"io"
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

// TomlLoader handels a toml configuration file
type TomlLoader struct {
	file io.Reader
}

// Load implementation of the Loader interface to load the config from the file
func (tl TomlLoader) Load() (Data, error) {
	tomlData, err := ioutil.ReadAll(tl.file)
	if err != nil {
		return Data{}, err
	}

	data := Data{}
	err = toml.Unmarshal(tomlData, &data)
	if err != nil {
		return Data{}, err
	}

	return data, nil
}

// NewTomlFileLoader constructor for the TomlFileLoader struct
func NewTomlFileLoader(file io.Reader) TomlLoader {
	return TomlLoader{file: file}
}
