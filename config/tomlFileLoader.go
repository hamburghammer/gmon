package config

import (
	"io/ioutil"
	"os"
)

// TomlFileLoader handels a toml configuration file
type TomlFileLoader struct {
	file *os.File
}

// Load implementation of the Loader interface to load the config from the file
func (tfl TomlFileLoader) Load() (Data, error) {
	data, err := ioutil.ReadAll(tfl.file)
	if err != nil {
		return Data{}, err
	}
	return tomlLoader{toml: &data}.Load()
}

// NewTomlFileLoader constructor for the TomlFileLoader struct
func NewTomlFileLoader(file *os.File) TomlFileLoader {
	return TomlFileLoader{file: file}
}
