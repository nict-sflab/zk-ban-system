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

var Period = today
