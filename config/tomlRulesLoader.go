package config

import (
	"io"

	"github.com/pelletier/go-toml"
)

// TOMLRulesLoader handels the reading of rules.
type TOMLRulesLoader struct {
	reader io.Reader
}

// Load reads the rules and returns them.
func (trl TOMLRulesLoader) Load() (Rules, error) {
	rules := Rules{}
	err := toml.NewDecoder(trl.reader).Decode(&rules)
	if err != nil {
		return Rules{}, err
	}

	return rules, nil
}

// NewTOMLRulesLoader is a constructor for the TOMLRulesLoader.
func NewTOMLRulesLoader(reader io.Reader) TOMLRulesLoader {
	return TOMLRulesLoader{reader: reader}
}
