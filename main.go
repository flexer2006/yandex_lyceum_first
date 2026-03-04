package main

import (
	"strconv"
	"time"
)

const (
	minute = int64(60)
	hour   = 60 * minute
	day    = 24 * hour
	month  = 30 * day
	year   = 365 * day
)

func TimeAgo(pastTime time.Time) string {
	seconds := int64(time.Since(pastTime) / time.Second)
	if seconds < 60 {
		return "just now"
	}

	var value int64
	var unit string

	switch {
	case seconds >= year:
		value, unit = seconds/year, "year"
	case seconds >= month:
		value, unit = seconds/month, "month"
	case seconds >= day:
		value, unit = seconds/day, "day"
	case seconds >= hour:
		value, unit = seconds/hour, "hour"
	default:
		value, unit = seconds/minute, "minute"
	}

	n := strconv.FormatInt(value, 10)
	if value == 1 {
		return n + " " + unit + " ago"
	}

	return n + " " + unit + "s ago"
}
