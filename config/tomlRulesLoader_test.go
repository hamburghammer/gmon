package config

import (
	"testing"

	"github.com/hamburghammer/gmon/analyse"
	"github.com/stretchr/testify/require"
)

func TestYAMLRuleLoader_LoadCPURules(t *testing.T) {
	reader := &mockReader{content: []byte(`[[CPU]]
  Alert = 0.1
  Compare = "<"
  Deactivated = true
  Description = "testing"
  Name = "foo"
  Warning = 1.1`)}

	got, err := TOMLRulesLoader{reader: reader}.Load()
	want := Rules{CPU: []analyse.CPURule{{Rule: analyse.Rule{Name: "foo", Description: "testing", Compare: "<", Deactivated: true}, Alert: 0.1, Warning: 1.1}}}

	require.NoError(t, err)
	require.Equal(t, want, got)
}
