package core

import (
	"fmt"
	"time"
)

func FormatUserFriendlyDuration(d time.Duration) string {
	if d < time.Minute {
		return "just now"
	}

	units := []struct {
		name  string
		value time.Duration
	}{
		{"y", 365 * 24 * time.Hour},
		{"mo", 30 * 24 * time.Hour},
		{"w", 7 * 24 * time.Hour},
		{"d", 24 * time.Hour},
		{"h", time.Hour},
		{"m", time.Minute},
	}

	parts := make([]string, 0, 2)
	for _, unit := range units {
		if d >= unit.value {
			count := int(d / unit.value)
			parts = append(parts, fmt.Sprintf("%d%s", count, unit.name))
			d -= time.Duration(count) * unit.value
			if len(parts) == 2 {
				break
			}
		}
	}

	if len(parts) == 0 {
		return "0m"
	}

	return fmt.Sprintf("%s%s", parts[0], parts[1])
}
