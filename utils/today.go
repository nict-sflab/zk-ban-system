package utils

import (
	"time"
)

func today() int64 {
	t := time.Now()
	day := time.Hour * 24

	today := t.UnixNano() / int64(day)
	return today
}

func shortPeriod() int64 {
	t := time.Now()
	halfMinutes := time.Second * 30

	today := t.UnixNano() / int64(halfMinutes)
	return today
}

var Period = shortPeriod
