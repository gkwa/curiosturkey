package timespan

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var timespanRegex = regexp.MustCompile(`(\d+(?:\.\d+)?)([yMwdhm])`)

func Parse(s string) (time.Duration, error) {
	matches := timespanRegex.FindAllStringSubmatch(s, -1)
	if matches == nil {
		return 0, fmt.Errorf("invalid timespan format: %s", s)
	}

	var totalDuration time.Duration

	for _, match := range matches {
		value, err := strconv.ParseFloat(match[1], 64)
		if err != nil {
			return 0, fmt.Errorf("invalid timespan value: %s", match[1])
		}

		unit := match[2]
		switch unit {
		case "y":
			totalDuration += time.Duration(value * float64(365*24*time.Hour))
		case "M":
			totalDuration += time.Duration(value * float64(30*24*time.Hour))
		case "w":
			totalDuration += time.Duration(value * float64(7*24*time.Hour))
		case "d":
			totalDuration += time.Duration(value * float64(24*time.Hour))
		case "h":
			totalDuration += time.Duration(value * float64(time.Hour))
		case "m":
			totalDuration += time.Duration(value * float64(time.Minute))
		default:
			return 0, fmt.Errorf("invalid timespan unit: %s", unit)
		}
	}

	return totalDuration, nil
}
