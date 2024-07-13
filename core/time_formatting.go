package core

import (
	"fmt"
	"time"
)

func FormatUserFriendlyDuration(d time.Duration) string {
	if d < time.Minute {
		return "just now"
	}

	parts := []string{}

	years := int(d.Hours() / (24 * 365))
	if years > 0 {
		parts = append(parts, fmt.Sprintf("%dy", years))
		d -= time.Duration(years) * 24 * 365 * time.Hour
	}

	months := int(d.Hours() / (24 * 30))
	if months > 0 {
		parts = append(parts, fmt.Sprintf("%dmo", months))
		d -= time.Duration(months) * 24 * 30 * time.Hour
	}

	weeks := int(d.Hours() / (24 * 7))
	if weeks > 0 {
		parts = append(parts, fmt.Sprintf("%dw", weeks))
		d -= time.Duration(weeks) * 24 * 7 * time.Hour
	}

	days := int(d.Hours() / 24)
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dd", days))
	}

	if len(parts) == 0 {
		hours := int(d.Hours())
		if hours > 0 {
			parts = append(parts, fmt.Sprintf("%dh", hours))
		} else {
			minutes := int(d.Minutes())
			parts = append(parts, fmt.Sprintf("%dm", minutes))
		}
	}

	return parts[0]
}
