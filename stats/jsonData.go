package stats

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type jsonData struct {
	Date      string
	CPU       float64
	Processes []Process
	Disk      string
	Mem       string
}

func (jd jsonData) transformToData() (Data, error) {
	date, err := jd.parseISODateString(jd.Date)
	if err != nil {
		return Data{}, fmt.Errorf("transforming the JSON Data object produced an error: %w", err)
	}
	return Data{Date: date}, nil
}

func (jd jsonData) parseISODateString(dateStr string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return time.Now(), fmt.Errorf("parsing the json Date string produced an error: %w", err)
	}
	return t, nil
}

// parseSpaceString parses a string with following structure to an Space object: "100/1000"
func (jd jsonData) parseSpaceString(s string) (Space, error) {
	parts := strings.Split(s, "/")
	if len(parts) != 2 {
		return Space{}, fmt.Errorf("the string is missing the separating slash like 100/1000: %s", s)
	}
	var partsAsInt [2]int
	for i, part := range parts {
		num, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return Space{}, fmt.Errorf("parsing parts of a string representing a Space struct: %w", err)
		}
		partsAsInt[i] = int(num)
	}
	return Space{Used: partsAsInt[0], Total: partsAsInt[1]}, nil
}
