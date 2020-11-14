package analyse

import (
	"testing"

	"github.com/hamburghammer/gmon/stats"
	"github.com/stretchr/testify/require"
)

func TestCPURule(t *testing.T) {
	tests := []struct {
		Name    string
		Compare string
		CPU     float64
		Warning float64
		Alert   float64
		Want    AlertStatus
	}{
		{
			Name:    "should be smaller as alert",
			Compare: "<",
			CPU:     0.5,
			Warning: 2,
			Alert:   1,
			Want:    StatusAlert,
		},
		{
			Name:    "should be smaller as warning",
			Compare: "<",
			CPU:     1.5,
			Warning: 2,
			Alert:   1,
			Want:    StatusWarning,
		},
		{
			Name:    "should not match any smaller alert lvl",
			Compare: "<",
			CPU:     3,
			Warning: 2,
			Alert:   1,
			Want:    StatusOK,
		},
		{
			Name:    "should be greater as alert",
			Compare: ">",
			CPU:     3,
			Warning: 2,
			Alert:   1,
			Want:    StatusAlert,
		},
		{
			Name:    "should be greater as warning",
			Compare: ">",
			CPU:     3,
			Warning: 2,
			Alert:   4,
			Want:    StatusWarning,
		},
		{
			Name:    "should not match any grater alert lvl",
			Compare: ">",
			CPU:     1,
			Warning: 2,
			Alert:   3,
			Want:    StatusOK,
		},
		{
			Name:    "should be equal as alert",
			Compare: "=",
			CPU:     2,
			Warning: 1,
			Alert:   2,
			Want:    StatusAlert,
		},
		{
			Name:    "should be equal as warning",
			Compare: "=",
			CPU:     2,
			Warning: 2,
			Alert:   4,
			Want:    StatusWarning,
		},
		{
			Name:    "should not match any equal alert lvl",
			Compare: "=",
			CPU:     1,
			Warning: 2,
			Alert:   3,
			Want:    StatusOK,
		},
		{
			Name:    "should be not equal as alert",
			Compare: "!=",
			CPU:     2,
			Warning: 2,
			Alert:   1,
			Want:    StatusAlert,
		},
		{
			Name:    "should be not equal as warning",
			Compare: "!=",
			CPU:     2,
			Warning: 4,
			Alert:   2,
			Want:    StatusWarning,
		},
		{
			Name:    "should not match any not equal alert lvl",
			Compare: "!=",
			CPU:     1,
			Warning: 1,
			Alert:   1,
			Want:    StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			cpuRule := &CPURule{Rule: Rule{Compare: test.Compare}, Alert: test.Alert, Warning: test.Warning}

			got, err := cpuRule.Analyse(stats.Data{CPU: test.CPU})

			require.NoError(t, err)
			require.Equal(t, test.Want, got.AlertStatus)
		})
	}

	t.Run("throw error if no compare character does not match", func(t *testing.T) {
		cpuRule := &CPURule{Rule: Rule{Compare: "foo", Name: "wrong compare"}}

		_, got := cpuRule.Analyse(stats.Data{})
		want := "CPU rule 'wrong compare': The compare character does not match neither '>', '<', '=' or '!='"

		require.Equal(t, want, got.Error())
	})
}
