package utils

import (
	"time"
)

var PeriodUnit = time.Hour * 24

func period() int64 {
	t := time.Now()
	today := t.UnixNano() / int64(PeriodUnit)
	return today
}

var Period = period
