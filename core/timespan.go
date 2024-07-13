package core

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseTimespan(s string) (time.Duration, error) {
	s = strings.TrimSpace(s)
	if len(s) < 2 {
		return 0, fmt.Errorf("invalid timespan format: %s", s)
	}

	value, err := strconv.Atoi(s[:len(s)-1])
	if err != nil {
		return 0, fmt.Errorf("invalid timespan format: %s", s)
	}

	unit := strings.ToLower(s[len(s)-1:])
	switch unit {
	case "y":
		return time.Duration(value) * 365 * 24 * time.Hour, nil
	case "m":
		return time.Duration(value) * 30 * 24 * time.Hour, nil
	case "w":
		return time.Duration(value) * 7 * 24 * time.Hour, nil
	case "d":
		return time.Duration(value) * 24 * time.Hour, nil
	case "h":
		return time.Duration(value) * time.Hour, nil
	default:
		return 0, fmt.Errorf("invalid timespan unit: %s", unit)
	}
}
