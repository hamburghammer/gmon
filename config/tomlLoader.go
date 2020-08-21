package config

import (
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml"
)

type tomlLoader struct {
	toml *[]byte
}

func (tp tomlLoader) Load() (Data, error) {
	data := Data{}
	err := toml.Unmarshal(*tp.toml, &data)
	if err != nil {
		return Data{}, err
	}

	return data, nil
}

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
