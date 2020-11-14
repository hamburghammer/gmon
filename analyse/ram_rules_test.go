package analyse

import (
	"testing"

	"github.com/hamburghammer/gmon/stats"
	"github.com/stretchr/testify/require"
)

func TestRAMRule(t *testing.T) {
	tests := []struct {
		Name    string
		Compare string
		Disk    int
		Warning int
		Alert   int
		Want    AlertStatus
	}{
		{
			Name:    "should be smaller as alert",
			Compare: "<",
			Disk:    0,
			Warning: 3,
			Alert:   1,
			Want:    StatusAlert,
		},
		{
			Name:    "should be smaller as warning",
			Compare: "<",
			Disk:    2,
			Warning: 3,
			Alert:   1,
			Want:    StatusWarning,
		},
		{
			Name:    "should not match any smaller alert lvl",
			Compare: "<",
			Disk:    3,
			Warning: 2,
			Alert:   1,
			Want:    StatusOK,
		},
		{
			Name:    "should be greater as alert",
			Compare: ">",
			Disk:    3,
			Warning: 2,
			Alert:   1,
			Want:    StatusAlert,
		},
		{
			Name:    "should be greater as warning",
			Compare: ">",
			Disk:    3,
			Warning: 2,
			Alert:   4,
			Want:    StatusWarning,
		},
		{
			Name:    "should not match any grater alert lvl",
			Compare: ">",
			Disk:    1,
			Warning: 2,
			Alert:   3,
			Want:    StatusOK,
		},
		{
			Name:    "should be equal as alert",
			Compare: "=",
			Disk:    2,
			Warning: 1,
			Alert:   2,
			Want:    StatusAlert,
		},
		{
			Name:    "should be equal as warning",
			Compare: "=",
			Disk:    2,
			Warning: 2,
			Alert:   4,
			Want:    StatusWarning,
		},
		{
			Name:    "should not match any equal alert lvl",
			Compare: "=",
			Disk:    1,
			Warning: 2,
			Alert:   3,
			Want:    StatusOK,
		},
		{
			Name:    "should be not equal as alert",
			Compare: "!=",
			Disk:    2,
			Warning: 2,
			Alert:   1,
			Want:    StatusAlert,
		},
		{
			Name:    "should be not equal as warning",
			Compare: "!=",
			Disk:    2,
			Warning: 4,
			Alert:   2,
			Want:    StatusWarning,
		},
		{
			Name:    "should not match any not equal alert lvl",
			Compare: "!=",
			Disk:    1,
			Warning: 1,
			Alert:   1,
			Want:    StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			ramRule := &RAMRule{Rule: Rule{Compare: test.Compare}, Alert: test.Alert, Warning: test.Warning}

			got, err := ramRule.Analyse(stats.Data{Mem: stats.Memory{Used: test.Disk}})

			require.NoError(t, err)
			require.Equal(t, test.Want, got.AlertStatus)
		})
	}

	t.Run("throw error if no compare character does not match", func(t *testing.T) {
		ramRule := &RAMRule{Rule: Rule{Compare: "foo", Name: "wrong compare"}}

		_, got := ramRule.Analyse(stats.Data{})
		want := "RAM rule 'wrong compare': The compare character does not match neither '>', '<', '=' or '!='"

		require.Equal(t, want, got.Error())
	})
}
