package config

import (
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
