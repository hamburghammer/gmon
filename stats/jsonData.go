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
	errorMessage := "transforming jsonData to Data"

	date, err := jd.parseDateToTime()
	if err != nil {
		return Data{}, fmt.Errorf("%s: %w", errorMessage, err)
	}

	disk, err := jd.parseDiskToMemory()
	if err != nil {
		return Data{}, fmt.Errorf("%s: %w", errorMessage, err)
	}

	mem, err := jd.parseMemToMemory()
	if err != nil {
		return Data{}, fmt.Errorf("%s: %w", errorMessage, err)
	}

	return Data{Date: date, Disk: disk, Mem: mem, CPU: jd.CPU, Processes: jd.Processes}, nil
}

func (jd jsonData) parseDateToTime() (time.Time, error) {
	t, err := time.Parse(time.RFC3339, jd.Date)
	if err != nil {
		return time.Now(), fmt.Errorf("parsing the json Date string produced an error: %w", err)
	}
	return t, nil
}

func (jd jsonData) parseDiskToMemory() (Memory, error) {
	disk, err := jd.parseMemoryString(jd.Disk)
	if err != nil {
		return Memory{}, fmt.Errorf("disk string: %w", err)
	}
	return disk, err
}

func (jd jsonData) parseMemToMemory() (Memory, error) {
	mem, err := jd.parseMemoryString(jd.Mem)
	if err != nil {
		return Memory{}, fmt.Errorf("mem string: %w", err)
	}
	return mem, err
}

// parseMemoryString parses a string with following structure to an Memory object: "100/1000"
func (jd jsonData) parseMemoryString(s string) (Memory, error) {
	parts := strings.Split(s, "/")
	if len(parts) != 2 {
		return Memory{}, fmt.Errorf("the string is missing the separating slash like 100/1000: %s", s)
	}
	var partsAsInt [2]int
	for i, part := range parts {
		num, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return Memory{}, fmt.Errorf("parsing parts of a the string for the Memory struct: %w", err)
		}
		partsAsInt[i] = int(num)
	}
	return Memory{Used: partsAsInt[0], Total: partsAsInt[1]}, nil
}
