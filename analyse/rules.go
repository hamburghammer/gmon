package analyse

import (
	"errors"
)

// ErrCompareMatching should be thrown if the compare character can not be used.
var ErrCompareMatching = errors.New("The compare character does not match neither '>', '<', '=' or '!='")

type compareFloatFunc func(want float64) bool
type compareIntFunc func(want int) bool

// Rule represents a generic rule that can be extended by a more specified rule.
type Rule struct {
	Name        string
	Description string
	Deactivated bool
	// Compare has to be '>', '<', '=' or '!='
	Compare string
}
