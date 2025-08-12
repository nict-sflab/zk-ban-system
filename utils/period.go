package utils

import (
	"time"
)

var PeriodUnit = time.Hour * 24

var OneDay = time.Hour * 24
var HalfMinutes = time.Minute / 2

func period() int64 {
	t := time.Now()
	today := t.UnixNano() / int64(PeriodUnit)
	return today
}

var Period = period
