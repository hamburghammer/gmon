package stats

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransformation(t *testing.T) {
	t.Run("should transform iso date string into a time struct", func(t *testing.T) {
		currentTime := time.Now().Format(time.RFC3339)

		want := currentTime
		got, err := jsonData{Date: currentTime}.transformToData()
		if err != nil {
			t.Errorf("An unexpected error happened: %s", err.Error())
		}

		assert.Equal(t, want, got.Date.Format(time.RFC3339))
	})

}

func TestParseISODateString(t *testing.T) {
	t.Run("should parse correct formatted string", func(t *testing.T) {
		currentTime := time.Now().Format(time.RFC3339)

		want := currentTime
		got, err := jsonData{Date: currentTime}.parseDateToTime()
		if err != nil {
			t.Errorf("An unexpected error happened: %s", err.Error())
		}

		assert.Equal(t, want, got.Format(time.RFC3339))
	})

	t.Run("should return error on wrong string", func(t *testing.T) {
		currentTime := "2020-08-19T15:18:06+02:00xxx"

		want := "parsing the json Date string produced an error: parsing time \"2020-08-19T15:18:06+02:00xxx\": extra text: xxx"
		_, err := jsonData{Date: currentTime}.parseDateToTime()

		assert.NotNil(t, err)
		assert.Equal(t, want, err.Error())
	})
}

func TestParseMemoryString(t *testing.T) {
	t.Run("should build a Space struct from a valid/simple string", func(t *testing.T) {
		want := Memory{Used: 100, Total: 1000}
		got, err := jsonData{}.parseMemoryString("100/1000")
		if err != nil {
			t.Errorf("An unexpected error happened: %s", err.Error())
		}

		assert.Equal(t, want, got)
	})

	t.Run("should return error if not seperated by a slash", func(t *testing.T) {
		want := "the string is missing the separating slash like 100/1000: 100 1000"
		_, err := jsonData{}.parseMemoryString("100 1000")

		assert.NotNil(t, err)
		assert.Equal(t, want, err.Error())
	})

	t.Run("should return error if not seperated by a slash", func(t *testing.T) {
		want := "parsing parts of a the string for the Memory struct: strconv.ParseInt: parsing \"foo\": invalid syntax"
		_, err := jsonData{}.parseMemoryString("foo/1000")

		assert.NotNil(t, err)
		assert.Equal(t, want, err.Error())
	})
}
