package analyse

import (
	"errors"
)

var CompareMatchingError = errors.New("The compare character does not match neither '>', '<', '=' or '!='")

type compareFloatFunc func(want float64) bool
type compareIntFunc func(want int) bool

type Rule struct {
	Name        string
	Description string
	Deactivated bool
	// Compare has to be '>', '<', '=' or '!='
	Compare string
}

type RAMRule struct {
	Rule
	DiskRule
}
