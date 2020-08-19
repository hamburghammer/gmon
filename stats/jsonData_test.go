package stats

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func getJSONData() jsonData {
	return jsonData{Date: "2020-08-19T15:18:06+02:00", Mem: "100/1000", Disk: "1000/10000"}
}

func TestTransformation(t *testing.T) {

	t.Run("should transform iso date string into a time struct", func(t *testing.T) {
		want := getJSONData().Date
		got, err := getJSONData().transformToData()

		assert.Nil(t, err)
		assert.Equal(t, want, got.Date.Format(time.RFC3339))
	})

	t.Run("should wrap transforming iso date error", func(t *testing.T) {
		wrongDateJSONData := getJSONData()
		wrongDateJSONData.Date = "2020-08-19T15:18:06+02:00xxx"

		want := "transforming jsonData to Data: parsing the json Date string produced an error: parsing time \"2020-08-19T15:18:06+02:00xxx\": extra text: xxx"
		_, got := wrongDateJSONData.transformToData()

		assert.Equal(t, want, got.Error())
	})

	t.Run("should transform mem string into Memory struct", func(t *testing.T) {
		want := Memory{Used: 100, Total: 1000}
		got, err := getJSONData().transformToData()

		assert.Nil(t, err)
		assert.Equal(t, want, got.Mem)
	})

	t.Run("should wrap paring Mem error", func(t *testing.T) {
		wrongMemJSONData := getJSONData()
		wrongMemJSONData.Mem = "100"

		want := "transforming jsonData to Data: mem string: the string is missing the separating slash like 100/1000: 100"
		_, got := wrongMemJSONData.transformToData()

		assert.Equal(t, want, got.Error())
	})

	t.Run("should transform disk string into Memory struct", func(t *testing.T) {
		want := Memory{Used: 1000, Total: 10000}
		got, err := getJSONData().transformToData()

		assert.Nil(t, err)
		assert.Equal(t, want, got.Disk)
	})

	t.Run("should wrap paring Mem error", func(t *testing.T) {
		wrongMemJSONData := getJSONData()
		wrongMemJSONData.Disk = "100"

		want := "transforming jsonData to Data: disk string: the string is missing the separating slash like 100/1000: 100"
		_, got := wrongMemJSONData.transformToData()

		assert.Equal(t, want, got.Error())
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

func TestParseDiskToMemory(t *testing.T) {
	t.Run("should parse the disk field to Memory", func(t *testing.T) {
		want := Memory{Used: 100, Total: 1000}
		got, err := jsonData{Disk: "100/1000"}.parseDiskToMemory()

		assert.Nil(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("should wrap error with disk parsing context", func(t *testing.T) {
		want := "disk string: parsing parts of a the string for the Memory struct: strconv.ParseInt: parsing \"foo\": invalid syntax"
		_, got := jsonData{Disk: "100/foo"}.parseDiskToMemory()

		assert.NotNil(t, got)
		assert.Equal(t, want, got.Error())
	})
}

func TestParseMemToMemory(t *testing.T) {
	t.Run("should parse the Mem field to Memory", func(t *testing.T) {
		want := Memory{Used: 100, Total: 1000}
		got, err := jsonData{Mem: "100/1000"}.parseMemToMemory()

		assert.Nil(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("should wrap error with disk parsing context", func(t *testing.T) {
		want := "mem string: parsing parts of a the string for the Memory struct: strconv.ParseInt: parsing \"foo\": invalid syntax"
		_, got := jsonData{Mem: "100/foo"}.parseMemToMemory()

		assert.NotNil(t, got)
		assert.Equal(t, want, got.Error())
	})
}

func TestParseMemoryString(t *testing.T) {
	t.Run("should build a Memory struct from a valid/simple string", func(t *testing.T) {
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
