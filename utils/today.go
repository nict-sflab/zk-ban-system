package utils

import (
	"time"
)

func Today() int64 {
	t := time.Now()
	day := time.Hour * 24

	today := t.UnixNano() / int64(day)
	return today
}
